// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: redpanda/api/dataplane/v1alpha1/kafka_connect.proto

package dataplanev1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	KafkaConnectService_ListConnectClusters_FullMethodName  = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/ListConnectClusters"
	KafkaConnectService_GetConnectCluster_FullMethodName    = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/GetConnectCluster"
	KafkaConnectService_ListConnectors_FullMethodName       = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/ListConnectors"
	KafkaConnectService_CreateConnector_FullMethodName      = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/CreateConnector"
	KafkaConnectService_RestartConnector_FullMethodName     = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/RestartConnector"
	KafkaConnectService_GetConnector_FullMethodName         = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/GetConnector"
	KafkaConnectService_GetConnectorStatus_FullMethodName   = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/GetConnectorStatus"
	KafkaConnectService_PauseConnector_FullMethodName       = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/PauseConnector"
	KafkaConnectService_ResumeConnector_FullMethodName      = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/ResumeConnector"
	KafkaConnectService_StopConnector_FullMethodName        = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/StopConnector"
	KafkaConnectService_DeleteConnector_FullMethodName      = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/DeleteConnector"
	KafkaConnectService_UpsertConnector_FullMethodName      = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/UpsertConnector"
	KafkaConnectService_GetConnectorConfig_FullMethodName   = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/GetConnectorConfig"
	KafkaConnectService_ListConnectorTopics_FullMethodName  = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/ListConnectorTopics"
	KafkaConnectService_ResetConnectorTopics_FullMethodName = "/redpanda.api.dataplane.v1alpha1.KafkaConnectService/ResetConnectorTopics"
)

