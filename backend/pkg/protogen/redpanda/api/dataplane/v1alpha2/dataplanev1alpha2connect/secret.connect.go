// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1alpha2/secret.proto

package dataplanev1alpha2connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha2 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1alpha2"
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
	// SecretServiceName is the fully-qualified name of the SecretService service.
	SecretServiceName = "redpanda.api.dataplane.v1alpha2.SecretService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SecretServiceGetSecretProcedure is the fully-qualified name of the SecretService's GetSecret RPC.
	SecretServiceGetSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/GetSecret"
	// SecretServiceListSecretsProcedure is the fully-qualified name of the SecretService's ListSecrets
	// RPC.
	SecretServiceListSecretsProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/ListSecrets"
	// SecretServiceCreateSecretProcedure is the fully-qualified name of the SecretService's
	// CreateSecret RPC.
	SecretServiceCreateSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/CreateSecret"
	// SecretServiceUpdateSecretProcedure is the fully-qualified name of the SecretService's
	// UpdateSecret RPC.
	SecretServiceUpdateSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/UpdateSecret"
	// SecretServiceDeleteSecretProcedure is the fully-qualified name of the SecretService's
	// DeleteSecret RPC.
	SecretServiceDeleteSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/DeleteSecret"
	// SecretServiceListSecretScopesProcedure is the fully-qualified name of the SecretService's
	// ListSecretScopes RPC.
	SecretServiceListSecretScopesProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/ListSecretScopes"
	// SecretServiceGetConnectSecretProcedure is the fully-qualified name of the SecretService's
	// GetConnectSecret RPC.
	SecretServiceGetConnectSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/GetConnectSecret"
	// SecretServiceListConnectSecretsProcedure is the fully-qualified name of the SecretService's
	// ListConnectSecrets RPC.
	SecretServiceListConnectSecretsProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/ListConnectSecrets"
	// SecretServiceCreateConnectSecretProcedure is the fully-qualified name of the SecretService's
	// CreateConnectSecret RPC.
	SecretServiceCreateConnectSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/CreateConnectSecret"
	// SecretServiceUpdateConnectSecretProcedure is the fully-qualified name of the SecretService's
	// UpdateConnectSecret RPC.
	SecretServiceUpdateConnectSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/UpdateConnectSecret"
	// SecretServiceDeleteConnectSecretProcedure is the fully-qualified name of the SecretService's
	// DeleteConnectSecret RPC.
	SecretServiceDeleteConnectSecretProcedure = "/redpanda.api.dataplane.v1alpha2.SecretService/DeleteConnectSecret"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	secretServiceServiceDescriptor                   = v1alpha2.File_redpanda_api_dataplane_v1alpha2_secret_proto.Services().ByName("SecretService")
	secretServiceGetSecretMethodDescriptor           = secretServiceServiceDescriptor.Methods().ByName("GetSecret")
	secretServiceListSecretsMethodDescriptor         = secretServiceServiceDescriptor.Methods().ByName("ListSecrets")
	secretServiceCreateSecretMethodDescriptor        = secretServiceServiceDescriptor.Methods().ByName("CreateSecret")
	secretServiceUpdateSecretMethodDescriptor        = secretServiceServiceDescriptor.Methods().ByName("UpdateSecret")
	secretServiceDeleteSecretMethodDescriptor        = secretServiceServiceDescriptor.Methods().ByName("DeleteSecret")
	secretServiceListSecretScopesMethodDescriptor    = secretServiceServiceDescriptor.Methods().ByName("ListSecretScopes")
	secretServiceGetConnectSecretMethodDescriptor    = secretServiceServiceDescriptor.Methods().ByName("GetConnectSecret")
	secretServiceListConnectSecretsMethodDescriptor  = secretServiceServiceDescriptor.Methods().ByName("ListConnectSecrets")
	secretServiceCreateConnectSecretMethodDescriptor = secretServiceServiceDescriptor.Methods().ByName("CreateConnectSecret")
	secretServiceUpdateConnectSecretMethodDescriptor = secretServiceServiceDescriptor.Methods().ByName("UpdateConnectSecret")
	secretServiceDeleteConnectSecretMethodDescriptor = secretServiceServiceDescriptor.Methods().ByName("DeleteConnectSecret")
)

