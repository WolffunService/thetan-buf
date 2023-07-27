// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/gateway/v1/gs_sync.proto

package thetan_gateway_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ThetanGatewayTracking_RoomAllocated_FullMethodName = "/thetan.gateway.v1.ThetanGatewayTracking/RoomAllocated"
	ThetanGatewayTracking_RoomRelease_FullMethodName   = "/thetan.gateway.v1.ThetanGatewayTracking/RoomRelease"
)

// ThetanGatewayTrackingClient is the client API for ThetanGatewayTracking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThetanGatewayTrackingClient interface {
	RoomAllocated(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RoomRelease(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type thetanGatewayTrackingClient struct {
	cc grpc.ClientConnInterface
}

func NewThetanGatewayTrackingClient(cc grpc.ClientConnInterface) ThetanGatewayTrackingClient {
	return &thetanGatewayTrackingClient{cc}
}

func (c *thetanGatewayTrackingClient) RoomAllocated(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ThetanGatewayTracking_RoomAllocated_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanGatewayTrackingClient) RoomRelease(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ThetanGatewayTracking_RoomRelease_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetanGatewayTrackingServer is the server API for ThetanGatewayTracking service.
// All implementations must embed UnimplementedThetanGatewayTrackingServer
// for forward compatibility
type ThetanGatewayTrackingServer interface {
	RoomAllocated(context.Context, *RoomRequest) (*emptypb.Empty, error)
	RoomRelease(context.Context, *RoomRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedThetanGatewayTrackingServer()
}

// UnimplementedThetanGatewayTrackingServer must be embedded to have forward compatible implementations.
type UnimplementedThetanGatewayTrackingServer struct {
}

func (UnimplementedThetanGatewayTrackingServer) RoomAllocated(context.Context, *RoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomAllocated not implemented")
}
func (UnimplementedThetanGatewayTrackingServer) RoomRelease(context.Context, *RoomRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoomRelease not implemented")
}
func (UnimplementedThetanGatewayTrackingServer) mustEmbedUnimplementedThetanGatewayTrackingServer() {}

// UnsafeThetanGatewayTrackingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThetanGatewayTrackingServer will
// result in compilation errors.
type UnsafeThetanGatewayTrackingServer interface {
	mustEmbedUnimplementedThetanGatewayTrackingServer()
}

func RegisterThetanGatewayTrackingServer(s grpc.ServiceRegistrar, srv ThetanGatewayTrackingServer) {
	s.RegisterService(&ThetanGatewayTracking_ServiceDesc, srv)
}

func _ThetanGatewayTracking_RoomAllocated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanGatewayTrackingServer).RoomAllocated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanGatewayTracking_RoomAllocated_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanGatewayTrackingServer).RoomAllocated(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanGatewayTracking_RoomRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanGatewayTrackingServer).RoomRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanGatewayTracking_RoomRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanGatewayTrackingServer).RoomRelease(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThetanGatewayTracking_ServiceDesc is the grpc.ServiceDesc for ThetanGatewayTracking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThetanGatewayTracking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.gateway.v1.ThetanGatewayTracking",
	HandlerType: (*ThetanGatewayTrackingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoomAllocated",
			Handler:    _ThetanGatewayTracking_RoomAllocated_Handler,
		},
		{
			MethodName: "RoomRelease",
			Handler:    _ThetanGatewayTracking_RoomRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/gateway/v1/gs_sync.proto",
}

const (
	ThetanGatewayImmortal_AllocateGameServer_FullMethodName = "/thetan.gateway.v1.ThetanGatewayImmortal/AllocateGameServer"
)

// ThetanGatewayImmortalClient is the client API for ThetanGatewayImmortal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThetanGatewayImmortalClient interface {
	AllocateGameServer(ctx context.Context, in *ImmortalAllocateRequest, opts ...grpc.CallOption) (*ImmortalRoomAllocationResp, error)
}

type thetanGatewayImmortalClient struct {
	cc grpc.ClientConnInterface
}

func NewThetanGatewayImmortalClient(cc grpc.ClientConnInterface) ThetanGatewayImmortalClient {
	return &thetanGatewayImmortalClient{cc}
}

func (c *thetanGatewayImmortalClient) AllocateGameServer(ctx context.Context, in *ImmortalAllocateRequest, opts ...grpc.CallOption) (*ImmortalRoomAllocationResp, error) {
	out := new(ImmortalRoomAllocationResp)
	err := c.cc.Invoke(ctx, ThetanGatewayImmortal_AllocateGameServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetanGatewayImmortalServer is the server API for ThetanGatewayImmortal service.
// All implementations must embed UnimplementedThetanGatewayImmortalServer
// for forward compatibility
type ThetanGatewayImmortalServer interface {
	AllocateGameServer(context.Context, *ImmortalAllocateRequest) (*ImmortalRoomAllocationResp, error)
	mustEmbedUnimplementedThetanGatewayImmortalServer()
}

// UnimplementedThetanGatewayImmortalServer must be embedded to have forward compatible implementations.
type UnimplementedThetanGatewayImmortalServer struct {
}

func (UnimplementedThetanGatewayImmortalServer) AllocateGameServer(context.Context, *ImmortalAllocateRequest) (*ImmortalRoomAllocationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateGameServer not implemented")
}
func (UnimplementedThetanGatewayImmortalServer) mustEmbedUnimplementedThetanGatewayImmortalServer() {}

// UnsafeThetanGatewayImmortalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThetanGatewayImmortalServer will
// result in compilation errors.
type UnsafeThetanGatewayImmortalServer interface {
	mustEmbedUnimplementedThetanGatewayImmortalServer()
}

func RegisterThetanGatewayImmortalServer(s grpc.ServiceRegistrar, srv ThetanGatewayImmortalServer) {
	s.RegisterService(&ThetanGatewayImmortal_ServiceDesc, srv)
}

func _ThetanGatewayImmortal_AllocateGameServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImmortalAllocateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanGatewayImmortalServer).AllocateGameServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanGatewayImmortal_AllocateGameServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanGatewayImmortalServer).AllocateGameServer(ctx, req.(*ImmortalAllocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThetanGatewayImmortal_ServiceDesc is the grpc.ServiceDesc for ThetanGatewayImmortal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThetanGatewayImmortal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.gateway.v1.ThetanGatewayImmortal",
	HandlerType: (*ThetanGatewayImmortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocateGameServer",
			Handler:    _ThetanGatewayImmortal_AllocateGameServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/gateway/v1/gs_sync.proto",
}
