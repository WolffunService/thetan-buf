// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/rivals/v1/player.proto

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
	ThetanRivalsPlayerService_CreatePlayersInfo_FullMethodName              = "/thetan.rivals.v1.ThetanRivalsPlayerService/CreatePlayersInfo"
	ThetanRivalsPlayerService_UpdatePlayersInfo_FullMethodName              = "/thetan.rivals.v1.ThetanRivalsPlayerService/UpdatePlayersInfo"
	ThetanRivalsPlayerService_FindEligibleGuildForPlayerJoin_FullMethodName = "/thetan.rivals.v1.ThetanRivalsPlayerService/FindEligibleGuildForPlayerJoin"
	ThetanRivalsPlayerService_GuildAction_FullMethodName                    = "/thetan.rivals.v1.ThetanRivalsPlayerService/GuildAction"
)

// ThetanRivalsPlayerServiceClient is the client API for ThetanRivalsPlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThetanRivalsPlayerServiceClient interface {
	CreatePlayersInfo(ctx context.Context, in *CreatePlayersInfoRequest, opts ...grpc.CallOption) (*CreatePlayersInfoResponse, error)
	UpdatePlayersInfo(ctx context.Context, in *UpdatePlayersInfoRequest, opts ...grpc.CallOption) (*UpdatePlayersInfoResponse, error)
	FindEligibleGuildForPlayerJoin(ctx context.Context, in *FindEligibleGuildForPlayerJoinRequest, opts ...grpc.CallOption) (*FindEligibleGuildForPlayerJoinResponse, error)
	GuildAction(ctx context.Context, in *GuildActionRequest, opts ...grpc.CallOption) (*GuildActionResponse, error)
}

type thetanRivalsPlayerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThetanRivalsPlayerServiceClient(cc grpc.ClientConnInterface) ThetanRivalsPlayerServiceClient {
	return &thetanRivalsPlayerServiceClient{cc}
}

