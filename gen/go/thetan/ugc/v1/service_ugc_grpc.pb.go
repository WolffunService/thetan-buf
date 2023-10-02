// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/ugc/v1/service_ugc.proto

package thetan_ugc_v1

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
	ThetanUGCService_GetUserColorings_FullMethodName   = "/thetan.ugc.v1.ThetanUGCService/GetUserColorings"
	ThetanUGCService_GetOneUserColoring_FullMethodName = "/thetan.ugc.v1.ThetanUGCService/GetOneUserColoring"
)

// ThetanUGCServiceClient is the client API for ThetanUGCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThetanUGCServiceClient interface {
	// Colorings
	GetUserColorings(ctx context.Context, in *UserColoringRequest, opts ...grpc.CallOption) (*UserColoringResponse, error)
	GetOneUserColoring(ctx context.Context, in *UserOneColoringRequest, opts ...grpc.CallOption) (*UserColoring, error)
}

type thetanUGCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewThetanUGCServiceClient(cc grpc.ClientConnInterface) ThetanUGCServiceClient {
	return &thetanUGCServiceClient{cc}
}

func (c *thetanUGCServiceClient) GetUserColorings(ctx context.Context, in *UserColoringRequest, opts ...grpc.CallOption) (*UserColoringResponse, error) {
	out := new(UserColoringResponse)
	err := c.cc.Invoke(ctx, ThetanUGCService_GetUserColorings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thetanUGCServiceClient) GetOneUserColoring(ctx context.Context, in *UserOneColoringRequest, opts ...grpc.CallOption) (*UserColoring, error) {
	out := new(UserColoring)
	err := c.cc.Invoke(ctx, ThetanUGCService_GetOneUserColoring_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThetanUGCServiceServer is the server API for ThetanUGCService service.
// All implementations must embed UnimplementedThetanUGCServiceServer
// for forward compatibility
type ThetanUGCServiceServer interface {
	// Colorings
	GetUserColorings(context.Context, *UserColoringRequest) (*UserColoringResponse, error)
	GetOneUserColoring(context.Context, *UserOneColoringRequest) (*UserColoring, error)
	mustEmbedUnimplementedThetanUGCServiceServer()
}

// UnimplementedThetanUGCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedThetanUGCServiceServer struct {
}

func (UnimplementedThetanUGCServiceServer) GetUserColorings(context.Context, *UserColoringRequest) (*UserColoringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserColorings not implemented")
}
func (UnimplementedThetanUGCServiceServer) GetOneUserColoring(context.Context, *UserOneColoringRequest) (*UserColoring, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneUserColoring not implemented")
}
func (UnimplementedThetanUGCServiceServer) mustEmbedUnimplementedThetanUGCServiceServer() {}

// UnsafeThetanUGCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThetanUGCServiceServer will
// result in compilation errors.
type UnsafeThetanUGCServiceServer interface {
	mustEmbedUnimplementedThetanUGCServiceServer()
}

func RegisterThetanUGCServiceServer(s grpc.ServiceRegistrar, srv ThetanUGCServiceServer) {
	s.RegisterService(&ThetanUGCService_ServiceDesc, srv)
}

func _ThetanUGCService_GetUserColorings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserColoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanUGCServiceServer).GetUserColorings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanUGCService_GetUserColorings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanUGCServiceServer).GetUserColorings(ctx, req.(*UserColoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ThetanUGCService_GetOneUserColoring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserOneColoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThetanUGCServiceServer).GetOneUserColoring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ThetanUGCService_GetOneUserColoring_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThetanUGCServiceServer).GetOneUserColoring(ctx, req.(*UserOneColoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ThetanUGCService_ServiceDesc is the grpc.ServiceDesc for ThetanUGCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ThetanUGCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.ugc.v1.ThetanUGCService",
	HandlerType: (*ThetanUGCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserColorings",
			Handler:    _ThetanUGCService_GetUserColorings_Handler,
		},
		{
			MethodName: "GetOneUserColoring",
			Handler:    _ThetanUGCService_GetOneUserColoring_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/ugc/v1/service_ugc.proto",
}
