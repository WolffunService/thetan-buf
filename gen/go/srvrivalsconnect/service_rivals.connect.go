// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: service_rivals.proto

package srvrivalsconnect

import (
	srvrivals "./srvrivals"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ThetanRivalServiceName is the fully-qualified name of the ThetanRivalService service.
	ThetanRivalServiceName = "services.ThetanRivalService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ThetanRivalServiceGetUserProfileProcedure is the fully-qualified name of the ThetanRivalService's
	// GetUserProfile RPC.
	ThetanRivalServiceGetUserProfileProcedure = "/services.ThetanRivalService/GetUserProfile"
	// ThetanRivalServiceGetUserMinionsProcedure is the fully-qualified name of the ThetanRivalService's
	// GetUserMinions RPC.
	ThetanRivalServiceGetUserMinionsProcedure = "/services.ThetanRivalService/GetUserMinions"
	// ThetanRivalServiceGetUserSelectedMinionProcedure is the fully-qualified name of the
	// ThetanRivalService's GetUserSelectedMinion RPC.
	ThetanRivalServiceGetUserSelectedMinionProcedure = "/services.ThetanRivalService/GetUserSelectedMinion"
	// ThetanRivalServiceGetMinionProcedure is the fully-qualified name of the ThetanRivalService's
	// GetMinion RPC.
	ThetanRivalServiceGetMinionProcedure = "/services.ThetanRivalService/GetMinion"
	// ThetanRivalServiceCreateMinionProcedure is the fully-qualified name of the ThetanRivalService's
	// CreateMinion RPC.
	ThetanRivalServiceCreateMinionProcedure = "/services.ThetanRivalService/CreateMinion"
	// ThetanRivalServiceGetListFriendsProcedure is the fully-qualified name of the ThetanRivalService's
	// GetListFriends RPC.
	ThetanRivalServiceGetListFriendsProcedure = "/services.ThetanRivalService/GetListFriends"
	// ThetanRivalServiceGetMatchInfoOnboardingProcedure is the fully-qualified name of the
	// ThetanRivalService's GetMatchInfoOnboarding RPC.
	ThetanRivalServiceGetMatchInfoOnboardingProcedure = "/services.ThetanRivalService/GetMatchInfoOnboarding"
	// ThetanRivalServiceGetLatestLobbyActivityInfoProcedure is the fully-qualified name of the
	// ThetanRivalService's GetLatestLobbyActivityInfo RPC.
	ThetanRivalServiceGetLatestLobbyActivityInfoProcedure = "/services.ThetanRivalService/GetLatestLobbyActivityInfo"
)

// ThetanRivalServiceClient is a client for the services.ThetanRivalService service.
type ThetanRivalServiceClient interface {
	// Profile
	GetUserProfile(context.Context, *connect_go.Request[srvrivals.UserProfileRequest]) (*connect_go.Response[srvrivals.UserProfileResponse], error)
	// Minions
	GetUserMinions(context.Context, *connect_go.Request[srvrivals.UserMinionsRequest]) (*connect_go.Response[srvrivals.UserMinionsResponse], error)
	GetUserSelectedMinion(context.Context, *connect_go.Request[srvrivals.UserSelectedMinionRequest]) (*connect_go.Response[srvrivals.UserSelectedMinionResponse], error)
	GetMinion(context.Context, *connect_go.Request[srvrivals.MinionRequest]) (*connect_go.Response[srvrivals.MinionResponse], error)
	CreateMinion(context.Context, *connect_go.Request[srvrivals.CreateMinionRequest]) (*connect_go.Response[srvrivals.CreateMinionResponse], error)
	// Friends
	GetListFriends(context.Context, *connect_go.Request[srvrivals.GetUserFriendRequest]) (*connect_go.Response[srvrivals.GetUserFriendResponse], error)
	GetMatchInfoOnboarding(context.Context, *connect_go.Request[srvrivals.GetMatchInfoRequest]) (*connect_go.Response[srvrivals.GetMatchInfoResponse], error)
	// Lobby
	GetLatestLobbyActivityInfo(context.Context, *connect_go.Request[srvrivals.GetActivityRequest]) (*connect_go.Response[srvrivals.GetActivityResponse], error)
}

