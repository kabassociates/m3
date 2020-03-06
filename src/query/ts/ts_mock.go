// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/query/ts (interfaces: Values)

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package ts is a generated GoMock package.
package ts

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/query/models"

	"github.com/golang/mock/gomock"
)

// MockValues is a mock of Values interface
type MockValues struct {
	ctrl     *gomock.Controller
	recorder *MockValuesMockRecorder
}

// MockValuesMockRecorder is the mock recorder for MockValues
type MockValuesMockRecorder struct {
	mock *MockValues
}

// NewMockValues creates a new mock instance
func NewMockValues(ctrl *gomock.Controller) *MockValues {
	mock := &MockValues{ctrl: ctrl}
	mock.recorder = &MockValuesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValues) EXPECT() *MockValuesMockRecorder {
	return m.recorder
}

// AlignToBounds mocks base method
func (m *MockValues) AlignToBounds(arg0 models.Bounds, arg1 time.Duration, arg2 AlignedDatapoints) AlignedDatapoints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlignToBounds", arg0, arg1, arg2)
	ret0, _ := ret[0].(AlignedDatapoints)
	return ret0
}

// AlignToBounds indicates an expected call of AlignToBounds
func (mr *MockValuesMockRecorder) AlignToBounds(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlignToBounds", reflect.TypeOf((*MockValues)(nil).AlignToBounds), arg0, arg1, arg2)
}

// AlignToBoundsNoWriteForward mocks base method
func (m *MockValues) AlignToBoundsNoWriteForward(arg0 models.Bounds, arg1 time.Duration, arg2 AlignedDatapoints) AlignedDatapoints {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlignToBoundsNoWriteForward", arg0, arg1, arg2)
	ret0, _ := ret[0].(AlignedDatapoints)
	return ret0
}

// AlignToBoundsNoWriteForward indicates an expected call of AlignToBoundsNoWriteForward
func (mr *MockValuesMockRecorder) AlignToBoundsNoWriteForward(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlignToBoundsNoWriteForward", reflect.TypeOf((*MockValues)(nil).AlignToBoundsNoWriteForward), arg0, arg1, arg2)
}

// DatapointAt mocks base method
func (m *MockValues) DatapointAt(arg0 int) Datapoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatapointAt", arg0)
	ret0, _ := ret[0].(Datapoint)
	return ret0
}

// DatapointAt indicates an expected call of DatapointAt
func (mr *MockValuesMockRecorder) DatapointAt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatapointAt", reflect.TypeOf((*MockValues)(nil).DatapointAt), arg0)
}

// Datapoints mocks base method
func (m *MockValues) Datapoints() []Datapoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datapoints")
	ret0, _ := ret[0].([]Datapoint)
	return ret0
}

// Datapoints indicates an expected call of Datapoints
func (mr *MockValuesMockRecorder) Datapoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datapoints", reflect.TypeOf((*MockValues)(nil).Datapoints))
}

// Len mocks base method
func (m *MockValues) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len
func (mr *MockValuesMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockValues)(nil).Len))
}

// ValueAt mocks base method
func (m *MockValues) ValueAt(arg0 int) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueAt", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// ValueAt indicates an expected call of ValueAt
func (mr *MockValuesMockRecorder) ValueAt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueAt", reflect.TypeOf((*MockValues)(nil).ValueAt), arg0)
}