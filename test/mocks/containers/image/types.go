// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/containers/image/v5/types (interfaces: ImageCloser)

// Package imagetypesmock is a generated GoMock package.
package imagetypesmock

import (
	context "context"
	reference "github.com/containers/image/v5/docker/reference"
	types "github.com/containers/image/v5/types"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	reflect "reflect"
)

// MockImageCloser is a mock of ImageCloser interface
type MockImageCloser struct {
	ctrl     *gomock.Controller
	recorder *MockImageCloserMockRecorder
}

// MockImageCloserMockRecorder is the mock recorder for MockImageCloser
type MockImageCloserMockRecorder struct {
	mock *MockImageCloser
}

// NewMockImageCloser creates a new mock instance
func NewMockImageCloser(ctrl *gomock.Controller) *MockImageCloser {
	mock := &MockImageCloser{ctrl: ctrl}
	mock.recorder = &MockImageCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageCloser) EXPECT() *MockImageCloserMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockImageCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockImageCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockImageCloser)(nil).Close))
}

// ConfigBlob mocks base method
func (m *MockImageCloser) ConfigBlob(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigBlob", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfigBlob indicates an expected call of ConfigBlob
func (mr *MockImageCloserMockRecorder) ConfigBlob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigBlob", reflect.TypeOf((*MockImageCloser)(nil).ConfigBlob), arg0)
}

// ConfigInfo mocks base method
func (m *MockImageCloser) ConfigInfo() types.BlobInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigInfo")
	ret0, _ := ret[0].(types.BlobInfo)
	return ret0
}

// ConfigInfo indicates an expected call of ConfigInfo
func (mr *MockImageCloserMockRecorder) ConfigInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigInfo", reflect.TypeOf((*MockImageCloser)(nil).ConfigInfo))
}

// EmbeddedDockerReferenceConflicts mocks base method
func (m *MockImageCloser) EmbeddedDockerReferenceConflicts(arg0 reference.Named) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmbeddedDockerReferenceConflicts", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// EmbeddedDockerReferenceConflicts indicates an expected call of EmbeddedDockerReferenceConflicts
func (mr *MockImageCloserMockRecorder) EmbeddedDockerReferenceConflicts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmbeddedDockerReferenceConflicts", reflect.TypeOf((*MockImageCloser)(nil).EmbeddedDockerReferenceConflicts), arg0)
}

// Inspect mocks base method
func (m *MockImageCloser) Inspect(arg0 context.Context) (*types.ImageInspectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", arg0)
	ret0, _ := ret[0].(*types.ImageInspectInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockImageCloserMockRecorder) Inspect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockImageCloser)(nil).Inspect), arg0)
}

// LayerInfos mocks base method
func (m *MockImageCloser) LayerInfos() []types.BlobInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerInfos")
	ret0, _ := ret[0].([]types.BlobInfo)
	return ret0
}

// LayerInfos indicates an expected call of LayerInfos
func (mr *MockImageCloserMockRecorder) LayerInfos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerInfos", reflect.TypeOf((*MockImageCloser)(nil).LayerInfos))
}

// LayerInfosForCopy mocks base method
func (m *MockImageCloser) LayerInfosForCopy(arg0 context.Context) ([]types.BlobInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerInfosForCopy", arg0)
	ret0, _ := ret[0].([]types.BlobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayerInfosForCopy indicates an expected call of LayerInfosForCopy
func (mr *MockImageCloserMockRecorder) LayerInfosForCopy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerInfosForCopy", reflect.TypeOf((*MockImageCloser)(nil).LayerInfosForCopy), arg0)
}

// Manifest mocks base method
func (m *MockImageCloser) Manifest(arg0 context.Context) ([]byte, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifest", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Manifest indicates an expected call of Manifest
func (mr *MockImageCloserMockRecorder) Manifest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*MockImageCloser)(nil).Manifest), arg0)
}

// OCIConfig mocks base method
func (m *MockImageCloser) OCIConfig(arg0 context.Context) (*v1.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OCIConfig", arg0)
	ret0, _ := ret[0].(*v1.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OCIConfig indicates an expected call of OCIConfig
func (mr *MockImageCloserMockRecorder) OCIConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OCIConfig", reflect.TypeOf((*MockImageCloser)(nil).OCIConfig), arg0)
}

// Reference mocks base method
func (m *MockImageCloser) Reference() types.ImageReference {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reference")
	ret0, _ := ret[0].(types.ImageReference)
	return ret0
}

// Reference indicates an expected call of Reference
func (mr *MockImageCloserMockRecorder) Reference() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reference", reflect.TypeOf((*MockImageCloser)(nil).Reference))
}

// Signatures mocks base method
func (m *MockImageCloser) Signatures(arg0 context.Context) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signatures", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signatures indicates an expected call of Signatures
func (mr *MockImageCloserMockRecorder) Signatures(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signatures", reflect.TypeOf((*MockImageCloser)(nil).Signatures), arg0)
}

// Size mocks base method
func (m *MockImageCloser) Size() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Size indicates an expected call of Size
func (mr *MockImageCloserMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockImageCloser)(nil).Size))
}

// SupportsEncryption mocks base method
func (m *MockImageCloser) SupportsEncryption(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsEncryption", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SupportsEncryption indicates an expected call of SupportsEncryption
func (mr *MockImageCloserMockRecorder) SupportsEncryption(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsEncryption", reflect.TypeOf((*MockImageCloser)(nil).SupportsEncryption), arg0)
}

// UpdatedImage mocks base method
func (m *MockImageCloser) UpdatedImage(arg0 context.Context, arg1 types.ManifestUpdateOptions) (types.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatedImage", arg0, arg1)
	ret0, _ := ret[0].(types.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatedImage indicates an expected call of UpdatedImage
func (mr *MockImageCloserMockRecorder) UpdatedImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedImage", reflect.TypeOf((*MockImageCloser)(nil).UpdatedImage), arg0, arg1)
}

// UpdatedImageNeedsLayerDiffIDs mocks base method
func (m *MockImageCloser) UpdatedImageNeedsLayerDiffIDs(arg0 types.ManifestUpdateOptions) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatedImageNeedsLayerDiffIDs", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdatedImageNeedsLayerDiffIDs indicates an expected call of UpdatedImageNeedsLayerDiffIDs
func (mr *MockImageCloserMockRecorder) UpdatedImageNeedsLayerDiffIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedImageNeedsLayerDiffIDs", reflect.TypeOf((*MockImageCloser)(nil).UpdatedImageNeedsLayerDiffIDs), arg0)
}
