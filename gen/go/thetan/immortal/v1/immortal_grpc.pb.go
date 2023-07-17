// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/immortal/v1/immortal.proto

package thetan_immortal_v1

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
	ImmortalService_SearchPlayerInfo_FullMethodName = "/thetan.immortal.v1.ImmortalService/SearchPlayerInfo"
	ImmortalService_GetUserProfile_FullMethodName   = "/thetan.immortal.v1.ImmortalService/GetUserProfile"
)

// ImmortalServiceClient is the client API for ImmortalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImmortalServiceClient interface {
	SearchPlayerInfo(ctx context.Context, in *SearchPlayerInfoRequest, opts ...grpc.CallOption) (*SearchPlayerInfoResponse, error)
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error)
}

type immortalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImmortalServiceClient(cc grpc.ClientConnInterface) ImmortalServiceClient {
	return &immortalServiceClient{cc}
}

func (c *immortalServiceClient) SearchPlayerInfo(ctx context.Context, in *SearchPlayerInfoRequest, opts ...grpc.CallOption) (*SearchPlayerInfoResponse, error) {
	out := new(SearchPlayerInfoResponse)
	err := c.cc.Invoke(ctx, ImmortalService_SearchPlayerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error) {
	out := new(GetUserProfileResponse)
	err := c.cc.Invoke(ctx, ImmortalService_GetUserProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImmortalServiceServer is the server API for ImmortalService service.
// All implementations must embed UnimplementedImmortalServiceServer
// for forward compatibility
type ImmortalServiceServer interface {
	SearchPlayerInfo(context.Context, *SearchPlayerInfoRequest) (*SearchPlayerInfoResponse, error)
	GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error)
	mustEmbedUnimplementedImmortalServiceServer()
}

// UnimplementedImmortalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImmortalServiceServer struct {
}

func (UnimplementedImmortalServiceServer) SearchPlayerInfo(context.Context, *SearchPlayerInfoRequest) (*SearchPlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlayerInfo not implemented")
}
func (UnimplementedImmortalServiceServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
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

func _ImmortalService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalServiceServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImmortalService_ServiceDesc is the grpc.ServiceDesc for ImmortalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImmortalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.immortal.v1.ImmortalService",
	HandlerType: (*ImmortalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchPlayerInfo",
			Handler:    _ImmortalService_SearchPlayerInfo_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _ImmortalService_GetUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/immortal/v1/immortal.proto",
}