func (c *thetanRivalsPlayerServiceClient) CreatePlayersInfo(ctx context.Context, in *CreatePlayersInfoRequest, opts ...grpc.CallOption) (*CreatePlayersInfoResponse, error) {
	out := new(CreatePlayersInfoResponse)
	err := c.cc.Invoke(ctx, ThetanRivalsPlayerService_CreatePlayersInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalsPlayerServiceClient) UpdatePlayersInfo(ctx context.Context, in *UpdatePlayersInfoRequest, opts ...grpc.CallOption) (*UpdatePlayersInfoResponse, error) {
	out := new(UpdatePlayersInfoResponse)
	err := c.cc.Invoke(ctx, ThetanRivalsPlayerService_UpdatePlayersInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalsPlayerServiceClient) FindEligibleGuildForPlayerJoin(ctx context.Context, in *FindEligibleGuildForPlayerJoinRequest, opts ...grpc.CallOption) (*FindEligibleGuildForPlayerJoinResponse, error) {
	out := new(FindEligibleGuildForPlayerJoinResponse)
	err := c.cc.Invoke(ctx, ThetanRivalsPlayerService_FindEligibleGuildForPlayerJoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanRivalsPlayerServiceClient) GuildAction(ctx context.Context, in *GuildActionRequest, opts ...grpc.CallOption) (*GuildActionResponse, error) {
	out := new(GuildActionResponse)
	err := c.cc.Invoke(ctx, ThetanRivalsPlayerService_GuildAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetanRivalsPlayerServiceServer is the server API for ThetanRivalsPlayerService service.
// All implementations must embed UnimplementedThetanRivalsPlayerServiceServer
// for forward compatibility
type ThetanRivalsPlayerServiceServer interface {
	CreatePlayersInfo(context.Context, *CreatePlayersInfoRequest) (*CreatePlayersInfoResponse, error)
	UpdatePlayersInfo(context.Context, *UpdatePlayersInfoRequest) (*UpdatePlayersInfoResponse, error)
	FindEligibleGuildForPlayerJoin(context.Context, *FindEligibleGuildForPlayerJoinRequest) (*FindEligibleGuildForPlayerJoinResponse, error)
	GuildAction(context.Context, *GuildActionRequest) (*GuildActionResponse, error)
	mustEmbedUnimplementedThetanRivalsPlayerServiceServer()
}

// UnimplementedThetanRivalsPlayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThetanRivalsPlayerServiceServer struct {
}

func (UnimplementedThetanRivalsPlayerServiceServer) CreatePlayersInfo(context.Context, *CreatePlayersInfoRequest) (*CreatePlayersInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayersInfo not implemented")
}
func (UnimplementedThetanRivalsPlayerServiceServer) UpdatePlayersInfo(context.Context, *UpdatePlayersInfoRequest) (*UpdatePlayersInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlayersInfo not implemented")
}
func (UnimplementedThetanRivalsPlayerServiceServer) FindEligibleGuildForPlayerJoin(context.Context, *FindEligibleGuildForPlayerJoinRequest) (*FindEligibleGuildForPlayerJoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEligibleGuildForPlayerJoin not implemented")
}
func (UnimplementedThetanRivalsPlayerServiceServer) GuildAction(context.Context, *GuildActionRequest) (*GuildActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuildAction not implemented")
}
func (UnimplementedThetanRivalsPlayerServiceServer) mustEmbedUnimplementedThetanRivalsPlayerServiceServer() {
}

// UnsafeThetanRivalsPlayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThetanRivalsPlayerServiceServer will
// result in compilation errors.
type UnsafeThetanRivalsPlayerServiceServer interface {
	mustEmbedUnimplementedThetanRivalsPlayerServiceServer()
}

func RegisterThetanRivalsPlayerServiceServer(s grpc.ServiceRegistrar, srv ThetanRivalsPlayerServiceServer) {
	s.RegisterService(&ThetanRivalsPlayerService_ServiceDesc, srv)
}

func _ThetanRivalsPlayerService_CreatePlayersInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayersInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalsPlayerServiceServer).CreatePlayersInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalsPlayerService_CreatePlayersInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalsPlayerServiceServer).CreatePlayersInfo(ctx, req.(*CreatePlayersInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalsPlayerService_UpdatePlayersInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlayersInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalsPlayerServiceServer).UpdatePlayersInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalsPlayerService_UpdatePlayersInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalsPlayerServiceServer).UpdatePlayersInfo(ctx, req.(*UpdatePlayersInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalsPlayerService_FindEligibleGuildForPlayerJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEligibleGuildForPlayerJoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalsPlayerServiceServer).FindEligibleGuildForPlayerJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalsPlayerService_FindEligibleGuildForPlayerJoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalsPlayerServiceServer).FindEligibleGuildForPlayerJoin(ctx, req.(*FindEligibleGuildForPlayerJoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanRivalsPlayerService_GuildAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuildActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanRivalsPlayerServiceServer).GuildAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanRivalsPlayerService_GuildAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanRivalsPlayerServiceServer).GuildAction(ctx, req.(*GuildActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThetanRivalsPlayerService_ServiceDesc is the grpc.ServiceDesc for ThetanRivalsPlayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThetanRivalsPlayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.rivals.v1.ThetanRivalsPlayerService",
	HandlerType: (*ThetanRivalsPlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlayersInfo",
			Handler:    _ThetanRivalsPlayerService_CreatePlayersInfo_Handler,
		},
		{
			MethodName: "UpdatePlayersInfo",
			Handler:    _ThetanRivalsPlayerService_UpdatePlayersInfo_Handler,
		},
		{
			MethodName: "FindEligibleGuildForPlayerJoin",
			Handler:    _ThetanRivalsPlayerService_FindEligibleGuildForPlayerJoin_Handler,
		},
		{
			MethodName: "GuildAction",
			Handler:    _ThetanRivalsPlayerService_GuildAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/rivals/v1/player.proto",
}
