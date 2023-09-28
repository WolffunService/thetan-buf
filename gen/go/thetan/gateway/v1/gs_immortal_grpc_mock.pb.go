// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: thetan/gateway/v1/gs_immortal.proto

package thetan_gateway_v1

import (
	context "context"
	reflect "reflect"
	thetan_immortal_v1 "thetan-buf/gen/go/thetan/immortal/v1"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockThetanGatewayTrackingClient is a mock of ThetanGatewayTrackingClient interface.
type MockThetanGatewayTrackingClient struct {
	ctrl     *gomock.Controller
	recorder *MockThetanGatewayTrackingClientMockRecorder
}

// MockThetanGatewayTrackingClientMockRecorder is the mock recorder for MockThetanGatewayTrackingClient.
type MockThetanGatewayTrackingClientMockRecorder struct {
	mock *MockThetanGatewayTrackingClient
}

// NewMockThetanGatewayTrackingClient creates a new mock instance.
func NewMockThetanGatewayTrackingClient(ctrl *gomock.Controller) *MockThetanGatewayTrackingClient {
	mock := &MockThetanGatewayTrackingClient{ctrl: ctrl}
	mock.recorder = &MockThetanGatewayTrackingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanGatewayTrackingClient) EXPECT() *MockThetanGatewayTrackingClientMockRecorder {
	return m.recorder
}

// RoomAllocated mocks base method.
func (m *MockThetanGatewayTrackingClient) RoomAllocated(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RoomAllocated", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomAllocated indicates an expected call of RoomAllocated.
func (mr *MockThetanGatewayTrackingClientMockRecorder) RoomAllocated(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomAllocated", reflect.TypeOf((*MockThetanGatewayTrackingClient)(nil).RoomAllocated), varargs...)
}

// RoomRelease mocks base method.
func (m *MockThetanGatewayTrackingClient) RoomRelease(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RoomRelease", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomRelease indicates an expected call of RoomRelease.
func (mr *MockThetanGatewayTrackingClientMockRecorder) RoomRelease(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomRelease", reflect.TypeOf((*MockThetanGatewayTrackingClient)(nil).RoomRelease), varargs...)
}

// MockThetanGatewayTrackingServer is a mock of ThetanGatewayTrackingServer interface.
type MockThetanGatewayTrackingServer struct {
	ctrl     *gomock.Controller
	recorder *MockThetanGatewayTrackingServerMockRecorder
}

// MockThetanGatewayTrackingServerMockRecorder is the mock recorder for MockThetanGatewayTrackingServer.
type MockThetanGatewayTrackingServerMockRecorder struct {
	mock *MockThetanGatewayTrackingServer
}

// NewMockThetanGatewayTrackingServer creates a new mock instance.
func NewMockThetanGatewayTrackingServer(ctrl *gomock.Controller) *MockThetanGatewayTrackingServer {
	mock := &MockThetanGatewayTrackingServer{ctrl: ctrl}
	mock.recorder = &MockThetanGatewayTrackingServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanGatewayTrackingServer) EXPECT() *MockThetanGatewayTrackingServerMockRecorder {
	return m.recorder
}

// RoomAllocated mocks base method.
func (m *MockThetanGatewayTrackingServer) RoomAllocated(ctx context.Context, in *RoomRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoomAllocated", ctx, in)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomAllocated indicates an expected call of RoomAllocated.
func (mr *MockThetanGatewayTrackingServerMockRecorder) RoomAllocated(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomAllocated", reflect.TypeOf((*MockThetanGatewayTrackingServer)(nil).RoomAllocated), ctx, in)
}

// RoomRelease mocks base method.
func (m *MockThetanGatewayTrackingServer) RoomRelease(ctx context.Context, in *RoomRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoomRelease", ctx, in)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomRelease indicates an expected call of RoomRelease.
func (mr *MockThetanGatewayTrackingServerMockRecorder) RoomRelease(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomRelease", reflect.TypeOf((*MockThetanGatewayTrackingServer)(nil).RoomRelease), ctx, in)
}

// MockThetanGatewayImmortalClient is a mock of ThetanGatewayImmortalClient interface.
type MockThetanGatewayImmortalClient struct {
	ctrl     *gomock.Controller
	recorder *MockThetanGatewayImmortalClientMockRecorder
}

// MockThetanGatewayImmortalClientMockRecorder is the mock recorder for MockThetanGatewayImmortalClient.
type MockThetanGatewayImmortalClientMockRecorder struct {
	mock *MockThetanGatewayImmortalClient
}

// NewMockThetanGatewayImmortalClient creates a new mock instance.
func NewMockThetanGatewayImmortalClient(ctrl *gomock.Controller) *MockThetanGatewayImmortalClient {
	mock := &MockThetanGatewayImmortalClient{ctrl: ctrl}
	mock.recorder = &MockThetanGatewayImmortalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanGatewayImmortalClient) EXPECT() *MockThetanGatewayImmortalClientMockRecorder {
	return m.recorder
}

// AllocateGameServer mocks base method.
func (m *MockThetanGatewayImmortalClient) AllocateGameServer(ctx context.Context, in *thetan_immortal_v1.ImmortalMatchFoundResponseProto, opts ...grpc.CallOption) (*ImmortalRoomAllocationResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AllocateGameServer", varargs...)
	ret0, _ := ret[0].(*ImmortalRoomAllocationResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllocateGameServer indicates an expected call of AllocateGameServer.
func (mr *MockThetanGatewayImmortalClientMockRecorder) AllocateGameServer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateGameServer", reflect.TypeOf((*MockThetanGatewayImmortalClient)(nil).AllocateGameServer), varargs...)
}

// MockThetanGatewayImmortalServer is a mock of ThetanGatewayImmortalServer interface.
type MockThetanGatewayImmortalServer struct {
	ctrl     *gomock.Controller
	recorder *MockThetanGatewayImmortalServerMockRecorder
}

// MockThetanGatewayImmortalServerMockRecorder is the mock recorder for MockThetanGatewayImmortalServer.
type MockThetanGatewayImmortalServerMockRecorder struct {
	mock *MockThetanGatewayImmortalServer
}

// NewMockThetanGatewayImmortalServer creates a new mock instance.
func NewMockThetanGatewayImmortalServer(ctrl *gomock.Controller) *MockThetanGatewayImmortalServer {
	mock := &MockThetanGatewayImmortalServer{ctrl: ctrl}
	mock.recorder = &MockThetanGatewayImmortalServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanGatewayImmortalServer) EXPECT() *MockThetanGatewayImmortalServerMockRecorder {
	return m.recorder
}

// AllocateGameServer mocks base method.
func (m *MockThetanGatewayImmortalServer) AllocateGameServer(ctx context.Context, in *thetan_immortal_v1.ImmortalMatchFoundResponseProto) (*ImmortalRoomAllocationResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocateGameServer", ctx, in)
	ret0, _ := ret[0].(*ImmortalRoomAllocationResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllocateGameServer indicates an expected call of AllocateGameServer.
func (mr *MockThetanGatewayImmortalServerMockRecorder) AllocateGameServer(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateGameServer", reflect.TypeOf((*MockThetanGatewayImmortalServer)(nil).AllocateGameServer), ctx, in)
}
