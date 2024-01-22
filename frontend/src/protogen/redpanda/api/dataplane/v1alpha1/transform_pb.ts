// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1alpha1/transform.proto (package redpanda.api.dataplane.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.TransformMetadata
 */
export class TransformMetadata extends Message<TransformMetadata> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string input_topic_name = 2;
   */
  inputTopicName = "";

  /**
   * @generated from field: repeated string output_topic_names = 3;
   */
  outputTopicNames: string[] = [];

  /**
   * @generated from field: repeated redpanda.api.dataplane.v1alpha1.PartitionTransformStatus status = 4;
   */
  status: PartitionTransformStatus[] = [];

  constructor(data?: PartialMessage<TransformMetadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.TransformMetadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "input_topic_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "output_topic_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "status", kind: "message", T: PartitionTransformStatus, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TransformMetadata {
    return new TransformMetadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TransformMetadata {
    return new TransformMetadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TransformMetadata {
    return new TransformMetadata().fromJsonString(jsonString, options);
  }

  static equals(a: TransformMetadata | PlainMessage<TransformMetadata> | undefined, b: TransformMetadata | PlainMessage<TransformMetadata> | undefined): boolean {
    return proto3.util.equals(TransformMetadata, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.PartitionTransformStatus
 */
export class PartitionTransformStatus extends Message<PartitionTransformStatus> {
  /**
   * @generated from field: int32 node_id = 1;
   */
  nodeId = 0;

  /**
   * @generated from field: int32 partition = 2;
   */
  partition = 0;

  /**
   * @generated from field: redpanda.api.dataplane.v1alpha1.PartitionTransformStatus.PartitionStatus status = 3;
   */
  status = PartitionTransformStatus_PartitionStatus.UNSPECIFIED;

  /**
   * @generated from field: int32 lag = 4;
   */
  lag = 0;

  constructor(data?: PartialMessage<PartitionTransformStatus>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.PartitionTransformStatus";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "node_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "partition", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(PartitionTransformStatus_PartitionStatus) },
    { no: 4, name: "lag", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PartitionTransformStatus {
    return new PartitionTransformStatus().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PartitionTransformStatus {
    return new PartitionTransformStatus().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PartitionTransformStatus {
    return new PartitionTransformStatus().fromJsonString(jsonString, options);
  }

  static equals(a: PartitionTransformStatus | PlainMessage<PartitionTransformStatus> | undefined, b: PartitionTransformStatus | PlainMessage<PartitionTransformStatus> | undefined): boolean {
    return proto3.util.equals(PartitionTransformStatus, a, b);
  }
}

/**
 * @generated from enum redpanda.api.dataplane.v1alpha1.PartitionTransformStatus.PartitionStatus
 */
export enum PartitionTransformStatus_PartitionStatus {
  /**
   * @generated from enum value: PARTITION_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PARTITION_STATUS_RUNNING = 1;
   */
  RUNNING = 1,

  /**
   * @generated from enum value: PARTITION_STATUS_INACTIVE = 2;
   */
  INACTIVE = 2,

  /**
   * @generated from enum value: PARTITION_STATUS_ERRORED = 3;
   */
  ERRORED = 3,

  /**
   * @generated from enum value: PARTITION_STATUS_UNKNOWN = 4;
   */
  UNKNOWN = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(PartitionTransformStatus_PartitionStatus)
proto3.util.setEnumType(PartitionTransformStatus_PartitionStatus, "redpanda.api.dataplane.v1alpha1.PartitionTransformStatus.PartitionStatus", [
  { no: 0, name: "PARTITION_STATUS_UNSPECIFIED" },
  { no: 1, name: "PARTITION_STATUS_RUNNING" },
  { no: 2, name: "PARTITION_STATUS_INACTIVE" },
  { no: 3, name: "PARTITION_STATUS_ERRORED" },
  { no: 4, name: "PARTITION_STATUS_UNKNOWN" },
]);

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.ListTransformsRequest
 */
export class ListTransformsRequest extends Message<ListTransformsRequest> {
  /**
   * @generated from field: redpanda.api.dataplane.v1alpha1.ListTransformsRequest.Filter filter = 1;
   */
  filter?: ListTransformsRequest_Filter;

  /**
   * @generated from field: string page_token = 2;
   */
  pageToken = "";

  constructor(data?: PartialMessage<ListTransformsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.ListTransformsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "filter", kind: "message", T: ListTransformsRequest_Filter },
    { no: 2, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListTransformsRequest {
    return new ListTransformsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListTransformsRequest {
    return new ListTransformsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListTransformsRequest {
    return new ListTransformsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListTransformsRequest | PlainMessage<ListTransformsRequest> | undefined, b: ListTransformsRequest | PlainMessage<ListTransformsRequest> | undefined): boolean {
    return proto3.util.equals(ListTransformsRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.ListTransformsRequest.Filter
 */
export class ListTransformsRequest_Filter extends Message<ListTransformsRequest_Filter> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<ListTransformsRequest_Filter>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.ListTransformsRequest.Filter";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListTransformsRequest_Filter {
    return new ListTransformsRequest_Filter().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListTransformsRequest_Filter {
    return new ListTransformsRequest_Filter().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListTransformsRequest_Filter {
    return new ListTransformsRequest_Filter().fromJsonString(jsonString, options);
  }

  static equals(a: ListTransformsRequest_Filter | PlainMessage<ListTransformsRequest_Filter> | undefined, b: ListTransformsRequest_Filter | PlainMessage<ListTransformsRequest_Filter> | undefined): boolean {
    return proto3.util.equals(ListTransformsRequest_Filter, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.ListTransformsResponse
 */
export class ListTransformsResponse extends Message<ListTransformsResponse> {
  /**
   * @generated from field: string next_page_token = 1;
   */
  nextPageToken = "";

  /**
   * @generated from field: repeated redpanda.api.dataplane.v1alpha1.TransformMetadata transforms = 2;
   */
  transforms: TransformMetadata[] = [];

  constructor(data?: PartialMessage<ListTransformsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.ListTransformsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "transforms", kind: "message", T: TransformMetadata, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListTransformsResponse {
    return new ListTransformsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListTransformsResponse {
    return new ListTransformsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListTransformsResponse {
    return new ListTransformsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListTransformsResponse | PlainMessage<ListTransformsResponse> | undefined, b: ListTransformsResponse | PlainMessage<ListTransformsResponse> | undefined): boolean {
    return proto3.util.equals(ListTransformsResponse, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.GetTransformRequest
 */
export class GetTransformRequest extends Message<GetTransformRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<GetTransformRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.GetTransformRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTransformRequest {
    return new GetTransformRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTransformRequest {
    return new GetTransformRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTransformRequest {
    return new GetTransformRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetTransformRequest | PlainMessage<GetTransformRequest> | undefined, b: GetTransformRequest | PlainMessage<GetTransformRequest> | undefined): boolean {
    return proto3.util.equals(GetTransformRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.GetTransformResponse
 */
export class GetTransformResponse extends Message<GetTransformResponse> {
  /**
   * @generated from field: redpanda.api.dataplane.v1alpha1.TransformMetadata transform = 1;
   */
  transform?: TransformMetadata;

  constructor(data?: PartialMessage<GetTransformResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.GetTransformResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "transform", kind: "message", T: TransformMetadata },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTransformResponse {
    return new GetTransformResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTransformResponse {
    return new GetTransformResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTransformResponse {
    return new GetTransformResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetTransformResponse | PlainMessage<GetTransformResponse> | undefined, b: GetTransformResponse | PlainMessage<GetTransformResponse> | undefined): boolean {
    return proto3.util.equals(GetTransformResponse, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.DeleteTransformRequest
 */
export class DeleteTransformRequest extends Message<DeleteTransformRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<DeleteTransformRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.DeleteTransformRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTransformRequest {
    return new DeleteTransformRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTransformRequest {
    return new DeleteTransformRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTransformRequest {
    return new DeleteTransformRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTransformRequest | PlainMessage<DeleteTransformRequest> | undefined, b: DeleteTransformRequest | PlainMessage<DeleteTransformRequest> | undefined): boolean {
    return proto3.util.equals(DeleteTransformRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.DeleteTransformResponse
 */
export class DeleteTransformResponse extends Message<DeleteTransformResponse> {
  constructor(data?: PartialMessage<DeleteTransformResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.DeleteTransformResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTransformResponse {
    return new DeleteTransformResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTransformResponse {
    return new DeleteTransformResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTransformResponse {
    return new DeleteTransformResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTransformResponse | PlainMessage<DeleteTransformResponse> | undefined, b: DeleteTransformResponse | PlainMessage<DeleteTransformResponse> | undefined): boolean {
    return proto3.util.equals(DeleteTransformResponse, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.DeployTransformRequest
 */
export class DeployTransformRequest extends Message<DeployTransformRequest> {
  /**
   * 'name' is the unique identifier for the transform. It must be a non-empty string,
   * with a maximum length of 128 bytes. It should contain only characters that are
   * valid UTF-8 and not control or format characters.
   *
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * WASM binary for the transform. The binary must be a valid WASM binary
   *
   * @generated from field: bytes wasm_binary = 2;
   */
  wasmBinary = new Uint8Array(0);

  /**
   * 'input_topic_name' specifies the name of the input topic for the transform.
   *
   * @generated from field: string input_topic_name = 3;
   */
  inputTopicName = "";

  /**
   * 'output_topic_names' is a list of output topics for the transform. This field can contain
   * multiple strings, each representing the name of an output topic.
   *
   * @generated from field: repeated string output_topic_names = 4;
   */
  outputTopicNames: string[] = [];

  /**
   * 'environment' is a map representing key-value pairs of environment configurations.
   * Each key must be a non-empty string, not exceeding 128 bytes, and matching the
   * specified pattern. Each value must be a string not exceeding 2048 bytes (2 KiB).
   * The map can have at most 128 key-value pairs.
   *
   * @generated from field: map<string, string> environment = 5;
   */
  environment: { [key: string]: string } = {};

  constructor(data?: PartialMessage<DeployTransformRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.DeployTransformRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "wasm_binary", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "input_topic_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "output_topic_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "environment", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeployTransformRequest {
    return new DeployTransformRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeployTransformRequest {
    return new DeployTransformRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeployTransformRequest {
    return new DeployTransformRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeployTransformRequest | PlainMessage<DeployTransformRequest> | undefined, b: DeployTransformRequest | PlainMessage<DeployTransformRequest> | undefined): boolean {
    return proto3.util.equals(DeployTransformRequest, a, b);
  }
}

/**
 * @generated from message redpanda.api.dataplane.v1alpha1.DeployTransformResponse
 */
export class DeployTransformResponse extends Message<DeployTransformResponse> {
  /**
   * 'transform' contains the metadata for the deployed transform. This includes
   * details like the name, input and output topics, status
   *
   * @generated from field: redpanda.api.dataplane.v1alpha1.TransformMetadata transform = 1;
   */
  transform?: TransformMetadata;

  constructor(data?: PartialMessage<DeployTransformResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1alpha1.DeployTransformResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "transform", kind: "message", T: TransformMetadata },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeployTransformResponse {
    return new DeployTransformResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeployTransformResponse {
    return new DeployTransformResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeployTransformResponse {
    return new DeployTransformResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeployTransformResponse | PlainMessage<DeployTransformResponse> | undefined, b: DeployTransformResponse | PlainMessage<DeployTransformResponse> | undefined): boolean {
    return proto3.util.equals(DeployTransformResponse, a, b);
  }
}

