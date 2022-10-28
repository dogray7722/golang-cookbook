// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	recipes_db "github.com/golang-cookbook/datasources/postgres/recipes_db/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateRecipe mocks base method.
func (m *MockStore) CreateRecipe(arg0 context.Context, arg1 recipes_db.CreateRecipeParams) (recipes_db.Recipe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecipe", arg0, arg1)
	ret0, _ := ret[0].(recipes_db.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRecipe indicates an expected call of CreateRecipe.
func (mr *MockStoreMockRecorder) CreateRecipe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecipe", reflect.TypeOf((*MockStore)(nil).CreateRecipe), arg0, arg1)
}

// DeleteRecipe mocks base method.
func (m *MockStore) DeleteRecipe(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecipe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecipe indicates an expected call of DeleteRecipe.
func (mr *MockStoreMockRecorder) DeleteRecipe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecipe", reflect.TypeOf((*MockStore)(nil).DeleteRecipe), arg0, arg1)
}

// GetRecipe mocks base method.
func (m *MockStore) GetRecipe(arg0 context.Context, arg1 int32) (recipes_db.Recipe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecipe", arg0, arg1)
	ret0, _ := ret[0].(recipes_db.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecipe indicates an expected call of GetRecipe.
func (mr *MockStoreMockRecorder) GetRecipe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipe", reflect.TypeOf((*MockStore)(nil).GetRecipe), arg0, arg1)
}

// ListRecipes mocks base method.
func (m *MockStore) ListRecipes(arg0 context.Context, arg1 recipes_db.ListRecipesParams) ([]recipes_db.Recipe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecipes", arg0, arg1)
	ret0, _ := ret[0].([]recipes_db.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecipes indicates an expected call of ListRecipes.
func (mr *MockStoreMockRecorder) ListRecipes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecipes", reflect.TypeOf((*MockStore)(nil).ListRecipes), arg0, arg1)
}

// UpdateRecipe mocks base method.
func (m *MockStore) UpdateRecipe(arg0 context.Context, arg1 recipes_db.UpdateRecipeParams) (recipes_db.Recipe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRecipe", arg0, arg1)
	ret0, _ := ret[0].(recipes_db.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRecipe indicates an expected call of UpdateRecipe.
func (mr *MockStoreMockRecorder) UpdateRecipe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRecipe", reflect.TypeOf((*MockStore)(nil).UpdateRecipe), arg0, arg1)
}
