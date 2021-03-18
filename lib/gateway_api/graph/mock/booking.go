// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cobbinma/booking-platform/lib/protobuf/autogen/lang/go/booking/api (interfaces: BookingAPIClient)

// Package mock_resolver is a generated GoMock package.
package mock_resolver

import (
	context "context"
	reflect "reflect"

	api "github.com/cobbinma/booking-platform/lib/protobuf/autogen/lang/go/booking/api"
	models "github.com/cobbinma/booking-platform/lib/protobuf/autogen/lang/go/booking/models"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBookingAPIClient is a mock of BookingAPIClient interface.
type MockBookingAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockBookingAPIClientMockRecorder
}

// MockBookingAPIClientMockRecorder is the mock recorder for MockBookingAPIClient.
type MockBookingAPIClientMockRecorder struct {
	mock *MockBookingAPIClient
}

// NewMockBookingAPIClient creates a new mock instance.
func NewMockBookingAPIClient(ctrl *gomock.Controller) *MockBookingAPIClient {
	mock := &MockBookingAPIClient{ctrl: ctrl}
	mock.recorder = &MockBookingAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookingAPIClient) EXPECT() *MockBookingAPIClientMockRecorder {
	return m.recorder
}

// CancelBooking mocks base method.
func (m *MockBookingAPIClient) CancelBooking(arg0 context.Context, arg1 *api.CancelBookingRequest, arg2 ...grpc.CallOption) (*models.Booking, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelBooking", varargs...)
	ret0, _ := ret[0].(*models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelBooking indicates an expected call of CancelBooking.
func (mr *MockBookingAPIClientMockRecorder) CancelBooking(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelBooking", reflect.TypeOf((*MockBookingAPIClient)(nil).CancelBooking), varargs...)
}

// CreateBooking mocks base method.
func (m *MockBookingAPIClient) CreateBooking(arg0 context.Context, arg1 *models.SlotInput, arg2 ...grpc.CallOption) (*models.Booking, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBooking", varargs...)
	ret0, _ := ret[0].(*models.Booking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBooking indicates an expected call of CreateBooking.
func (mr *MockBookingAPIClientMockRecorder) CreateBooking(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBooking", reflect.TypeOf((*MockBookingAPIClient)(nil).CreateBooking), varargs...)
}

// GetBookings mocks base method.
func (m *MockBookingAPIClient) GetBookings(arg0 context.Context, arg1 *api.GetBookingsRequest, arg2 ...grpc.CallOption) (*api.GetBookingsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBookings", varargs...)
	ret0, _ := ret[0].(*api.GetBookingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookings indicates an expected call of GetBookings.
func (mr *MockBookingAPIClientMockRecorder) GetBookings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookings", reflect.TypeOf((*MockBookingAPIClient)(nil).GetBookings), varargs...)
}

// GetSlot mocks base method.
func (m *MockBookingAPIClient) GetSlot(arg0 context.Context, arg1 *models.SlotInput, arg2 ...grpc.CallOption) (*api.GetSlotResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSlot", varargs...)
	ret0, _ := ret[0].(*api.GetSlotResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSlot indicates an expected call of GetSlot.
func (mr *MockBookingAPIClientMockRecorder) GetSlot(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSlot", reflect.TypeOf((*MockBookingAPIClient)(nil).GetSlot), varargs...)
}