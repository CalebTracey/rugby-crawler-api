// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/calebtracey/rugby-crawler-api/internal/facade (interfaces: APIFacadeI)

// Package facade is a generated GoMock package.
package facade

import (
	context "context"
	reflect "reflect"

	request "github.com/calebtracey/rugby-crawler-api/external/models/request"
	response "github.com/calebtracey/rugby-crawler-api/external/models/response"
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

// CompetitionCrawlData mocks base method.
func (m *MockAPIFacadeI) CompetitionCrawlData(arg0 context.Context, arg1 request.CrawlLeaderboardRequest) response.CrawlLeaderboardResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompetitionCrawlData", arg0, arg1)
	ret0, _ := ret[0].(response.CrawlLeaderboardResponse)
	return ret0
}

// CompetitionCrawlData indicates an expected call of CompetitionCrawlData.
func (mr *MockAPIFacadeIMockRecorder) CompetitionCrawlData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompetitionCrawlData", reflect.TypeOf((*MockAPIFacadeI)(nil).CompetitionCrawlData), arg0, arg1)
}
