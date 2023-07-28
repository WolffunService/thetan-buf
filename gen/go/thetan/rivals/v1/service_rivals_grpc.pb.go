// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/rivals/v1/service_rivals.proto

package thetan_rivals_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ThetanRivalService_GetUserProfile_FullMethodName                = "/thetan.rivals.v1.ThetanRivalService/GetUserProfile"
	ThetanRivalService_GetUserMinions_FullMethodName                = "/thetan.rivals.v1.ThetanRivalService/GetUserMinions"
	ThetanRivalService_GetUserSelectedMinion_FullMethodName         = "/thetan.rivals.v1.ThetanRivalService/GetUserSelectedMinion"
	ThetanRivalService_GetMinion_FullMethodName                     = "/thetan.rivals.v1.ThetanRivalService/GetMinion"
	ThetanRivalService_CreateMinion_FullMethodName                  = "/thetan.rivals.v1.ThetanRivalService/CreateMinion"
	ThetanRivalService_GetListFriends_FullMethodName                = "/thetan.rivals.v1.ThetanRivalService/GetListFriends"
	ThetanRivalService_GetMatchInfoOnboarding_FullMethodName        = "/thetan.rivals.v1.ThetanRivalService/GetMatchInfoOnboarding"
	ThetanRivalService_GetLatestLobbyActivityInfo_FullMethodName    = "/thetan.rivals.v1.ThetanRivalService/GetLatestLobbyActivityInfo"
	ThetanRivalService_GetTownUser_FullMethodName                   = "/thetan.rivals.v1.ThetanRivalService/GetTownUser"
	ThetanRivalService_PickTownForUser_FullMethodName               = "/thetan.rivals.v1.ThetanRivalService/PickTownForUser"
	ThetanRivalService_TrackSession_FullMethodName                  = "/thetan.rivals.v1.ThetanRivalService/TrackSession"
	ThetanRivalService_GetFindMatchInfo_FullMethodName              = "/thetan.rivals.v1.ThetanRivalService/GetFindMatchInfo"
	ThetanRivalService_GetConfigForBot_FullMethodName               = "/thetan.rivals.v1.ThetanRivalService/GetConfigForBot"
	ThetanRivalService_GetActiveTournaments_FullMethodName          = "/thetan.rivals.v1.ThetanRivalService/GetActiveTournaments"
	ThetanRivalService_GetFindMatchInfoForTournament_FullMethodName = "/thetan.rivals.v1.ThetanRivalService/GetFindMatchInfoForTournament"
	ThetanRivalService_MatchFoundTournament_FullMethodName          = "/thetan.rivals.v1.ThetanRivalService/MatchFoundTournament"
)

// ThetanRivalServiceClient is the client API for ThetanRivalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThetanRivalServiceClient interface {
	// Profile
	GetUserProfile(ctx context.Context, in *UserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error)
	// Minions
	GetUserMinions(ctx context.Context, in *UserMinionsRequest, opts ...grpc.CallOption) (*UserMinionsResponse, error)
	GetUserSelectedMinion(ctx context.Context, in *UserSelectedMinionRequest, opts ...grpc.CallOption) (*UserSelectedMinionResponse, error)
	GetMinion(ctx context.Context, in *MinionRequest, opts ...grpc.CallOption) (*MinionResponse, error)
	CreateMinion(ctx context.Context, in *CreateMinionRequest, opts ...grpc.CallOption) (*CreateMinionResponse, error)
	// Friends
	GetListFriends(ctx context.Context, in *GetUserFriendRequest, opts ...grpc.CallOption) (*GetUserFriendResponse, error)
	GetMatchInfoOnboarding(ctx context.Context, in *GetMatchInfoRequest, opts ...grpc.CallOption) (*GetMatchInfoResponse, error)
	// Lobby
	GetLatestLobbyActivityInfo(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityResponse, error)
	GetTownUser(ctx context.Context, in *GetTownUserRequest, opts ...grpc.CallOption) (*LobbyTown, error)
	PickTownForUser(ctx context.Context, in *PickTownUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Track Action
	TrackSession(ctx context.Context, in *TrackSessionRequest, opts ...grpc.CallOption) (*TrackSessionResponse, error)
	// GetFindMatchInfo
	GetFindMatchInfo(ctx context.Context, in *FindMatchInfoRequest, opts ...grpc.CallOption) (*FindMatchInfoResponse, error)
	// Get config for bot
	GetConfigForBot(ctx context.Context, in *GetConfigForBotRequest, opts ...grpc.CallOption) (*GetConfigForBotResponse, error)
	// Tournament
	GetActiveTournaments(ctx context.Context, in *GetActiveTournamentsRequest, opts ...grpc.CallOption) (*GetActiveTournamentsResponse, error)
	GetFindMatchInfoForTournament(ctx context.Context, in *FindMatchInfoForTournamentRequest, opts ...grpc.CallOption) (*FindMatchInfoResponse, error)
	MatchFoundTournament(ctx context.Context, in *MatchFoundTournamentRequest, opts ...grpc.CallOption) (*MatchFoundTournamentResponse, error)
}

type thetanRivalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThetanRivalServiceClient(cc grpc.ClientConnInterface) ThetanRivalServiceClient {
	return &thetanRivalServiceClient{cc}
}

func (c *thetanRivalServiceClient) GetUserProfile(ctx context.Context, in *UserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error) {
	out := new(UserProfileResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetUserProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetUserMinions(ctx context.Context, in *UserMinionsRequest, opts ...grpc.CallOption) (*UserMinionsResponse, error) {
	out := new(UserMinionsResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetUserMinions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetUserSelectedMinion(ctx context.Context, in *UserSelectedMinionRequest, opts ...grpc.CallOption) (*UserSelectedMinionResponse, error) {
	out := new(UserSelectedMinionResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetUserSelectedMinion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetMinion(ctx context.Context, in *MinionRequest, opts ...grpc.CallOption) (*MinionResponse, error) {
	out := new(MinionResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetMinion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) CreateMinion(ctx context.Context, in *CreateMinionRequest, opts ...grpc.CallOption) (*CreateMinionResponse, error) {
	out := new(CreateMinionResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_CreateMinion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetListFriends(ctx context.Context, in *GetUserFriendRequest, opts ...grpc.CallOption) (*GetUserFriendResponse, error) {
	out := new(GetUserFriendResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetListFriends_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetMatchInfoOnboarding(ctx context.Context, in *GetMatchInfoRequest, opts ...grpc.CallOption) (*GetMatchInfoResponse, error) {
	out := new(GetMatchInfoResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetMatchInfoOnboarding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetLatestLobbyActivityInfo(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityResponse, error) {
	out := new(GetActivityResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetLatestLobbyActivityInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetTownUser(ctx context.Context, in *GetTownUserRequest, opts ...grpc.CallOption) (*LobbyTown, error) {
	out := new(LobbyTown)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetTownUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) PickTownForUser(ctx context.Context, in *PickTownUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_PickTownForUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) TrackSession(ctx context.Context, in *TrackSessionRequest, opts ...grpc.CallOption) (*TrackSessionResponse, error) {
	out := new(TrackSessionResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_TrackSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetFindMatchInfo(ctx context.Context, in *FindMatchInfoRequest, opts ...grpc.CallOption) (*FindMatchInfoResponse, error) {
	out := new(FindMatchInfoResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetFindMatchInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetConfigForBot(ctx context.Context, in *GetConfigForBotRequest, opts ...grpc.CallOption) (*GetConfigForBotResponse, error) {
	out := new(GetConfigForBotResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetConfigForBot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetActiveTournaments(ctx context.Context, in *GetActiveTournamentsRequest, opts ...grpc.CallOption) (*GetActiveTournamentsResponse, error) {
	out := new(GetActiveTournamentsResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetActiveTournaments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) GetFindMatchInfoForTournament(ctx context.Context, in *FindMatchInfoForTournamentRequest, opts ...grpc.CallOption) (*FindMatchInfoResponse, error) {
	out := new(FindMatchInfoResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_GetFindMatchInfoForTournament_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalServiceClient) MatchFoundTournament(ctx context.Context, in *MatchFoundTournamentRequest, opts ...grpc.CallOption) (*MatchFoundTournamentResponse, error) {
	out := new(MatchFoundTournamentResponse)
	err := c.cc.Invoke(ctx, ThetanRivalService_MatchFoundTournament_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetanRivalServiceServer is the server API for ThetanRivalService service.
// All implementations must embed UnimplementedThetanRivalServiceServer
// for forward compatibility
type ThetanRivalServiceServer interface {
	// Profile
	GetUserProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error)
	// Minions
	GetUserMinions(context.Context, *UserMinionsRequest) (*UserMinionsResponse, error)
	GetUserSelectedMinion(context.Context, *UserSelectedMinionRequest) (*UserSelectedMinionResponse, error)
	GetMinion(context.Context, *MinionRequest) (*MinionResponse, error)
	CreateMinion(context.Context, *CreateMinionRequest) (*CreateMinionResponse, error)
	// Friends
	GetListFriends(context.Context, *GetUserFriendRequest) (*GetUserFriendResponse, error)
	GetMatchInfoOnboarding(context.Context, *GetMatchInfoRequest) (*GetMatchInfoResponse, error)
	// Lobby
	GetLatestLobbyActivityInfo(context.Context, *GetActivityRequest) (*GetActivityResponse, error)
	GetTownUser(context.Context, *GetTownUserRequest) (*LobbyTown, error)
	PickTownForUser(context.Context, *PickTownUserRequest) (*EmptyResponse, error)
	// Track Action
	TrackSession(context.Context, *TrackSessionRequest) (*TrackSessionResponse, error)
	// GetFindMatchInfo
	GetFindMatchInfo(context.Context, *FindMatchInfoRequest) (*FindMatchInfoResponse, error)
	// Get config for bot
	GetConfigForBot(context.Context, *GetConfigForBotRequest) (*GetConfigForBotResponse, error)
	// Tournament
	GetActiveTournaments(context.Context, *GetActiveTournamentsRequest) (*GetActiveTournamentsResponse, error)
	GetFindMatchInfoForTournament(context.Context, *FindMatchInfoForTournamentRequest) (*FindMatchInfoResponse, error)
	MatchFoundTournament(context.Context, *MatchFoundTournamentRequest) (*MatchFoundTournamentResponse, error)
	mustEmbedUnimplementedThetanRivalServiceServer()
}

// UnimplementedThetanRivalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThetanRivalServiceServer struct {
}

func (UnimplementedThetanRivalServiceServer) GetUserProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetUserMinions(context.Context, *UserMinionsRequest) (*UserMinionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMinions not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetUserSelectedMinion(context.Context, *UserSelectedMinionRequest) (*UserSelectedMinionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSelectedMinion not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetMinion(context.Context, *MinionRequest) (*MinionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinion not implemented")
}
func (UnimplementedThetanRivalServiceServer) CreateMinion(context.Context, *CreateMinionRequest) (*CreateMinionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMinion not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetListFriends(context.Context, *GetUserFriendRequest) (*GetUserFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFriends not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetMatchInfoOnboarding(context.Context, *GetMatchInfoRequest) (*GetMatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchInfoOnboarding not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetLatestLobbyActivityInfo(context.Context, *GetActivityRequest) (*GetActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestLobbyActivityInfo not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetTownUser(context.Context, *GetTownUserRequest) (*LobbyTown, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTownUser not implemented")
}
func (UnimplementedThetanRivalServiceServer) PickTownForUser(context.Context, *PickTownUserRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickTownForUser not implemented")
}
func (UnimplementedThetanRivalServiceServer) TrackSession(context.Context, *TrackSessionRequest) (*TrackSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackSession not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetFindMatchInfo(context.Context, *FindMatchInfoRequest) (*FindMatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFindMatchInfo not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetConfigForBot(context.Context, *GetConfigForBotRequest) (*GetConfigForBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigForBot not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetActiveTournaments(context.Context, *GetActiveTournamentsRequest) (*GetActiveTournamentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveTournaments not implemented")
}
func (UnimplementedThetanRivalServiceServer) GetFindMatchInfoForTournament(context.Context, *FindMatchInfoForTournamentRequest) (*FindMatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFindMatchInfoForTournament not implemented")
}
func (UnimplementedThetanRivalServiceServer) MatchFoundTournament(context.Context, *MatchFoundTournamentRequest) (*MatchFoundTournamentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatchFoundTournament not implemented")
}
func (UnimplementedThetanRivalServiceServer) mustEmbedUnimplementedThetanRivalServiceServer() {}

// UnsafeThetanRivalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThetanRivalServiceServer will
// result in compilation errors.
type UnsafeThetanRivalServiceServer interface {
	mustEmbedUnimplementedThetanRivalServiceServer()
}

func RegisterThetanRivalServiceServer(s grpc.ServiceRegistrar, srv ThetanRivalServiceServer) {
	s.RegisterService(&ThetanRivalService_ServiceDesc, srv)
}

func _ThetanRivalService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetUserProfile(ctx, req.(*UserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetUserMinions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMinionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetUserMinions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetUserMinions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetUserMinions(ctx, req.(*UserMinionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetUserSelectedMinion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSelectedMinionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetUserSelectedMinion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetUserSelectedMinion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetUserSelectedMinion(ctx, req.(*UserSelectedMinionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetMinion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetMinion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetMinion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetMinion(ctx, req.(*MinionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_CreateMinion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMinionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).CreateMinion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_CreateMinion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).CreateMinion(ctx, req.(*CreateMinionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetListFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetListFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetListFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetListFriends(ctx, req.(*GetUserFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetMatchInfoOnboarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetMatchInfoOnboarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetMatchInfoOnboarding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetMatchInfoOnboarding(ctx, req.(*GetMatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetLatestLobbyActivityInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetLatestLobbyActivityInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetLatestLobbyActivityInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetLatestLobbyActivityInfo(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetTownUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTownUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetTownUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetTownUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetTownUser(ctx, req.(*GetTownUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_PickTownForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PickTownUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).PickTownForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_PickTownForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).PickTownForUser(ctx, req.(*PickTownUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_TrackSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).TrackSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_TrackSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).TrackSession(ctx, req.(*TrackSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetFindMatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetFindMatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetFindMatchInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetFindMatchInfo(ctx, req.(*FindMatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetConfigForBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigForBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetConfigForBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetConfigForBot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetConfigForBot(ctx, req.(*GetConfigForBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetActiveTournaments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveTournamentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetActiveTournaments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetActiveTournaments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetActiveTournaments(ctx, req.(*GetActiveTournamentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_GetFindMatchInfoForTournament_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMatchInfoForTournamentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).GetFindMatchInfoForTournament(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_GetFindMatchInfoForTournament_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).GetFindMatchInfoForTournament(ctx, req.(*FindMatchInfoForTournamentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalService_MatchFoundTournament_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchFoundTournamentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalServiceServer).MatchFoundTournament(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalService_MatchFoundTournament_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalServiceServer).MatchFoundTournament(ctx, req.(*MatchFoundTournamentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThetanRivalService_ServiceDesc is the grpc.ServiceDesc for ThetanRivalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThetanRivalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.rivals.v1.ThetanRivalService",
	HandlerType: (*ThetanRivalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserProfile",
			Handler:    _ThetanRivalService_GetUserProfile_Handler,
		},
		{
			MethodName: "GetUserMinions",
			Handler:    _ThetanRivalService_GetUserMinions_Handler,
		},
		{
			MethodName: "GetUserSelectedMinion",
			Handler:    _ThetanRivalService_GetUserSelectedMinion_Handler,
		},
		{
			MethodName: "GetMinion",
			Handler:    _ThetanRivalService_GetMinion_Handler,
		},
		{
			MethodName: "CreateMinion",
			Handler:    _ThetanRivalService_CreateMinion_Handler,
		},
		{
			MethodName: "GetListFriends",
			Handler:    _ThetanRivalService_GetListFriends_Handler,
		},
		{
			MethodName: "GetMatchInfoOnboarding",
			Handler:    _ThetanRivalService_GetMatchInfoOnboarding_Handler,
		},
		{
			MethodName: "GetLatestLobbyActivityInfo",
			Handler:    _ThetanRivalService_GetLatestLobbyActivityInfo_Handler,
		},
		{
			MethodName: "GetTownUser",
			Handler:    _ThetanRivalService_GetTownUser_Handler,
		},
		{
			MethodName: "PickTownForUser",
			Handler:    _ThetanRivalService_PickTownForUser_Handler,
		},
		{
			MethodName: "TrackSession",
			Handler:    _ThetanRivalService_TrackSession_Handler,
		},
		{
			MethodName: "GetFindMatchInfo",
			Handler:    _ThetanRivalService_GetFindMatchInfo_Handler,
		},
		{
			MethodName: "GetConfigForBot",
			Handler:    _ThetanRivalService_GetConfigForBot_Handler,
		},
		{
			MethodName: "GetActiveTournaments",
			Handler:    _ThetanRivalService_GetActiveTournaments_Handler,
		},
		{
			MethodName: "GetFindMatchInfoForTournament",
			Handler:    _ThetanRivalService_GetFindMatchInfoForTournament_Handler,
		},
		{
			MethodName: "MatchFoundTournament",
			Handler:    _ThetanRivalService_MatchFoundTournament_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/rivals/v1/service_rivals.proto",
}
