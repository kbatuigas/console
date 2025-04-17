// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: redpanda/api/console/v1alpha1/console_service.proto

package consolev1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ConsoleService_ListMessages_FullMethodName   = "/redpanda.api.console.v1alpha1.ConsoleService/ListMessages"
	ConsoleService_PublishMessage_FullMethodName = "/redpanda.api.console.v1alpha1.ConsoleService/PublishMessage"
)

// ConsoleServiceClient is the client API for ConsoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ConsoleService represents the Console API service.
type ConsoleServiceClient interface {
	// ListMessages lists the messages according to the requested query.
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListMessagesResponse], error)
	// PublishMessage publishes message.
	PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error)
}

type consoleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsoleServiceClient(cc grpc.ClientConnInterface) ConsoleServiceClient {
	return &consoleServiceClient{cc}
}

func (c *consoleServiceClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListMessagesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ConsoleService_ServiceDesc.Streams[0], ConsoleService_ListMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListMessagesRequest, ListMessagesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ConsoleService_ListMessagesClient = grpc.ServerStreamingClient[ListMessagesResponse]

func (c *consoleServiceClient) PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishMessageResponse)
	err := c.cc.Invoke(ctx, ConsoleService_PublishMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsoleServiceServer is the server API for ConsoleService service.
// All implementations must embed UnimplementedConsoleServiceServer
// for forward compatibility.
//
// ConsoleService represents the Console API service.
type ConsoleServiceServer interface {
	// ListMessages lists the messages according to the requested query.
	ListMessages(*ListMessagesRequest, grpc.ServerStreamingServer[ListMessagesResponse]) error
	// PublishMessage publishes message.
	PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error)
	mustEmbedUnimplementedConsoleServiceServer()
}

// UnimplementedConsoleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConsoleServiceServer struct{}

func (UnimplementedConsoleServiceServer) ListMessages(*ListMessagesRequest, grpc.ServerStreamingServer[ListMessagesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ListMessages not implemented")
}
func (UnimplementedConsoleServiceServer) PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishMessage not implemented")
}
func (UnimplementedConsoleServiceServer) mustEmbedUnimplementedConsoleServiceServer() {}
func (UnimplementedConsoleServiceServer) testEmbeddedByValue()                        {}

// UnsafeConsoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsoleServiceServer will
// result in compilation errors.
type UnsafeConsoleServiceServer interface {
	mustEmbedUnimplementedConsoleServiceServer()
}

func RegisterConsoleServiceServer(s grpc.ServiceRegistrar, srv ConsoleServiceServer) {
	// If the following call pancis, it indicates UnimplementedConsoleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConsoleService_ServiceDesc, srv)
}

func _ConsoleService_ListMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsoleServiceServer).ListMessages(m, &grpc.GenericServerStream[ListMessagesRequest, ListMessagesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ConsoleService_ListMessagesServer = grpc.ServerStreamingServer[ListMessagesResponse]

func _ConsoleService_PublishMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsoleServiceServer).PublishMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsoleService_PublishMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsoleServiceServer).PublishMessage(ctx, req.(*PublishMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsoleService_ServiceDesc is the grpc.ServiceDesc for ConsoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "redpanda.api.console.v1alpha1.ConsoleService",
	HandlerType: (*ConsoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishMessage",
			Handler:    _ConsoleService_PublishMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMessages",
			Handler:       _ConsoleService_ListMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "redpanda/api/console/v1alpha1/console_service.proto",
}
