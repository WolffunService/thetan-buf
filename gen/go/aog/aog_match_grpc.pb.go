// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aog_match.proto

package aog

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
	AogMatchDirectorService_CancelTicket_FullMethodName = "/aog.AogMatchDirectorService/CancelTicket"
)

// AogMatchDirectorServiceClient is the client API for AogMatchDirectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AogMatchDirectorServiceClient interface {
	CancelTicket(ctx context.Context, in *AogCancelTicketRequest, opts ...grpc.CallOption) (*AogCancelTicketResponse, error)
}

type aogMatchDirectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAogMatchDirectorServiceClient(cc grpc.ClientConnInterface) AogMatchDirectorServiceClient {
	return &aogMatchDirectorServiceClient{cc}
}

func (c *aogMatchDirectorServiceClient) CancelTicket(ctx context.Context, in *AogCancelTicketRequest, opts ...grpc.CallOption) (*AogCancelTicketResponse, error) {
	out := new(AogCancelTicketResponse)
	err := c.cc.Invoke(ctx, AogMatchDirectorService_CancelTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AogMatchDirectorServiceServer is the server API for AogMatchDirectorService service.
// All implementations must embed UnimplementedAogMatchDirectorServiceServer
// for forward compatibility
type AogMatchDirectorServiceServer interface {
	CancelTicket(context.Context, *AogCancelTicketRequest) (*AogCancelTicketResponse, error)
	mustEmbedUnimplementedAogMatchDirectorServiceServer()
}

// UnimplementedAogMatchDirectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAogMatchDirectorServiceServer struct {
}

func (UnimplementedAogMatchDirectorServiceServer) CancelTicket(context.Context, *AogCancelTicketRequest) (*AogCancelTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTicket not implemented")
}
func (UnimplementedAogMatchDirectorServiceServer) mustEmbedUnimplementedAogMatchDirectorServiceServer() {
}

// UnsafeAogMatchDirectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AogMatchDirectorServiceServer will
// result in compilation errors.
type UnsafeAogMatchDirectorServiceServer interface {
	mustEmbedUnimplementedAogMatchDirectorServiceServer()
}

func RegisterAogMatchDirectorServiceServer(s grpc.ServiceRegistrar, srv AogMatchDirectorServiceServer) {
	s.RegisterService(&AogMatchDirectorService_ServiceDesc, srv)
}

func _AogMatchDirectorService_CancelTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AogCancelTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AogMatchDirectorServiceServer).CancelTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AogMatchDirectorService_CancelTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AogMatchDirectorServiceServer).CancelTicket(ctx, req.(*AogCancelTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AogMatchDirectorService_ServiceDesc is the grpc.ServiceDesc for AogMatchDirectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AogMatchDirectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aog.AogMatchDirectorService",
	HandlerType: (*AogMatchDirectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelTicket",
			Handler:    _AogMatchDirectorService_CancelTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aog_match.proto",
}
