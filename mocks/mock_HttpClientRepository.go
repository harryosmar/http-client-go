// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/harryosmar/http-client-go (interfaces: HttpClientRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	bytes "bytes"
	context "context"
	url "net/url"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	library_http_client_go "github.com/harryosmar/http-client-go"
	ctx "github.com/harryosmar/http-client-go/ctx"
)

// MockHttpClientRepository is a mock of HttpClientRepository interface.
type MockHttpClientRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHttpClientRepositoryMockRecorder
}

// MockHttpClientRepositoryMockRecorder is the mock recorder for MockHttpClientRepository.
type MockHttpClientRepositoryMockRecorder struct {
	mock *MockHttpClientRepository
}

// NewMockHttpClientRepository creates a new mock instance.
func NewMockHttpClientRepository(ctrl *gomock.Controller) *MockHttpClientRepository {
	mock := &MockHttpClientRepository{ctrl: ctrl}
	mock.recorder = &MockHttpClientRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpClientRepository) EXPECT() *MockHttpClientRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockHttpClientRepository) Delete(arg0 context.Context, arg1 string, arg2 map[string]string) (*library_http_client_go.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(*library_http_client_go.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockHttpClientRepositoryMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHttpClientRepository)(nil).Delete), arg0, arg1, arg2)
}

// DeleteX mocks base method.
func (m *MockHttpClientRepository) DeleteX(arg0 context.Context, arg1 string, arg2 interface{}, arg3 map[string]string) (*library_http_client_go.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteX", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*library_http_client_go.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteX indicates an expected call of DeleteX.
func (mr *MockHttpClientRepositoryMockRecorder) DeleteX(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteX", reflect.TypeOf((*MockHttpClientRepository)(nil).DeleteX), arg0, arg1, arg2, arg3)
}

// DisableDebug mocks base method.
func (m *MockHttpClientRepository) DisableDebug() library_http_client_go.HttpClientRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableDebug")
	ret0, _ := ret[0].(library_http_client_go.HttpClientRepository)
	return ret0
}

// DisableDebug indicates an expected call of DisableDebug.
func (mr *MockHttpClientRepositoryMockRecorder) DisableDebug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableDebug", reflect.TypeOf((*MockHttpClientRepository)(nil).DisableDebug))
}

// EnableDebug mocks base method.
func (m *MockHttpClientRepository) EnableDebug() library_http_client_go.HttpClientRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableDebug")
	ret0, _ := ret[0].(library_http_client_go.HttpClientRepository)
	return ret0
}

// EnableDebug indicates an expected call of EnableDebug.
func (mr *MockHttpClientRepositoryMockRecorder) EnableDebug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDebug", reflect.TypeOf((*MockHttpClientRepository)(nil).EnableDebug))
}

// Get mocks base method.
func (m *MockHttpClientRepository) Get(arg0 context.Context, arg1 string, arg2 map[string][]string, arg3 map[string]string) (*library_http_client_go.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*library_http_client_go.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHttpClientRepositoryMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHttpClientRepository)(nil).Get), arg0, arg1, arg2, arg3)
}

// Post mocks base method.
func (m *MockHttpClientRepository) Post(arg0 context.Context, arg1 string, arg2 *bytes.Buffer, arg3 map[string]string) (*library_http_client_go.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*library_http_client_go.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockHttpClientRepositoryMockRecorder) Post(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHttpClientRepository)(nil).Post), arg0, arg1, arg2, arg3)
}

// PostFormUrlEncoded mocks base method.
func (m *MockHttpClientRepository) PostFormUrlEncoded(arg0 context.Context, arg1 string, arg2 url.Values, arg3 map[string]string) (*library_http_client_go.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostFormUrlEncoded", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*library_http_client_go.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostFormUrlEncoded indicates an expected call of PostFormUrlEncoded.
func (mr *MockHttpClientRepositoryMockRecorder) PostFormUrlEncoded(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostFormUrlEncoded", reflect.TypeOf((*MockHttpClientRepository)(nil).PostFormUrlEncoded), arg0, arg1, arg2, arg3)
}

// PostMultipart mocks base method.
func (m *MockHttpClientRepository) PostMultipart(arg0 context.Context, arg1 string, arg2 *os.File, arg3 map[string]string) (*library_http_client_go.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMultipart", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*library_http_client_go.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostMultipart indicates an expected call of PostMultipart.
func (mr *MockHttpClientRepositoryMockRecorder) PostMultipart(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMultipart", reflect.TypeOf((*MockHttpClientRepository)(nil).PostMultipart), arg0, arg1, arg2, arg3)
}

// Put mocks base method.
func (m *MockHttpClientRepository) Put(arg0 context.Context, arg1 string, arg2 *bytes.Buffer, arg3 map[string]string) (*library_http_client_go.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*library_http_client_go.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockHttpClientRepositoryMockRecorder) Put(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockHttpClientRepository)(nil).Put), arg0, arg1, arg2, arg3)
}

// SetLogger mocks base method.
func (m *MockHttpClientRepository) SetLogger(arg0 ctx.LoggerCtx) library_http_client_go.HttpClientRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLogger", arg0)
	ret0, _ := ret[0].(library_http_client_go.HttpClientRepository)
	return ret0
}

// SetLogger indicates an expected call of SetLogger.
func (mr *MockHttpClientRepositoryMockRecorder) SetLogger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockHttpClientRepository)(nil).SetLogger), arg0)
}
