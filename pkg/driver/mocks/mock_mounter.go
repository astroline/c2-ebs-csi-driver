// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/astroline/aws-ebs-csi-driver/pkg/driver (interfaces: Mounter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mount "k8s.io/kubernetes/pkg/util/mount"
	os "os"
	reflect "reflect"
)

// MockMounter is a mock of Mounter interface
type MockMounter struct {
	ctrl     *gomock.Controller
	recorder *MockMounterMockRecorder
}

// MockMounterMockRecorder is the mock recorder for MockMounter
type MockMounterMockRecorder struct {
	mock *MockMounter
}

// NewMockMounter creates a new mock instance
func NewMockMounter(ctrl *gomock.Controller) *MockMounter {
	mock := &MockMounter{ctrl: ctrl}
	mock.recorder = &MockMounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMounter) EXPECT() *MockMounterMockRecorder {
	return m.recorder
}

// CleanSubPaths mocks base method
func (m *MockMounter) CleanSubPaths(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "CleanSubPaths", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanSubPaths indicates an expected call of CleanSubPaths
func (mr *MockMounterMockRecorder) CleanSubPaths(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanSubPaths", reflect.TypeOf((*MockMounter)(nil).CleanSubPaths), arg0, arg1)
}

// DeviceOpened mocks base method
func (m *MockMounter) DeviceOpened(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "DeviceOpened", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeviceOpened indicates an expected call of DeviceOpened
func (mr *MockMounterMockRecorder) DeviceOpened(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeviceOpened", reflect.TypeOf((*MockMounter)(nil).DeviceOpened), arg0)
}

// EvalHostSymlinks mocks base method
func (m *MockMounter) EvalHostSymlinks(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "EvalHostSymlinks", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvalHostSymlinks indicates an expected call of EvalHostSymlinks
func (mr *MockMounterMockRecorder) EvalHostSymlinks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalHostSymlinks", reflect.TypeOf((*MockMounter)(nil).EvalHostSymlinks), arg0)
}

// ExistsPath mocks base method
func (m *MockMounter) ExistsPath(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "ExistsPath", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsPath indicates an expected call of ExistsPath
func (mr *MockMounterMockRecorder) ExistsPath(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsPath", reflect.TypeOf((*MockMounter)(nil).ExistsPath), arg0)
}

// FormatAndMount mocks base method
func (m *MockMounter) FormatAndMount(arg0, arg1, arg2 string, arg3 []string) error {
	ret := m.ctrl.Call(m, "FormatAndMount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FormatAndMount indicates an expected call of FormatAndMount
func (mr *MockMounterMockRecorder) FormatAndMount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatAndMount", reflect.TypeOf((*MockMounter)(nil).FormatAndMount), arg0, arg1, arg2, arg3)
}

// GetDeviceName mocks base method
func (m *MockMounter) GetDeviceName(arg0 string) (string, int, error) {
	ret := m.ctrl.Call(m, "GetDeviceName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDeviceName indicates an expected call of GetDeviceName
func (mr *MockMounterMockRecorder) GetDeviceName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceName", reflect.TypeOf((*MockMounter)(nil).GetDeviceName), arg0)
}

// GetDeviceNameFromMount mocks base method
func (m *MockMounter) GetDeviceNameFromMount(arg0, arg1 string) (string, error) {
	ret := m.ctrl.Call(m, "GetDeviceNameFromMount", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceNameFromMount indicates an expected call of GetDeviceNameFromMount
func (mr *MockMounterMockRecorder) GetDeviceNameFromMount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceNameFromMount", reflect.TypeOf((*MockMounter)(nil).GetDeviceNameFromMount), arg0, arg1)
}

// GetFSGroup mocks base method
func (m *MockMounter) GetFSGroup(arg0 string) (int64, error) {
	ret := m.ctrl.Call(m, "GetFSGroup", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFSGroup indicates an expected call of GetFSGroup
func (mr *MockMounterMockRecorder) GetFSGroup(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFSGroup", reflect.TypeOf((*MockMounter)(nil).GetFSGroup), arg0)
}

// GetFileType mocks base method
func (m *MockMounter) GetFileType(arg0 string) (mount.FileType, error) {
	ret := m.ctrl.Call(m, "GetFileType", arg0)
	ret0, _ := ret[0].(mount.FileType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileType indicates an expected call of GetFileType
func (mr *MockMounterMockRecorder) GetFileType(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileType", reflect.TypeOf((*MockMounter)(nil).GetFileType), arg0)
}

// GetMode mocks base method
func (m *MockMounter) GetMode(arg0 string) (os.FileMode, error) {
	ret := m.ctrl.Call(m, "GetMode", arg0)
	ret0, _ := ret[0].(os.FileMode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMode indicates an expected call of GetMode
func (mr *MockMounterMockRecorder) GetMode(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMode", reflect.TypeOf((*MockMounter)(nil).GetMode), arg0)
}

// GetMountRefs mocks base method
func (m *MockMounter) GetMountRefs(arg0 string) ([]string, error) {
	ret := m.ctrl.Call(m, "GetMountRefs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMountRefs indicates an expected call of GetMountRefs
func (mr *MockMounterMockRecorder) GetMountRefs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMountRefs", reflect.TypeOf((*MockMounter)(nil).GetMountRefs), arg0)
}

// GetSELinuxSupport mocks base method
func (m *MockMounter) GetSELinuxSupport(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "GetSELinuxSupport", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSELinuxSupport indicates an expected call of GetSELinuxSupport
func (mr *MockMounterMockRecorder) GetSELinuxSupport(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSELinuxSupport", reflect.TypeOf((*MockMounter)(nil).GetSELinuxSupport), arg0)
}

// IsLikelyNotMountPoint mocks base method
func (m *MockMounter) IsLikelyNotMountPoint(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "IsLikelyNotMountPoint", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLikelyNotMountPoint indicates an expected call of IsLikelyNotMountPoint
func (mr *MockMounterMockRecorder) IsLikelyNotMountPoint(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLikelyNotMountPoint", reflect.TypeOf((*MockMounter)(nil).IsLikelyNotMountPoint), arg0)
}

// IsMountPointMatch mocks base method
func (m *MockMounter) IsMountPointMatch(arg0 mount.MountPoint, arg1 string) bool {
	ret := m.ctrl.Call(m, "IsMountPointMatch", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsMountPointMatch indicates an expected call of IsMountPointMatch
func (mr *MockMounterMockRecorder) IsMountPointMatch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMountPointMatch", reflect.TypeOf((*MockMounter)(nil).IsMountPointMatch), arg0, arg1)
}

// IsNotMountPoint mocks base method
func (m *MockMounter) IsNotMountPoint(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "IsNotMountPoint", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsNotMountPoint indicates an expected call of IsNotMountPoint
func (mr *MockMounterMockRecorder) IsNotMountPoint(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotMountPoint", reflect.TypeOf((*MockMounter)(nil).IsNotMountPoint), arg0)
}

// List mocks base method
func (m *MockMounter) List() ([]mount.MountPoint, error) {
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]mount.MountPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockMounterMockRecorder) List() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMounter)(nil).List))
}

// MakeDir mocks base method
func (m *MockMounter) MakeDir(arg0 string) error {
	ret := m.ctrl.Call(m, "MakeDir", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir
func (mr *MockMounterMockRecorder) MakeDir(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockMounter)(nil).MakeDir), arg0)
}

// MakeFile mocks base method
func (m *MockMounter) MakeFile(arg0 string) error {
	ret := m.ctrl.Call(m, "MakeFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeFile indicates an expected call of MakeFile
func (mr *MockMounterMockRecorder) MakeFile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeFile", reflect.TypeOf((*MockMounter)(nil).MakeFile), arg0)
}

// MakeRShared mocks base method
func (m *MockMounter) MakeRShared(arg0 string) error {
	ret := m.ctrl.Call(m, "MakeRShared", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeRShared indicates an expected call of MakeRShared
func (mr *MockMounterMockRecorder) MakeRShared(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRShared", reflect.TypeOf((*MockMounter)(nil).MakeRShared), arg0)
}

// Mount mocks base method
func (m *MockMounter) Mount(arg0, arg1, arg2 string, arg3 []string) error {
	ret := m.ctrl.Call(m, "Mount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount
func (mr *MockMounterMockRecorder) Mount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockMounter)(nil).Mount), arg0, arg1, arg2, arg3)
}

// PathIsDevice mocks base method
func (m *MockMounter) PathIsDevice(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "PathIsDevice", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PathIsDevice indicates an expected call of PathIsDevice
func (mr *MockMounterMockRecorder) PathIsDevice(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PathIsDevice", reflect.TypeOf((*MockMounter)(nil).PathIsDevice), arg0)
}

// PrepareSafeSubpath mocks base method
func (m *MockMounter) PrepareSafeSubpath(arg0 mount.Subpath) (string, func(), error) {
	ret := m.ctrl.Call(m, "PrepareSafeSubpath", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(func())
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PrepareSafeSubpath indicates an expected call of PrepareSafeSubpath
func (mr *MockMounterMockRecorder) PrepareSafeSubpath(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareSafeSubpath", reflect.TypeOf((*MockMounter)(nil).PrepareSafeSubpath), arg0)
}

// Run mocks base method
func (m *MockMounter) Run(arg0 string, arg1 ...string) ([]byte, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockMounterMockRecorder) Run(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMounter)(nil).Run), varargs...)
}

// SafeMakeDir mocks base method
func (m *MockMounter) SafeMakeDir(arg0, arg1 string, arg2 os.FileMode) error {
	ret := m.ctrl.Call(m, "SafeMakeDir", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SafeMakeDir indicates an expected call of SafeMakeDir
func (mr *MockMounterMockRecorder) SafeMakeDir(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SafeMakeDir", reflect.TypeOf((*MockMounter)(nil).SafeMakeDir), arg0, arg1, arg2)
}

// Unmount mocks base method
func (m *MockMounter) Unmount(arg0 string) error {
	ret := m.ctrl.Call(m, "Unmount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount
func (mr *MockMounterMockRecorder) Unmount(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockMounter)(nil).Unmount), arg0)
}
