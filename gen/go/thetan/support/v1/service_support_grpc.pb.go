// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/support/v1/service_support.proto

package thetan_support_v1

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
	SupportService_SearchBots_FullMethodName = "/thetan.support.v1.SupportService/SearchBots"
)

// SupportServiceClient is the client API for SupportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupportServiceClient interface {
	SearchBots(ctx context.Context, in *SearchBotsRequest, opts ...grpc.CallOption) (*SearchBotsResponse, error)
}

type supportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupportServiceClient(cc grpc.ClientConnInterface) SupportServiceClient {
	return &supportServiceClient{cc}
}

func (c *supportServiceClient) SearchBots(ctx context.Context, in *SearchBotsRequest, opts ...grpc.CallOption) (*SearchBotsResponse, error) {
	out := new(SearchBotsResponse)
	err := c.cc.Invoke(ctx, SupportService_SearchBots_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupportServiceServer is the server API for SupportService service.
// All implementations must embed UnimplementedSupportServiceServer
// for forward compatibility
type SupportServiceServer interface {
	SearchBots(context.Context, *SearchBotsRequest) (*SearchBotsResponse, error)
	mustEmbedUnimplementedSupportServiceServer()
}

// UnimplementedSupportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupportServiceServer struct {
}

func (UnimplementedSupportServiceServer) SearchBots(context.Context, *SearchBotsRequest) (*SearchBotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBots not implemented")
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

func _SupportService_SearchBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupportServiceServer).SearchBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupportService_SearchBots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupportServiceServer).SearchBots(ctx, req.(*SearchBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupportService_ServiceDesc is the grpc.ServiceDesc for SupportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.support.v1.SupportService",
	HandlerType: (*SupportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchBots",
			Handler:    _SupportService_SearchBots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/support/v1/service_support.proto",
}
