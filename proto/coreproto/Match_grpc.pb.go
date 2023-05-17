// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: proto/Match.proto

package coreproto

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

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchServiceClient interface {
	CreateMatchWithVersion(ctx context.Context, in *MatchProtoVersionPackage, opts ...grpc.CallOption) (MatchService_CreateMatchWithVersionClient, error)
	//register de nhan su kien match found
	RegisterMatchFound(ctx context.Context, in *MatchProtoPackage, opts ...grpc.CallOption) (MatchService_RegisterMatchFoundClient, error)
	CancelMatchMaking(ctx context.Context, in *MatchProtoPackage, opts ...grpc.CallOption) (*MatchProtoPackage, error)
}

type matchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchServiceClient(cc grpc.ClientConnInterface) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) CreateMatchWithVersion(ctx context.Context, in *MatchProtoVersionPackage, opts ...grpc.CallOption) (MatchService_CreateMatchWithVersionClient, error) {
	stream, err := c.cc.NewStream(ctx, &MatchService_ServiceDesc.Streams[0], "/core.proto.MatchService/CreateMatchWithVersion", opts...)
	if err != nil {
		return nil, err
	}
	x := &matchServiceCreateMatchWithVersionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MatchService_CreateMatchWithVersionClient interface {
	Recv() (*MatchProtoPackage, error)
	grpc.ClientStream
}

type matchServiceCreateMatchWithVersionClient struct {
	grpc.ClientStream
}

func (x *matchServiceCreateMatchWithVersionClient) Recv() (*MatchProtoPackage, error) {
	m := new(MatchProtoPackage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *matchServiceClient) RegisterMatchFound(ctx context.Context, in *MatchProtoPackage, opts ...grpc.CallOption) (MatchService_RegisterMatchFoundClient, error) {
	stream, err := c.cc.NewStream(ctx, &MatchService_ServiceDesc.Streams[1], "/core.proto.MatchService/RegisterMatchFound", opts...)
	if err != nil {
		return nil, err
	}
	x := &matchServiceRegisterMatchFoundClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MatchService_RegisterMatchFoundClient interface {
	Recv() (*MatchProtoPackage, error)
	grpc.ClientStream
}

type matchServiceRegisterMatchFoundClient struct {
	grpc.ClientStream
}

func (x *matchServiceRegisterMatchFoundClient) Recv() (*MatchProtoPackage, error) {
	m := new(MatchProtoPackage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *matchServiceClient) CancelMatchMaking(ctx context.Context, in *MatchProtoPackage, opts ...grpc.CallOption) (*MatchProtoPackage, error) {
	out := new(MatchProtoPackage)
	err := c.cc.Invoke(ctx, "/core.proto.MatchService/CancelMatchMaking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
// All implementations must embed UnimplementedMatchServiceServer
// for forward compatibility
type MatchServiceServer interface {
	CreateMatchWithVersion(*MatchProtoVersionPackage, MatchService_CreateMatchWithVersionServer) error
	//register de nhan su kien match found
	RegisterMatchFound(*MatchProtoPackage, MatchService_RegisterMatchFoundServer) error
	CancelMatchMaking(context.Context, *MatchProtoPackage) (*MatchProtoPackage, error)
	mustEmbedUnimplementedMatchServiceServer()
}

// UnimplementedMatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (UnimplementedMatchServiceServer) CreateMatchWithVersion(*MatchProtoVersionPackage, MatchService_CreateMatchWithVersionServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateMatchWithVersion not implemented")
}
func (UnimplementedMatchServiceServer) RegisterMatchFound(*MatchProtoPackage, MatchService_RegisterMatchFoundServer) error {
	return status.Errorf(codes.Unimplemented, "method RegisterMatchFound not implemented")
}
func (UnimplementedMatchServiceServer) CancelMatchMaking(context.Context, *MatchProtoPackage) (*MatchProtoPackage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelMatchMaking not implemented")
}
func (UnimplementedMatchServiceServer) mustEmbedUnimplementedMatchServiceServer() {}

// UnsafeMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchServiceServer will
// result in compilation errors.
type UnsafeMatchServiceServer interface {
	mustEmbedUnimplementedMatchServiceServer()
}

func RegisterMatchServiceServer(s grpc.ServiceRegistrar, srv MatchServiceServer) {
	s.RegisterService(&MatchService_ServiceDesc, srv)
}

func _MatchService_CreateMatchWithVersion_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchProtoVersionPackage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchServiceServer).CreateMatchWithVersion(m, &matchServiceCreateMatchWithVersionServer{stream})
}

type MatchService_CreateMatchWithVersionServer interface {
	Send(*MatchProtoPackage) error
	grpc.ServerStream
}

type matchServiceCreateMatchWithVersionServer struct {
	grpc.ServerStream
}

func (x *matchServiceCreateMatchWithVersionServer) Send(m *MatchProtoPackage) error {
	return x.ServerStream.SendMsg(m)
}

func _MatchService_RegisterMatchFound_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchProtoPackage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchServiceServer).RegisterMatchFound(m, &matchServiceRegisterMatchFoundServer{stream})
}

type MatchService_RegisterMatchFoundServer interface {
	Send(*MatchProtoPackage) error
	grpc.ServerStream
}

type matchServiceRegisterMatchFoundServer struct {
	grpc.ServerStream
}

func (x *matchServiceRegisterMatchFoundServer) Send(m *MatchProtoPackage) error {
	return x.ServerStream.SendMsg(m)
}

func _MatchService_CancelMatchMaking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchProtoPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).CancelMatchMaking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.proto.MatchService/CancelMatchMaking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).CancelMatchMaking(ctx, req.(*MatchProtoPackage))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchService_ServiceDesc is the grpc.ServiceDesc for MatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.proto.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelMatchMaking",
			Handler:    _MatchService_CancelMatchMaking_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateMatchWithVersion",
			Handler:       _MatchService_CreateMatchWithVersion_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RegisterMatchFound",
			Handler:       _MatchService_RegisterMatchFound_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/Match.proto",
}

