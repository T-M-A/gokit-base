// Automatically generated by MockGen. DO NOT EDIT!
// Source: users/service.go

package users

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *_MockServiceRecorder
}

// Recorder for MockService (not exported)
type _MockServiceRecorder struct {
	mock *MockService
}

func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &_MockServiceRecorder{mock}
	return mock
}

func (_m *MockService) EXPECT() *_MockServiceRecorder {
	return _m.recorder
}

func (_m *MockService) CreateUser(id int, fname string, lname string, color string) (int, error) {
	ret := _m.ctrl.Call(_m, "CreateUser", id, fname, lname, color)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceRecorder) CreateUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUser", arg0, arg1, arg2, arg3)
}

func (_m *MockService) ReadUser(id int) (User, error) {
	ret := _m.ctrl.Call(_m, "ReadUser", id)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockServiceRecorder) ReadUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadUser", arg0)
}

func (_m *MockService) UpdateUserColor(id int, color string) error {
	ret := _m.ctrl.Call(_m, "UpdateUserColor", id, color)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockServiceRecorder) UpdateUserColor(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateUserColor", arg0, arg1)
}

func (_m *MockService) Users() []*User {
	ret := _m.ctrl.Call(_m, "Users")
	ret0, _ := ret[0].([]*User)
	return ret0
}

func (_mr *_MockServiceRecorder) Users() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Users")
}