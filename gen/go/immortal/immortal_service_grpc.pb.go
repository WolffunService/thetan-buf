// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: immortal_service.proto

package immortal

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
	ImmortalService_SearchUserRanking_FullMethodName     = "/immortal.ImmortalService/SearchUserRanking"
	ImmortalService_CreateManyUserRanking_FullMethodName = "/immortal.ImmortalService/CreateManyUserRanking"
	ImmortalService_SearchPlayerInfo_FullMethodName      = "/immortal.ImmortalService/SearchPlayerInfo"
)

// ImmortalServiceClient is the client API for ImmortalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImmortalServiceClient interface {
	// For bot
	SearchUserRanking(ctx context.Context, in *SearchUserRankingRequest, opts ...grpc.CallOption) (*SearchUserRankingResponse, error)
	CreateManyUserRanking(ctx context.Context, in *CreateManyUserRankingRequest, opts ...grpc.CallOption) (*CreateManyUserRankingResponse, error)
	// Search player info
	SearchPlayerInfo(ctx context.Context, in *SearchPlayerInfoRequest, opts ...grpc.CallOption) (*SearchPlayerInfoResponse, error)
}

type immortalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImmortalServiceClient(cc grpc.ClientConnInterface) ImmortalServiceClient {
	return &immortalServiceClient{cc}
}

func (c *immortalServiceClient) SearchUserRanking(ctx context.Context, in *SearchUserRankingRequest, opts ...grpc.CallOption) (*SearchUserRankingResponse, error) {
	out := new(SearchUserRankingResponse)
	err := c.cc.Invoke(ctx, ImmortalService_SearchUserRanking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) CreateManyUserRanking(ctx context.Context, in *CreateManyUserRankingRequest, opts ...grpc.CallOption) (*CreateManyUserRankingResponse, error) {
	out := new(CreateManyUserRankingResponse)
	err := c.cc.Invoke(ctx, ImmortalService_CreateManyUserRanking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) SearchPlayerInfo(ctx context.Context, in *SearchPlayerInfoRequest, opts ...grpc.CallOption) (*SearchPlayerInfoResponse, error) {
	out := new(SearchPlayerInfoResponse)
	err := c.cc.Invoke(ctx, ImmortalService_SearchPlayerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImmortalServiceServer is the server API for ImmortalService service.
// All implementations must embed UnimplementedImmortalServiceServer
// for forward compatibility
type ImmortalServiceServer interface {
	// For bot
	SearchUserRanking(context.Context, *SearchUserRankingRequest) (*SearchUserRankingResponse, error)
	CreateManyUserRanking(context.Context, *CreateManyUserRankingRequest) (*CreateManyUserRankingResponse, error)
	// Search player info
	SearchPlayerInfo(context.Context, *SearchPlayerInfoRequest) (*SearchPlayerInfoResponse, error)
	mustEmbedUnimplementedImmortalServiceServer()
}

// UnimplementedImmortalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImmortalServiceServer struct {
}

func (UnimplementedImmortalServiceServer) SearchUserRanking(context.Context, *SearchUserRankingRequest) (*SearchUserRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserRanking not implemented")
}
func (UnimplementedImmortalServiceServer) CreateManyUserRanking(context.Context, *CreateManyUserRankingRequest) (*CreateManyUserRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManyUserRanking not implemented")
}
func (UnimplementedImmortalServiceServer) SearchPlayerInfo(context.Context, *SearchPlayerInfoRequest) (*SearchPlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlayerInfo not implemented")
}
func (UnimplementedImmortalServiceServer) mustEmbedUnimplementedImmortalServiceServer() {}

// UnsafeImmortalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImmortalServiceServer will
// result in compilation errors.
type UnsafeImmortalServiceServer interface {
	mustEmbedUnimplementedImmortalServiceServer()
}

func RegisterImmortalServiceServer(s grpc.ServiceRegistrar, srv ImmortalServiceServer) {
	s.RegisterService(&ImmortalService_ServiceDesc, srv)
}

func _ImmortalService_SearchUserRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).SearchUserRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_SearchUserRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).SearchUserRanking(ctx, req.(*SearchUserRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_CreateManyUserRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateManyUserRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).CreateManyUserRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_CreateManyUserRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).CreateManyUserRanking(ctx, req.(*CreateManyUserRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalService_SearchPlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).SearchPlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_SearchPlayerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).SearchPlayerInfo(ctx, req.(*SearchPlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImmortalService_ServiceDesc is the grpc.ServiceDesc for ImmortalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImmortalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "immortal.ImmortalService",
	HandlerType: (*ImmortalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUserRanking",
			Handler:    _ImmortalService_SearchUserRanking_Handler,
		},
		{
			MethodName: "CreateManyUserRanking",
			Handler:    _ImmortalService_CreateManyUserRanking_Handler,
		},
		{
			MethodName: "SearchPlayerInfo",
			Handler:    _ImmortalService_SearchPlayerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "immortal_service.proto",
}
