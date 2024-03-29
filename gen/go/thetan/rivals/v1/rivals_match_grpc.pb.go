// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/rivals/v1/rivals_match.proto

package thetan_rivals_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "thetan-buf/gen/go/thetan/shared/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RivalMatchDirectorService_CancelTicket_FullMethodName           = "/thetan.rivals.v1.RivalMatchDirectorService/CancelTicket"
	RivalMatchDirectorService_CreateMatchOnboard_FullMethodName     = "/thetan.rivals.v1.RivalMatchDirectorService/CreateMatchOnboard"
	RivalMatchDirectorService_CreateMatchNonMatching_FullMethodName = "/thetan.rivals.v1.RivalMatchDirectorService/CreateMatchNonMatching"
	RivalMatchDirectorService_CreateMatchTutorial_FullMethodName    = "/thetan.rivals.v1.RivalMatchDirectorService/CreateMatchTutorial"
)

// RivalMatchDirectorServiceClient is the client API for RivalMatchDirectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RivalMatchDirectorServiceClient interface {
	CancelTicket(ctx context.Context, in *RivalCancelTicketRequest, opts ...grpc.CallOption) (*RivalCancelTicketResponse, error)
	// Deprecated: Do not use.
	CreateMatchOnboard(ctx context.Context, in *GetMatchInfoRequest, opts ...grpc.CallOption) (*v1.MatchFoundResponseProto, error)
	CreateMatchNonMatching(ctx context.Context, in *CreateMatchNonMatchingRequest, opts ...grpc.CallOption) (*CreateMatchNonMatchingResponse, error)
	CreateMatchTutorial(ctx context.Context, in *CreateMatchTutorialRequest, opts ...grpc.CallOption) (*v1.MatchFoundResponseProto, error)
}

type rivalMatchDirectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRivalMatchDirectorServiceClient(cc grpc.ClientConnInterface) RivalMatchDirectorServiceClient {
	return &rivalMatchDirectorServiceClient{cc}
}

