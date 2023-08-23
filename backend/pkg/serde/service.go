// Copyright 2023 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package serde

import (
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/redpanda-data/console/backend/pkg/msgpack"
	"github.com/redpanda-data/console/backend/pkg/proto"
	"github.com/redpanda-data/console/backend/pkg/schema"
)

// Service is the struct that holds all dependencies that are required to deserialize
// a record.
type Service struct {
	SerDes []Serde
}

const defaultMaxPayloadSize = 1_000_000 // 1 MB

func NewService(schemaService *schema.Service, protoSvc *proto.Service, msgPackSvc *msgpack.Service) *Service {
	return &Service{
		SerDes: []Serde{
			NoneSerde{},
			JsonSerde{},
			JsonSchemaSerde{},
			XMLSerde{},
			AvroSerde{SchemaSvc: schemaService},
			ProtobufSerde{ProtoSvc: protoSvc},
			ProtobufSchemaSerde{ProtoSvc: protoSvc},
			MsgPackSerde{MsgPackService: msgPackSvc},
			SmileSerde{},
			UTF8Serde{},
			TextSerde{},
		},
	}
}

// DeserializeRecord tries to deserialize a Kafka record into a struct that
// can be processed by the Frontend.
func (s *Service) DeserializeRecord(record *kgo.Record, opts DeserializationOptions) *Record {
	// defaults
	if opts.MaxPayloadSize <= 0 {
		opts.MaxPayloadSize = defaultMaxPayloadSize
	}

	// 1. Test if it's a known binary Format
	if record.Topic == "__consumer_offsets" {
		rec, err := s.deserializeConsumerOffset(record)
		if err == nil {
			return rec
		}
	}

	// 2. Deserialize key & value separately
	key := s.deserializePayload(record, PayloadTypeKey, &opts)
	val := s.deserializePayload(record, PayloadTypeValue, &opts)
	headers := recordHeaders(record)

	return &Record{
		Key:     key,
		Value:   val,
		Headers: headers,
	}
}

// deserializePayload deserializes either the key or value of a Kafka record by trying
// the pre-defined deserialization strategies.
func (s *Service) deserializePayload(record *kgo.Record, payloadType PayloadType, opts *DeserializationOptions) RecordPayload {
	payload := payloadFromRecord(record, payloadType)

	// Check if payload is empty
	if len(payload) == 0 {
		return RecordPayload{
			OriginalPayload:  payload,
			IsPayloadNull:    payload == nil,
			PayloadSizeBytes: 0,
			Encoding:         PayloadEncodingNone,
		}
	}

	troubleshooting := make([]TroubleshootingReport, 0)

	// Try all registered SerDes in the order they were registered
	for _, serde := range s.SerDes {
		rp, err := serde.DeserializePayload(record, payloadType)
		if err != nil {
			troubleshooting = append(troubleshooting, TroubleshootingReport{
				SerdeName: string(serde.Name()),
				Message:   err.Error(),
			})
		} else {
			// Serde deserialized successfully, let's add fields that always shall
			// be set, regardless of the SerDe used.
			rp.PayloadSizeBytes = len(payload)
			rp.IsPayloadNull = payload == nil
			if len(payload) > opts.MaxPayloadSize {
				rp.OriginalPayload = nil
				rp.IsPayloadTooLarge = true
			} else {
				rp.OriginalPayload = payload
				rp.IsPayloadTooLarge = false
			}

			if opts.Troubleshoot {
				rp.Troubleshooting = troubleshooting
			}

			return rp
		}
	}

	// Anything else is considered binary
	rp := RecordPayload{
		PayloadSizeBytes: len(payload),
		IsPayloadNull:    payload == nil,
		Encoding:         PayloadEncodingBinary,
	}

	if len(payload) > opts.MaxPayloadSize {
		rp.OriginalPayload = nil
		rp.IsPayloadTooLarge = true
	} else {
		rp.OriginalPayload = payload
		rp.IsPayloadTooLarge = false
	}

	if opts.Troubleshoot {
		rp.Troubleshooting = troubleshooting
	}

	return rp
}

// DeserializationOptions that can be provided by the requester to influence
// the deserialization.
type DeserializationOptions struct {
	// KeyEncoding may be specified by the frontend to indicate that this
	// encoding type shall be used to deserialize the key. This is helpful,
	// if the requester knows that a primitive type like int16 is used, which couldn't
	// be guessed automatically. No other encoding method will be tried if this
	// method has failed. Troubleshoot should always be set to true in this case.
	KeyEncoding PayloadEncoding

	// PreferredValueEncoding may be specified by the frontend to indicate that this
	// encoding type shall be used to deserialize the value. This is helpful,
	// if the requester knows that a primitive type like int16 is used, which couldn't
	// be guessed automatically. No other encoding method will be tried if this
	// method has failed. Troubleshoot should always be set to true in this case.
	ValueEncoding PayloadEncoding

	// Troubleshoot can be enabled to return additional information which reports
	// why each performed deserialization strategy has failed. If the first
	// tested encoding method worked successfully no troubleshooting information
	// is returned.
	Troubleshoot bool

	// MaxPayloadSize is the maximum size of the payload.
	MaxPayloadSize int
}

func (s *Service) SerializeRecord(input SerializeInput) (*SerializeOutput, error) {
	var keySerResult RecordPayloadSerializeResult
	var valueSerResult RecordPayloadSerializeResult

	var err error
	sr := SerializeOutput{}

	// key
	keyTS := make([]TroubleshootingReport, 0)
	found := false
	for _, serde := range s.SerDes {
		if input.Key.Encoding == serde.Name() {
			found = true

			bytes, serErr := serde.SerializeObject(input.Key.Payload, PayloadTypeKey, input.Key.Options...)
			if serErr != nil {
				keyTS = append(keyTS, TroubleshootingReport{
					SerdeName: string(serde.Name()),
					Message:   serErr.Error(),
				})
			} else {
				keySerResult.Encoding = serde.Name()
				keySerResult.Payload = bytes
				keySerResult.Troubleshooting = keyTS
			}
		}
	}

	sr.Key = &keySerResult

	if !found {
		err = fmt.Errorf("invalid encoding for key: %s", input.Key.Encoding)
	}

	if err != nil {
		return &sr, err
	}

	valueTS := make([]TroubleshootingReport, 0)
	found = false
	for _, serde := range s.SerDes {
		if input.Value.Encoding == serde.Name() {
			found = true

			bytes, serErr := serde.SerializeObject(input.Value.Payload, PayloadTypeValue, input.Value.Options...)
			if serErr != nil {
				valueTS = append(valueTS, TroubleshootingReport{
					SerdeName: string(serde.Name()),
					Message:   serErr.Error(),
				})
			} else {
				valueSerResult.Encoding = serde.Name()
				valueSerResult.Payload = bytes
				valueSerResult.Troubleshooting = valueTS
			}
		}
	}

	sr.Value = &valueSerResult

	if !found {
		err = fmt.Errorf("invalid encoding for value: %s", input.Key.Encoding)
	}

	return &sr, err
}

func payloadFromRecord(record *kgo.Record, payloadType PayloadType) []byte {
	if payloadType == PayloadTypeValue {
		return record.Value
	}
	return record.Key
}