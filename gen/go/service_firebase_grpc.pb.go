// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: service_firebase.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FirebaseService_TrackPlayerStat_FullMethodName = "/services.FirebaseService/TrackPlayerStat"
)

// FirebaseServiceClient is the client API for FirebaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FirebaseServiceClient interface {
	TrackPlayerStat(ctx context.Context, in *TrackPlayerStatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type firebaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFirebaseServiceClient(cc grpc.ClientConnInterface) FirebaseServiceClient {
	return &firebaseServiceClient{cc}
}

func (c *firebaseServiceClient) TrackPlayerStat(ctx context.Context, in *TrackPlayerStatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FirebaseService_TrackPlayerStat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirebaseServiceServer is the server API for FirebaseService service.
// All implementations must embed UnimplementedFirebaseServiceServer
// for forward compatibility
type FirebaseServiceServer interface {
	TrackPlayerStat(context.Context, *TrackPlayerStatRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFirebaseServiceServer()
}

// UnimplementedFirebaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFirebaseServiceServer struct {
}

func (UnimplementedFirebaseServiceServer) TrackPlayerStat(context.Context, *TrackPlayerStatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackPlayerStat not implemented")
}
func (UnimplementedFirebaseServiceServer) mustEmbedUnimplementedFirebaseServiceServer() {}

// UnsafeFirebaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FirebaseServiceServer will
// result in compilation errors.
type UnsafeFirebaseServiceServer interface {
	mustEmbedUnimplementedFirebaseServiceServer()
}

func RegisterFirebaseServiceServer(s grpc.ServiceRegistrar, srv FirebaseServiceServer) {
	s.RegisterService(&FirebaseService_ServiceDesc, srv)
}

func _FirebaseService_TrackPlayerStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackPlayerStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirebaseServiceServer).TrackPlayerStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FirebaseService_TrackPlayerStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirebaseServiceServer).TrackPlayerStat(ctx, req.(*TrackPlayerStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FirebaseService_ServiceDesc is the grpc.ServiceDesc for FirebaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FirebaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.FirebaseService",
	HandlerType: (*FirebaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TrackPlayerStat",
			Handler:    _FirebaseService_TrackPlayerStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_firebase.proto",
}
