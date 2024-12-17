// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/amir/GolangProjects/golang-template/internal/database/arango/connection.go

// Package mock_arango is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	arangodb "github.com/arangodb/go-driver/v2/arangodb"
	gomock "go.uber.org/mock/gomock"
)

// MockArangoDB is a mock of ArangoDB interface.
type MockArangoDB struct {
	ctrl     *gomock.Controller
	recorder *MockArangoDBMockRecorder
}

// MockArangoDBMockRecorder is the mock recorder for MockArangoDB.
type MockArangoDBMockRecorder struct {
	mock *MockArangoDB
}

// NewMockArangoDB creates a new mock instance.
func NewMockArangoDB(ctrl *gomock.Controller) *MockArangoDB {
	mock := &MockArangoDB{ctrl: ctrl}
	mock.recorder = &MockArangoDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArangoDB) EXPECT() *MockArangoDBMockRecorder {
	return m.recorder
}

// Database mocks base method.
func (m *MockArangoDB) Database(ctx context.Context) arangodb.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Database", ctx)
	ret0, _ := ret[0].(arangodb.Database)
	return ret0
}

// Database indicates an expected call of Database.
func (mr *MockArangoDBMockRecorder) Database(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Database", reflect.TypeOf((*MockArangoDB)(nil).Database), ctx)
}

// GetCollection mocks base method.
func (m *MockArangoDB) GetCollection(ctx context.Context, name string) (arangodb.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", ctx, name)
	ret0, _ := ret[0].(arangodb.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockArangoDBMockRecorder) GetCollection(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockArangoDB)(nil).GetCollection), ctx, name)
}

// Ping mocks base method.
func (m *MockArangoDB) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockArangoDBMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockArangoDB)(nil).Ping), ctx)
}
