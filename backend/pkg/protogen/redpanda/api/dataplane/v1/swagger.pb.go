// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: redpanda/api/dataplane/v1/swagger.proto

package dataplanev1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_redpanda_api_dataplane_v1_swagger_proto protoreflect.FileDescriptor

var file_redpanda_api_dataplane_v1_swagger_proto_rawDesc = []byte{
	0x0a, 0x27, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x77, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x42, 0x97, 0x03, 0x92, 0x41, 0x82, 0x01, 0x52, 0x31, 0x0a, 0x03, 0x34,
	0x30, 0x31, 0x12, 0x2a, 0x0a, 0x10, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x12, 0x16, 0x0a, 0x14, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x4d,
	0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x46, 0x0a, 0x2c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x20,
	0x52, 0x65, 0x61, 0x63, 0x68, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x12, 0x16, 0x0a, 0x14, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x0a, 0x1d, 0x63,
	0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x77,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x41, 0x44, 0xaa,
	0x02, 0x19, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x52, 0x65,
	0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x25, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e,
	0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1c, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_redpanda_api_dataplane_v1_swagger_proto_goTypes = []any{}
var file_redpanda_api_dataplane_v1_swagger_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_redpanda_api_dataplane_v1_swagger_proto_init() }
func file_redpanda_api_dataplane_v1_swagger_proto_init() {
	if File_redpanda_api_dataplane_v1_swagger_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_dataplane_v1_swagger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redpanda_api_dataplane_v1_swagger_proto_goTypes,
		DependencyIndexes: file_redpanda_api_dataplane_v1_swagger_proto_depIdxs,
	}.Build()
	File_redpanda_api_dataplane_v1_swagger_proto = out.File
	file_redpanda_api_dataplane_v1_swagger_proto_rawDesc = nil
	file_redpanda_api_dataplane_v1_swagger_proto_goTypes = nil
	file_redpanda_api_dataplane_v1_swagger_proto_depIdxs = nil
}
