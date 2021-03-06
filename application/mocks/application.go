// Code generated by MockGen. DO NOT EDIT.
// Source: application/product.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	application "github.com/stamatogabriel/go-hexagonal/application"
)

// MockProductInterface is a mock of ProductInterface interface.
type MockProductInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductInterfaceMockRecorder
}

// MockProductInterfaceMockRecorder is the mock recorder for MockProductInterface.
type MockProductInterfaceMockRecorder struct {
	mock *MockProductInterface
}

// NewMockProductInterface creates a new mock instance.
func NewMockProductInterface(ctrl *gomock.Controller) *MockProductInterface {
	mock := &MockProductInterface{ctrl: ctrl}
	mock.recorder = &MockProductInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductInterface) EXPECT() *MockProductInterfaceMockRecorder {
	return m.recorder
}

// Disable mocks base method.
func (m *MockProductInterface) Disable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockProductInterfaceMockRecorder) Disable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockProductInterface)(nil).Disable))
}

// Enable mocks base method.
func (m *MockProductInterface) Enable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockProductInterfaceMockRecorder) Enable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockProductInterface)(nil).Enable))
}

// GetId mocks base method.
func (m *MockProductInterface) GetId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockProductInterfaceMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockProductInterface)(nil).GetId))
}

// GetName mocks base method.
func (m *MockProductInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockProductInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockProductInterface)(nil).GetName))
}

// GetPrice mocks base method.
func (m *MockProductInterface) GetPrice() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockProductInterfaceMockRecorder) GetPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockProductInterface)(nil).GetPrice))
}

// GetStatus mocks base method.
func (m *MockProductInterface) GetStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockProductInterfaceMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockProductInterface)(nil).GetStatus))
}

// IsValid mocks base method.
func (m *MockProductInterface) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockProductInterfaceMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockProductInterface)(nil).IsValid))
}

// MockProductServiceInterface is a mock of ProductServiceInterface interface.
type MockProductServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceInterfaceMockRecorder
}

// MockProductServiceInterfaceMockRecorder is the mock recorder for MockProductServiceInterface.
type MockProductServiceInterfaceMockRecorder struct {
	mock *MockProductServiceInterface
}

// NewMockProductServiceInterface creates a new mock instance.
func NewMockProductServiceInterface(ctrl *gomock.Controller) *MockProductServiceInterface {
	mock := &MockProductServiceInterface{ctrl: ctrl}
	mock.recorder = &MockProductServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceInterface) EXPECT() *MockProductServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductServiceInterface) Create(name string, price float64) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceInterfaceMockRecorder) Create(name, price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductServiceInterface)(nil).Create), name, price)
}

// Disabled mocks base method.
func (m *MockProductServiceInterface) Disabled(arg0 application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disabled", arg0)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disabled indicates an expected call of Disabled.
func (mr *MockProductServiceInterfaceMockRecorder) Disabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disabled", reflect.TypeOf((*MockProductServiceInterface)(nil).Disabled), arg0)
}

// Enabled mocks base method.
func (m *MockProductServiceInterface) Enabled(arg0 application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled", arg0)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enabled indicates an expected call of Enabled.
func (mr *MockProductServiceInterfaceMockRecorder) Enabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockProductServiceInterface)(nil).Enabled), arg0)
}

// Get mocks base method.
func (m *MockProductServiceInterface) Get(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductServiceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductServiceInterface)(nil).Get), id)
}

// MockProductReaderInterface is a mock of ProductReaderInterface interface.
type MockProductReaderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductReaderInterfaceMockRecorder
}

// MockProductReaderInterfaceMockRecorder is the mock recorder for MockProductReaderInterface.
type MockProductReaderInterfaceMockRecorder struct {
	mock *MockProductReaderInterface
}

// NewMockProductReaderInterface creates a new mock instance.
func NewMockProductReaderInterface(ctrl *gomock.Controller) *MockProductReaderInterface {
	mock := &MockProductReaderInterface{ctrl: ctrl}
	mock.recorder = &MockProductReaderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductReaderInterface) EXPECT() *MockProductReaderInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProductReaderInterface) Get(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductReaderInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductReaderInterface)(nil).Get), id)
}

// MockProductWritterInterface is a mock of ProductWritterInterface interface.
type MockProductWritterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductWritterInterfaceMockRecorder
}

// MockProductWritterInterfaceMockRecorder is the mock recorder for MockProductWritterInterface.
type MockProductWritterInterfaceMockRecorder struct {
	mock *MockProductWritterInterface
}

// NewMockProductWritterInterface creates a new mock instance.
func NewMockProductWritterInterface(ctrl *gomock.Controller) *MockProductWritterInterface {
	mock := &MockProductWritterInterface{ctrl: ctrl}
	mock.recorder = &MockProductWritterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductWritterInterface) EXPECT() *MockProductWritterInterfaceMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockProductWritterInterface) Save(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductWritterInterfaceMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductWritterInterface)(nil).Save), product)
}

// MockProductPersistenceInterface is a mock of ProductPersistenceInterface interface.
type MockProductPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductPersistenceInterfaceMockRecorder
}

// MockProductPersistenceInterfaceMockRecorder is the mock recorder for MockProductPersistenceInterface.
type MockProductPersistenceInterfaceMockRecorder struct {
	mock *MockProductPersistenceInterface
}

// NewMockProductPersistenceInterface creates a new mock instance.
func NewMockProductPersistenceInterface(ctrl *gomock.Controller) *MockProductPersistenceInterface {
	mock := &MockProductPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockProductPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductPersistenceInterface) EXPECT() *MockProductPersistenceInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProductPersistenceInterface) Get(id string) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductPersistenceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductPersistenceInterface)(nil).Get), id)
}

// Save mocks base method.
func (m *MockProductPersistenceInterface) Save(product application.ProductInterface) (application.ProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.ProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockProductPersistenceInterfaceMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductPersistenceInterface)(nil).Save), product)
}
