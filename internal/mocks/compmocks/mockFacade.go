// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/calebtracey/rugby-crawler-api/internal/facade/comp (interfaces: FacadeI)

// Package compmocks is a generated GoMock package.
package compmocks

import (
	context "context"
	reflect "reflect"

	request "github.com/calebtracey/rugby-models/pkg/dtos/request"
	response "github.com/calebtracey/rugby-models/pkg/dtos/response"
	gomock "github.com/golang/mock/gomock"
)

// MockFacadeI is a mock of FacadeI interface.
type MockFacadeI struct {
	ctrl     *gomock.Controller
	recorder *MockFacadeIMockRecorder
}

// MockFacadeIMockRecorder is the mock recorder for MockFacadeI.
type MockFacadeIMockRecorder struct {
	mock *MockFacadeI
}

// NewMockFacadeI creates a new mock instance.
func NewMockFacadeI(ctrl *gomock.Controller) *MockFacadeI {
	mock := &MockFacadeI{ctrl: ctrl}
	mock.recorder = &MockFacadeIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacadeI) EXPECT() *MockFacadeIMockRecorder {
	return m.recorder
}

// CrawlAllLeaderboards mocks base method.
func (m *MockFacadeI) CrawlAllLeaderboards(arg0 context.Context) response.AllLeaderboardsResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlAllLeaderboards", arg0)
	ret0, _ := ret[0].(response.AllLeaderboardsResponse)
	return ret0
}

// CrawlAllLeaderboards indicates an expected call of CrawlAllLeaderboards.
func (mr *MockFacadeIMockRecorder) CrawlAllLeaderboards(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlAllLeaderboards", reflect.TypeOf((*MockFacadeI)(nil).CrawlAllLeaderboards), arg0)
}

// CrawlLeaderboard mocks base method.
func (m *MockFacadeI) CrawlLeaderboard(arg0 context.Context, arg1 request.LeaderboardRequest) response.LeaderboardResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlLeaderboard", arg0, arg1)
	ret0, _ := ret[0].(response.LeaderboardResponse)
	return ret0
}

// CrawlLeaderboard indicates an expected call of CrawlLeaderboard.
func (mr *MockFacadeIMockRecorder) CrawlLeaderboard(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlLeaderboard", reflect.TypeOf((*MockFacadeI)(nil).CrawlLeaderboard), arg0, arg1)
}