func (c *rivalMatchDirectorServiceClient) CancelTicket(ctx context.Context, in *RivalCancelTicketRequest, opts ...grpc.CallOption) (*RivalCancelTicketResponse, error) {
	out := new(RivalCancelTicketResponse)
	err := c.cc.Invoke(ctx, RivalMatchDirectorService_CancelTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *rivalMatchDirectorServiceClient) CreateMatchOnboard(ctx context.Context, in *GetMatchInfoRequest, opts ...grpc.CallOption) (*v1.MatchFoundResponseProto, error) {
	out := new(v1.MatchFoundResponseProto)
	err := c.cc.Invoke(ctx, RivalMatchDirectorService_CreateMatchOnboard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rivalMatchDirectorServiceClient) CreateMatchNonMatching(ctx context.Context, in *CreateMatchNonMatchingRequest, opts ...grpc.CallOption) (*CreateMatchNonMatchingResponse, error) {
	out := new(CreateMatchNonMatchingResponse)
	err := c.cc.Invoke(ctx, RivalMatchDirectorService_CreateMatchNonMatching_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rivalMatchDirectorServiceClient) CreateMatchTutorial(ctx context.Context, in *CreateMatchTutorialRequest, opts ...grpc.CallOption) (*v1.MatchFoundResponseProto, error) {
	out := new(v1.MatchFoundResponseProto)
	err := c.cc.Invoke(ctx, RivalMatchDirectorService_CreateMatchTutorial_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RivalMatchDirectorServiceServer is the server API for RivalMatchDirectorService service.
// All implementations must embed UnimplementedRivalMatchDirectorServiceServer
// for forward compatibility
type RivalMatchDirectorServiceServer interface {
	CancelTicket(context.Context, *RivalCancelTicketRequest) (*RivalCancelTicketResponse, error)
	// Deprecated: Do not use.
	CreateMatchOnboard(context.Context, *GetMatchInfoRequest) (*v1.MatchFoundResponseProto, error)
	CreateMatchNonMatching(context.Context, *CreateMatchNonMatchingRequest) (*CreateMatchNonMatchingResponse, error)
	CreateMatchTutorial(context.Context, *CreateMatchTutorialRequest) (*v1.MatchFoundResponseProto, error)
	mustEmbedUnimplementedRivalMatchDirectorServiceServer()
}

// UnimplementedRivalMatchDirectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRivalMatchDirectorServiceServer struct {
}

func (UnimplementedRivalMatchDirectorServiceServer) CancelTicket(context.Context, *RivalCancelTicketRequest) (*RivalCancelTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTicket not implemented")
}
func (UnimplementedRivalMatchDirectorServiceServer) CreateMatchOnboard(context.Context, *GetMatchInfoRequest) (*v1.MatchFoundResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatchOnboard not implemented")
}
func (UnimplementedRivalMatchDirectorServiceServer) CreateMatchNonMatching(context.Context, *CreateMatchNonMatchingRequest) (*CreateMatchNonMatchingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatchNonMatching not implemented")
}
func (UnimplementedRivalMatchDirectorServiceServer) CreateMatchTutorial(context.Context, *CreateMatchTutorialRequest) (*v1.MatchFoundResponseProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatchTutorial not implemented")
}
func (UnimplementedRivalMatchDirectorServiceServer) mustEmbedUnimplementedRivalMatchDirectorServiceServer() {
}

// UnsafeRivalMatchDirectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RivalMatchDirectorServiceServer will
// result in compilation errors.
type UnsafeRivalMatchDirectorServiceServer interface {
	mustEmbedUnimplementedRivalMatchDirectorServiceServer()
}

func RegisterRivalMatchDirectorServiceServer(s grpc.ServiceRegistrar, srv RivalMatchDirectorServiceServer) {
	s.RegisterService(&RivalMatchDirectorService_ServiceDesc, srv)
}

func _RivalMatchDirectorService_CancelTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RivalCancelTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RivalMatchDirectorServiceServer).CancelTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RivalMatchDirectorService_CancelTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RivalMatchDirectorServiceServer).CancelTicket(ctx, req.(*RivalCancelTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RivalMatchDirectorService_CreateMatchOnboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RivalMatchDirectorServiceServer).CreateMatchOnboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RivalMatchDirectorService_CreateMatchOnboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RivalMatchDirectorServiceServer).CreateMatchOnboard(ctx, req.(*GetMatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RivalMatchDirectorService_CreateMatchNonMatching_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMatchNonMatchingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RivalMatchDirectorServiceServer).CreateMatchNonMatching(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RivalMatchDirectorService_CreateMatchNonMatching_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RivalMatchDirectorServiceServer).CreateMatchNonMatching(ctx, req.(*CreateMatchNonMatchingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RivalMatchDirectorService_CreateMatchTutorial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMatchTutorialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RivalMatchDirectorServiceServer).CreateMatchTutorial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RivalMatchDirectorService_CreateMatchTutorial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RivalMatchDirectorServiceServer).CreateMatchTutorial(ctx, req.(*CreateMatchTutorialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RivalMatchDirectorService_ServiceDesc is the grpc.ServiceDesc for RivalMatchDirectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RivalMatchDirectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.rivals.v1.RivalMatchDirectorService",
	HandlerType: (*RivalMatchDirectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelTicket",
			Handler:    _RivalMatchDirectorService_CancelTicket_Handler,
		},
		{
			MethodName: "CreateMatchOnboard",
			Handler:    _RivalMatchDirectorService_CreateMatchOnboard_Handler,
		},
		{
			MethodName: "CreateMatchNonMatching",
			Handler:    _RivalMatchDirectorService_CreateMatchNonMatching_Handler,
		},
		{
			MethodName: "CreateMatchTutorial",
			Handler:    _RivalMatchDirectorService_CreateMatchTutorial_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/rivals/v1/rivals_match.proto",
}
