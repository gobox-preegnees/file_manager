// Code generated by MockGen. DO NOT EDIT.
// Source: owner.go

// Package service_message_owner is a generated GoMock package.
package service_message_owner

import (
	context "context"
	reflect "reflect"

	repo "github.com/gobox-preegnees/file_manager/internal/adapters/dao"
	entity "github.com/gobox-preegnees/file_manager/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIDaoOwner is a mock of IDaoOwner interface.
type MockIDaoOwner struct {
	ctrl     *gomock.Controller
	recorder *MockIDaoOwnerMockRecorder
}

// MockIDaoOwnerMockRecorder is the mock recorder for MockIDaoOwner.
type MockIDaoOwnerMockRecorder struct {
	mock *MockIDaoOwner
}

// NewMockIDaoOwner creates a new mock instance.
func NewMockIDaoOwner(ctrl *gomock.Controller) *MockIDaoOwner {
	mock := &MockIDaoOwner{ctrl: ctrl}
	mock.recorder = &MockIDaoOwnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDaoOwner) EXPECT() *MockIDaoOwnerMockRecorder {
	return m.recorder
}

// DeleteOwner mocks base method.
func (m *MockIDaoOwner) DeleteOwner(ctx context.Context, deleteOwnerDTO repo.DeleteOwnerDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOwner", ctx, deleteOwnerDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOwner indicates an expected call of DeleteOwner.
func (mr *MockIDaoOwnerMockRecorder) DeleteOwner(ctx, deleteOwnerDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOwner", reflect.TypeOf((*MockIDaoOwner)(nil).DeleteOwner), ctx, deleteOwnerDTO)
}

// SaveOwner mocks base method.
func (m *MockIDaoOwner) SaveOwner(ctx context.Context, saveOwnerDTO repo.SaveOwnerDTO) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOwner", ctx, saveOwnerDTO)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveOwner indicates an expected call of SaveOwner.
func (mr *MockIDaoOwnerMockRecorder) SaveOwner(ctx, saveOwnerDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOwner", reflect.TypeOf((*MockIDaoOwner)(nil).SaveOwner), ctx, saveOwnerDTO)
}

// MockIServiceMessageOwner is a mock of IServiceMessageOwner interface.
type MockIServiceMessageOwner struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMessageOwnerMockRecorder
}

// MockIServiceMessageOwnerMockRecorder is the mock recorder for MockIServiceMessageOwner.
type MockIServiceMessageOwnerMockRecorder struct {
	mock *MockIServiceMessageOwner
}

// NewMockIServiceMessageOwner creates a new mock instance.
func NewMockIServiceMessageOwner(ctrl *gomock.Controller) *MockIServiceMessageOwner {
	mock := &MockIServiceMessageOwner{ctrl: ctrl}
	mock.recorder = &MockIServiceMessageOwnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIServiceMessageOwner) EXPECT() *MockIServiceMessageOwnerMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockIServiceMessageOwner) SendMessage(message entity.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockIServiceMessageOwnerMockRecorder) SendMessage(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockIServiceMessageOwner)(nil).SendMessage), message)
}
