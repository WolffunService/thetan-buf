// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/arena/v1/service.proto

package thetan_arena_v1

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
	ThetanArenaService_DisableEarningTHCHero_FullMethodName = "/thetan.arena.v1.ThetanArenaService/DisableEarningTHCHero"
)

// ThetanArenaServiceClient is the client API for ThetanArenaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThetanArenaServiceClient interface {
	DisableEarningTHCHero(ctx context.Context, in *DisableEarningTHCHeroRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type thetanArenaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThetanArenaServiceClient(cc grpc.ClientConnInterface) ThetanArenaServiceClient {
	return &thetanArenaServiceClient{cc}
}

func (c *thetanArenaServiceClient) DisableEarningTHCHero(ctx context.Context, in *DisableEarningTHCHeroRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ThetanArenaService_DisableEarningTHCHero_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetanArenaServiceServer is the server API for ThetanArenaService service.
// All implementations must embed UnimplementedThetanArenaServiceServer
// for forward compatibility
type ThetanArenaServiceServer interface {
	DisableEarningTHCHero(context.Context, *DisableEarningTHCHeroRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedThetanArenaServiceServer()
}

// UnimplementedThetanArenaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThetanArenaServiceServer struct {
}

func (UnimplementedThetanArenaServiceServer) DisableEarningTHCHero(context.Context, *DisableEarningTHCHeroRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableEarningTHCHero not implemented")
}
func (UnimplementedThetanArenaServiceServer) mustEmbedUnimplementedThetanArenaServiceServer() {}

// UnsafeThetanArenaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThetanArenaServiceServer will
// result in compilation errors.
type UnsafeThetanArenaServiceServer interface {
	mustEmbedUnimplementedThetanArenaServiceServer()
}

func RegisterThetanArenaServiceServer(s grpc.ServiceRegistrar, srv ThetanArenaServiceServer) {
	s.RegisterService(&ThetanArenaService_ServiceDesc, srv)
}

func _ThetanArenaService_DisableEarningTHCHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableEarningTHCHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanArenaServiceServer).DisableEarningTHCHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanArenaService_DisableEarningTHCHero_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanArenaServiceServer).DisableEarningTHCHero(ctx, req.(*DisableEarningTHCHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThetanArenaService_ServiceDesc is the grpc.ServiceDesc for ThetanArenaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThetanArenaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.arena.v1.ThetanArenaService",
	HandlerType: (*ThetanArenaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DisableEarningTHCHero",
			Handler:    _ThetanArenaService_DisableEarningTHCHero_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/arena/v1/service.proto",
}
