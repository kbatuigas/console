// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1/acl.proto

package dataplanev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ACLServiceName is the fully-qualified name of the ACLService service.
	ACLServiceName = "redpanda.api.dataplane.v1.ACLService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ACLServiceListACLsProcedure is the fully-qualified name of the ACLService's ListACLs RPC.
	ACLServiceListACLsProcedure = "/redpanda.api.dataplane.v1.ACLService/ListACLs"
	// ACLServiceCreateACLProcedure is the fully-qualified name of the ACLService's CreateACL RPC.
	ACLServiceCreateACLProcedure = "/redpanda.api.dataplane.v1.ACLService/CreateACL"
	// ACLServiceDeleteACLsProcedure is the fully-qualified name of the ACLService's DeleteACLs RPC.
	ACLServiceDeleteACLsProcedure = "/redpanda.api.dataplane.v1.ACLService/DeleteACLs"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	aCLServiceServiceDescriptor          = v1.File_redpanda_api_dataplane_v1_acl_proto.Services().ByName("ACLService")
	aCLServiceListACLsMethodDescriptor   = aCLServiceServiceDescriptor.Methods().ByName("ListACLs")
	aCLServiceCreateACLMethodDescriptor  = aCLServiceServiceDescriptor.Methods().ByName("CreateACL")
	aCLServiceDeleteACLsMethodDescriptor = aCLServiceServiceDescriptor.Methods().ByName("DeleteACLs")
)

// ACLServiceClient is a client for the redpanda.api.dataplane.v1.ACLService service.
type ACLServiceClient interface {
	ListACLs(context.Context, *connect.Request[v1.ListACLsRequest]) (*connect.Response[v1.ListACLsResponse], error)
	CreateACL(context.Context, *connect.Request[v1.CreateACLRequest]) (*connect.Response[v1.CreateACLResponse], error)
	DeleteACLs(context.Context, *connect.Request[v1.DeleteACLsRequest]) (*connect.Response[v1.DeleteACLsResponse], error)
}

// NewACLServiceClient constructs a client for the redpanda.api.dataplane.v1.ACLService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewACLServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ACLServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &aCLServiceClient{
		listACLs: connect.NewClient[v1.ListACLsRequest, v1.ListACLsResponse](
			httpClient,
			baseURL+ACLServiceListACLsProcedure,
			connect.WithSchema(aCLServiceListACLsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createACL: connect.NewClient[v1.CreateACLRequest, v1.CreateACLResponse](
			httpClient,
			baseURL+ACLServiceCreateACLProcedure,
			connect.WithSchema(aCLServiceCreateACLMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteACLs: connect.NewClient[v1.DeleteACLsRequest, v1.DeleteACLsResponse](
			httpClient,
			baseURL+ACLServiceDeleteACLsProcedure,
			connect.WithSchema(aCLServiceDeleteACLsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// aCLServiceClient implements ACLServiceClient.
type aCLServiceClient struct {
	listACLs   *connect.Client[v1.ListACLsRequest, v1.ListACLsResponse]
	createACL  *connect.Client[v1.CreateACLRequest, v1.CreateACLResponse]
	deleteACLs *connect.Client[v1.DeleteACLsRequest, v1.DeleteACLsResponse]
}

// ListACLs calls redpanda.api.dataplane.v1.ACLService.ListACLs.
func (c *aCLServiceClient) ListACLs(ctx context.Context, req *connect.Request[v1.ListACLsRequest]) (*connect.Response[v1.ListACLsResponse], error) {
	return c.listACLs.CallUnary(ctx, req)
}

// CreateACL calls redpanda.api.dataplane.v1.ACLService.CreateACL.
func (c *aCLServiceClient) CreateACL(ctx context.Context, req *connect.Request[v1.CreateACLRequest]) (*connect.Response[v1.CreateACLResponse], error) {
	return c.createACL.CallUnary(ctx, req)
}

// DeleteACLs calls redpanda.api.dataplane.v1.ACLService.DeleteACLs.
func (c *aCLServiceClient) DeleteACLs(ctx context.Context, req *connect.Request[v1.DeleteACLsRequest]) (*connect.Response[v1.DeleteACLsResponse], error) {
	return c.deleteACLs.CallUnary(ctx, req)
}

// ACLServiceHandler is an implementation of the redpanda.api.dataplane.v1.ACLService service.
type ACLServiceHandler interface {
	ListACLs(context.Context, *connect.Request[v1.ListACLsRequest]) (*connect.Response[v1.ListACLsResponse], error)
	CreateACL(context.Context, *connect.Request[v1.CreateACLRequest]) (*connect.Response[v1.CreateACLResponse], error)
	DeleteACLs(context.Context, *connect.Request[v1.DeleteACLsRequest]) (*connect.Response[v1.DeleteACLsResponse], error)
}

// NewACLServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewACLServiceHandler(svc ACLServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	aCLServiceListACLsHandler := connect.NewUnaryHandler(
		ACLServiceListACLsProcedure,
		svc.ListACLs,
		connect.WithSchema(aCLServiceListACLsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	aCLServiceCreateACLHandler := connect.NewUnaryHandler(
		ACLServiceCreateACLProcedure,
		svc.CreateACL,
		connect.WithSchema(aCLServiceCreateACLMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	aCLServiceDeleteACLsHandler := connect.NewUnaryHandler(
		ACLServiceDeleteACLsProcedure,
		svc.DeleteACLs,
		connect.WithSchema(aCLServiceDeleteACLsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.dataplane.v1.ACLService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ACLServiceListACLsProcedure:
			aCLServiceListACLsHandler.ServeHTTP(w, r)
		case ACLServiceCreateACLProcedure:
			aCLServiceCreateACLHandler.ServeHTTP(w, r)
		case ACLServiceDeleteACLsProcedure:
			aCLServiceDeleteACLsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedACLServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedACLServiceHandler struct{}

func (UnimplementedACLServiceHandler) ListACLs(context.Context, *connect.Request[v1.ListACLsRequest]) (*connect.Response[v1.ListACLsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.ACLService.ListACLs is not implemented"))
}

func (UnimplementedACLServiceHandler) CreateACL(context.Context, *connect.Request[v1.CreateACLRequest]) (*connect.Response[v1.CreateACLResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.ACLService.CreateACL is not implemented"))
}

func (UnimplementedACLServiceHandler) DeleteACLs(context.Context, *connect.Request[v1.DeleteACLsRequest]) (*connect.Response[v1.DeleteACLsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.ACLService.DeleteACLs is not implemented"))
}
