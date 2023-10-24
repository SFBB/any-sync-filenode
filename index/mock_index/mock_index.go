// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync-filenode/index (interfaces: Index)
//
// Generated by this command:
//
//	mockgen -destination mock_index/mock_index.go github.com/anyproto/any-sync-filenode/index Index
//
// Package mock_index is a generated GoMock package.
package mock_index

import (
	context "context"
	reflect "reflect"

	index "github.com/anyproto/any-sync-filenode/index"
	app "github.com/anyproto/any-sync/app"
	blocks "github.com/ipfs/go-block-format"
	cid "github.com/ipfs/go-cid"
	gomock "go.uber.org/mock/gomock"
)

// MockIndex is a mock of Index interface.
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *MockIndexMockRecorder
}

// MockIndexMockRecorder is the mock recorder for MockIndex.
type MockIndexMockRecorder struct {
	mock *MockIndex
}

// NewMockIndex creates a new mock instance.
func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &MockIndexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndex) EXPECT() *MockIndexMockRecorder {
	return m.recorder
}

// BlocksAdd mocks base method.
func (m *MockIndex) BlocksAdd(arg0 context.Context, arg1 []blocks.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlocksAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BlocksAdd indicates an expected call of BlocksAdd.
func (mr *MockIndexMockRecorder) BlocksAdd(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlocksAdd", reflect.TypeOf((*MockIndex)(nil).BlocksAdd), arg0, arg1)
}

// BlocksGetNonExistent mocks base method.
func (m *MockIndex) BlocksGetNonExistent(arg0 context.Context, arg1 []blocks.Block) ([]blocks.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlocksGetNonExistent", arg0, arg1)
	ret0, _ := ret[0].([]blocks.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlocksGetNonExistent indicates an expected call of BlocksGetNonExistent.
func (mr *MockIndexMockRecorder) BlocksGetNonExistent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlocksGetNonExistent", reflect.TypeOf((*MockIndex)(nil).BlocksGetNonExistent), arg0, arg1)
}

// BlocksLock mocks base method.
func (m *MockIndex) BlocksLock(arg0 context.Context, arg1 []blocks.Block) (func(), error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlocksLock", arg0, arg1)
	ret0, _ := ret[0].(func())
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlocksLock indicates an expected call of BlocksLock.
func (mr *MockIndexMockRecorder) BlocksLock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlocksLock", reflect.TypeOf((*MockIndex)(nil).BlocksLock), arg0, arg1)
}

// CidEntries mocks base method.
func (m *MockIndex) CidEntries(arg0 context.Context, arg1 []cid.Cid) (*index.CidEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CidEntries", arg0, arg1)
	ret0, _ := ret[0].(*index.CidEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CidEntries indicates an expected call of CidEntries.
func (mr *MockIndexMockRecorder) CidEntries(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CidEntries", reflect.TypeOf((*MockIndex)(nil).CidEntries), arg0, arg1)
}

// CidEntriesByBlocks mocks base method.
func (m *MockIndex) CidEntriesByBlocks(arg0 context.Context, arg1 []blocks.Block) (*index.CidEntries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CidEntriesByBlocks", arg0, arg1)
	ret0, _ := ret[0].(*index.CidEntries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CidEntriesByBlocks indicates an expected call of CidEntriesByBlocks.
func (mr *MockIndexMockRecorder) CidEntriesByBlocks(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CidEntriesByBlocks", reflect.TypeOf((*MockIndex)(nil).CidEntriesByBlocks), arg0, arg1)
}

// CidExists mocks base method.
func (m *MockIndex) CidExists(arg0 context.Context, arg1 cid.Cid) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CidExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CidExists indicates an expected call of CidExists.
func (mr *MockIndexMockRecorder) CidExists(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CidExists", reflect.TypeOf((*MockIndex)(nil).CidExists), arg0, arg1)
}

// CidExistsInSpace mocks base method.
func (m *MockIndex) CidExistsInSpace(arg0 context.Context, arg1 index.Key, arg2 []cid.Cid) ([]cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CidExistsInSpace", arg0, arg1, arg2)
	ret0, _ := ret[0].([]cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CidExistsInSpace indicates an expected call of CidExistsInSpace.
func (mr *MockIndexMockRecorder) CidExistsInSpace(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CidExistsInSpace", reflect.TypeOf((*MockIndex)(nil).CidExistsInSpace), arg0, arg1, arg2)
}

// Close mocks base method.
func (m *MockIndex) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIndexMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIndex)(nil).Close), arg0)
}

// FileBind mocks base method.
func (m *MockIndex) FileBind(arg0 context.Context, arg1 index.Key, arg2 string, arg3 *index.CidEntries) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileBind", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileBind indicates an expected call of FileBind.
func (mr *MockIndexMockRecorder) FileBind(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileBind", reflect.TypeOf((*MockIndex)(nil).FileBind), arg0, arg1, arg2, arg3)
}

// FileInfo mocks base method.
func (m *MockIndex) FileInfo(arg0 context.Context, arg1 index.Key, arg2 ...string) ([]index.FileInfo, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FileInfo", varargs...)
	ret0, _ := ret[0].([]index.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileInfo indicates an expected call of FileInfo.
func (mr *MockIndexMockRecorder) FileInfo(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileInfo", reflect.TypeOf((*MockIndex)(nil).FileInfo), varargs...)
}

// FileUnbind mocks base method.
func (m *MockIndex) FileUnbind(arg0 context.Context, arg1 index.Key, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileUnbind", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileUnbind indicates an expected call of FileUnbind.
func (mr *MockIndexMockRecorder) FileUnbind(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileUnbind", reflect.TypeOf((*MockIndex)(nil).FileUnbind), arg0, arg1, arg2)
}

// GroupInfo mocks base method.
func (m *MockIndex) GroupInfo(arg0 context.Context, arg1 string) (index.GroupInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupInfo", arg0, arg1)
	ret0, _ := ret[0].(index.GroupInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GroupInfo indicates an expected call of GroupInfo.
func (mr *MockIndexMockRecorder) GroupInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupInfo", reflect.TypeOf((*MockIndex)(nil).GroupInfo), arg0, arg1)
}

// Init mocks base method.
func (m *MockIndex) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockIndexMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockIndex)(nil).Init), arg0)
}

// Migrate mocks base method.
func (m *MockIndex) Migrate(arg0 context.Context, arg1 index.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockIndexMockRecorder) Migrate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockIndex)(nil).Migrate), arg0, arg1)
}

// Name mocks base method.
func (m *MockIndex) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockIndexMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockIndex)(nil).Name))
}

// Run mocks base method.
func (m *MockIndex) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockIndexMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockIndex)(nil).Run), arg0)
}

// SpaceInfo mocks base method.
func (m *MockIndex) SpaceInfo(arg0 context.Context, arg1 index.Key) (index.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceInfo", arg0, arg1)
	ret0, _ := ret[0].(index.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceInfo indicates an expected call of SpaceInfo.
func (mr *MockIndexMockRecorder) SpaceInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceInfo", reflect.TypeOf((*MockIndex)(nil).SpaceInfo), arg0, arg1)
}
