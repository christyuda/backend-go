// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	bundlecounseling "FinalProject/features/bundle_counseling"

	mock "github.com/stretchr/testify/mock"
)

// BundleCounselingDataInterface is an autogenerated mock type for the BundleCounselingDataInterface type
type BundleCounselingDataInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *BundleCounselingDataInterface) Create(input bundlecounseling.BundleCounseling) (*bundlecounseling.BundleCounseling, error) {
	ret := _m.Called(input)

	var r0 *bundlecounseling.BundleCounseling
	var r1 error
	if rf, ok := ret.Get(0).(func(bundlecounseling.BundleCounseling) (*bundlecounseling.BundleCounseling, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(bundlecounseling.BundleCounseling) *bundlecounseling.BundleCounseling); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bundlecounseling.BundleCounseling)
		}
	}

	if rf, ok := ret.Get(1).(func(bundlecounseling.BundleCounseling) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *BundleCounselingDataInterface) Delete(id int) (bool, error) {
	ret := _m.Called(id)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *BundleCounselingDataInterface) GetAll() ([]bundlecounseling.BundleCounselingInfo, error) {
	ret := _m.Called()

	var r0 []bundlecounseling.BundleCounselingInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]bundlecounseling.BundleCounselingInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []bundlecounseling.BundleCounselingInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bundlecounseling.BundleCounselingInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllFilter provides a mock function with given fields: jenis
func (_m *BundleCounselingDataInterface) GetAllFilter(jenis string) ([]bundlecounseling.BundleCounselingInfo, error) {
	ret := _m.Called(jenis)

	var r0 []bundlecounseling.BundleCounselingInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]bundlecounseling.BundleCounselingInfo, error)); ok {
		return rf(jenis)
	}
	if rf, ok := ret.Get(0).(func(string) []bundlecounseling.BundleCounselingInfo); ok {
		r0 = rf(jenis)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bundlecounseling.BundleCounselingInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jenis)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *BundleCounselingDataInterface) GetById(id int) (*bundlecounseling.BundleCounseling, error) {
	ret := _m.Called(id)

	var r0 *bundlecounseling.BundleCounseling
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*bundlecounseling.BundleCounseling, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *bundlecounseling.BundleCounseling); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bundlecounseling.BundleCounseling)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HargaDurasi provides a mock function with given fields: id
func (_m *BundleCounselingDataInterface) HargaDurasi(id int) (uint, error) {
	ret := _m.Called(id)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (uint, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) uint); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HargaMetode provides a mock function with given fields: id
func (_m *BundleCounselingDataInterface) HargaMetode(id int) (uint, error) {
	ret := _m.Called(id)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (uint, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) uint); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, newData
func (_m *BundleCounselingDataInterface) Update(id int, newData bundlecounseling.BundleCounseling) (bool, error) {
	ret := _m.Called(id, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, bundlecounseling.BundleCounseling) (bool, error)); ok {
		return rf(id, newData)
	}
	if rf, ok := ret.Get(0).(func(int, bundlecounseling.BundleCounseling) bool); ok {
		r0 = rf(id, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, bundlecounseling.BundleCounseling) error); ok {
		r1 = rf(id, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBundleCounselingDataInterface creates a new instance of BundleCounselingDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBundleCounselingDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BundleCounselingDataInterface {
	mock := &BundleCounselingDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}