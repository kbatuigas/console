// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha2/topic.proto

package dataplanev1alpha2connect

import (
	context "context"
	fmt "fmt"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1alpha2 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha2"
	connect_gateway "go.vallahaye.net/connect-gateway"
)

// TopicServiceGatewayServer implements the gRPC server API for the TopicService service.
type TopicServiceGatewayServer struct {
	v1alpha2.UnimplementedTopicServiceServer
	createTopic               connect_gateway.UnaryHandler[v1alpha2.CreateTopicRequest, v1alpha2.CreateTopicResponse]
	listTopics                connect_gateway.UnaryHandler[v1alpha2.ListTopicsRequest, v1alpha2.ListTopicsResponse]
	deleteTopic               connect_gateway.UnaryHandler[v1alpha2.DeleteTopicRequest, v1alpha2.DeleteTopicResponse]
	getTopicConfigurations    connect_gateway.UnaryHandler[v1alpha2.GetTopicConfigurationsRequest, v1alpha2.GetTopicConfigurationsResponse]
	updateTopicConfigurations connect_gateway.UnaryHandler[v1alpha2.UpdateTopicConfigurationsRequest, v1alpha2.UpdateTopicConfigurationsResponse]
	setTopicConfigurations    connect_gateway.UnaryHandler[v1alpha2.SetTopicConfigurationsRequest, v1alpha2.SetTopicConfigurationsResponse]
}

// NewTopicServiceGatewayServer constructs a Connect-Gateway gRPC server for the TopicService
// service.
func NewTopicServiceGatewayServer(svc TopicServiceHandler, opts ...connect_gateway.HandlerOption) *TopicServiceGatewayServer {
	return &TopicServiceGatewayServer{
		createTopic:               connect_gateway.NewUnaryHandler(TopicServiceCreateTopicProcedure, svc.CreateTopic, opts...),
		listTopics:                connect_gateway.NewUnaryHandler(TopicServiceListTopicsProcedure, svc.ListTopics, opts...),
		deleteTopic:               connect_gateway.NewUnaryHandler(TopicServiceDeleteTopicProcedure, svc.DeleteTopic, opts...),
		getTopicConfigurations:    connect_gateway.NewUnaryHandler(TopicServiceGetTopicConfigurationsProcedure, svc.GetTopicConfigurations, opts...),
		updateTopicConfigurations: connect_gateway.NewUnaryHandler(TopicServiceUpdateTopicConfigurationsProcedure, svc.UpdateTopicConfigurations, opts...),
		setTopicConfigurations:    connect_gateway.NewUnaryHandler(TopicServiceSetTopicConfigurationsProcedure, svc.SetTopicConfigurations, opts...),
	}
}

func (s *TopicServiceGatewayServer) CreateTopic(ctx context.Context, req *v1alpha2.CreateTopicRequest) (*v1alpha2.CreateTopicResponse, error) {
	return s.createTopic(ctx, req)
}

func (s *TopicServiceGatewayServer) ListTopics(ctx context.Context, req *v1alpha2.ListTopicsRequest) (*v1alpha2.ListTopicsResponse, error) {
	return s.listTopics(ctx, req)
}

func (s *TopicServiceGatewayServer) DeleteTopic(ctx context.Context, req *v1alpha2.DeleteTopicRequest) (*v1alpha2.DeleteTopicResponse, error) {
	return s.deleteTopic(ctx, req)
}

func (s *TopicServiceGatewayServer) GetTopicConfigurations(ctx context.Context, req *v1alpha2.GetTopicConfigurationsRequest) (*v1alpha2.GetTopicConfigurationsResponse, error) {
	return s.getTopicConfigurations(ctx, req)
}

func (s *TopicServiceGatewayServer) UpdateTopicConfigurations(ctx context.Context, req *v1alpha2.UpdateTopicConfigurationsRequest) (*v1alpha2.UpdateTopicConfigurationsResponse, error) {
	return s.updateTopicConfigurations(ctx, req)
}

func (s *TopicServiceGatewayServer) SetTopicConfigurations(ctx context.Context, req *v1alpha2.SetTopicConfigurationsRequest) (*v1alpha2.SetTopicConfigurationsResponse, error) {
	return s.setTopicConfigurations(ctx, req)
}

// RegisterTopicServiceHandlerGatewayServer registers the Connect handlers for the TopicService
// "svc" to "mux".
func RegisterTopicServiceHandlerGatewayServer(mux *runtime.ServeMux, svc TopicServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1alpha2.RegisterTopicServiceHandlerServer(context.TODO(), mux, NewTopicServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}
