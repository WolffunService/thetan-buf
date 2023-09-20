// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: thetan/fishnet/rivals/v1/spectator_rivals.proto

package thetan_fishnet_spectator_rivals_v1

import (
	context "context"
	reflect "reflect"
	v1 "thetan-buf/gen/go/thetan/shared/v1"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockThetanSpectatorRivalsClient is a mock of ThetanSpectatorRivalsClient interface.
type MockThetanSpectatorRivalsClient struct {
	ctrl     *gomock.Controller
	recorder *MockThetanSpectatorRivalsClientMockRecorder
}

// MockThetanSpectatorRivalsClientMockRecorder is the mock recorder for MockThetanSpectatorRivalsClient.
type MockThetanSpectatorRivalsClientMockRecorder struct {
	mock *MockThetanSpectatorRivalsClient
}

// NewMockThetanSpectatorRivalsClient creates a new mock instance.
func NewMockThetanSpectatorRivalsClient(ctrl *gomock.Controller) *MockThetanSpectatorRivalsClient {
	mock := &MockThetanSpectatorRivalsClient{ctrl: ctrl}
	mock.recorder = &MockThetanSpectatorRivalsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanSpectatorRivalsClient) EXPECT() *MockThetanSpectatorRivalsClientMockRecorder {
	return m.recorder
}

// GameServerInfo mocks base method.
func (m *MockThetanSpectatorRivalsClient) GameServerInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GameServerInfoResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GameServerInfo", varargs...)
	ret0, _ := ret[0].(*GameServerInfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GameServerInfo indicates an expected call of GameServerInfo.
func (mr *MockThetanSpectatorRivalsClientMockRecorder) GameServerInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GameServerInfo", reflect.TypeOf((*MockThetanSpectatorRivalsClient)(nil).GameServerInfo), varargs...)
}

// RoomAllocation mocks base method.
func (m *MockThetanSpectatorRivalsClient) RoomAllocation(ctx context.Context, in *v1.MatchFoundResponseProto, opts ...grpc.CallOption) (*RoomAllocationResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RoomAllocation", varargs...)
	ret0, _ := ret[0].(*RoomAllocationResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomAllocation indicates an expected call of RoomAllocation.
func (mr *MockThetanSpectatorRivalsClientMockRecorder) RoomAllocation(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomAllocation", reflect.TypeOf((*MockThetanSpectatorRivalsClient)(nil).RoomAllocation), varargs...)
}

// Shutdown mocks base method.
func (m *MockThetanSpectatorRivalsClient) Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Shutdown", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockThetanSpectatorRivalsClientMockRecorder) Shutdown(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockThetanSpectatorRivalsClient)(nil).Shutdown), varargs...)
}

// MockThetanSpectatorRivalsServer is a mock of ThetanSpectatorRivalsServer interface.
type MockThetanSpectatorRivalsServer struct {
	ctrl     *gomock.Controller
	recorder *MockThetanSpectatorRivalsServerMockRecorder
}

// MockThetanSpectatorRivalsServerMockRecorder is the mock recorder for MockThetanSpectatorRivalsServer.
type MockThetanSpectatorRivalsServerMockRecorder struct {
	mock *MockThetanSpectatorRivalsServer
}

// NewMockThetanSpectatorRivalsServer creates a new mock instance.
func NewMockThetanSpectatorRivalsServer(ctrl *gomock.Controller) *MockThetanSpectatorRivalsServer {
	mock := &MockThetanSpectatorRivalsServer{ctrl: ctrl}
	mock.recorder = &MockThetanSpectatorRivalsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanSpectatorRivalsServer) EXPECT() *MockThetanSpectatorRivalsServerMockRecorder {
	return m.recorder
}

// GameServerInfo mocks base method.
func (m *MockThetanSpectatorRivalsServer) GameServerInfo(ctx context.Context, in *emptypb.Empty) (*GameServerInfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GameServerInfo", ctx, in)
	ret0, _ := ret[0].(*GameServerInfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GameServerInfo indicates an expected call of GameServerInfo.
func (mr *MockThetanSpectatorRivalsServerMockRecorder) GameServerInfo(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GameServerInfo", reflect.TypeOf((*MockThetanSpectatorRivalsServer)(nil).GameServerInfo), ctx, in)
}

// RoomAllocation mocks base method.
func (m *MockThetanSpectatorRivalsServer) RoomAllocation(ctx context.Context, in *v1.MatchFoundResponseProto) (*RoomAllocationResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoomAllocation", ctx, in)
	ret0, _ := ret[0].(*RoomAllocationResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomAllocation indicates an expected call of RoomAllocation.
func (mr *MockThetanSpectatorRivalsServerMockRecorder) RoomAllocation(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomAllocation", reflect.TypeOf((*MockThetanSpectatorRivalsServer)(nil).RoomAllocation), ctx, in)
}

// Shutdown mocks base method.
func (m *MockThetanSpectatorRivalsServer) Shutdown(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", ctx, in)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockThetanSpectatorRivalsServerMockRecorder) Shutdown(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockThetanSpectatorRivalsServer)(nil).Shutdown), ctx, in)
}
