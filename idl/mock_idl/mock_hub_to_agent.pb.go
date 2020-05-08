// Code generated by MockGen. DO NOT EDIT.
// Source: hub_to_agent.pb.go

// Package mock_idl is a generated GoMock package.
package mock_idl

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	idl "github.com/greenplum-db/gpupgrade/idl"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockAgentClient is a mock of AgentClient interface
type MockAgentClient struct {
	ctrl     *gomock.Controller
	recorder *MockAgentClientMockRecorder
}

// MockAgentClientMockRecorder is the mock recorder for MockAgentClient
type MockAgentClientMockRecorder struct {
	mock *MockAgentClient
}

// NewMockAgentClient creates a new mock instance
func NewMockAgentClient(ctrl *gomock.Controller) *MockAgentClient {
	mock := &MockAgentClient{ctrl: ctrl}
	mock.recorder = &MockAgentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentClient) EXPECT() *MockAgentClientMockRecorder {
	return m.recorder
}

// CheckDiskSpace mocks base method
func (m *MockAgentClient) CheckDiskSpace(ctx context.Context, in *idl.CheckSegmentDiskSpaceRequest, opts ...grpc.CallOption) (*idl.CheckDiskSpaceReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckDiskSpace", varargs...)
	ret0, _ := ret[0].(*idl.CheckDiskSpaceReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskSpace indicates an expected call of CheckDiskSpace
func (mr *MockAgentClientMockRecorder) CheckDiskSpace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskSpace", reflect.TypeOf((*MockAgentClient)(nil).CheckDiskSpace), varargs...)
}

// UpgradePrimaries mocks base method
func (m *MockAgentClient) UpgradePrimaries(ctx context.Context, in *idl.UpgradePrimariesRequest, opts ...grpc.CallOption) (*idl.UpgradePrimariesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradePrimaries", varargs...)
	ret0, _ := ret[0].(*idl.UpgradePrimariesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradePrimaries indicates an expected call of UpgradePrimaries
func (mr *MockAgentClientMockRecorder) UpgradePrimaries(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradePrimaries", reflect.TypeOf((*MockAgentClient)(nil).UpgradePrimaries), varargs...)
}

// RenameDirectories mocks base method
func (m *MockAgentClient) RenameDirectories(ctx context.Context, in *idl.RenameDirectoriesRequest, opts ...grpc.CallOption) (*idl.RenameDirectoriesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenameDirectories", varargs...)
	ret0, _ := ret[0].(*idl.RenameDirectoriesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameDirectories indicates an expected call of RenameDirectories
func (mr *MockAgentClientMockRecorder) RenameDirectories(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameDirectories", reflect.TypeOf((*MockAgentClient)(nil).RenameDirectories), varargs...)
}

// StopAgent mocks base method
func (m *MockAgentClient) StopAgent(ctx context.Context, in *idl.StopAgentRequest, opts ...grpc.CallOption) (*idl.StopAgentReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopAgent", varargs...)
	ret0, _ := ret[0].(*idl.StopAgentReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopAgent indicates an expected call of StopAgent
func (mr *MockAgentClientMockRecorder) StopAgent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAgent", reflect.TypeOf((*MockAgentClient)(nil).StopAgent), varargs...)
}

// DeleteDataDirectories mocks base method
func (m *MockAgentClient) DeleteDataDirectories(ctx context.Context, in *idl.DeleteDataDirectoriesRequest, opts ...grpc.CallOption) (*idl.DeleteDataDirectoriesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDataDirectories", varargs...)
	ret0, _ := ret[0].(*idl.DeleteDataDirectoriesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDataDirectories indicates an expected call of DeleteDataDirectories
func (mr *MockAgentClientMockRecorder) DeleteDataDirectories(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataDirectories", reflect.TypeOf((*MockAgentClient)(nil).DeleteDataDirectories), varargs...)
}

// DeleteStateDirectory mocks base method
func (m *MockAgentClient) DeleteStateDirectory(ctx context.Context, in *idl.DeleteStateDirectoryRequest, opts ...grpc.CallOption) (*idl.DeleteStateDirectoryReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteStateDirectory", varargs...)
	ret0, _ := ret[0].(*idl.DeleteStateDirectoryReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStateDirectory indicates an expected call of DeleteStateDirectory
func (mr *MockAgentClientMockRecorder) DeleteStateDirectory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStateDirectory", reflect.TypeOf((*MockAgentClient)(nil).DeleteStateDirectory), varargs...)
}

// ArchiveLogDirectory mocks base method
func (m *MockAgentClient) ArchiveLogDirectory(ctx context.Context, in *idl.ArchiveLogDirectoryRequest, opts ...grpc.CallOption) (*idl.ArchiveLogDirectoryReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ArchiveLogDirectory", varargs...)
	ret0, _ := ret[0].(*idl.ArchiveLogDirectoryReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveLogDirectory indicates an expected call of ArchiveLogDirectory
func (mr *MockAgentClientMockRecorder) ArchiveLogDirectory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveLogDirectory", reflect.TypeOf((*MockAgentClient)(nil).ArchiveLogDirectory), varargs...)
}

// MockAgentServer is a mock of AgentServer interface
type MockAgentServer struct {
	ctrl     *gomock.Controller
	recorder *MockAgentServerMockRecorder
}

// MockAgentServerMockRecorder is the mock recorder for MockAgentServer
type MockAgentServerMockRecorder struct {
	mock *MockAgentServer
}

// NewMockAgentServer creates a new mock instance
func NewMockAgentServer(ctrl *gomock.Controller) *MockAgentServer {
	mock := &MockAgentServer{ctrl: ctrl}
	mock.recorder = &MockAgentServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentServer) EXPECT() *MockAgentServerMockRecorder {
	return m.recorder
}

// CheckDiskSpace mocks base method
func (m *MockAgentServer) CheckDiskSpace(arg0 context.Context, arg1 *idl.CheckSegmentDiskSpaceRequest) (*idl.CheckDiskSpaceReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDiskSpace", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckDiskSpaceReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskSpace indicates an expected call of CheckDiskSpace
func (mr *MockAgentServerMockRecorder) CheckDiskSpace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskSpace", reflect.TypeOf((*MockAgentServer)(nil).CheckDiskSpace), arg0, arg1)
}

// UpgradePrimaries mocks base method
func (m *MockAgentServer) UpgradePrimaries(arg0 context.Context, arg1 *idl.UpgradePrimariesRequest) (*idl.UpgradePrimariesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradePrimaries", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradePrimariesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradePrimaries indicates an expected call of UpgradePrimaries
func (mr *MockAgentServerMockRecorder) UpgradePrimaries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradePrimaries", reflect.TypeOf((*MockAgentServer)(nil).UpgradePrimaries), arg0, arg1)
}

// RenameDirectories mocks base method
func (m *MockAgentServer) RenameDirectories(arg0 context.Context, arg1 *idl.RenameDirectoriesRequest) (*idl.RenameDirectoriesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameDirectories", arg0, arg1)
	ret0, _ := ret[0].(*idl.RenameDirectoriesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameDirectories indicates an expected call of RenameDirectories
func (mr *MockAgentServerMockRecorder) RenameDirectories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameDirectories", reflect.TypeOf((*MockAgentServer)(nil).RenameDirectories), arg0, arg1)
}

// StopAgent mocks base method
func (m *MockAgentServer) StopAgent(arg0 context.Context, arg1 *idl.StopAgentRequest) (*idl.StopAgentReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAgent", arg0, arg1)
	ret0, _ := ret[0].(*idl.StopAgentReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopAgent indicates an expected call of StopAgent
func (mr *MockAgentServerMockRecorder) StopAgent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAgent", reflect.TypeOf((*MockAgentServer)(nil).StopAgent), arg0, arg1)
}

// DeleteDataDirectories mocks base method
func (m *MockAgentServer) DeleteDataDirectories(arg0 context.Context, arg1 *idl.DeleteDataDirectoriesRequest) (*idl.DeleteDataDirectoriesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDataDirectories", arg0, arg1)
	ret0, _ := ret[0].(*idl.DeleteDataDirectoriesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDataDirectories indicates an expected call of DeleteDataDirectories
func (mr *MockAgentServerMockRecorder) DeleteDataDirectories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataDirectories", reflect.TypeOf((*MockAgentServer)(nil).DeleteDataDirectories), arg0, arg1)
}

// DeleteStateDirectory mocks base method
func (m *MockAgentServer) DeleteStateDirectory(arg0 context.Context, arg1 *idl.DeleteStateDirectoryRequest) (*idl.DeleteStateDirectoryReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStateDirectory", arg0, arg1)
	ret0, _ := ret[0].(*idl.DeleteStateDirectoryReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStateDirectory indicates an expected call of DeleteStateDirectory
func (mr *MockAgentServerMockRecorder) DeleteStateDirectory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStateDirectory", reflect.TypeOf((*MockAgentServer)(nil).DeleteStateDirectory), arg0, arg1)
}

// ArchiveLogDirectory mocks base method
func (m *MockAgentServer) ArchiveLogDirectory(arg0 context.Context, arg1 *idl.ArchiveLogDirectoryRequest) (*idl.ArchiveLogDirectoryReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArchiveLogDirectory", arg0, arg1)
	ret0, _ := ret[0].(*idl.ArchiveLogDirectoryReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveLogDirectory indicates an expected call of ArchiveLogDirectory
func (mr *MockAgentServerMockRecorder) ArchiveLogDirectory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveLogDirectory", reflect.TypeOf((*MockAgentServer)(nil).ArchiveLogDirectory), arg0, arg1)
}
