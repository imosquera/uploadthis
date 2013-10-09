// Automatically generated by MockGen. DO NOT EDIT!
// Source: ./commands/commands.go

package mocks

import (
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of Commander interface
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *_MockCommanderRecorder
}

// Recorder for MockCommander (not exported)
type _MockCommanderRecorder struct {
	mock *MockCommander
}

func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &_MockCommanderRecorder{mock}
	return mock
}

func (_m *MockCommander) EXPECT() *_MockCommanderRecorder {
	return _m.recorder
}

func (_m *MockCommander) SetUploadFiles(uploadFiles []string) {
	_m.ctrl.Call(_m, "SetUploadFiles", uploadFiles)
}

func (_mr *_MockCommanderRecorder) SetUploadFiles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetUploadFiles", arg0)
}

func (_m *MockCommander) Prepare(workDir string) {
	_m.ctrl.Call(_m, "Prepare", workDir)
}

func (_mr *_MockCommanderRecorder) Prepare(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prepare", arg0)
}

func (_m *MockCommander) Run() ([]string, error) {
	ret := _m.ctrl.Call(_m, "Run")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCommanderRecorder) Run() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run")
}