// KafkaConnectServiceClient is the client API for KafkaConnectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// KafkaConnectService is the service for the Kafka connect, it exposes the
// Kafka Connect API, you can set multiple Kafka connect services and all of
// them can be managed using this service definition, the request is not only
// proxied but also enriched with better error handling and custom
// documentation and configuration
//
// Deprecated: Do not use.
type KafkaConnectServiceClient interface {
	// Deprecated: Do not use.
	// ListConnectClusters implements the list clusters method, list connect
	// clusters available in the console configuration
	ListConnectClusters(ctx context.Context, in *ListConnectClustersRequest, opts ...grpc.CallOption) (*ListConnectClustersResponse, error)
	// Deprecated: Do not use.
	// GetConnectCluster implements the get cluster info method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnectCluster(ctx context.Context, in *GetConnectClusterRequest, opts ...grpc.CallOption) (*GetConnectClusterResponse, error)
	// Deprecated: Do not use.
	// ListConnectors implements the list connectors method, exposes a Kafka
	// Connect equivalent REST endpoint
	ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error)
	// Deprecated: Do not use.
	// CreateConnector implements the create connector method, and exposes an
	// equivalent REST endpoint as the Kafka connect API endpoint
	CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*CreateConnectorResponse, error)
	// Deprecated: Do not use.
	// RestartConnector implements the restart connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	RestartConnector(ctx context.Context, in *RestartConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// GetConnector implements the get connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*GetConnectorResponse, error)
	// Deprecated: Do not use.
	// GetConnectorStatus implement the get status method, Gets the current status of the connector, including:
	// Whether it is running or restarting, or if it has failed or paused
	// Which worker it is assigned to
	// Error information if it has failed
	// The state of all its tasks
	GetConnectorStatus(ctx context.Context, in *GetConnectorStatusRequest, opts ...grpc.CallOption) (*GetConnectorStatusResponse, error)
	// Deprecated: Do not use.
	// PauseConnector implements the pause connector method, exposes a Kafka
	// connect equivalent REST endpoint
	PauseConnector(ctx context.Context, in *PauseConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// ResumeConnector implements the resume connector method, exposes a Kafka
	// connect equivalent REST endpoint
	ResumeConnector(ctx context.Context, in *ResumeConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// StopConnector implements the stop connector method, exposes a Kafka
	// connect equivalent REST endpoint it stops the connector but does not
	// delete the connector. All tasks for the connector are shut down completely
	StopConnector(ctx context.Context, in *StopConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteConnector implements the delete connector method, exposes a Kafka
	// connect equivalent REST endpoint
	DeleteConnector(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// UpsertConector implements the update or create connector method, it
	// exposes a kafka connect equivalent REST endpoint
	UpsertConnector(ctx context.Context, in *UpsertConnectorRequest, opts ...grpc.CallOption) (*UpsertConnectorResponse, error)
	// Deprecated: Do not use.
	// GetConnectorConfig implements the get connector configuration method, expose a kafka connect equivalent REST endpoint
	GetConnectorConfig(ctx context.Context, in *GetConnectorConfigRequest, opts ...grpc.CallOption) (*GetConnectorConfigResponse, error)
	// Deprecated: Do not use.
	// ListConnectorTopics implements the list connector topics method, expose a kafka connect equivalent REST endpoint
	ListConnectorTopics(ctx context.Context, in *ListConnectorTopicsRequest, opts ...grpc.CallOption) (*ListConnectorTopicsResponse, error)
	// Deprecated: Do not use.
	// ResetConnectorTopics implements the reset connector topics method, expose a kafka connect equivalent REST endpoint
	// the request body is empty.
	ResetConnectorTopics(ctx context.Context, in *ResetConnectorTopicsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type kafkaConnectServiceClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewKafkaConnectServiceClient(cc grpc.ClientConnInterface) KafkaConnectServiceClient {
	return &kafkaConnectServiceClient{cc}
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) ListConnectClusters(ctx context.Context, in *ListConnectClustersRequest, opts ...grpc.CallOption) (*ListConnectClustersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConnectClustersResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_ListConnectClusters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) GetConnectCluster(ctx context.Context, in *GetConnectClusterRequest, opts ...grpc.CallOption) (*GetConnectClusterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConnectClusterResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_GetConnectCluster_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) ListConnectors(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConnectorsResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_ListConnectors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) CreateConnector(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*CreateConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateConnectorResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_CreateConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) RestartConnector(ctx context.Context, in *RestartConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KafkaConnectService_RestartConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) GetConnector(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*GetConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConnectorResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_GetConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) GetConnectorStatus(ctx context.Context, in *GetConnectorStatusRequest, opts ...grpc.CallOption) (*GetConnectorStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConnectorStatusResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_GetConnectorStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) PauseConnector(ctx context.Context, in *PauseConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KafkaConnectService_PauseConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) ResumeConnector(ctx context.Context, in *ResumeConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KafkaConnectService_ResumeConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) StopConnector(ctx context.Context, in *StopConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KafkaConnectService_StopConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaConnectServiceClient) DeleteConnector(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KafkaConnectService_DeleteConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) UpsertConnector(ctx context.Context, in *UpsertConnectorRequest, opts ...grpc.CallOption) (*UpsertConnectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpsertConnectorResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_UpsertConnector_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) GetConnectorConfig(ctx context.Context, in *GetConnectorConfigRequest, opts ...grpc.CallOption) (*GetConnectorConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConnectorConfigResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_GetConnectorConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) ListConnectorTopics(ctx context.Context, in *ListConnectorTopicsRequest, opts ...grpc.CallOption) (*ListConnectorTopicsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConnectorTopicsResponse)
	err := c.cc.Invoke(ctx, KafkaConnectService_ListConnectorTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *kafkaConnectServiceClient) ResetConnectorTopics(ctx context.Context, in *ResetConnectorTopicsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KafkaConnectService_ResetConnectorTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaConnectServiceServer is the server API for KafkaConnectService service.
// All implementations must embed UnimplementedKafkaConnectServiceServer
// for forward compatibility.
//
// KafkaConnectService is the service for the Kafka connect, it exposes the
// Kafka Connect API, you can set multiple Kafka connect services and all of
// them can be managed using this service definition, the request is not only
// proxied but also enriched with better error handling and custom
// documentation and configuration
//
// Deprecated: Do not use.
type KafkaConnectServiceServer interface {
	// Deprecated: Do not use.
	// ListConnectClusters implements the list clusters method, list connect
	// clusters available in the console configuration
	ListConnectClusters(context.Context, *ListConnectClustersRequest) (*ListConnectClustersResponse, error)
	// Deprecated: Do not use.
	// GetConnectCluster implements the get cluster info method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnectCluster(context.Context, *GetConnectClusterRequest) (*GetConnectClusterResponse, error)
	// Deprecated: Do not use.
	// ListConnectors implements the list connectors method, exposes a Kafka
	// Connect equivalent REST endpoint
	ListConnectors(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error)
	// Deprecated: Do not use.
	// CreateConnector implements the create connector method, and exposes an
	// equivalent REST endpoint as the Kafka connect API endpoint
	CreateConnector(context.Context, *CreateConnectorRequest) (*CreateConnectorResponse, error)
	// Deprecated: Do not use.
	// RestartConnector implements the restart connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	RestartConnector(context.Context, *RestartConnectorRequest) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// GetConnector implements the get connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnector(context.Context, *GetConnectorRequest) (*GetConnectorResponse, error)
	// Deprecated: Do not use.
	// GetConnectorStatus implement the get status method, Gets the current status of the connector, including:
	// Whether it is running or restarting, or if it has failed or paused
	// Which worker it is assigned to
	// Error information if it has failed
	// The state of all its tasks
	GetConnectorStatus(context.Context, *GetConnectorStatusRequest) (*GetConnectorStatusResponse, error)
	// Deprecated: Do not use.
	// PauseConnector implements the pause connector method, exposes a Kafka
	// connect equivalent REST endpoint
	PauseConnector(context.Context, *PauseConnectorRequest) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// ResumeConnector implements the resume connector method, exposes a Kafka
	// connect equivalent REST endpoint
	ResumeConnector(context.Context, *ResumeConnectorRequest) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// StopConnector implements the stop connector method, exposes a Kafka
	// connect equivalent REST endpoint it stops the connector but does not
	// delete the connector. All tasks for the connector are shut down completely
	StopConnector(context.Context, *StopConnectorRequest) (*emptypb.Empty, error)
	// DeleteConnector implements the delete connector method, exposes a Kafka
	// connect equivalent REST endpoint
	DeleteConnector(context.Context, *DeleteConnectorRequest) (*emptypb.Empty, error)
	// Deprecated: Do not use.
	// UpsertConector implements the update or create connector method, it
	// exposes a kafka connect equivalent REST endpoint
	UpsertConnector(context.Context, *UpsertConnectorRequest) (*UpsertConnectorResponse, error)
	// Deprecated: Do not use.
	// GetConnectorConfig implements the get connector configuration method, expose a kafka connect equivalent REST endpoint
	GetConnectorConfig(context.Context, *GetConnectorConfigRequest) (*GetConnectorConfigResponse, error)
	// Deprecated: Do not use.
	// ListConnectorTopics implements the list connector topics method, expose a kafka connect equivalent REST endpoint
	ListConnectorTopics(context.Context, *ListConnectorTopicsRequest) (*ListConnectorTopicsResponse, error)
	// Deprecated: Do not use.
	// ResetConnectorTopics implements the reset connector topics method, expose a kafka connect equivalent REST endpoint
	// the request body is empty.
	ResetConnectorTopics(context.Context, *ResetConnectorTopicsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedKafkaConnectServiceServer()
}

// UnimplementedKafkaConnectServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKafkaConnectServiceServer struct{}

func (UnimplementedKafkaConnectServiceServer) ListConnectClusters(context.Context, *ListConnectClustersRequest) (*ListConnectClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectClusters not implemented")
}
func (UnimplementedKafkaConnectServiceServer) GetConnectCluster(context.Context, *GetConnectClusterRequest) (*GetConnectClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectCluster not implemented")
}
func (UnimplementedKafkaConnectServiceServer) ListConnectors(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectors not implemented")
}
func (UnimplementedKafkaConnectServiceServer) CreateConnector(context.Context, *CreateConnectorRequest) (*CreateConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) RestartConnector(context.Context, *RestartConnectorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) GetConnector(context.Context, *GetConnectorRequest) (*GetConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) GetConnectorStatus(context.Context, *GetConnectorStatusRequest) (*GetConnectorStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectorStatus not implemented")
}
func (UnimplementedKafkaConnectServiceServer) PauseConnector(context.Context, *PauseConnectorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) ResumeConnector(context.Context, *ResumeConnectorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) StopConnector(context.Context, *StopConnectorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) DeleteConnector(context.Context, *DeleteConnectorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) UpsertConnector(context.Context, *UpsertConnectorRequest) (*UpsertConnectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertConnector not implemented")
}
func (UnimplementedKafkaConnectServiceServer) GetConnectorConfig(context.Context, *GetConnectorConfigRequest) (*GetConnectorConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectorConfig not implemented")
}
func (UnimplementedKafkaConnectServiceServer) ListConnectorTopics(context.Context, *ListConnectorTopicsRequest) (*ListConnectorTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectorTopics not implemented")
}
func (UnimplementedKafkaConnectServiceServer) ResetConnectorTopics(context.Context, *ResetConnectorTopicsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetConnectorTopics not implemented")
}
func (UnimplementedKafkaConnectServiceServer) mustEmbedUnimplementedKafkaConnectServiceServer() {}
func (UnimplementedKafkaConnectServiceServer) testEmbeddedByValue()                             {}

// UnsafeKafkaConnectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaConnectServiceServer will
// result in compilation errors.
type UnsafeKafkaConnectServiceServer interface {
	mustEmbedUnimplementedKafkaConnectServiceServer()
}

// Deprecated: Do not use.
func RegisterKafkaConnectServiceServer(s grpc.ServiceRegistrar, srv KafkaConnectServiceServer) {
	// If the following call pancis, it indicates UnimplementedKafkaConnectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KafkaConnectService_ServiceDesc, srv)
}

func _KafkaConnectService_ListConnectClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).ListConnectClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_ListConnectClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).ListConnectClusters(ctx, req.(*ListConnectClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_GetConnectCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).GetConnectCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_GetConnectCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).GetConnectCluster(ctx, req.(*GetConnectClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_ListConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).ListConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_ListConnectors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).ListConnectors(ctx, req.(*ListConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_CreateConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).CreateConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_CreateConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).CreateConnector(ctx, req.(*CreateConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_RestartConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).RestartConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_RestartConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).RestartConnector(ctx, req.(*RestartConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_GetConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).GetConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_GetConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).GetConnector(ctx, req.(*GetConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_GetConnectorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).GetConnectorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_GetConnectorStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).GetConnectorStatus(ctx, req.(*GetConnectorStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_PauseConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).PauseConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_PauseConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).PauseConnector(ctx, req.(*PauseConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_ResumeConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).ResumeConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_ResumeConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).ResumeConnector(ctx, req.(*ResumeConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_StopConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).StopConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_StopConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).StopConnector(ctx, req.(*StopConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_DeleteConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).DeleteConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_DeleteConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).DeleteConnector(ctx, req.(*DeleteConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_UpsertConnector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).UpsertConnector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_UpsertConnector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).UpsertConnector(ctx, req.(*UpsertConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_GetConnectorConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).GetConnectorConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_GetConnectorConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).GetConnectorConfig(ctx, req.(*GetConnectorConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_ListConnectorTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).ListConnectorTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_ListConnectorTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).ListConnectorTopics(ctx, req.(*ListConnectorTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaConnectService_ResetConnectorTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetConnectorTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaConnectServiceServer).ResetConnectorTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KafkaConnectService_ResetConnectorTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaConnectServiceServer).ResetConnectorTopics(ctx, req.(*ResetConnectorTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaConnectService_ServiceDesc is the grpc.ServiceDesc for KafkaConnectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaConnectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "redpanda.api.dataplane.v1alpha1.KafkaConnectService",
	HandlerType: (*KafkaConnectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListConnectClusters",
			Handler:    _KafkaConnectService_ListConnectClusters_Handler,
		},
		{
			MethodName: "GetConnectCluster",
			Handler:    _KafkaConnectService_GetConnectCluster_Handler,
		},
		{
			MethodName: "ListConnectors",
			Handler:    _KafkaConnectService_ListConnectors_Handler,
		},
		{
			MethodName: "CreateConnector",
			Handler:    _KafkaConnectService_CreateConnector_Handler,
		},
		{
			MethodName: "RestartConnector",
			Handler:    _KafkaConnectService_RestartConnector_Handler,
		},
		{
			MethodName: "GetConnector",
			Handler:    _KafkaConnectService_GetConnector_Handler,
		},
		{
			MethodName: "GetConnectorStatus",
			Handler:    _KafkaConnectService_GetConnectorStatus_Handler,
		},
		{
			MethodName: "PauseConnector",
			Handler:    _KafkaConnectService_PauseConnector_Handler,
		},
		{
			MethodName: "ResumeConnector",
			Handler:    _KafkaConnectService_ResumeConnector_Handler,
		},
		{
			MethodName: "StopConnector",
			Handler:    _KafkaConnectService_StopConnector_Handler,
		},
		{
			MethodName: "DeleteConnector",
			Handler:    _KafkaConnectService_DeleteConnector_Handler,
		},
		{
			MethodName: "UpsertConnector",
			Handler:    _KafkaConnectService_UpsertConnector_Handler,
		},
		{
			MethodName: "GetConnectorConfig",
			Handler:    _KafkaConnectService_GetConnectorConfig_Handler,
		},
		{
			MethodName: "ListConnectorTopics",
			Handler:    _KafkaConnectService_ListConnectorTopics_Handler,
		},
		{
			MethodName: "ResetConnectorTopics",
			Handler:    _KafkaConnectService_ResetConnectorTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redpanda/api/dataplane/v1alpha1/kafka_connect.proto",
}
