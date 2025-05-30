// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/immortal/v1/immortal.proto

package thetan_immortal_v1

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
	ImmortalService_SearchPlayerInfo_FullMethodName  = "/thetan.immortal.v1.ImmortalService/SearchPlayerInfo"
	ImmortalService_GetUserProfile_FullMethodName    = "/thetan.immortal.v1.ImmortalService/GetUserProfile"
	ImmortalService_BattleEnd_FullMethodName         = "/thetan.immortal.v1.ImmortalService/BattleEnd"
	ImmortalService_GetHeroConfig_FullMethodName     = "/thetan.immortal.v1.ImmortalService/GetHeroConfig"
	ImmortalService_GetSkillConfig_FullMethodName    = "/thetan.immortal.v1.ImmortalService/GetSkillConfig"
	ImmortalService_GetListFriends_FullMethodName    = "/thetan.immortal.v1.ImmortalService/GetListFriends"
	ImmortalService_GetSeasonal_FullMethodName       = "/thetan.immortal.v1.ImmortalService/GetSeasonal"
	ImmortalService_TrackStartSession_FullMethodName = "/thetan.immortal.v1.ImmortalService/TrackStartSession"
	ImmortalService_GetGameData_FullMethodName       = "/thetan.immortal.v1.ImmortalService/GetGameData"
)

// ImmortalServiceClient is the client API for ImmortalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImmortalServiceClient interface {
	SearchPlayerInfo(ctx context.Context, in *SearchPlayerInfoRequest, opts ...grpc.CallOption) (*SearchPlayerInfoResponse, error)
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error)
	BattleEnd(ctx context.Context, in *BattleEndRequest, opts ...grpc.CallOption) (ImmortalService_BattleEndClient, error)
	GetHeroConfig(ctx context.Context, in *GetHeroConfigRequest, opts ...grpc.CallOption) (*GetHeroConfigResponse, error)
	GetSkillConfig(ctx context.Context, in *GetSkillConfigRequest, opts ...grpc.CallOption) (*GetSkillConfigResponse, error)
	// Friends
	GetListFriends(ctx context.Context, in *GetUserFriendRequest, opts ...grpc.CallOption) (*GetUserFriendResponse, error)
	// Seasonal
	GetSeasonal(ctx context.Context, in *GetSeasonalRequest, opts ...grpc.CallOption) (*GetSeasonalResponse, error)
	TrackStartSession(ctx context.Context, in *TrackSessionRequest, opts ...grpc.CallOption) (*TrackSessionResponse, error)
	GetGameData(ctx context.Context, in *GetGameDataRequest, opts ...grpc.CallOption) (*GetGameDataResponse, error)
}

type immortalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImmortalServiceClient(cc grpc.ClientConnInterface) ImmortalServiceClient {
	return &immortalServiceClient{cc}
}