// SecretServiceClient is a client for the redpanda.api.dataplane.v1alpha2.SecretService service.
type SecretServiceClient interface {
	// GetSecret retrieves the specific secret.
	GetSecret(context.Context, *connect.Request[v1alpha2.GetSecretRequest]) (*connect.Response[v1alpha2.GetSecretResponse], error)
	// ListSecrets lists the secrets based on optional filter.
	ListSecrets(context.Context, *connect.Request[v1alpha2.ListSecretsRequest]) (*connect.Response[v1alpha2.ListSecretsResponse], error)
	// CreateSecret creates the secret.
	CreateSecret(context.Context, *connect.Request[v1alpha2.CreateSecretRequest]) (*connect.Response[v1alpha2.CreateSecretResponse], error)
	// UpdateSecret updates the secret.
	UpdateSecret(context.Context, *connect.Request[v1alpha2.UpdateSecretRequest]) (*connect.Response[v1alpha2.UpdateSecretResponse], error)
	// DeleteSecret deletes the secret.
	DeleteSecret(context.Context, *connect.Request[v1alpha2.DeleteSecretRequest]) (*connect.Response[v1alpha2.DeleteSecretResponse], error)
	// ListSecretScopes lists the supported secret scopes.
	ListSecretScopes(context.Context, *connect.Request[v1alpha2.ListSecretScopesRequest]) (*connect.Response[v1alpha2.ListSecretScopesResponse], error)
	// GetConnectSecret retrieves the specific secret for a specific Connect.
	GetConnectSecret(context.Context, *connect.Request[v1alpha2.GetConnectSecretRequest]) (*connect.Response[v1alpha2.GetConnectSecretResponse], error)
	// ListConnectSecrets lists the Connect secrets based on optional filter.
	ListConnectSecrets(context.Context, *connect.Request[v1alpha2.ListConnectSecretsRequest]) (*connect.Response[v1alpha2.ListConnectSecretsResponse], error)
	// CreateConnectSecret creates the secret for a Connect.
	CreateConnectSecret(context.Context, *connect.Request[v1alpha2.CreateConnectSecretRequest]) (*connect.Response[v1alpha2.CreateConnectSecretResponse], error)
	// UpdateConnectSecret updates the Connect secret.
	UpdateConnectSecret(context.Context, *connect.Request[v1alpha2.UpdateConnectSecretRequest]) (*connect.Response[v1alpha2.UpdateConnectSecretResponse], error)
	// DeleteSecret deletes the secret.
	DeleteConnectSecret(context.Context, *connect.Request[v1alpha2.DeleteConnectSecretRequest]) (*connect.Response[v1alpha2.DeleteConnectSecretResponse], error)
}

