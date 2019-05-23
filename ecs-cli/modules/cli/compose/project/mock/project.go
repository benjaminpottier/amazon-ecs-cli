// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/cli/compose/project (interfaces: Project)

// Package mock_project is a generated GoMock package.
package mock_project

import (
	reflect "reflect"

	adapter "github.com/aws/amazon-ecs-cli/ecs-cli/modules/cli/compose/adapter"
	context "github.com/aws/amazon-ecs-cli/ecs-cli/modules/cli/compose/context"
	entity "github.com/aws/amazon-ecs-cli/ecs-cli/modules/cli/compose/entity"
	project "github.com/docker/libcompose/project"
	gomock "github.com/golang/mock/gomock"
)

// MockProject is a mock of Project interface
type MockProject struct {
	ctrl     *gomock.Controller
	recorder *MockProjectMockRecorder
}

// MockProjectMockRecorder is the mock recorder for MockProject
type MockProjectMockRecorder struct {
	mock *MockProject
}

// NewMockProject creates a new mock instance
func NewMockProject(ctrl *gomock.Controller) *MockProject {
	mock := &MockProject{ctrl: ctrl}
	mock.recorder = &MockProjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProject) EXPECT() *MockProjectMockRecorder {
	return m.recorder
}

// ContainerConfigs mocks base method
func (m *MockProject) ContainerConfigs() []adapter.ContainerConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerConfigs")
	ret0, _ := ret[0].([]adapter.ContainerConfig)
	return ret0
}

// ContainerConfigs indicates an expected call of ContainerConfigs
func (mr *MockProjectMockRecorder) ContainerConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerConfigs", reflect.TypeOf((*MockProject)(nil).ContainerConfigs))
}

// Context mocks base method
func (m *MockProject) Context() *context.ECSContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(*context.ECSContext)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockProjectMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockProject)(nil).Context))
}

// Create mocks base method
func (m *MockProject) Create() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockProjectMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProject)(nil).Create))
}

// Down mocks base method
func (m *MockProject) Down() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Down")
	ret0, _ := ret[0].(error)
	return ret0
}

// Down indicates an expected call of Down
func (mr *MockProjectMockRecorder) Down() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Down", reflect.TypeOf((*MockProject)(nil).Down))
}

// Entity mocks base method
func (m *MockProject) Entity() entity.ProjectEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entity")
	ret0, _ := ret[0].(entity.ProjectEntity)
	return ret0
}

// Entity indicates an expected call of Entity
func (mr *MockProjectMockRecorder) Entity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entity", reflect.TypeOf((*MockProject)(nil).Entity))
}

// Info mocks base method
func (m *MockProject) Info(arg0 string) (project.InfoSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(project.InfoSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockProjectMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockProject)(nil).Info), arg0)
}

// Name mocks base method
func (m *MockProject) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockProjectMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockProject)(nil).Name))
}

// Parse mocks base method
func (m *MockProject) Parse() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse")
	ret0, _ := ret[0].(error)
	return ret0
}

// Parse indicates an expected call of Parse
func (mr *MockProjectMockRecorder) Parse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockProject)(nil).Parse))
}

// Run mocks base method
func (m *MockProject) Run(arg0 map[string][]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockProjectMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockProject)(nil).Run), arg0)
}

// Scale mocks base method
func (m *MockProject) Scale(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scale indicates an expected call of Scale
func (mr *MockProjectMockRecorder) Scale(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockProject)(nil).Scale), arg0)
}

// Start mocks base method
func (m *MockProject) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockProjectMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProject)(nil).Start))
}

// Stop mocks base method
func (m *MockProject) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockProjectMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockProject)(nil).Stop))
}

// Up mocks base method
func (m *MockProject) Up() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up")
	ret0, _ := ret[0].(error)
	return ret0
}

// Up indicates an expected call of Up
func (mr *MockProjectMockRecorder) Up() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockProject)(nil).Up))
}

// VolumeConfigs mocks base method
func (m *MockProject) VolumeConfigs() *adapter.Volumes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeConfigs")
	ret0, _ := ret[0].(*adapter.Volumes)
	return ret0
}

// VolumeConfigs indicates an expected call of VolumeConfigs
func (mr *MockProjectMockRecorder) VolumeConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeConfigs", reflect.TypeOf((*MockProject)(nil).VolumeConfigs))
}
