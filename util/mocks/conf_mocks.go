// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/imosquera/uploadthis/conf (interfaces: ConfigLoader,LoggerConfig)

package mocks

import (
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of ConfigLoader interface
type MockConfigLoader struct {
	ctrl     *gomock.Controller
	recorder *_MockConfigLoaderRecorder
}

// Recorder for MockConfigLoader (not exported)
type _MockConfigLoaderRecorder struct {
	mock *MockConfigLoader
}

func NewMockConfigLoader(ctrl *gomock.Controller) *MockConfigLoader {
	mock := &MockConfigLoader{ctrl: ctrl}
	mock.recorder = &_MockConfigLoaderRecorder{mock}
	return mock
}

func (_m *MockConfigLoader) EXPECT() *_MockConfigLoaderRecorder {
	return _m.recorder
}

func (_m *MockConfigLoader) LoadConfig(_param0 string, _param1 interface{}) {
	_m.ctrl.Call(_m, "LoadConfig", _param0, _param1)
}

func (_mr *_MockConfigLoaderRecorder) LoadConfig(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoadConfig", arg0, arg1)
}

// Mock of LoggerConfig interface
type MockLoggerConfig struct {
	ctrl     *gomock.Controller
	recorder *_MockLoggerConfigRecorder
}

// Recorder for MockLoggerConfig (not exported)
type _MockLoggerConfigRecorder struct {
	mock *MockLoggerConfig
}

func NewMockLoggerConfig(ctrl *gomock.Controller) *MockLoggerConfig {
	mock := &MockLoggerConfig{ctrl: ctrl}
	mock.recorder = &_MockLoggerConfigRecorder{mock}
	return mock
}

func (_m *MockLoggerConfig) EXPECT() *_MockLoggerConfigRecorder {
	return _m.recorder
}

func (_m *MockLoggerConfig) ConfigLogger(_param0 string) {
	_m.ctrl.Call(_m, "ConfigLogger", _param0)
}

func (_mr *_MockLoggerConfigRecorder) ConfigLogger(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigLogger", arg0)
}
