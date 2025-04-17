// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha2/cloud_storage.proto

package dataplanev1alpha2connect

import (
	context "context"
	fmt "fmt"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1alpha2 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha2"
	connect_gateway "go.vallahaye.net/connect-gateway"
)

// CloudStorageServiceGatewayServer implements the gRPC server API for the CloudStorageService
// service.
type CloudStorageServiceGatewayServer struct {
	v1alpha2.UnimplementedCloudStorageServiceServer
	mountTopics         connect_gateway.UnaryHandler[v1alpha2.MountTopicsRequest, v1alpha2.MountTopicsResponse]
	unmountTopics       connect_gateway.UnaryHandler[v1alpha2.UnmountTopicsRequest, v1alpha2.UnmountTopicsResponse]
	listMountableTopics connect_gateway.UnaryHandler[v1alpha2.ListMountableTopicsRequest, v1alpha2.ListMountableTopicsResponse]
	listMountTasks      connect_gateway.UnaryHandler[v1alpha2.ListMountTasksRequest, v1alpha2.ListMountTasksResponse]
	getMountTask        connect_gateway.UnaryHandler[v1alpha2.GetMountTaskRequest, v1alpha2.GetMountTaskResponse]
	deleteMountTask     connect_gateway.UnaryHandler[v1alpha2.DeleteMountTaskRequest, v1alpha2.DeleteMountTaskResponse]
	updateMountTask     connect_gateway.UnaryHandler[v1alpha2.UpdateMountTaskRequest, v1alpha2.UpdateMountTaskResponse]
}

// NewCloudStorageServiceGatewayServer constructs a Connect-Gateway gRPC server for the
// CloudStorageService service.
func NewCloudStorageServiceGatewayServer(svc CloudStorageServiceHandler, opts ...connect_gateway.HandlerOption) *CloudStorageServiceGatewayServer {
	return &CloudStorageServiceGatewayServer{
		mountTopics:         connect_gateway.NewUnaryHandler(CloudStorageServiceMountTopicsProcedure, svc.MountTopics, opts...),
		unmountTopics:       connect_gateway.NewUnaryHandler(CloudStorageServiceUnmountTopicsProcedure, svc.UnmountTopics, opts...),
		listMountableTopics: connect_gateway.NewUnaryHandler(CloudStorageServiceListMountableTopicsProcedure, svc.ListMountableTopics, opts...),
		listMountTasks:      connect_gateway.NewUnaryHandler(CloudStorageServiceListMountTasksProcedure, svc.ListMountTasks, opts...),
		getMountTask:        connect_gateway.NewUnaryHandler(CloudStorageServiceGetMountTaskProcedure, svc.GetMountTask, opts...),
		deleteMountTask:     connect_gateway.NewUnaryHandler(CloudStorageServiceDeleteMountTaskProcedure, svc.DeleteMountTask, opts...),
		updateMountTask:     connect_gateway.NewUnaryHandler(CloudStorageServiceUpdateMountTaskProcedure, svc.UpdateMountTask, opts...),
	}
}

func (s *CloudStorageServiceGatewayServer) MountTopics(ctx context.Context, req *v1alpha2.MountTopicsRequest) (*v1alpha2.MountTopicsResponse, error) {
	return s.mountTopics(ctx, req)
}

func (s *CloudStorageServiceGatewayServer) UnmountTopics(ctx context.Context, req *v1alpha2.UnmountTopicsRequest) (*v1alpha2.UnmountTopicsResponse, error) {
	return s.unmountTopics(ctx, req)
}

func (s *CloudStorageServiceGatewayServer) ListMountableTopics(ctx context.Context, req *v1alpha2.ListMountableTopicsRequest) (*v1alpha2.ListMountableTopicsResponse, error) {
	return s.listMountableTopics(ctx, req)
}

func (s *CloudStorageServiceGatewayServer) ListMountTasks(ctx context.Context, req *v1alpha2.ListMountTasksRequest) (*v1alpha2.ListMountTasksResponse, error) {
	return s.listMountTasks(ctx, req)
}

func (s *CloudStorageServiceGatewayServer) GetMountTask(ctx context.Context, req *v1alpha2.GetMountTaskRequest) (*v1alpha2.GetMountTaskResponse, error) {
	return s.getMountTask(ctx, req)
}

func (s *CloudStorageServiceGatewayServer) DeleteMountTask(ctx context.Context, req *v1alpha2.DeleteMountTaskRequest) (*v1alpha2.DeleteMountTaskResponse, error) {
	return s.deleteMountTask(ctx, req)
}

func (s *CloudStorageServiceGatewayServer) UpdateMountTask(ctx context.Context, req *v1alpha2.UpdateMountTaskRequest) (*v1alpha2.UpdateMountTaskResponse, error) {
	return s.updateMountTask(ctx, req)
}

// RegisterCloudStorageServiceHandlerGatewayServer registers the Connect handlers for the
// CloudStorageService "svc" to "mux".
func RegisterCloudStorageServiceHandlerGatewayServer(mux *runtime.ServeMux, svc CloudStorageServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1alpha2.RegisterCloudStorageServiceHandlerServer(context.TODO(), mux, NewCloudStorageServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}
