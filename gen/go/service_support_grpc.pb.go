// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: service_support.proto

package gen

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
	SupportService_ProfaneContains_FullMethodName = "/services.SupportService/ProfaneContains"
)

// SupportServiceClient is the client API for SupportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupportServiceClient interface {
	ProfaneContains(ctx context.Context, in *ProfaneRequest, opts ...grpc.CallOption) (*ProfaneResponse, error)
}

type supportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupportServiceClient(cc grpc.ClientConnInterface) SupportServiceClient {
	return &supportServiceClient{cc}
}

func (c *supportServiceClient) ProfaneContains(ctx context.Context, in *ProfaneRequest, opts ...grpc.CallOption) (*ProfaneResponse, error) {
	out := new(ProfaneResponse)
	err := c.cc.Invoke(ctx, SupportService_ProfaneContains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupportServiceServer is the server API for SupportService service.
// All implementations must embed UnimplementedSupportServiceServer
// for forward compatibility
type SupportServiceServer interface {
	ProfaneContains(context.Context, *ProfaneRequest) (*ProfaneResponse, error)
	mustEmbedUnimplementedSupportServiceServer()
}

// UnimplementedSupportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupportServiceServer struct {
}

func (UnimplementedSupportServiceServer) ProfaneContains(context.Context, *ProfaneRequest) (*ProfaneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfaneContains not implemented")
}
func (UnimplementedSupportServiceServer) mustEmbedUnimplementedSupportServiceServer() {}

// UnsafeSupportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupportServiceServer will
// result in compilation errors.
type UnsafeSupportServiceServer interface {
	mustEmbedUnimplementedSupportServiceServer()
}

func RegisterSupportServiceServer(s grpc.ServiceRegistrar, srv SupportServiceServer) {
	s.RegisterService(&SupportService_ServiceDesc, srv)
}

func _SupportService_ProfaneContains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfaneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportServiceServer).ProfaneContains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupportService_ProfaneContains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportServiceServer).ProfaneContains(ctx, req.(*ProfaneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupportService_ServiceDesc is the grpc.ServiceDesc for SupportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.SupportService",
	HandlerType: (*SupportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfaneContains",
			Handler:    _SupportService_ProfaneContains_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_support.proto",
}
