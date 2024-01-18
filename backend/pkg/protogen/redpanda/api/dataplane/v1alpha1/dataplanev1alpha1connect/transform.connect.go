// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha1/transform.proto

package dataplanev1alpha1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"

	v1alpha1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TransformServiceName is the fully-qualified name of the TransformService service.
	TransformServiceName = "redpanda.api.dataplane.v1alpha1.TransformService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TransformServiceDeployTransformProcedure is the fully-qualified name of the TransformService's
	// DeployTransform RPC.
	TransformServiceDeployTransformProcedure = "/redpanda.api.dataplane.v1alpha1.TransformService/DeployTransform"
	// TransformServiceListTransformsProcedure is the fully-qualified name of the TransformService's
	// ListTransforms RPC.
	TransformServiceListTransformsProcedure = "/redpanda.api.dataplane.v1alpha1.TransformService/ListTransforms"
	// TransformServiceGetTransformProcedure is the fully-qualified name of the TransformService's
	// GetTransform RPC.
	TransformServiceGetTransformProcedure = "/redpanda.api.dataplane.v1alpha1.TransformService/GetTransform"
	// TransformServiceDeleteTransformProcedure is the fully-qualified name of the TransformService's
	// DeleteTransform RPC.
	TransformServiceDeleteTransformProcedure = "/redpanda.api.dataplane.v1alpha1.TransformService/DeleteTransform"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	transformServiceServiceDescriptor               = v1alpha1.File_redpanda_api_dataplane_v1alpha1_transform_proto.Services().ByName("TransformService")
	transformServiceDeployTransformMethodDescriptor = transformServiceServiceDescriptor.Methods().ByName("DeployTransform")
	transformServiceListTransformsMethodDescriptor  = transformServiceServiceDescriptor.Methods().ByName("ListTransforms")
	transformServiceGetTransformMethodDescriptor    = transformServiceServiceDescriptor.Methods().ByName("GetTransform")
	transformServiceDeleteTransformMethodDescriptor = transformServiceServiceDescriptor.Methods().ByName("DeleteTransform")
)

// TransformServiceClient is a client for the redpanda.api.dataplane.v1alpha1.TransformService
// service.
type TransformServiceClient interface {
	DeployTransform(context.Context, *connect.Request[v1alpha1.DeployTransformRequest]) (*connect.Response[v1alpha1.DeployTransformResponse], error)
	ListTransforms(context.Context, *connect.Request[v1alpha1.ListTransformsRequest]) (*connect.Response[v1alpha1.ListTransformsResponse], error)
	GetTransform(context.Context, *connect.Request[v1alpha1.GetTransformRequest]) (*connect.Response[v1alpha1.GetTransformResponse], error)
	DeleteTransform(context.Context, *connect.Request[v1alpha1.DeleteTransformRequest]) (*connect.Response[v1alpha1.DeleteTransformResponse], error)
}