// NewThetanRivalServiceClient constructs a client for the services.ThetanRivalService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewThetanRivalServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ThetanRivalServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &thetanRivalServiceClient{
		getUserProfile: connect_go.NewClient[srvrivals.UserProfileRequest, srvrivals.UserProfileResponse](
			httpClient,
			baseURL+ThetanRivalServiceGetUserProfileProcedure,
			opts...,
		),
		getUserMinions: connect_go.NewClient[srvrivals.UserMinionsRequest, srvrivals.UserMinionsResponse](
			httpClient,
			baseURL+ThetanRivalServiceGetUserMinionsProcedure,
			opts...,
		),
		getUserSelectedMinion: connect_go.NewClient[srvrivals.UserSelectedMinionRequest, srvrivals.UserSelectedMinionResponse](
			httpClient,
			baseURL+ThetanRivalServiceGetUserSelectedMinionProcedure,
			opts...,
		),
		getMinion: connect_go.NewClient[srvrivals.MinionRequest, srvrivals.MinionResponse](
			httpClient,
			baseURL+ThetanRivalServiceGetMinionProcedure,
			opts...,
		),
		createMinion: connect_go.NewClient[srvrivals.CreateMinionRequest, srvrivals.CreateMinionResponse](
			httpClient,
			baseURL+ThetanRivalServiceCreateMinionProcedure,
			opts...,
		),
		getListFriends: connect_go.NewClient[srvrivals.GetUserFriendRequest, srvrivals.GetUserFriendResponse](
			httpClient,
			baseURL+ThetanRivalServiceGetListFriendsProcedure,
			opts...,
		),
		getMatchInfoOnboarding: connect_go.NewClient[srvrivals.GetMatchInfoRequest, srvrivals.GetMatchInfoResponse](
			httpClient,
			baseURL+ThetanRivalServiceGetMatchInfoOnboardingProcedure,
			opts...,
		),
		getLatestLobbyActivityInfo: connect_go.NewClient[srvrivals.GetActivityRequest, srvrivals.GetActivityResponse](
			httpClient,
			baseURL+ThetanRivalServiceGetLatestLobbyActivityInfoProcedure,
			opts...,
		),
	}
}

// thetanRivalServiceClient implements ThetanRivalServiceClient.
type thetanRivalServiceClient struct {
	getUserProfile             *connect_go.Client[srvrivals.UserProfileRequest, srvrivals.UserProfileResponse]
	getUserMinions             *connect_go.Client[srvrivals.UserMinionsRequest, srvrivals.UserMinionsResponse]
	getUserSelectedMinion      *connect_go.Client[srvrivals.UserSelectedMinionRequest, srvrivals.UserSelectedMinionResponse]
	getMinion                  *connect_go.Client[srvrivals.MinionRequest, srvrivals.MinionResponse]
	createMinion               *connect_go.Client[srvrivals.CreateMinionRequest, srvrivals.CreateMinionResponse]
	getListFriends             *connect_go.Client[srvrivals.GetUserFriendRequest, srvrivals.GetUserFriendResponse]
	getMatchInfoOnboarding     *connect_go.Client[srvrivals.GetMatchInfoRequest, srvrivals.GetMatchInfoResponse]
	getLatestLobbyActivityInfo *connect_go.Client[srvrivals.GetActivityRequest, srvrivals.GetActivityResponse]
}

// GetUserProfile calls services.ThetanRivalService.GetUserProfile.
func (c *thetanRivalServiceClient) GetUserProfile(ctx context.Context, req *connect_go.Request[srvrivals.UserProfileRequest]) (*connect_go.Response[srvrivals.UserProfileResponse], error) {
	return c.getUserProfile.CallUnary(ctx, req)
}

