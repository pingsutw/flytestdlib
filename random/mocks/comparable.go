// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import random "github.com/lyft/flytestdlib/random"

// Comparable is an autogenerated mock type for the Comparable type
type Comparable struct {
	mock.Mock
}

// Compare provides a mock function with given fields: to
func (_m *Comparable) Compare(to random.Comparable) bool {
	ret := _m.Called(to)

	var r0 bool
	if rf, ok := ret.Get(0).(func(random.Comparable) bool); ok {
		r0 = rf(to)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
