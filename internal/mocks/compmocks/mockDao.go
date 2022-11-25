// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/calebtracey/rugby-crawler-api/internal/dao/comp (interfaces: DAOI)

// Package compmocks is a generated GoMock package.
package compmocks

import (
	response "github.com/calebtracey/rugby-models/pkg/dtos/response"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDAOI is a mock of DAOI interface.
type MockDAOI struct {
	ctrl     *gomock.Controller
	recorder *MockDAOIMockRecorder
}

// MockDAOIMockRecorder is the mock recorder for MockDAOI.
type MockDAOIMockRecorder struct {
	mock *MockDAOI
}

// NewMockDAOI creates a new mock instance.
func NewMockDAOI(ctrl *gomock.Controller) *MockDAOI {
	mock := &MockDAOI{ctrl: ctrl}
	mock.recorder = &MockDAOIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDAOI) EXPECT() *MockDAOIMockRecorder {
	return m.recorder
}

// CrawlLeaderboardData mocks base method.
func (m *MockDAOI) CrawlLeaderboardData(arg0 string) (response.LeaderboardResponse, *response.ErrorLog) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlLeaderboardData", arg0)
	ret0, _ := ret[0].(response.LeaderboardResponse)
	ret1, _ := ret[1].(*response.ErrorLog)
	return ret0, ret1
}

// CrawlLeaderboardData indicates an expected call of CrawlLeaderboardData.
func (mr *MockDAOIMockRecorder) CrawlLeaderboardData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlLeaderboardData", reflect.TypeOf((*MockDAOI)(nil).CrawlLeaderboardData), arg0)
}
