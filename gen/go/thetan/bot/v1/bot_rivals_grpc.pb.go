// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/bot/v1/bot_rivals.proto

package thetan_bot_v1

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
	BotRivalsService_FetchLobbyBots_FullMethodName = "/thetan.rivals.v1.BotRivalsService/FetchLobbyBots"
	BotRivalsService_LobbyBotAction_FullMethodName = "/thetan.rivals.v1.BotRivalsService/LobbyBotAction"
)

// BotRivalsServiceClient is the client API for BotRivalsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotRivalsServiceClient interface {
	FetchLobbyBots(ctx context.Context, in *FetchLobbyBotsRequest, opts ...grpc.CallOption) (*FetchLobbyBotsResponse, error)
	LobbyBotAction(ctx context.Context, in *LobbyBotActionRequest, opts ...grpc.CallOption) (*LobbyBotActionResponse, error)
}

type botRivalsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBotRivalsServiceClient(cc grpc.ClientConnInterface) BotRivalsServiceClient {
	return &botRivalsServiceClient{cc}
}

func (c *botRivalsServiceClient) FetchLobbyBots(ctx context.Context, in *FetchLobbyBotsRequest, opts ...grpc.CallOption) (*FetchLobbyBotsResponse, error) {
	out := new(FetchLobbyBotsResponse)
	err := c.cc.Invoke(ctx, BotRivalsService_FetchLobbyBots_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botRivalsServiceClient) LobbyBotAction(ctx context.Context, in *LobbyBotActionRequest, opts ...grpc.CallOption) (*LobbyBotActionResponse, error) {
	out := new(LobbyBotActionResponse)
	err := c.cc.Invoke(ctx, BotRivalsService_LobbyBotAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotRivalsServiceServer is the server API for BotRivalsService service.
// All implementations must embed UnimplementedBotRivalsServiceServer
// for forward compatibility
type BotRivalsServiceServer interface {
	FetchLobbyBots(context.Context, *FetchLobbyBotsRequest) (*FetchLobbyBotsResponse, error)
	LobbyBotAction(context.Context, *LobbyBotActionRequest) (*LobbyBotActionResponse, error)
	mustEmbedUnimplementedBotRivalsServiceServer()
}

// UnimplementedBotRivalsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBotRivalsServiceServer struct {
}

func (UnimplementedBotRivalsServiceServer) FetchLobbyBots(context.Context, *FetchLobbyBotsRequest) (*FetchLobbyBotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchLobbyBots not implemented")
}
func (UnimplementedBotRivalsServiceServer) LobbyBotAction(context.Context, *LobbyBotActionRequest) (*LobbyBotActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LobbyBotAction not implemented")
}
func (UnimplementedBotRivalsServiceServer) mustEmbedUnimplementedBotRivalsServiceServer() {}

// UnsafeBotRivalsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotRivalsServiceServer will
// result in compilation errors.
type UnsafeBotRivalsServiceServer interface {
	mustEmbedUnimplementedBotRivalsServiceServer()
}

func RegisterBotRivalsServiceServer(s grpc.ServiceRegistrar, srv BotRivalsServiceServer) {
	s.RegisterService(&BotRivalsService_ServiceDesc, srv)
}

func _BotRivalsService_FetchLobbyBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchLobbyBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotRivalsServiceServer).FetchLobbyBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotRivalsService_FetchLobbyBots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotRivalsServiceServer).FetchLobbyBots(ctx, req.(*FetchLobbyBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotRivalsService_LobbyBotAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LobbyBotActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotRivalsServiceServer).LobbyBotAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotRivalsService_LobbyBotAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotRivalsServiceServer).LobbyBotAction(ctx, req.(*LobbyBotActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BotRivalsService_ServiceDesc is the grpc.ServiceDesc for BotRivalsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotRivalsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.rivals.v1.BotRivalsService",
	HandlerType: (*BotRivalsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchLobbyBots",
			Handler:    _BotRivalsService_FetchLobbyBots_Handler,
		},
		{
			MethodName: "LobbyBotAction",
			Handler:    _BotRivalsService_LobbyBotAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/bot/v1/bot_rivals.proto",
}