// MatchHandleServiceClient is the client API for MatchHandleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchHandleServiceClient interface {
	HandlePlayAgain(ctx context.Context, in *DataPlayAgainSuccess, opts ...grpc.CallOption) (*EmptyResponse, error)
	HandleDeleteTicket(ctx context.Context, in *DeleteTicketSuccess, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type matchHandleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchHandleServiceClient(cc grpc.ClientConnInterface) MatchHandleServiceClient {
	return &matchHandleServiceClient{cc}
}

func (c *matchHandleServiceClient) HandlePlayAgain(ctx context.Context, in *DataPlayAgainSuccess, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/core.proto.MatchHandleService/HandlePlayAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchHandleServiceClient) HandleDeleteTicket(ctx context.Context, in *DeleteTicketSuccess, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/core.proto.MatchHandleService/HandleDeleteTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchHandleServiceServer is the server API for MatchHandleService service.
// All implementations must embed UnimplementedMatchHandleServiceServer
// for forward compatibility
type MatchHandleServiceServer interface {
	HandlePlayAgain(context.Context, *DataPlayAgainSuccess) (*EmptyResponse, error)
	HandleDeleteTicket(context.Context, *DeleteTicketSuccess) (*EmptyResponse, error)
	mustEmbedUnimplementedMatchHandleServiceServer()
}

// UnimplementedMatchHandleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchHandleServiceServer struct {
}

func (UnimplementedMatchHandleServiceServer) HandlePlayAgain(context.Context, *DataPlayAgainSuccess) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlePlayAgain not implemented")
}
func (UnimplementedMatchHandleServiceServer) HandleDeleteTicket(context.Context, *DeleteTicketSuccess) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleDeleteTicket not implemented")
}
func (UnimplementedMatchHandleServiceServer) mustEmbedUnimplementedMatchHandleServiceServer() {}

// UnsafeMatchHandleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchHandleServiceServer will
// result in compilation errors.
type UnsafeMatchHandleServiceServer interface {
	mustEmbedUnimplementedMatchHandleServiceServer()
}

func RegisterMatchHandleServiceServer(s grpc.ServiceRegistrar, srv MatchHandleServiceServer) {
	s.RegisterService(&MatchHandleService_ServiceDesc, srv)
}

func _MatchHandleService_HandlePlayAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataPlayAgainSuccess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchHandleServiceServer).HandlePlayAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.proto.MatchHandleService/HandlePlayAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchHandleServiceServer).HandlePlayAgain(ctx, req.(*DataPlayAgainSuccess))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchHandleService_HandleDeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketSuccess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchHandleServiceServer).HandleDeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.proto.MatchHandleService/HandleDeleteTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchHandleServiceServer).HandleDeleteTicket(ctx, req.(*DeleteTicketSuccess))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchHandleService_ServiceDesc is the grpc.ServiceDesc for MatchHandleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchHandleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.proto.MatchHandleService",
	HandlerType: (*MatchHandleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandlePlayAgain",
			Handler:    _MatchHandleService_HandlePlayAgain_Handler,
		},
		{
			MethodName: "HandleDeleteTicket",
			Handler:    _MatchHandleService_HandleDeleteTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Match.proto",
}

