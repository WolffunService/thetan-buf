// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: thetan/immortal/v1/immortal_player.proto

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
	ImmortalPlayerService_SearchPlayers_FullMethodName = "/thetan.immortal.v1.ImmortalPlayerService/SearchPlayers"
	ImmortalPlayerService_GetBotNames_FullMethodName   = "/thetan.immortal.v1.ImmortalPlayerService/GetBotNames"
)

// ImmortalPlayerServiceClient is the client API for ImmortalPlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImmortalPlayerServiceClient interface {
	SearchPlayers(ctx context.Context, in *SearchPlayersRequest, opts ...grpc.CallOption) (*SearchPlayersResponse, error)
	GetBotNames(ctx context.Context, in *GetBotNamesRequest, opts ...grpc.CallOption) (*GetBotNamesResponse, error)
}

type immortalPlayerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImmortalPlayerServiceClient(cc grpc.ClientConnInterface) ImmortalPlayerServiceClient {
	return &immortalPlayerServiceClient{cc}
}

func (c *immortalPlayerServiceClient) SearchPlayers(ctx context.Context, in *SearchPlayersRequest, opts ...grpc.CallOption) (*SearchPlayersResponse, error) {
	out := new(SearchPlayersResponse)
	err := c.cc.Invoke(ctx, ImmortalPlayerService_SearchPlayers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immortalPlayerServiceClient) GetBotNames(ctx context.Context, in *GetBotNamesRequest, opts ...grpc.CallOption) (*GetBotNamesResponse, error) {
	out := new(GetBotNamesResponse)
	err := c.cc.Invoke(ctx, ImmortalPlayerService_GetBotNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImmortalPlayerServiceServer is the server API for ImmortalPlayerService service.
// All implementations must embed UnimplementedImmortalPlayerServiceServer
// for forward compatibility
type ImmortalPlayerServiceServer interface {
	SearchPlayers(context.Context, *SearchPlayersRequest) (*SearchPlayersResponse, error)
	GetBotNames(context.Context, *GetBotNamesRequest) (*GetBotNamesResponse, error)
	mustEmbedUnimplementedImmortalPlayerServiceServer()
}

// UnimplementedImmortalPlayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImmortalPlayerServiceServer struct {
}

func (UnimplementedImmortalPlayerServiceServer) SearchPlayers(context.Context, *SearchPlayersRequest) (*SearchPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlayers not implemented")
}
func (UnimplementedImmortalPlayerServiceServer) GetBotNames(context.Context, *GetBotNamesRequest) (*GetBotNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBotNames not implemented")
}
func (UnimplementedImmortalPlayerServiceServer) mustEmbedUnimplementedImmortalPlayerServiceServer() {}

// UnsafeImmortalPlayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImmortalPlayerServiceServer will
// result in compilation errors.
type UnsafeImmortalPlayerServiceServer interface {
	mustEmbedUnimplementedImmortalPlayerServiceServer()
}

func RegisterImmortalPlayerServiceServer(s grpc.ServiceRegistrar, srv ImmortalPlayerServiceServer) {
	s.RegisterService(&ImmortalPlayerService_ServiceDesc, srv)
}

func _ImmortalPlayerService_SearchPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalPlayerServiceServer).SearchPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalPlayerService_SearchPlayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalPlayerServiceServer).SearchPlayers(ctx, req.(*SearchPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmortalPlayerService_GetBotNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmortalPlayerServiceServer).GetBotNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImmortalPlayerService_GetBotNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmortalPlayerServiceServer).GetBotNames(ctx, req.(*GetBotNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImmortalPlayerService_ServiceDesc is the grpc.ServiceDesc for ImmortalPlayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImmortalPlayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thetan.immortal.v1.ImmortalPlayerService",
	HandlerType: (*ImmortalPlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchPlayers",
			Handler:    _ImmortalPlayerService_SearchPlayers_Handler,
		},
		{
			MethodName: "GetBotNames",
			Handler:    _ImmortalPlayerService_GetBotNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thetan/immortal/v1/immortal_player.proto",
}
