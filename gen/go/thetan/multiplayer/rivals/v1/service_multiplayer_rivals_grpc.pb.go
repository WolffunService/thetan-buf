// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/multiplayer/rivals/v1/service_multiplayer_rivals.proto

package thetan_multiplayer_rivals_v1

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
	RivalsMultiplayerService_GetOnlineStatus_FullMethodName    = "/thetan.multiplayer.rivals.v1.RivalsMultiplayerService/GetOnlineStatus"
	RivalsMultiplayerService_Notify_FullMethodName             = "/thetan.multiplayer.rivals.v1.RivalsMultiplayerService/Notify"
	RivalsMultiplayerService_GetTownOnlineUsers_FullMethodName = "/thetan.multiplayer.rivals.v1.RivalsMultiplayerService/GetTownOnlineUsers"
)

// RivalsMultiplayerServiceClient is the client API for RivalsMultiplayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RivalsMultiplayerServiceClient interface {
	GetOnlineStatus(ctx context.Context, in *GetOnlineStatusRequest, opts ...grpc.CallOption) (*GetOnlineStatusResponse, error)
	// dùng để bắn event websocket trực tiếp đến user.
	// nếu cần bắn event thường xuyên, nên để user subscribe vào channel và bắn event vào channel đó.
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	GetTownOnlineUsers(ctx context.Context, in *GetTownOnlineUsersRequest, opts ...grpc.CallOption) (*GetTownOnlineUsersResponse, error)
}

type rivalsMultiplayerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRivalsMultiplayerServiceClient(cc grpc.ClientConnInterface) RivalsMultiplayerServiceClient {
	return &rivalsMultiplayerServiceClient{cc}
}

func (c *rivalsMultiplayerServiceClient) GetOnlineStatus(ctx context.Context, in *GetOnlineStatusRequest, opts ...grpc.CallOption) (*GetOnlineStatusResponse, error) {
	out := new(GetOnlineStatusResponse)
	err := c.cc.Invoke(ctx, RivalsMultiplayerService_GetOnlineStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rivalsMultiplayerServiceClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, RivalsMultiplayerService_Notify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rivalsMultiplayerServiceClient) GetTownOnlineUsers(ctx context.Context, in *GetTownOnlineUsersRequest, opts ...grpc.CallOption) (*GetTownOnlineUsersResponse, error) {
	out := new(GetTownOnlineUsersResponse)
	err := c.cc.Invoke(ctx, RivalsMultiplayerService_GetTownOnlineUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RivalsMultiplayerServiceServer is the server API for RivalsMultiplayerService service.
// All implementations must embed UnimplementedRivalsMultiplayerServiceServer
// for forward compatibility
type RivalsMultiplayerServiceServer interface {
	GetOnlineStatus(context.Context, *GetOnlineStatusRequest) (*GetOnlineStatusResponse, error)
	// dùng để bắn event websocket trực tiếp đến user.
	// nếu cần bắn event thường xuyên, nên để user subscribe vào channel và bắn event vào channel đó.
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
	GetTownOnlineUsers(context.Context, *GetTownOnlineUsersRequest) (*GetTownOnlineUsersResponse, error)
	mustEmbedUnimplementedRivalsMultiplayerServiceServer()
}

// UnimplementedRivalsMultiplayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRivalsMultiplayerServiceServer struct {
}

func (UnimplementedRivalsMultiplayerServiceServer) GetOnlineStatus(context.Context, *GetOnlineStatusRequest) (*GetOnlineStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineStatus not implemented")
}
func (UnimplementedRivalsMultiplayerServiceServer) Notify(context.Context, *NotifyRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedRivalsMultiplayerServiceServer) GetTownOnlineUsers(context.Context, *GetTownOnlineUsersRequest) (*GetTownOnlineUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTownOnlineUsers not implemented")
}
func (UnimplementedRivalsMultiplayerServiceServer) mustEmbedUnimplementedRivalsMultiplayerServiceServer() {
}

// UnsafeRivalsMultiplayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RivalsMultiplayerServiceServer will
// result in compilation errors.
type UnsafeRivalsMultiplayerServiceServer interface {
	mustEmbedUnimplementedRivalsMultiplayerServiceServer()
}

func RegisterRivalsMultiplayerServiceServer(s grpc.ServiceRegistrar, srv RivalsMultiplayerServiceServer) {
	s.RegisterService(&RivalsMultiplayerService_ServiceDesc, srv)
}

func _RivalsMultiplayerService_GetOnlineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOnlineStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RivalsMultiplayerServiceServer).GetOnlineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RivalsMultiplayerService_GetOnlineStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RivalsMultiplayerServiceServer).GetOnlineStatus(ctx, req.(*GetOnlineStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RivalsMultiplayerService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RivalsMultiplayerServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RivalsMultiplayerService_Notify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RivalsMultiplayerServiceServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RivalsMultiplayerService_GetTownOnlineUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTownOnlineUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RivalsMultiplayerServiceServer).GetTownOnlineUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RivalsMultiplayerService_GetTownOnlineUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RivalsMultiplayerServiceServer).GetTownOnlineUsers(ctx, req.(*GetTownOnlineUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RivalsMultiplayerService_ServiceDesc is the grpc.ServiceDesc for RivalsMultiplayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RivalsMultiplayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.multiplayer.rivals.v1.RivalsMultiplayerService",
	HandlerType: (*RivalsMultiplayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOnlineStatus",
			Handler:    _RivalsMultiplayerService_GetOnlineStatus_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _RivalsMultiplayerService_Notify_Handler,
		},
		{
			MethodName: "GetTownOnlineUsers",
			Handler:    _RivalsMultiplayerService_GetTownOnlineUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/multiplayer/rivals/v1/service_multiplayer_rivals.proto",
}
