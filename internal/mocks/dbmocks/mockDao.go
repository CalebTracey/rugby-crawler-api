// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/calebtracey/rugby-crawler-api/internal/dao/psql (interfaces: DAOI)

// Package dbmocks is a generated GoMock package.
package dbmocks

import (
	context "context"
	sql "database/sql"
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

// InsertOne mocks base method.
func (m *MockDAOI) InsertOne(arg0 context.Context, arg1 string) (sql.Result, *response.ErrorLog) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", arg0, arg1)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(*response.ErrorLog)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockDAOIMockRecorder) InsertOne(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockDAOI)(nil).InsertOne), arg0, arg1)
}