// GetUserMinions calls services.ThetanRivalService.GetUserMinions.
func (c *thetanRivalServiceClient) GetUserMinions(ctx context.Context, req *connect_go.Request[srvrivals.UserMinionsRequest]) (*connect_go.Response[srvrivals.UserMinionsResponse], error) {
	return c.getUserMinions.CallUnary(ctx, req)
}

// GetUserSelectedMinion calls services.ThetanRivalService.GetUserSelectedMinion.
func (c *thetanRivalServiceClient) GetUserSelectedMinion(ctx context.Context, req *connect_go.Request[srvrivals.UserSelectedMinionRequest]) (*connect_go.Response[srvrivals.UserSelectedMinionResponse], error) {
	return c.getUserSelectedMinion.CallUnary(ctx, req)
}

// GetMinion calls services.ThetanRivalService.GetMinion.
func (c *thetanRivalServiceClient) GetMinion(ctx context.Context, req *connect_go.Request[srvrivals.MinionRequest]) (*connect_go.Response[srvrivals.MinionResponse], error) {
	return c.getMinion.CallUnary(ctx, req)
}

// CreateMinion calls services.ThetanRivalService.CreateMinion.
func (c *thetanRivalServiceClient) CreateMinion(ctx context.Context, req *connect_go.Request[srvrivals.CreateMinionRequest]) (*connect_go.Response[srvrivals.CreateMinionResponse], error) {
	return c.createMinion.CallUnary(ctx, req)
}

// GetListFriends calls services.ThetanRivalService.GetListFriends.
func (c *thetanRivalServiceClient) GetListFriends(ctx context.Context, req *connect_go.Request[srvrivals.GetUserFriendRequest]) (*connect_go.Response[srvrivals.GetUserFriendResponse], error) {
	return c.getListFriends.CallUnary(ctx, req)
}

// GetMatchInfoOnboarding calls services.ThetanRivalService.GetMatchInfoOnboarding.
func (c *thetanRivalServiceClient) GetMatchInfoOnboarding(ctx context.Context, req *connect_go.Request[srvrivals.GetMatchInfoRequest]) (*connect_go.Response[srvrivals.GetMatchInfoResponse], error) {
	return c.getMatchInfoOnboarding.CallUnary(ctx, req)
}

// GetLatestLobbyActivityInfo calls services.ThetanRivalService.GetLatestLobbyActivityInfo.
func (c *thetanRivalServiceClient) GetLatestLobbyActivityInfo(ctx context.Context, req *connect_go.Request[srvrivals.GetActivityRequest]) (*connect_go.Response[srvrivals.GetActivityResponse], error) {
	return c.getLatestLobbyActivityInfo.CallUnary(ctx, req)
}

// ThetanRivalServiceHandler is an implementation of the services.ThetanRivalService service.
type ThetanRivalServiceHandler interface {
	// Profile
	GetUserProfile(context.Context, *connect_go.Request[srvrivals.UserProfileRequest]) (*connect_go.Response[srvrivals.UserProfileResponse], error)
	// Minions
	GetUserMinions(context.Context, *connect_go.Request[srvrivals.UserMinionsRequest]) (*connect_go.Response[srvrivals.UserMinionsResponse], error)
	GetUserSelectedMinion(context.Context, *connect_go.Request[srvrivals.UserSelectedMinionRequest]) (*connect_go.Response[srvrivals.UserSelectedMinionResponse], error)
	GetMinion(context.Context, *connect_go.Request[srvrivals.MinionRequest]) (*connect_go.Response[srvrivals.MinionResponse], error)
	CreateMinion(context.Context, *connect_go.Request[srvrivals.CreateMinionRequest]) (*connect_go.Response[srvrivals.CreateMinionResponse], error)
	// Friends
	GetListFriends(context.Context, *connect_go.Request[srvrivals.GetUserFriendRequest]) (*connect_go.Response[srvrivals.GetUserFriendResponse], error)
	GetMatchInfoOnboarding(context.Context, *connect_go.Request[srvrivals.GetMatchInfoRequest]) (*connect_go.Response[srvrivals.GetMatchInfoResponse], error)
	// Lobby
	GetLatestLobbyActivityInfo(context.Context, *connect_go.Request[srvrivals.GetActivityRequest]) (*connect_go.Response[srvrivals.GetActivityResponse], error)
}

// NewThetanRivalServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewThetanRivalServiceHandler(svc ThetanRivalServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ThetanRivalServiceGetUserProfileProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceGetUserProfileProcedure,
		svc.GetUserProfile,
		opts...,
	))
	mux.Handle(ThetanRivalServiceGetUserMinionsProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceGetUserMinionsProcedure,
		svc.GetUserMinions,
		opts...,
	))
	mux.Handle(ThetanRivalServiceGetUserSelectedMinionProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceGetUserSelectedMinionProcedure,
		svc.GetUserSelectedMinion,
		opts...,
	))
	mux.Handle(ThetanRivalServiceGetMinionProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceGetMinionProcedure,
		svc.GetMinion,
		opts...,
	))
	mux.Handle(ThetanRivalServiceCreateMinionProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceCreateMinionProcedure,
		svc.CreateMinion,
		opts...,
	))
	mux.Handle(ThetanRivalServiceGetListFriendsProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceGetListFriendsProcedure,
		svc.GetListFriends,
		opts...,
	))
	mux.Handle(ThetanRivalServiceGetMatchInfoOnboardingProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceGetMatchInfoOnboardingProcedure,
		svc.GetMatchInfoOnboarding,
		opts...,
	))
	mux.Handle(ThetanRivalServiceGetLatestLobbyActivityInfoProcedure, connect_go.NewUnaryHandler(
		ThetanRivalServiceGetLatestLobbyActivityInfoProcedure,
		svc.GetLatestLobbyActivityInfo,
		opts...,
	))
	return "/services.ThetanRivalService/", mux
}

// UnimplementedThetanRivalServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedThetanRivalServiceHandler struct{}

func (UnimplementedThetanRivalServiceHandler) GetUserProfile(context.Context, *connect_go.Request[srvrivals.UserProfileRequest]) (*connect_go.Response[srvrivals.UserProfileResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.GetUserProfile is not implemented"))
}

func (UnimplementedThetanRivalServiceHandler) GetUserMinions(context.Context, *connect_go.Request[srvrivals.UserMinionsRequest]) (*connect_go.Response[srvrivals.UserMinionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.GetUserMinions is not implemented"))
}

func (UnimplementedThetanRivalServiceHandler) GetUserSelectedMinion(context.Context, *connect_go.Request[srvrivals.UserSelectedMinionRequest]) (*connect_go.Response[srvrivals.UserSelectedMinionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.GetUserSelectedMinion is not implemented"))
}

func (UnimplementedThetanRivalServiceHandler) GetMinion(context.Context, *connect_go.Request[srvrivals.MinionRequest]) (*connect_go.Response[srvrivals.MinionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.GetMinion is not implemented"))
}

func (UnimplementedThetanRivalServiceHandler) CreateMinion(context.Context, *connect_go.Request[srvrivals.CreateMinionRequest]) (*connect_go.Response[srvrivals.CreateMinionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.CreateMinion is not implemented"))
}

func (UnimplementedThetanRivalServiceHandler) GetListFriends(context.Context, *connect_go.Request[srvrivals.GetUserFriendRequest]) (*connect_go.Response[srvrivals.GetUserFriendResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.GetListFriends is not implemented"))
}

func (UnimplementedThetanRivalServiceHandler) GetMatchInfoOnboarding(context.Context, *connect_go.Request[srvrivals.GetMatchInfoRequest]) (*connect_go.Response[srvrivals.GetMatchInfoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.GetMatchInfoOnboarding is not implemented"))
}

func (UnimplementedThetanRivalServiceHandler) GetLatestLobbyActivityInfo(context.Context, *connect_go.Request[srvrivals.GetActivityRequest]) (*connect_go.Response[srvrivals.GetActivityResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.ThetanRivalService.GetLatestLobbyActivityInfo is not implemented"))
}
