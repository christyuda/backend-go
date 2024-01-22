// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	counselingmethods "FinalProject/features/counseling_methods"

	mock "github.com/stretchr/testify/mock"
)

// CounselingMethodDataInterface is an autogenerated mock type for the CounselingMethodDataInterface type
type CounselingMethodDataInterface struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *CounselingMethodDataInterface) GetAll() ([]counselingmethods.CounselingMethodInfo, error) {
	ret := _m.Called()

	var r0 []counselingmethods.CounselingMethodInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]counselingmethods.CounselingMethodInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []counselingmethods.CounselingMethodInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]counselingmethods.CounselingMethodInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *CounselingMethodDataInterface) GetByID(id int) ([]counselingmethods.CounselingMethodInfo, error) {
	ret := _m.Called(id)

	var r0 []counselingmethods.CounselingMethodInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]counselingmethods.CounselingMethodInfo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []counselingmethods.CounselingMethodInfo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]counselingmethods.CounselingMethodInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCounselingMethodDataInterface creates a new instance of CounselingMethodDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCounselingMethodDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CounselingMethodDataInterface {
	mock := &CounselingMethodDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}