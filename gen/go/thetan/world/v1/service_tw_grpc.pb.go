// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/world/v1/service_tw.proto

package thetan_world_v1

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
	ThetanWorldService_Test_FullMethodName = "/thetan.world.v1.ThetanWorldService/Test"
)

// ThetanWorldServiceClient is the client API for ThetanWorldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThetanWorldServiceClient interface {
	Test(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type thetanWorldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThetanWorldServiceClient(cc grpc.ClientConnInterface) ThetanWorldServiceClient {
	return &thetanWorldServiceClient{cc}
}

func (c *thetanWorldServiceClient) Test(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ThetanWorldService_Test_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetanWorldServiceServer is the server API for ThetanWorldService service.
// All implementations must embed UnimplementedThetanWorldServiceServer
// for forward compatibility
type ThetanWorldServiceServer interface {
	Test(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedThetanWorldServiceServer()
}

// UnimplementedThetanWorldServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThetanWorldServiceServer struct {
}

func (UnimplementedThetanWorldServiceServer) Test(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedThetanWorldServiceServer) mustEmbedUnimplementedThetanWorldServiceServer() {}

// UnsafeThetanWorldServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThetanWorldServiceServer will
// result in compilation errors.
type UnsafeThetanWorldServiceServer interface {
	mustEmbedUnimplementedThetanWorldServiceServer()
}

func RegisterThetanWorldServiceServer(s grpc.ServiceRegistrar, srv ThetanWorldServiceServer) {
	s.RegisterService(&ThetanWorldService_ServiceDesc, srv)
}

func _ThetanWorldService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanWorldServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanWorldService_Test_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanWorldServiceServer).Test(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ThetanWorldService_ServiceDesc is the grpc.ServiceDesc for ThetanWorldService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThetanWorldService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.world.v1.ThetanWorldService",
	HandlerType: (*ThetanWorldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _ThetanWorldService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/world/v1/service_tw.proto",
}