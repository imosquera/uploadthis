// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/imosquera/uploadthis/util (interfaces: OSFile)

package mocks

import (
	os "os"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of OSFile interface
type MockOSFile struct {
	ctrl     *gomock.Controller
	recorder *_MockOSFileRecorder
}

// Recorder for MockOSFile (not exported)
type _MockOSFileRecorder struct {
	mock *MockOSFile
}

func NewMockOSFile(ctrl *gomock.Controller) *MockOSFile {
	mock := &MockOSFile{ctrl: ctrl}
	mock.recorder = &_MockOSFileRecorder{mock}
	return mock
}

func (_m *MockOSFile) EXPECT() *_MockOSFileRecorder {
	return _m.recorder
}

func (_m *MockOSFile) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOSFileRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockOSFile) Read(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOSFileRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0)
}

func (_m *MockOSFile) ReadAt(_param0 []byte, _param1 int64) (int, error) {
	ret := _m.ctrl.Call(_m, "ReadAt", _param0, _param1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOSFileRecorder) ReadAt(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadAt", arg0, arg1)
}

func (_m *MockOSFile) Seek(_param0 int64, _param1 int) (int64, error) {
	ret := _m.ctrl.Call(_m, "Seek", _param0, _param1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOSFileRecorder) Seek(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Seek", arg0, arg1)
}

func (_m *MockOSFile) Stat() (os.FileInfo, error) {
	ret := _m.ctrl.Call(_m, "Stat")
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOSFileRecorder) Stat() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stat")
}
