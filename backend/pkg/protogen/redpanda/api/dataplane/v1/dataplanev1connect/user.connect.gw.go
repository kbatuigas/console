// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1/user.proto

package dataplanev1connect

import (
	context "context"
	fmt "fmt"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1"
	connect_gateway "go.vallahaye.net/connect-gateway"
)

// UserServiceGatewayServer implements the gRPC server API for the UserService service.
type UserServiceGatewayServer struct {
	v1.UnimplementedUserServiceServer
	createUser connect_gateway.UnaryHandler[v1.CreateUserRequest, v1.CreateUserResponse]
	updateUser connect_gateway.UnaryHandler[v1.UpdateUserRequest, v1.UpdateUserResponse]
	listUsers  connect_gateway.UnaryHandler[v1.ListUsersRequest, v1.ListUsersResponse]
	deleteUser connect_gateway.UnaryHandler[v1.DeleteUserRequest, v1.DeleteUserResponse]
}

// NewUserServiceGatewayServer constructs a Connect-Gateway gRPC server for the UserService service.
func NewUserServiceGatewayServer(svc UserServiceHandler, opts ...connect_gateway.HandlerOption) *UserServiceGatewayServer {
	return &UserServiceGatewayServer{
		createUser: connect_gateway.NewUnaryHandler(UserServiceCreateUserProcedure, svc.CreateUser, opts...),
		updateUser: connect_gateway.NewUnaryHandler(UserServiceUpdateUserProcedure, svc.UpdateUser, opts...),
		listUsers:  connect_gateway.NewUnaryHandler(UserServiceListUsersProcedure, svc.ListUsers, opts...),
		deleteUser: connect_gateway.NewUnaryHandler(UserServiceDeleteUserProcedure, svc.DeleteUser, opts...),
	}
}

func (s *UserServiceGatewayServer) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	return s.createUser(ctx, req)
}

func (s *UserServiceGatewayServer) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return s.updateUser(ctx, req)
}

func (s *UserServiceGatewayServer) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	return s.listUsers(ctx, req)
}

func (s *UserServiceGatewayServer) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return s.deleteUser(ctx, req)
}

// RegisterUserServiceHandlerGatewayServer registers the Connect handlers for the UserService "svc"
// to "mux".
func RegisterUserServiceHandlerGatewayServer(mux *runtime.ServeMux, svc UserServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1.RegisterUserServiceHandlerServer(context.TODO(), mux, NewUserServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}
