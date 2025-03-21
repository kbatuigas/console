// @generated by protoc-gen-es v1.6.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1/common.proto (package redpanda.api.dataplane.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum redpanda.api.dataplane.v1.ConfigSource
 */
export enum ConfigSource {
  /**
   * @generated from enum value: CONFIG_SOURCE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CONFIG_SOURCE_DYNAMIC_TOPIC_CONFIG = 1;
   */
  DYNAMIC_TOPIC_CONFIG = 1,

  /**
   * @generated from enum value: CONFIG_SOURCE_DYNAMIC_BROKER_CONFIG = 2;
   */
  DYNAMIC_BROKER_CONFIG = 2,

  /**
   * @generated from enum value: CONFIG_SOURCE_DYNAMIC_DEFAULT_BROKER_CONFIG = 3;
   */
  DYNAMIC_DEFAULT_BROKER_CONFIG = 3,

  /**
   * @generated from enum value: CONFIG_SOURCE_STATIC_BROKER_CONFIG = 4;
   */
  STATIC_BROKER_CONFIG = 4,

  /**
   * @generated from enum value: CONFIG_SOURCE_DEFAULT_CONFIG = 5;
   */
  DEFAULT_CONFIG = 5,

  /**
   * @generated from enum value: CONFIG_SOURCE_DYNAMIC_BROKER_LOGGER_CONFIG = 6;
   */
  DYNAMIC_BROKER_LOGGER_CONFIG = 6,
}
// Retrieve enum metadata with: proto3.getEnumType(ConfigSource)
proto3.util.setEnumType(ConfigSource, "redpanda.api.dataplane.v1.ConfigSource", [
  { no: 0, name: "CONFIG_SOURCE_UNSPECIFIED" },
  { no: 1, name: "CONFIG_SOURCE_DYNAMIC_TOPIC_CONFIG" },
  { no: 2, name: "CONFIG_SOURCE_DYNAMIC_BROKER_CONFIG" },
  { no: 3, name: "CONFIG_SOURCE_DYNAMIC_DEFAULT_BROKER_CONFIG" },
  { no: 4, name: "CONFIG_SOURCE_STATIC_BROKER_CONFIG" },
  { no: 5, name: "CONFIG_SOURCE_DEFAULT_CONFIG" },
  { no: 6, name: "CONFIG_SOURCE_DYNAMIC_BROKER_LOGGER_CONFIG" },
]);

/**
 * @generated from enum redpanda.api.dataplane.v1.ConfigType
 */
export enum ConfigType {
  /**
   * @generated from enum value: CONFIG_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CONFIG_TYPE_BOOLEAN = 1;
   */
  BOOLEAN = 1,

  /**
   * @generated from enum value: CONFIG_TYPE_STRING = 2;
   */
  STRING = 2,

  /**
   * @generated from enum value: CONFIG_TYPE_INT = 3;
   */
  INT = 3,

  /**
   * @generated from enum value: CONFIG_TYPE_SHORT = 4;
   */
  SHORT = 4,

  /**
   * @generated from enum value: CONFIG_TYPE_LONG = 5;
   */
  LONG = 5,

  /**
   * @generated from enum value: CONFIG_TYPE_DOUBLE = 6;
   */
  DOUBLE = 6,

  /**
   * @generated from enum value: CONFIG_TYPE_LIST = 7;
   */
  LIST = 7,

  /**
   * @generated from enum value: CONFIG_TYPE_CLASS = 8;
   */
  CLASS = 8,

  /**
   * @generated from enum value: CONFIG_TYPE_PASSWORD = 9;
   */
  PASSWORD = 9,
}
// Retrieve enum metadata with: proto3.getEnumType(ConfigType)
proto3.util.setEnumType(ConfigType, "redpanda.api.dataplane.v1.ConfigType", [
  { no: 0, name: "CONFIG_TYPE_UNSPECIFIED" },
  { no: 1, name: "CONFIG_TYPE_BOOLEAN" },
  { no: 2, name: "CONFIG_TYPE_STRING" },
  { no: 3, name: "CONFIG_TYPE_INT" },
  { no: 4, name: "CONFIG_TYPE_SHORT" },
  { no: 5, name: "CONFIG_TYPE_LONG" },
  { no: 6, name: "CONFIG_TYPE_DOUBLE" },
  { no: 7, name: "CONFIG_TYPE_LIST" },
  { no: 8, name: "CONFIG_TYPE_CLASS" },
  { no: 9, name: "CONFIG_TYPE_PASSWORD" },
]);

/**
 * @generated from enum redpanda.api.dataplane.v1.ConfigAlterOperation
 */
export enum ConfigAlterOperation {
  /**
   * @generated from enum value: CONFIG_ALTER_OPERATION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CONFIG_ALTER_OPERATION_SET = 1;
   */
  SET = 1,

  /**
   * @generated from enum value: CONFIG_ALTER_OPERATION_DELETE = 2;
   */
  DELETE = 2,

  /**
   * @generated from enum value: CONFIG_ALTER_OPERATION_APPEND = 3;
   */
  APPEND = 3,

  /**
   * @generated from enum value: CONFIG_ALTER_OPERATION_SUBTRACT = 4;
   */
  SUBTRACT = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(ConfigAlterOperation)
proto3.util.setEnumType(ConfigAlterOperation, "redpanda.api.dataplane.v1.ConfigAlterOperation", [
  { no: 0, name: "CONFIG_ALTER_OPERATION_UNSPECIFIED" },
  { no: 1, name: "CONFIG_ALTER_OPERATION_SET" },
  { no: 2, name: "CONFIG_ALTER_OPERATION_DELETE" },
  { no: 3, name: "CONFIG_ALTER_OPERATION_APPEND" },
  { no: 4, name: "CONFIG_ALTER_OPERATION_SUBTRACT" },
]);

/**
 * @generated from message redpanda.api.dataplane.v1.ConfigSynonym
 */
export class ConfigSynonym extends Message<ConfigSynonym> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: optional string value = 2;
   */
  value?: string;

  /**
   * @generated from field: redpanda.api.dataplane.v1.ConfigSource source = 3;
   */
  source = ConfigSource.UNSPECIFIED;

  constructor(data?: PartialMessage<ConfigSynonym>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "redpanda.api.dataplane.v1.ConfigSynonym";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "source", kind: "enum", T: proto3.getEnumType(ConfigSource) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConfigSynonym {
    return new ConfigSynonym().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConfigSynonym {
    return new ConfigSynonym().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConfigSynonym {
    return new ConfigSynonym().fromJsonString(jsonString, options);
  }

  static equals(a: ConfigSynonym | PlainMessage<ConfigSynonym> | undefined, b: ConfigSynonym | PlainMessage<ConfigSynonym> | undefined): boolean {
    return proto3.util.equals(ConfigSynonym, a, b);
  }
}

