// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/video_service.go
//
// Generated by this command:
//
//	mockgen -source=internal/services/video_service.go -destination=internal/mocks/video_service.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	dao "golang_template/handler/daos"
	dto "golang_template/handler/dtos"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockVideoService is a mock of VideoService interface.
type MockVideoService struct {
	ctrl     *gomock.Controller
	recorder *MockVideoServiceMockRecorder
	isgomock struct{}
}

// MockVideoServiceMockRecorder is the mock recorder for MockVideoService.
type MockVideoServiceMockRecorder struct {
	mock *MockVideoService
}

// NewMockVideoService creates a new mock instance.
func NewMockVideoService(ctrl *gomock.Controller) *MockVideoService {
	mock := &MockVideoService{ctrl: ctrl}
	mock.recorder = &MockVideoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVideoService) EXPECT() *MockVideoServiceMockRecorder {
	return m.recorder
}

// CreateVideo mocks base method.
func (m *MockVideoService) CreateVideo(video dto.Video) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVideo", video)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVideo indicates an expected call of CreateVideo.
func (mr *MockVideoServiceMockRecorder) CreateVideo(video any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVideo", reflect.TypeOf((*MockVideoService)(nil).CreateVideo), video)
}

// DeleteVideo mocks base method.
func (m *MockVideoService) DeleteVideo(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVideo", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVideo indicates an expected call of DeleteVideo.
func (mr *MockVideoServiceMockRecorder) DeleteVideo(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVideo", reflect.TypeOf((*MockVideoService)(nil).DeleteVideo), key)
}

// GetVideo mocks base method.
func (m *MockVideoService) GetVideo(key string) (*dao.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideo", key)
	ret0, _ := ret[0].(*dao.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideo indicates an expected call of GetVideo.
func (mr *MockVideoServiceMockRecorder) GetVideo(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideo", reflect.TypeOf((*MockVideoService)(nil).GetVideo), key)
}

// GetVideoByName mocks base method.
func (m *MockVideoService) GetVideoByName(name string) (*dao.VideoByName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideoByName", name)
	ret0, _ := ret[0].(*dao.VideoByName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideoByName indicates an expected call of GetVideoByName.
func (mr *MockVideoServiceMockRecorder) GetVideoByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideoByName", reflect.TypeOf((*MockVideoService)(nil).GetVideoByName), name)
}

// UpdateVideo mocks base method.
func (m *MockVideoService) UpdateVideo(videoUpdate dto.VideoUpdate) (*dao.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVideo", videoUpdate)
	ret0, _ := ret[0].(*dao.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVideo indicates an expected call of UpdateVideo.
func (mr *MockVideoServiceMockRecorder) UpdateVideo(videoUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVideo", reflect.TypeOf((*MockVideoService)(nil).UpdateVideo), videoUpdate)
}
