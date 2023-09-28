// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: thetan/fishnet/rivalslobby/v1/fishnet_rivals_lobby.proto

package thetan_fishnet_rivalslobby_v1

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockThetanFishNetRivalsLobbyClient is a mock of ThetanFishNetRivalsLobbyClient interface.
type MockThetanFishNetRivalsLobbyClient struct {
	ctrl     *gomock.Controller
	recorder *MockThetanFishNetRivalsLobbyClientMockRecorder
}

// MockThetanFishNetRivalsLobbyClientMockRecorder is the mock recorder for MockThetanFishNetRivalsLobbyClient.
type MockThetanFishNetRivalsLobbyClientMockRecorder struct {
	mock *MockThetanFishNetRivalsLobbyClient
}

// NewMockThetanFishNetRivalsLobbyClient creates a new mock instance.
func NewMockThetanFishNetRivalsLobbyClient(ctrl *gomock.Controller) *MockThetanFishNetRivalsLobbyClient {
	mock := &MockThetanFishNetRivalsLobbyClient{ctrl: ctrl}
	mock.recorder = &MockThetanFishNetRivalsLobbyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanFishNetRivalsLobbyClient) EXPECT() *MockThetanFishNetRivalsLobbyClientMockRecorder {
	return m.recorder
}

// GameServerInfo mocks base method.
func (m *MockThetanFishNetRivalsLobbyClient) GameServerInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GameServerInfoResp, error) {
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
func (mr *MockThetanFishNetRivalsLobbyClientMockRecorder) GameServerInfo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GameServerInfo", reflect.TypeOf((*MockThetanFishNetRivalsLobbyClient)(nil).GameServerInfo), varargs...)
}

// RoomAllocation mocks base method.
func (m *MockThetanFishNetRivalsLobbyClient) RoomAllocation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RoomAllocationResp, error) {
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
func (mr *MockThetanFishNetRivalsLobbyClientMockRecorder) RoomAllocation(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomAllocation", reflect.TypeOf((*MockThetanFishNetRivalsLobbyClient)(nil).RoomAllocation), varargs...)
}

// SetTown mocks base method.
func (m *MockThetanFishNetRivalsLobbyClient) SetTown(ctx context.Context, in *SetTownRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetTown", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTown indicates an expected call of SetTown.
func (mr *MockThetanFishNetRivalsLobbyClientMockRecorder) SetTown(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTown", reflect.TypeOf((*MockThetanFishNetRivalsLobbyClient)(nil).SetTown), varargs...)
}

// Shutdown mocks base method.
func (m *MockThetanFishNetRivalsLobbyClient) Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
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
func (mr *MockThetanFishNetRivalsLobbyClientMockRecorder) Shutdown(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockThetanFishNetRivalsLobbyClient)(nil).Shutdown), varargs...)
}

// MockThetanFishNetRivalsLobbyServer is a mock of ThetanFishNetRivalsLobbyServer interface.
type MockThetanFishNetRivalsLobbyServer struct {
	ctrl     *gomock.Controller
	recorder *MockThetanFishNetRivalsLobbyServerMockRecorder
}

// MockThetanFishNetRivalsLobbyServerMockRecorder is the mock recorder for MockThetanFishNetRivalsLobbyServer.
type MockThetanFishNetRivalsLobbyServerMockRecorder struct {
	mock *MockThetanFishNetRivalsLobbyServer
}

// NewMockThetanFishNetRivalsLobbyServer creates a new mock instance.
func NewMockThetanFishNetRivalsLobbyServer(ctrl *gomock.Controller) *MockThetanFishNetRivalsLobbyServer {
	mock := &MockThetanFishNetRivalsLobbyServer{ctrl: ctrl}
	mock.recorder = &MockThetanFishNetRivalsLobbyServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThetanFishNetRivalsLobbyServer) EXPECT() *MockThetanFishNetRivalsLobbyServerMockRecorder {
	return m.recorder
}

// GameServerInfo mocks base method.
func (m *MockThetanFishNetRivalsLobbyServer) GameServerInfo(ctx context.Context, in *emptypb.Empty) (*GameServerInfoResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GameServerInfo", ctx, in)
	ret0, _ := ret[0].(*GameServerInfoResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GameServerInfo indicates an expected call of GameServerInfo.
func (mr *MockThetanFishNetRivalsLobbyServerMockRecorder) GameServerInfo(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GameServerInfo", reflect.TypeOf((*MockThetanFishNetRivalsLobbyServer)(nil).GameServerInfo), ctx, in)
}

// RoomAllocation mocks base method.
func (m *MockThetanFishNetRivalsLobbyServer) RoomAllocation(ctx context.Context, in *emptypb.Empty) (*RoomAllocationResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoomAllocation", ctx, in)
	ret0, _ := ret[0].(*RoomAllocationResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RoomAllocation indicates an expected call of RoomAllocation.
func (mr *MockThetanFishNetRivalsLobbyServerMockRecorder) RoomAllocation(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoomAllocation", reflect.TypeOf((*MockThetanFishNetRivalsLobbyServer)(nil).RoomAllocation), ctx, in)
}

// SetTown mocks base method.
func (m *MockThetanFishNetRivalsLobbyServer) SetTown(ctx context.Context, in *SetTownRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTown", ctx, in)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTown indicates an expected call of SetTown.
func (mr *MockThetanFishNetRivalsLobbyServerMockRecorder) SetTown(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTown", reflect.TypeOf((*MockThetanFishNetRivalsLobbyServer)(nil).SetTown), ctx, in)
}

// Shutdown mocks base method.
func (m *MockThetanFishNetRivalsLobbyServer) Shutdown(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", ctx, in)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockThetanFishNetRivalsLobbyServerMockRecorder) Shutdown(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockThetanFishNetRivalsLobbyServer)(nil).Shutdown), ctx, in)
}
