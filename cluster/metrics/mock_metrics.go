/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: metrics.go

// Package mock_metrics is a generated GoMock package.
package metrics

import (
	reflect "reflect"
)

import (
	gomock "github.com/golang/mock/gomock"
)

import (
	common "dubbo.apache.org/dubbo-go/v3/common"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// GetInstanceMetrics mocks base method.
func (m *MockMetrics) GetInstanceMetrics(url *common.URL, key string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceMetrics", url, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceMetrics indicates an expected call of GetInstanceMetrics.
func (mr *MockMetricsMockRecorder) GetInstanceMetrics(url, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceMetrics", reflect.TypeOf((*MockMetrics)(nil).GetInstanceMetrics), url, key)
}

// GetInvokerMetrics mocks base method.
func (m *MockMetrics) GetInvokerMetrics(url *common.URL, key string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvokerMetrics", url, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvokerMetrics indicates an expected call of GetInvokerMetrics.
func (mr *MockMetricsMockRecorder) GetInvokerMetrics(url, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvokerMetrics", reflect.TypeOf((*MockMetrics)(nil).GetInvokerMetrics), url, key)
}

// GetMethodMetrics mocks base method.
func (m *MockMetrics) GetMethodMetrics(url *common.URL, methodName, key string) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMethodMetrics", url, methodName, key)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMethodMetrics indicates an expected call of GetMethodMetrics.
func (mr *MockMetricsMockRecorder) GetMethodMetrics(url, methodName, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMethodMetrics", reflect.TypeOf((*MockMetrics)(nil).GetMethodMetrics), url, methodName, key)
}

// SetInstanceMetrics mocks base method.
func (m *MockMetrics) SetInstanceMetrics(url *common.URL, key string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInstanceMetrics", url, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInstanceMetrics indicates an expected call of SetInstanceMetrics.
func (mr *MockMetricsMockRecorder) SetInstanceMetrics(url, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInstanceMetrics", reflect.TypeOf((*MockMetrics)(nil).SetInstanceMetrics), url, key, value)
}

// SetInvokerMetrics mocks base method.
func (m *MockMetrics) SetInvokerMetrics(url *common.URL, key string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInvokerMetrics", url, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInvokerMetrics indicates an expected call of SetInvokerMetrics.
func (mr *MockMetricsMockRecorder) SetInvokerMetrics(url, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInvokerMetrics", reflect.TypeOf((*MockMetrics)(nil).SetInvokerMetrics), url, key, value)
}

// SetMethodMetrics mocks base method.
func (m *MockMetrics) SetMethodMetrics(url *common.URL, methodName, key string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMethodMetrics", url, methodName, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMethodMetrics indicates an expected call of SetMethodMetrics.
func (mr *MockMetricsMockRecorder) SetMethodMetrics(url, methodName, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMethodMetrics", reflect.TypeOf((*MockMetrics)(nil).SetMethodMetrics), url, methodName, key, value)
}