// MatchDirectorServiceClient is the client API for MatchDirectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchDirectorServiceClient interface {
	CancelTicket(ctx context.Context, in *MatchProtoPackage, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetBots(ctx context.Context, in *GetBotsRequest, opts ...grpc.CallOption) (*BotsResponse, error)
	CreateMatchOnboard(ctx context.Context, in *CreateMatchOnboardRequest, opts ...grpc.CallOption) (*MatchFoundResponseExtProto, error)
}

type matchDirectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchDirectorServiceClient(cc grpc.ClientConnInterface) MatchDirectorServiceClient {
	return &matchDirectorServiceClient{cc}
}

func (c *matchDirectorServiceClient) CancelTicket(ctx context.Context, in *MatchProtoPackage, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/core.proto.MatchDirectorService/CancelTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchDirectorServiceClient) GetBots(ctx context.Context, in *GetBotsRequest, opts ...grpc.CallOption) (*BotsResponse, error) {
	out := new(BotsResponse)
	err := c.cc.Invoke(ctx, "/core.proto.MatchDirectorService/GetBots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchDirectorServiceClient) CreateMatchOnboard(ctx context.Context, in *CreateMatchOnboardRequest, opts ...grpc.CallOption) (*MatchFoundResponseExtProto, error) {
	out := new(MatchFoundResponseExtProto)
	err := c.cc.Invoke(ctx, "/core.proto.MatchDirectorService/CreateMatchOnboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchDirectorServiceServer is the server API for MatchDirectorService service.
// All implementations must embed UnimplementedMatchDirectorServiceServer
// for forward compatibility
type MatchDirectorServiceServer interface {
	CancelTicket(context.Context, *MatchProtoPackage) (*EmptyResponse, error)
	GetBots(context.Context, *GetBotsRequest) (*BotsResponse, error)
	CreateMatchOnboard(context.Context, *CreateMatchOnboardRequest) (*MatchFoundResponseExtProto, error)
	mustEmbedUnimplementedMatchDirectorServiceServer()
}

// UnimplementedMatchDirectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchDirectorServiceServer struct {
}

func (UnimplementedMatchDirectorServiceServer) CancelTicket(context.Context, *MatchProtoPackage) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTicket not implemented")
}
func (UnimplementedMatchDirectorServiceServer) GetBots(context.Context, *GetBotsRequest) (*BotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBots not implemented")
}
func (UnimplementedMatchDirectorServiceServer) CreateMatchOnboard(context.Context, *CreateMatchOnboardRequest) (*MatchFoundResponseExtProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatchOnboard not implemented")
}
func (UnimplementedMatchDirectorServiceServer) mustEmbedUnimplementedMatchDirectorServiceServer() {}

// UnsafeMatchDirectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchDirectorServiceServer will
// result in compilation errors.
type UnsafeMatchDirectorServiceServer interface {
	mustEmbedUnimplementedMatchDirectorServiceServer()
}

func RegisterMatchDirectorServiceServer(s grpc.ServiceRegistrar, srv MatchDirectorServiceServer) {
	s.RegisterService(&MatchDirectorService_ServiceDesc, srv)
}

func _MatchDirectorService_CancelTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchProtoPackage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchDirectorServiceServer).CancelTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.proto.MatchDirectorService/CancelTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchDirectorServiceServer).CancelTicket(ctx, req.(*MatchProtoPackage))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchDirectorService_GetBots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchDirectorServiceServer).GetBots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.proto.MatchDirectorService/GetBots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchDirectorServiceServer).GetBots(ctx, req.(*GetBotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchDirectorService_CreateMatchOnboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMatchOnboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchDirectorServiceServer).CreateMatchOnboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.proto.MatchDirectorService/CreateMatchOnboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchDirectorServiceServer).CreateMatchOnboard(ctx, req.(*CreateMatchOnboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchDirectorService_ServiceDesc is the grpc.ServiceDesc for MatchDirectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchDirectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.proto.MatchDirectorService",
	HandlerType: (*MatchDirectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelTicket",
			Handler:    _MatchDirectorService_CancelTicket_Handler,
		},
		{
			MethodName: "GetBots",
			Handler:    _MatchDirectorService_GetBots_Handler,
		},
		{
			MethodName: "CreateMatchOnboard",
			Handler:    _MatchDirectorService_CreateMatchOnboard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Match.proto",
}
