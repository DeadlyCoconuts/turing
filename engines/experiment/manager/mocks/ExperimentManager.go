// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	"github.com/gojek/turing/engines/experiment/manager"
	mock "github.com/stretchr/testify/mock"
)

// ExperimentManager is an autogenerated mock type for the ExperimentManager type
type ExperimentManager struct {
	mock.Mock
}

// GetEngineInfo provides a mock function with given fields:
func (_m *ExperimentManager) GetEngineInfo() manager.Engine {
	ret := _m.Called()

	var r0 manager.Engine
	if rf, ok := ret.Get(0).(func() manager.Engine); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(manager.Engine)
	}

	return r0
}
