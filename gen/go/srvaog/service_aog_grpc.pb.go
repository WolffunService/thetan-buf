// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: service_aog.proto

package srvaog

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
	AogService_SearchUserRanking_FullMethodName = "/services.AogService/SearchUserRanking"
	AogService_CreateUserRanking_FullMethodName = "/services.AogService/CreateUserRanking"
	AogService_GetFindMatchInfo_FullMethodName  = "/services.AogService/GetFindMatchInfo"
)

// AogServiceClient is the client API for AogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AogServiceClient interface {
	// For bot
	SearchUserRanking(ctx context.Context, in *SearchUserRankingRequest, opts ...grpc.CallOption) (*SearchUserRankingResponse, error)
	CreateUserRanking(ctx context.Context, in *CreateUserRankingRequest, opts ...grpc.CallOption) (*CreateUserRankingResponse, error)
	// Get find match info
	GetFindMatchInfo(ctx context.Context, in *FindMatchInfoRequest, opts ...grpc.CallOption) (*FindMatchInfoResponse, error)
}

type aogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAogServiceClient(cc grpc.ClientConnInterface) AogServiceClient {
	return &aogServiceClient{cc}
}

func (c *aogServiceClient) SearchUserRanking(ctx context.Context, in *SearchUserRankingRequest, opts ...grpc.CallOption) (*SearchUserRankingResponse, error) {
	out := new(SearchUserRankingResponse)
	err := c.cc.Invoke(ctx, AogService_SearchUserRanking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aogServiceClient) CreateUserRanking(ctx context.Context, in *CreateUserRankingRequest, opts ...grpc.CallOption) (*CreateUserRankingResponse, error) {
	out := new(CreateUserRankingResponse)
	err := c.cc.Invoke(ctx, AogService_CreateUserRanking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aogServiceClient) GetFindMatchInfo(ctx context.Context, in *FindMatchInfoRequest, opts ...grpc.CallOption) (*FindMatchInfoResponse, error) {
	out := new(FindMatchInfoResponse)
	err := c.cc.Invoke(ctx, AogService_GetFindMatchInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AogServiceServer is the server API for AogService service.
// All implementations must embed UnimplementedAogServiceServer
// for forward compatibility
type AogServiceServer interface {
	// For bot
	SearchUserRanking(context.Context, *SearchUserRankingRequest) (*SearchUserRankingResponse, error)
	CreateUserRanking(context.Context, *CreateUserRankingRequest) (*CreateUserRankingResponse, error)
	// Get find match info
	GetFindMatchInfo(context.Context, *FindMatchInfoRequest) (*FindMatchInfoResponse, error)
	mustEmbedUnimplementedAogServiceServer()
}

// UnimplementedAogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAogServiceServer struct {
}

func (UnimplementedAogServiceServer) SearchUserRanking(context.Context, *SearchUserRankingRequest) (*SearchUserRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserRanking not implemented")
}
func (UnimplementedAogServiceServer) CreateUserRanking(context.Context, *CreateUserRankingRequest) (*CreateUserRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserRanking not implemented")
}
func (UnimplementedAogServiceServer) GetFindMatchInfo(context.Context, *FindMatchInfoRequest) (*FindMatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFindMatchInfo not implemented")
}
func (UnimplementedAogServiceServer) mustEmbedUnimplementedAogServiceServer() {}

// UnsafeAogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AogServiceServer will
// result in compilation errors.
type UnsafeAogServiceServer interface {
	mustEmbedUnimplementedAogServiceServer()
}

func RegisterAogServiceServer(s grpc.ServiceRegistrar, srv AogServiceServer) {
	s.RegisterService(&AogService_ServiceDesc, srv)
}

func _AogService_SearchUserRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AogServiceServer).SearchUserRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AogService_SearchUserRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AogServiceServer).SearchUserRanking(ctx, req.(*SearchUserRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AogService_CreateUserRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AogServiceServer).CreateUserRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AogService_CreateUserRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AogServiceServer).CreateUserRanking(ctx, req.(*CreateUserRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AogService_GetFindMatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AogServiceServer).GetFindMatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AogService_GetFindMatchInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AogServiceServer).GetFindMatchInfo(ctx, req.(*FindMatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AogService_ServiceDesc is the grpc.ServiceDesc for AogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.AogService",
	HandlerType: (*AogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUserRanking",
			Handler:    _AogService_SearchUserRanking_Handler,
		},
		{
			MethodName: "CreateUserRanking",
			Handler:    _AogService_CreateUserRanking_Handler,
		},
		{
			MethodName: "GetFindMatchInfo",
			Handler:    _AogService_GetFindMatchInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_aog.proto",
}
