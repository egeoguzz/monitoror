// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	options "github.com/monitoror/monitoror/service/options"
	mock "github.com/stretchr/testify/mock"
)

// RouterOption is an autogenerated mock type for the RouterOption type
type RouterOption struct {
	mock.Mock
}

// Apply provides a mock function with given fields: s
func (_m *RouterOption) Apply(s *options.RouterSettings) {
	_m.Called(s)
}
