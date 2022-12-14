// Code generated by MockGen. DO NOT EDIT.
// Source: facade.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	leaderboard "github.com/calebtracey/rugby-models/pkg/dtos/leaderboard"
	gomock "github.com/golang/mock/gomock"
)

// MockAPIFacadeI is a mock of APIFacadeI interface.
type MockAPIFacadeI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIFacadeIMockRecorder
}

// MockAPIFacadeIMockRecorder is the mock recorder for MockAPIFacadeI.
type MockAPIFacadeIMockRecorder struct {
	mock *MockAPIFacadeI
}

// NewMockAPIFacadeI creates a new mock instance.
func NewMockAPIFacadeI(ctrl *gomock.Controller) *MockAPIFacadeI {
	mock := &MockAPIFacadeI{ctrl: ctrl}
	mock.recorder = &MockAPIFacadeIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIFacadeI) EXPECT() *MockAPIFacadeIMockRecorder {
	return m.recorder
}

// CrawlAllLeaderboardData mocks base method.
func (m *MockAPIFacadeI) CrawlAllLeaderboardData(ctx context.Context) leaderboard.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlAllLeaderboardData", ctx)
	ret0, _ := ret[0].(leaderboard.Response)
	return ret0
}

// CrawlAllLeaderboardData indicates an expected call of CrawlAllLeaderboardData.
func (mr *MockAPIFacadeIMockRecorder) CrawlAllLeaderboardData(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlAllLeaderboardData", reflect.TypeOf((*MockAPIFacadeI)(nil).CrawlAllLeaderboardData), ctx)
}

// CrawlLeaderboardData mocks base method.
func (m *MockAPIFacadeI) CrawlLeaderboardData(ctx context.Context, req leaderboard.Request) leaderboard.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlLeaderboardData", ctx, req)
	ret0, _ := ret[0].(leaderboard.Response)
	return ret0
}

// CrawlLeaderboardData indicates an expected call of CrawlLeaderboardData.
func (mr *MockAPIFacadeIMockRecorder) CrawlLeaderboardData(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlLeaderboardData", reflect.TypeOf((*MockAPIFacadeI)(nil).CrawlLeaderboardData), ctx, req)
}
