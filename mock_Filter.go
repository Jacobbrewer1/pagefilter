// Code generated by mockery. DO NOT EDIT.

package pagefilter

import mock "github.com/stretchr/testify/mock"

// MockFilter is an autogenerated mock type for the Filter type
type MockFilter struct {
	mock.Mock
}

// Join provides a mock function with no fields
func (_m *MockFilter) Join() (string, []interface{}) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Join")
	}

	var r0 string
	var r1 []interface{}
	if rf, ok := ret.Get(0).(func() (string, []interface{})); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() []interface{}); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	return r0, r1
}

// Where provides a mock function with no fields
func (_m *MockFilter) Where() (string, []interface{}) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Where")
	}

	var r0 string
	var r1 []interface{}
	if rf, ok := ret.Get(0).(func() (string, []interface{})); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() []interface{}); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	return r0, r1
}

// NewMockFilter creates a new instance of MockFilter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFilter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFilter {
	mock := &MockFilter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
