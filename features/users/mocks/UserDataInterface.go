// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	users "FinalProject/features/users"

	mock "github.com/stretchr/testify/mock"
)

// UserDataInterface is an autogenerated mock type for the UserDataInterface type
type UserDataInterface struct {
	mock.Mock
}

// DeleteCode provides a mock function with given fields: email
func (_m *UserDataInterface) DeleteCode(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByCode provides a mock function with given fields: code
func (_m *UserDataInterface) GetByCode(code string) (*users.UserResetPass, error) {
	ret := _m.Called(code)

	var r0 *users.UserResetPass
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*users.UserResetPass, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *users.UserResetPass); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UserResetPass)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserDataInterface) GetByEmail(email string) (*users.User, error) {
	ret := _m.Called(email)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*users.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *users.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertCode provides a mock function with given fields: email, code
func (_m *UserDataInterface) InsertCode(email string, code string) error {
	ret := _m.Called(email, code)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *UserDataInterface) Login(email string, password string) (*users.User, error) {
	ret := _m.Called(email, password)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*users.User, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *users.User); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newData
func (_m *UserDataInterface) Register(newData users.User) (*users.User, error) {
	ret := _m.Called(newData)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(users.User) (*users.User, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(users.User) *users.User); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(users.User) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: code, email, password
func (_m *UserDataInterface) ResetPassword(code string, email string, password string) error {
	ret := _m.Called(code, email, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(code, email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfile provides a mock function with given fields: id, newData
func (_m *UserDataInterface) UpdateProfile(id int, newData users.UpdateProfile) (bool, error) {
	ret := _m.Called(id, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, users.UpdateProfile) (bool, error)); ok {
		return rf(id, newData)
	}
	if rf, ok := ret.Get(0).(func(int, users.UpdateProfile) bool); ok {
		r0 = rf(id, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, users.UpdateProfile) error); ok {
		r1 = rf(id, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserDataInterface creates a new instance of UserDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserDataInterface {
	mock := &UserDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}