// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ExecDWriter is an autogenerated mock type for the ExecDWriter type
type ExecDWriter struct {
	mock.Mock
}

// Write provides a mock function with given fields: value
func (_m *ExecDWriter) Write(value map[string]string) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(map[string]string) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