func (c *immortalServiceClient) SearchPlayerInfo(ctx context.Context, in *SearchPlayerInfoRequest, opts ...grpc.CallOption) (*SearchPlayerInfoResponse, error) {
	out := new(SearchPlayerInfoResponse)
	err := c.cc.Invoke(ctx, ImmortalService_SearchPlayerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error) {
	out := new(GetUserProfileResponse)
	err := c.cc.Invoke(ctx, ImmortalService_GetUserProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) BattleEnd(ctx context.Context, in *BattleEndRequest, opts ...grpc.CallOption) (ImmortalService_BattleEndClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImmortalService_ServiceDesc.Streams[0], ImmortalService_BattleEnd_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &immortalServiceBattleEndClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ImmortalService_BattleEndClient interface {
	Recv() (*BattleEndResponse, error)
	grpc.ClientStream
}

type immortalServiceBattleEndClient struct {
	grpc.ClientStream
}

func (x *immortalServiceBattleEndClient) Recv() (*BattleEndResponse, error) {
	m := new(BattleEndResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *immortalServiceClient) GetHeroConfig(ctx context.Context, in *GetHeroConfigRequest, opts ...grpc.CallOption) (*GetHeroConfigResponse, error) {
	out := new(GetHeroConfigResponse)
	err := c.cc.Invoke(ctx, ImmortalService_GetHeroConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) GetSkillConfig(ctx context.Context, in *GetSkillConfigRequest, opts ...grpc.CallOption) (*GetSkillConfigResponse, error) {
	out := new(GetSkillConfigResponse)
	err := c.cc.Invoke(ctx, ImmortalService_GetSkillConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) GetListFriends(ctx context.Context, in *GetUserFriendRequest, opts ...grpc.CallOption) (*GetUserFriendResponse, error) {
	out := new(GetUserFriendResponse)
	err := c.cc.Invoke(ctx, ImmortalService_GetListFriends_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) GetSeasonal(ctx context.Context, in *GetSeasonalRequest, opts ...grpc.CallOption) (*GetSeasonalResponse, error) {
	out := new(GetSeasonalResponse)
	err := c.cc.Invoke(ctx, ImmortalService_GetSeasonal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) TrackStartSession(ctx context.Context, in *TrackSessionRequest, opts ...grpc.CallOption) (*TrackSessionResponse, error) {
	out := new(TrackSessionResponse)
	err := c.cc.Invoke(ctx, ImmortalService_TrackStartSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) GetGameData(ctx context.Context, in *GetGameDataRequest, opts ...grpc.CallOption) (*GetGameDataResponse, error) {
	out := new(GetGameDataResponse)
	err := c.cc.Invoke(ctx, ImmortalService_GetGameData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImmortalServiceServer is the server API for ImmortalService service.
// All implementations must embed UnimplementedImmortalServiceServer
// for forward compatibility
type ImmortalServiceServer interface {
	SearchPlayerInfo(context.Context, *SearchPlayerInfoRequest) (*SearchPlayerInfoResponse, error)
	GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error)
	BattleEnd(*BattleEndRequest, ImmortalService_BattleEndServer) error
	GetHeroConfig(context.Context, *GetHeroConfigRequest) (*GetHeroConfigResponse, error)
	GetSkillConfig(context.Context, *GetSkillConfigRequest) (*GetSkillConfigResponse, error)
	// Friends
	GetListFriends(context.Context, *GetUserFriendRequest) (*GetUserFriendResponse, error)
	// Seasonal
	GetSeasonal(context.Context, *GetSeasonalRequest) (*GetSeasonalResponse, error)
	TrackStartSession(context.Context, *TrackSessionRequest) (*TrackSessionResponse, error)
	GetGameData(context.Context, *GetGameDataRequest) (*GetGameDataResponse, error)
	mustEmbedUnimplementedImmortalServiceServer()
}

// UnimplementedImmortalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImmortalServiceServer struct {
}

func (UnimplementedImmortalServiceServer) SearchPlayerInfo(context.Context, *SearchPlayerInfoRequest) (*SearchPlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlayerInfo not implemented")
}
func (UnimplementedImmortalServiceServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedImmortalServiceServer) BattleEnd(*BattleEndRequest, ImmortalService_BattleEndServer) error {
	return status.Errorf(codes.Unimplemented, "method BattleEnd not implemented")
}
func (UnimplementedImmortalServiceServer) GetHeroConfig(context.Context, *GetHeroConfigRequest) (*GetHeroConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeroConfig not implemented")
}
func (UnimplementedImmortalServiceServer) GetSkillConfig(context.Context, *GetSkillConfigRequest) (*GetSkillConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkillConfig not implemented")
}
func (UnimplementedImmortalServiceServer) GetListFriends(context.Context, *GetUserFriendRequest) (*GetUserFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListFriends not implemented")
}
func (UnimplementedImmortalServiceServer) GetSeasonal(context.Context, *GetSeasonalRequest) (*GetSeasonalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeasonal not implemented")
}
func (UnimplementedImmortalServiceServer) TrackStartSession(context.Context, *TrackSessionRequest) (*TrackSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackStartSession not implemented")
}
func (UnimplementedImmortalServiceServer) GetGameData(context.Context, *GetGameDataRequest) (*GetGameDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameData not implemented")
}
func (UnimplementedImmortalServiceServer) mustEmbedUnimplementedImmortalServiceServer() {}

// UnsafeImmortalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImmortalServiceServer will
// result in compilation errors.
type UnsafeImmortalServiceServer interface {
	mustEmbedUnimplementedImmortalServiceServer()
}

func RegisterImmortalServiceServer(s grpc.ServiceRegistrar, srv ImmortalServiceServer) {
	s.RegisterService(&ImmortalService_ServiceDesc, srv)
}

func _ImmortalService_SearchPlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).SearchPlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_SearchPlayerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).SearchPlayerInfo(ctx, req.(*SearchPlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_BattleEnd_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BattleEndRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImmortalServiceServer).BattleEnd(m, &immortalServiceBattleEndServer{stream})
}

type ImmortalService_BattleEndServer interface {
	Send(*BattleEndResponse) error
	grpc.ServerStream
}

type immortalServiceBattleEndServer struct {
	grpc.ServerStream
}

func (x *immortalServiceBattleEndServer) Send(m *BattleEndResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ImmortalService_GetHeroConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeroConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).GetHeroConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_GetHeroConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).GetHeroConfig(ctx, req.(*GetHeroConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_GetSkillConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).GetSkillConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_GetSkillConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).GetSkillConfig(ctx, req.(*GetSkillConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_GetListFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).GetListFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_GetListFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).GetListFriends(ctx, req.(*GetUserFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_GetSeasonal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeasonalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).GetSeasonal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_GetSeasonal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).GetSeasonal(ctx, req.(*GetSeasonalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_TrackStartSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).TrackStartSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_TrackStartSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).TrackStartSession(ctx, req.(*TrackSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_GetGameData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).GetGameData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_GetGameData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).GetGameData(ctx, req.(*GetGameDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImmortalService_ServiceDesc is the grpc.ServiceDesc for ImmortalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImmortalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.immortal.v1.ImmortalService",
	HandlerType: (*ImmortalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchPlayerInfo",
			Handler:    _ImmortalService_SearchPlayerInfo_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _ImmortalService_GetUserProfile_Handler,
		},
		{
			MethodName: "GetHeroConfig",
			Handler:    _ImmortalService_GetHeroConfig_Handler,
		},
		{
			MethodName: "GetSkillConfig",
			Handler:    _ImmortalService_GetSkillConfig_Handler,
		},
		{
			MethodName: "GetListFriends",
			Handler:    _ImmortalService_GetListFriends_Handler,
		},
		{
			MethodName: "GetSeasonal",
			Handler:    _ImmortalService_GetSeasonal_Handler,
		},
		{
			MethodName: "TrackStartSession",
			Handler:    _ImmortalService_TrackStartSession_Handler,
		},
		{
			MethodName: "GetGameData",
			Handler:    _ImmortalService_GetGameData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BattleEnd",
			Handler:       _ImmortalService_BattleEnd_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "thetan/immortal/v1/immortal.proto",
}
