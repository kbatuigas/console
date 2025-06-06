// Copyright 2024 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package topic

import (
	"errors"
	"fmt"

	commonv1alpha2 "buf.build/gen/go/redpandadata/common/protocolbuffers/go/redpanda/api/common/v1alpha1"
	"connectrpc.com/connect"
	"github.com/twmb/franz-go/pkg/kerr"

	apierrors "github.com/redpanda-data/console/backend/pkg/api/connect/errors"
	v1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1"
)

// handleKafkaTopicError handles topic specific error codes, such as UNKNOWN_TOPIC_OR_PARTITION and
// translates these into a connect.Error with more information. If there's no error, nil will
// be returned.
func (*Service) handleKafkaTopicError(kafkaErrorCode int16, errorMessage *string) *connect.Error {
	if kafkaErrorCode == 0 {
		return nil
	}

	kafkaErr := kerr.ErrorForCode(kafkaErrorCode)
	switch {
	case errors.Is(kafkaErr, kerr.UnknownTopicOrPartition):
		return apierrors.NewConnectError(
			connect.CodeNotFound,
			errors.New("the requested topic does not exist"),
			apierrors.NewErrorInfo(
				commonv1alpha2.Reason_REASON_RESOURCE_NOT_FOUND.String(),
			))
	default:
		return apierrors.NewConnectErrorFromKafkaErrorCode(kafkaErrorCode, errorMessage)
	}
}

func (*Service) handleKafkaTopicPartitionError(kafkaErr error, errorMessage string) *connect.Error {
	var code connect.Code

	switch {
	case errors.Is(kafkaErr, kerr.UnknownTopicOrPartition):
		code = connect.CodeNotFound
	case errors.Is(kafkaErr, kerr.InvalidTopicException):
		code = connect.CodeNotFound
	case errors.Is(kafkaErr, kerr.InvalidPartitions):
		code = connect.CodeInvalidArgument
	case errors.Is(kafkaErr, kerr.InvalidRequest):
		code = connect.CodeInvalidArgument
	default:
		code = connect.CodeInternal
	}

	err := kafkaErr
	if errorMessage != "" {
		err = fmt.Errorf(errorMessage+": %w", err)
	}

	return apierrors.NewConnectError(
		code,
		err,
		apierrors.NewErrorInfo(v1.Reason_REASON_KAFKA_API_ERROR.String(), apierrors.KeyValsFromKafkaError(kafkaErr)...),
	)
}