// NewTransformServiceClient constructs a client for the
// redpanda.api.dataplane.v1alpha1.TransformService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTransformServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TransformServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &transformServiceClient{
		deployTransform: connect.NewClient[v1alpha1.DeployTransformRequest, v1alpha1.DeployTransformResponse](
			httpClient,
			baseURL+TransformServiceDeployTransformProcedure,
			connect.WithSchema(transformServiceDeployTransformMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listTransforms: connect.NewClient[v1alpha1.ListTransformsRequest, v1alpha1.ListTransformsResponse](
			httpClient,
			baseURL+TransformServiceListTransformsProcedure,
			connect.WithSchema(transformServiceListTransformsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getTransform: connect.NewClient[v1alpha1.GetTransformRequest, v1alpha1.GetTransformResponse](
			httpClient,
			baseURL+TransformServiceGetTransformProcedure,
			connect.WithSchema(transformServiceGetTransformMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteTransform: connect.NewClient[v1alpha1.DeleteTransformRequest, v1alpha1.DeleteTransformResponse](
			httpClient,
			baseURL+TransformServiceDeleteTransformProcedure,
			connect.WithSchema(transformServiceDeleteTransformMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// transformServiceClient implements TransformServiceClient.
type transformServiceClient struct {
	deployTransform *connect.Client[v1alpha1.DeployTransformRequest, v1alpha1.DeployTransformResponse]
	listTransforms  *connect.Client[v1alpha1.ListTransformsRequest, v1alpha1.ListTransformsResponse]
	getTransform    *connect.Client[v1alpha1.GetTransformRequest, v1alpha1.GetTransformResponse]
	deleteTransform *connect.Client[v1alpha1.DeleteTransformRequest, v1alpha1.DeleteTransformResponse]
}

// DeployTransform calls redpanda.api.dataplane.v1alpha1.TransformService.DeployTransform.
func (c *transformServiceClient) DeployTransform(ctx context.Context, req *connect.Request[v1alpha1.DeployTransformRequest]) (*connect.Response[v1alpha1.DeployTransformResponse], error) {
	return c.deployTransform.CallUnary(ctx, req)
}

// ListTransforms calls redpanda.api.dataplane.v1alpha1.TransformService.ListTransforms.
func (c *transformServiceClient) ListTransforms(ctx context.Context, req *connect.Request[v1alpha1.ListTransformsRequest]) (*connect.Response[v1alpha1.ListTransformsResponse], error) {
	return c.listTransforms.CallUnary(ctx, req)
}

// GetTransform calls redpanda.api.dataplane.v1alpha1.TransformService.GetTransform.
func (c *transformServiceClient) GetTransform(ctx context.Context, req *connect.Request[v1alpha1.GetTransformRequest]) (*connect.Response[v1alpha1.GetTransformResponse], error) {
	return c.getTransform.CallUnary(ctx, req)
}

// DeleteTransform calls redpanda.api.dataplane.v1alpha1.TransformService.DeleteTransform.
func (c *transformServiceClient) DeleteTransform(ctx context.Context, req *connect.Request[v1alpha1.DeleteTransformRequest]) (*connect.Response[v1alpha1.DeleteTransformResponse], error) {
	return c.deleteTransform.CallUnary(ctx, req)
}

// TransformServiceHandler is an implementation of the
// redpanda.api.dataplane.v1alpha1.TransformService service.
type TransformServiceHandler interface {
	DeployTransform(context.Context, *connect.Request[v1alpha1.DeployTransformRequest]) (*connect.Response[v1alpha1.DeployTransformResponse], error)
	ListTransforms(context.Context, *connect.Request[v1alpha1.ListTransformsRequest]) (*connect.Response[v1alpha1.ListTransformsResponse], error)
	GetTransform(context.Context, *connect.Request[v1alpha1.GetTransformRequest]) (*connect.Response[v1alpha1.GetTransformResponse], error)
	DeleteTransform(context.Context, *connect.Request[v1alpha1.DeleteTransformRequest]) (*connect.Response[v1alpha1.DeleteTransformResponse], error)
}

// NewTransformServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTransformServiceHandler(svc TransformServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	transformServiceDeployTransformHandler := connect.NewUnaryHandler(
		TransformServiceDeployTransformProcedure,
		svc.DeployTransform,
		connect.WithSchema(transformServiceDeployTransformMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	transformServiceListTransformsHandler := connect.NewUnaryHandler(
		TransformServiceListTransformsProcedure,
		svc.ListTransforms,
		connect.WithSchema(transformServiceListTransformsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	transformServiceGetTransformHandler := connect.NewUnaryHandler(
		TransformServiceGetTransformProcedure,
		svc.GetTransform,
		connect.WithSchema(transformServiceGetTransformMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	transformServiceDeleteTransformHandler := connect.NewUnaryHandler(
		TransformServiceDeleteTransformProcedure,
		svc.DeleteTransform,
		connect.WithSchema(transformServiceDeleteTransformMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.dataplane.v1alpha1.TransformService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TransformServiceDeployTransformProcedure:
			transformServiceDeployTransformHandler.ServeHTTP(w, r)
		case TransformServiceListTransformsProcedure:
			transformServiceListTransformsHandler.ServeHTTP(w, r)
		case TransformServiceGetTransformProcedure:
			transformServiceGetTransformHandler.ServeHTTP(w, r)
		case TransformServiceDeleteTransformProcedure:
			transformServiceDeleteTransformHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTransformServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTransformServiceHandler struct{}

func (UnimplementedTransformServiceHandler) DeployTransform(context.Context, *connect.Request[v1alpha1.DeployTransformRequest]) (*connect.Response[v1alpha1.DeployTransformResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TransformService.DeployTransform is not implemented"))
}

func (UnimplementedTransformServiceHandler) ListTransforms(context.Context, *connect.Request[v1alpha1.ListTransformsRequest]) (*connect.Response[v1alpha1.ListTransformsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TransformService.ListTransforms is not implemented"))
}

func (UnimplementedTransformServiceHandler) GetTransform(context.Context, *connect.Request[v1alpha1.GetTransformRequest]) (*connect.Response[v1alpha1.GetTransformResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TransformService.GetTransform is not implemented"))
}

func (UnimplementedTransformServiceHandler) DeleteTransform(context.Context, *connect.Request[v1alpha1.DeleteTransformRequest]) (*connect.Response[v1alpha1.DeleteTransformResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha1.TransformService.DeleteTransform is not implemented"))
}