// NewSecretServiceClient constructs a client for the redpanda.api.dataplane.v1alpha2.SecretService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSecretServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SecretServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &secretServiceClient{
		getSecret: connect.NewClient[v1alpha2.GetSecretRequest, v1alpha2.GetSecretResponse](
			httpClient,
			baseURL+SecretServiceGetSecretProcedure,
			connect.WithSchema(secretServiceGetSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSecrets: connect.NewClient[v1alpha2.ListSecretsRequest, v1alpha2.ListSecretsResponse](
			httpClient,
			baseURL+SecretServiceListSecretsProcedure,
			connect.WithSchema(secretServiceListSecretsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createSecret: connect.NewClient[v1alpha2.CreateSecretRequest, v1alpha2.CreateSecretResponse](
			httpClient,
			baseURL+SecretServiceCreateSecretProcedure,
			connect.WithSchema(secretServiceCreateSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateSecret: connect.NewClient[v1alpha2.UpdateSecretRequest, v1alpha2.UpdateSecretResponse](
			httpClient,
			baseURL+SecretServiceUpdateSecretProcedure,
			connect.WithSchema(secretServiceUpdateSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteSecret: connect.NewClient[v1alpha2.DeleteSecretRequest, v1alpha2.DeleteSecretResponse](
			httpClient,
			baseURL+SecretServiceDeleteSecretProcedure,
			connect.WithSchema(secretServiceDeleteSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSecretScopes: connect.NewClient[v1alpha2.ListSecretScopesRequest, v1alpha2.ListSecretScopesResponse](
			httpClient,
			baseURL+SecretServiceListSecretScopesProcedure,
			connect.WithSchema(secretServiceListSecretScopesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getConnectSecret: connect.NewClient[v1alpha2.GetConnectSecretRequest, v1alpha2.GetConnectSecretResponse](
			httpClient,
			baseURL+SecretServiceGetConnectSecretProcedure,
			connect.WithSchema(secretServiceGetConnectSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listConnectSecrets: connect.NewClient[v1alpha2.ListConnectSecretsRequest, v1alpha2.ListConnectSecretsResponse](
			httpClient,
			baseURL+SecretServiceListConnectSecretsProcedure,
			connect.WithSchema(secretServiceListConnectSecretsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createConnectSecret: connect.NewClient[v1alpha2.CreateConnectSecretRequest, v1alpha2.CreateConnectSecretResponse](
			httpClient,
			baseURL+SecretServiceCreateConnectSecretProcedure,
			connect.WithSchema(secretServiceCreateConnectSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateConnectSecret: connect.NewClient[v1alpha2.UpdateConnectSecretRequest, v1alpha2.UpdateConnectSecretResponse](
			httpClient,
			baseURL+SecretServiceUpdateConnectSecretProcedure,
			connect.WithSchema(secretServiceUpdateConnectSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteConnectSecret: connect.NewClient[v1alpha2.DeleteConnectSecretRequest, v1alpha2.DeleteConnectSecretResponse](
			httpClient,
			baseURL+SecretServiceDeleteConnectSecretProcedure,
			connect.WithSchema(secretServiceDeleteConnectSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// secretServiceClient implements SecretServiceClient.
type secretServiceClient struct {
	getSecret           *connect.Client[v1alpha2.GetSecretRequest, v1alpha2.GetSecretResponse]
	listSecrets         *connect.Client[v1alpha2.ListSecretsRequest, v1alpha2.ListSecretsResponse]
	createSecret        *connect.Client[v1alpha2.CreateSecretRequest, v1alpha2.CreateSecretResponse]
	updateSecret        *connect.Client[v1alpha2.UpdateSecretRequest, v1alpha2.UpdateSecretResponse]
	deleteSecret        *connect.Client[v1alpha2.DeleteSecretRequest, v1alpha2.DeleteSecretResponse]
	listSecretScopes    *connect.Client[v1alpha2.ListSecretScopesRequest, v1alpha2.ListSecretScopesResponse]
	getConnectSecret    *connect.Client[v1alpha2.GetConnectSecretRequest, v1alpha2.GetConnectSecretResponse]
	listConnectSecrets  *connect.Client[v1alpha2.ListConnectSecretsRequest, v1alpha2.ListConnectSecretsResponse]
	createConnectSecret *connect.Client[v1alpha2.CreateConnectSecretRequest, v1alpha2.CreateConnectSecretResponse]
	updateConnectSecret *connect.Client[v1alpha2.UpdateConnectSecretRequest, v1alpha2.UpdateConnectSecretResponse]
	deleteConnectSecret *connect.Client[v1alpha2.DeleteConnectSecretRequest, v1alpha2.DeleteConnectSecretResponse]
}

// GetSecret calls redpanda.api.dataplane.v1alpha2.SecretService.GetSecret.
func (c *secretServiceClient) GetSecret(ctx context.Context, req *connect.Request[v1alpha2.GetSecretRequest]) (*connect.Response[v1alpha2.GetSecretResponse], error) {
	return c.getSecret.CallUnary(ctx, req)
}

// ListSecrets calls redpanda.api.dataplane.v1alpha2.SecretService.ListSecrets.
func (c *secretServiceClient) ListSecrets(ctx context.Context, req *connect.Request[v1alpha2.ListSecretsRequest]) (*connect.Response[v1alpha2.ListSecretsResponse], error) {
	return c.listSecrets.CallUnary(ctx, req)
}

// CreateSecret calls redpanda.api.dataplane.v1alpha2.SecretService.CreateSecret.
func (c *secretServiceClient) CreateSecret(ctx context.Context, req *connect.Request[v1alpha2.CreateSecretRequest]) (*connect.Response[v1alpha2.CreateSecretResponse], error) {
	return c.createSecret.CallUnary(ctx, req)
}

// UpdateSecret calls redpanda.api.dataplane.v1alpha2.SecretService.UpdateSecret.
func (c *secretServiceClient) UpdateSecret(ctx context.Context, req *connect.Request[v1alpha2.UpdateSecretRequest]) (*connect.Response[v1alpha2.UpdateSecretResponse], error) {
	return c.updateSecret.CallUnary(ctx, req)
}

// DeleteSecret calls redpanda.api.dataplane.v1alpha2.SecretService.DeleteSecret.
func (c *secretServiceClient) DeleteSecret(ctx context.Context, req *connect.Request[v1alpha2.DeleteSecretRequest]) (*connect.Response[v1alpha2.DeleteSecretResponse], error) {
	return c.deleteSecret.CallUnary(ctx, req)
}

// ListSecretScopes calls redpanda.api.dataplane.v1alpha2.SecretService.ListSecretScopes.
func (c *secretServiceClient) ListSecretScopes(ctx context.Context, req *connect.Request[v1alpha2.ListSecretScopesRequest]) (*connect.Response[v1alpha2.ListSecretScopesResponse], error) {
	return c.listSecretScopes.CallUnary(ctx, req)
}

// GetConnectSecret calls redpanda.api.dataplane.v1alpha2.SecretService.GetConnectSecret.
func (c *secretServiceClient) GetConnectSecret(ctx context.Context, req *connect.Request[v1alpha2.GetConnectSecretRequest]) (*connect.Response[v1alpha2.GetConnectSecretResponse], error) {
	return c.getConnectSecret.CallUnary(ctx, req)
}

// ListConnectSecrets calls redpanda.api.dataplane.v1alpha2.SecretService.ListConnectSecrets.
func (c *secretServiceClient) ListConnectSecrets(ctx context.Context, req *connect.Request[v1alpha2.ListConnectSecretsRequest]) (*connect.Response[v1alpha2.ListConnectSecretsResponse], error) {
	return c.listConnectSecrets.CallUnary(ctx, req)
}

// CreateConnectSecret calls redpanda.api.dataplane.v1alpha2.SecretService.CreateConnectSecret.
func (c *secretServiceClient) CreateConnectSecret(ctx context.Context, req *connect.Request[v1alpha2.CreateConnectSecretRequest]) (*connect.Response[v1alpha2.CreateConnectSecretResponse], error) {
	return c.createConnectSecret.CallUnary(ctx, req)
}

// UpdateConnectSecret calls redpanda.api.dataplane.v1alpha2.SecretService.UpdateConnectSecret.
func (c *secretServiceClient) UpdateConnectSecret(ctx context.Context, req *connect.Request[v1alpha2.UpdateConnectSecretRequest]) (*connect.Response[v1alpha2.UpdateConnectSecretResponse], error) {
	return c.updateConnectSecret.CallUnary(ctx, req)
}

// DeleteConnectSecret calls redpanda.api.dataplane.v1alpha2.SecretService.DeleteConnectSecret.
func (c *secretServiceClient) DeleteConnectSecret(ctx context.Context, req *connect.Request[v1alpha2.DeleteConnectSecretRequest]) (*connect.Response[v1alpha2.DeleteConnectSecretResponse], error) {
	return c.deleteConnectSecret.CallUnary(ctx, req)
}

// SecretServiceHandler is an implementation of the redpanda.api.dataplane.v1alpha2.SecretService
// service.
type SecretServiceHandler interface {
	// GetSecret retrieves the specific secret.
	GetSecret(context.Context, *connect.Request[v1alpha2.GetSecretRequest]) (*connect.Response[v1alpha2.GetSecretResponse], error)
	// ListSecrets lists the secrets based on optional filter.
	ListSecrets(context.Context, *connect.Request[v1alpha2.ListSecretsRequest]) (*connect.Response[v1alpha2.ListSecretsResponse], error)
	// CreateSecret creates the secret.
	CreateSecret(context.Context, *connect.Request[v1alpha2.CreateSecretRequest]) (*connect.Response[v1alpha2.CreateSecretResponse], error)
	// UpdateSecret updates the secret.
	UpdateSecret(context.Context, *connect.Request[v1alpha2.UpdateSecretRequest]) (*connect.Response[v1alpha2.UpdateSecretResponse], error)
	// DeleteSecret deletes the secret.
	DeleteSecret(context.Context, *connect.Request[v1alpha2.DeleteSecretRequest]) (*connect.Response[v1alpha2.DeleteSecretResponse], error)
	// ListSecretScopes lists the supported secret scopes.
	ListSecretScopes(context.Context, *connect.Request[v1alpha2.ListSecretScopesRequest]) (*connect.Response[v1alpha2.ListSecretScopesResponse], error)
	// GetConnectSecret retrieves the specific secret for a specific Connect.
	GetConnectSecret(context.Context, *connect.Request[v1alpha2.GetConnectSecretRequest]) (*connect.Response[v1alpha2.GetConnectSecretResponse], error)
	// ListConnectSecrets lists the Connect secrets based on optional filter.
	ListConnectSecrets(context.Context, *connect.Request[v1alpha2.ListConnectSecretsRequest]) (*connect.Response[v1alpha2.ListConnectSecretsResponse], error)
	// CreateConnectSecret creates the secret for a Connect.
	CreateConnectSecret(context.Context, *connect.Request[v1alpha2.CreateConnectSecretRequest]) (*connect.Response[v1alpha2.CreateConnectSecretResponse], error)
	// UpdateConnectSecret updates the Connect secret.
	UpdateConnectSecret(context.Context, *connect.Request[v1alpha2.UpdateConnectSecretRequest]) (*connect.Response[v1alpha2.UpdateConnectSecretResponse], error)
	// DeleteSecret deletes the secret.
	DeleteConnectSecret(context.Context, *connect.Request[v1alpha2.DeleteConnectSecretRequest]) (*connect.Response[v1alpha2.DeleteConnectSecretResponse], error)
}

// NewSecretServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSecretServiceHandler(svc SecretServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	secretServiceGetSecretHandler := connect.NewUnaryHandler(
		SecretServiceGetSecretProcedure,
		svc.GetSecret,
		connect.WithSchema(secretServiceGetSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceListSecretsHandler := connect.NewUnaryHandler(
		SecretServiceListSecretsProcedure,
		svc.ListSecrets,
		connect.WithSchema(secretServiceListSecretsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceCreateSecretHandler := connect.NewUnaryHandler(
		SecretServiceCreateSecretProcedure,
		svc.CreateSecret,
		connect.WithSchema(secretServiceCreateSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceUpdateSecretHandler := connect.NewUnaryHandler(
		SecretServiceUpdateSecretProcedure,
		svc.UpdateSecret,
		connect.WithSchema(secretServiceUpdateSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceDeleteSecretHandler := connect.NewUnaryHandler(
		SecretServiceDeleteSecretProcedure,
		svc.DeleteSecret,
		connect.WithSchema(secretServiceDeleteSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceListSecretScopesHandler := connect.NewUnaryHandler(
		SecretServiceListSecretScopesProcedure,
		svc.ListSecretScopes,
		connect.WithSchema(secretServiceListSecretScopesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceGetConnectSecretHandler := connect.NewUnaryHandler(
		SecretServiceGetConnectSecretProcedure,
		svc.GetConnectSecret,
		connect.WithSchema(secretServiceGetConnectSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceListConnectSecretsHandler := connect.NewUnaryHandler(
		SecretServiceListConnectSecretsProcedure,
		svc.ListConnectSecrets,
		connect.WithSchema(secretServiceListConnectSecretsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceCreateConnectSecretHandler := connect.NewUnaryHandler(
		SecretServiceCreateConnectSecretProcedure,
		svc.CreateConnectSecret,
		connect.WithSchema(secretServiceCreateConnectSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceUpdateConnectSecretHandler := connect.NewUnaryHandler(
		SecretServiceUpdateConnectSecretProcedure,
		svc.UpdateConnectSecret,
		connect.WithSchema(secretServiceUpdateConnectSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceDeleteConnectSecretHandler := connect.NewUnaryHandler(
		SecretServiceDeleteConnectSecretProcedure,
		svc.DeleteConnectSecret,
		connect.WithSchema(secretServiceDeleteConnectSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.dataplane.v1alpha2.SecretService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SecretServiceGetSecretProcedure:
			secretServiceGetSecretHandler.ServeHTTP(w, r)
		case SecretServiceListSecretsProcedure:
			secretServiceListSecretsHandler.ServeHTTP(w, r)
		case SecretServiceCreateSecretProcedure:
			secretServiceCreateSecretHandler.ServeHTTP(w, r)
		case SecretServiceUpdateSecretProcedure:
			secretServiceUpdateSecretHandler.ServeHTTP(w, r)
		case SecretServiceDeleteSecretProcedure:
			secretServiceDeleteSecretHandler.ServeHTTP(w, r)
		case SecretServiceListSecretScopesProcedure:
			secretServiceListSecretScopesHandler.ServeHTTP(w, r)
		case SecretServiceGetConnectSecretProcedure:
			secretServiceGetConnectSecretHandler.ServeHTTP(w, r)
		case SecretServiceListConnectSecretsProcedure:
			secretServiceListConnectSecretsHandler.ServeHTTP(w, r)
		case SecretServiceCreateConnectSecretProcedure:
			secretServiceCreateConnectSecretHandler.ServeHTTP(w, r)
		case SecretServiceUpdateConnectSecretProcedure:
			secretServiceUpdateConnectSecretHandler.ServeHTTP(w, r)
		case SecretServiceDeleteConnectSecretProcedure:
			secretServiceDeleteConnectSecretHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSecretServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSecretServiceHandler struct{}

func (UnimplementedSecretServiceHandler) GetSecret(context.Context, *connect.Request[v1alpha2.GetSecretRequest]) (*connect.Response[v1alpha2.GetSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.GetSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) ListSecrets(context.Context, *connect.Request[v1alpha2.ListSecretsRequest]) (*connect.Response[v1alpha2.ListSecretsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.ListSecrets is not implemented"))
}

func (UnimplementedSecretServiceHandler) CreateSecret(context.Context, *connect.Request[v1alpha2.CreateSecretRequest]) (*connect.Response[v1alpha2.CreateSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.CreateSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) UpdateSecret(context.Context, *connect.Request[v1alpha2.UpdateSecretRequest]) (*connect.Response[v1alpha2.UpdateSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.UpdateSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) DeleteSecret(context.Context, *connect.Request[v1alpha2.DeleteSecretRequest]) (*connect.Response[v1alpha2.DeleteSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.DeleteSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) ListSecretScopes(context.Context, *connect.Request[v1alpha2.ListSecretScopesRequest]) (*connect.Response[v1alpha2.ListSecretScopesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.ListSecretScopes is not implemented"))
}

func (UnimplementedSecretServiceHandler) GetConnectSecret(context.Context, *connect.Request[v1alpha2.GetConnectSecretRequest]) (*connect.Response[v1alpha2.GetConnectSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.GetConnectSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) ListConnectSecrets(context.Context, *connect.Request[v1alpha2.ListConnectSecretsRequest]) (*connect.Response[v1alpha2.ListConnectSecretsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.ListConnectSecrets is not implemented"))
}

func (UnimplementedSecretServiceHandler) CreateConnectSecret(context.Context, *connect.Request[v1alpha2.CreateConnectSecretRequest]) (*connect.Response[v1alpha2.CreateConnectSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.CreateConnectSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) UpdateConnectSecret(context.Context, *connect.Request[v1alpha2.UpdateConnectSecretRequest]) (*connect.Response[v1alpha2.UpdateConnectSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.UpdateConnectSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) DeleteConnectSecret(context.Context, *connect.Request[v1alpha2.DeleteConnectSecretRequest]) (*connect.Response[v1alpha2.DeleteConnectSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1alpha2.SecretService.DeleteConnectSecret is not implemented"))
}
