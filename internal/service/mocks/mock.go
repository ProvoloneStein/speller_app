// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "Nexign/internal/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSpeller is a mock of Speller interface.
type MockSpeller struct {
	ctrl     *gomock.Controller
	recorder *MockSpellerMockRecorder
}

// MockSpellerMockRecorder is the mock recorder for MockSpeller.
type MockSpellerMockRecorder struct {
	mock *MockSpeller
}

// NewMockSpeller creates a new mock instance.
func NewMockSpeller(ctrl *gomock.Controller) *MockSpeller {
	mock := &MockSpeller{ctrl: ctrl}
	mock.recorder = &MockSpellerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpeller) EXPECT() *MockSpellerMockRecorder {
	return m.recorder
}

// CreateMany mocks base method.
func (m *MockSpeller) CreateMany(ctx context.Context, texts model.Spellers) ([][]model.SpellerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMany", ctx, texts)
	ret0, _ := ret[0].([][]model.SpellerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMany indicates an expected call of CreateMany.
func (mr *MockSpellerMockRecorder) CreateMany(ctx, texts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMany", reflect.TypeOf((*MockSpeller)(nil).CreateMany), ctx, texts)
}

// CreateOne mocks base method.
func (m *MockSpeller) CreateOne(ctx context.Context, text model.Speller) ([]model.SpellerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOne", ctx, text)
	ret0, _ := ret[0].([]model.SpellerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOne indicates an expected call of CreateOne.
func (mr *MockSpellerMockRecorder) CreateOne(ctx, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOne", reflect.TypeOf((*MockSpeller)(nil).CreateOne), ctx, text)
}
