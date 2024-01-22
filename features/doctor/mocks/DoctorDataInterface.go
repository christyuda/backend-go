// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	doctor "FinalProject/features/doctor"

	mock "github.com/stretchr/testify/mock"
)

// DoctorDataInterface is an autogenerated mock type for the DoctorDataInterface type
type DoctorDataInterface struct {
	mock.Mock
}

// ApproveDoctor provides a mock function with given fields: idDoctor
func (_m *DoctorDataInterface) ApproveDoctor(idDoctor int) (bool, error) {
	ret := _m.Called(idDoctor)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(idDoctor)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(idDoctor)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idDoctor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDoctor provides a mock function with given fields: doctorID
func (_m *DoctorDataInterface) DeleteDoctor(doctorID int) (bool, error) {
	ret := _m.Called(doctorID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(doctorID)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(doctorID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(doctorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDoctorEducation provides a mock function with given fields: doctorID
func (_m *DoctorDataInterface) DeleteDoctorEducation(doctorID int) (bool, error) {
	ret := _m.Called(doctorID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(doctorID)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(doctorID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(doctorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDoctorExperience provides a mock function with given fields: doctorID
func (_m *DoctorDataInterface) DeleteDoctorExperience(doctorID int) (bool, error) {
	ret := _m.Called(doctorID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(doctorID)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(doctorID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(doctorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDoctorRating provides a mock function with given fields: doctorID
func (_m *DoctorDataInterface) DeleteDoctorRating(doctorID int) (bool, error) {
	ret := _m.Called(doctorID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(doctorID)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(doctorID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(doctorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDoctorWorkdays provides a mock function with given fields: doctorID
func (_m *DoctorDataInterface) DeleteDoctorWorkdays(doctorID int) (bool, error) {
	ret := _m.Called(doctorID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(doctorID)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(doctorID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(doctorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DenyDoctor provides a mock function with given fields: idDoctor
func (_m *DoctorDataInterface) DenyDoctor(idDoctor int) (bool, error) {
	ret := _m.Called(idDoctor)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(idDoctor)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(idDoctor)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idDoctor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoctorDashboard provides a mock function with given fields: id
func (_m *DoctorDataInterface) DoctorDashboard(id int) (doctor.DoctorDashboard, error) {
	ret := _m.Called(id)

	var r0 doctor.DoctorDashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (doctor.DoctorDashboard, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) doctor.DoctorDashboard); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(doctor.DoctorDashboard)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoctorDashboardAdmin provides a mock function with given fields:
func (_m *DoctorDataInterface) DoctorDashboardAdmin() (doctor.DoctorDashboardAdmin, error) {
	ret := _m.Called()

	var r0 doctor.DoctorDashboardAdmin
	var r1 error
	if rf, ok := ret.Get(0).(func() (doctor.DoctorDashboardAdmin, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() doctor.DoctorDashboardAdmin); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(doctor.DoctorDashboardAdmin)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoctorDashboardPatient provides a mock function with given fields: id
func (_m *DoctorDataInterface) DoctorDashboardPatient(id int) ([]doctor.DoctorDashboardPatient, error) {
	ret := _m.Called(id)

	var r0 []doctor.DoctorDashboardPatient
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]doctor.DoctorDashboardPatient, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []doctor.DoctorDashboardPatient); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctor.DoctorDashboardPatient)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEmail provides a mock function with given fields: userID
func (_m *DoctorDataInterface) FindEmail(userID uint) (*string, error) {
	ret := _m.Called(userID)

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) *string); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: name
func (_m *DoctorDataInterface) GetAll(name string) ([]doctor.DoctorAll, error) {
	ret := _m.Called(name)

	var r0 []doctor.DoctorAll
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]doctor.DoctorAll, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) []doctor.DoctorAll); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctor.DoctorAll)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *DoctorDataInterface) GetByID(id int) (*doctor.DoctorAll, error) {
	ret := _m.Called(id)

	var r0 *doctor.DoctorAll
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*doctor.DoctorAll, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *doctor.DoctorAll); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.DoctorAll)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDEducation provides a mock function with given fields: id
func (_m *DoctorDataInterface) GetByIDEducation(id int) ([]doctor.DoctorEducation, error) {
	ret := _m.Called(id)

	var r0 []doctor.DoctorEducation
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]doctor.DoctorEducation, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []doctor.DoctorEducation); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctor.DoctorEducation)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDExperience provides a mock function with given fields: id
func (_m *DoctorDataInterface) GetByIDExperience(id int) ([]doctor.DoctorExperience, error) {
	ret := _m.Called(id)

	var r0 []doctor.DoctorExperience
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]doctor.DoctorExperience, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []doctor.DoctorExperience); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctor.DoctorExperience)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIDWorkadays provides a mock function with given fields: id
func (_m *DoctorDataInterface) GetByIDWorkadays(id int) ([]doctor.DoctorWorkdays, error) {
	ret := _m.Called(id)

	var r0 []doctor.DoctorWorkdays
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]doctor.DoctorWorkdays, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) []doctor.DoctorWorkdays); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctor.DoctorWorkdays)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDoctorByUserId provides a mock function with given fields: userID
func (_m *DoctorDataInterface) GetDoctorByUserId(userID int) (*doctor.DoctorAll, error) {
	ret := _m.Called(userID)

	var r0 *doctor.DoctorAll
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*doctor.DoctorAll, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) *doctor.DoctorAll); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.DoctorAll)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newData
func (_m *DoctorDataInterface) Insert(newData doctor.Doctor) (*doctor.Doctor, error) {
	ret := _m.Called(newData)

	var r0 *doctor.Doctor
	var r1 error
	if rf, ok := ret.Get(0).(func(doctor.Doctor) (*doctor.Doctor, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(doctor.Doctor) *doctor.Doctor); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.Doctor)
		}
	}

	if rf, ok := ret.Get(1).(func(doctor.Doctor) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertEducation provides a mock function with given fields: newData
func (_m *DoctorDataInterface) InsertEducation(newData doctor.DoctorEducation) (*doctor.DoctorEducation, error) {
	ret := _m.Called(newData)

	var r0 *doctor.DoctorEducation
	var r1 error
	if rf, ok := ret.Get(0).(func(doctor.DoctorEducation) (*doctor.DoctorEducation, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(doctor.DoctorEducation) *doctor.DoctorEducation); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.DoctorEducation)
		}
	}

	if rf, ok := ret.Get(1).(func(doctor.DoctorEducation) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertExperience provides a mock function with given fields: newData
func (_m *DoctorDataInterface) InsertExperience(newData doctor.DoctorExperience) (*doctor.DoctorExperience, error) {
	ret := _m.Called(newData)

	var r0 *doctor.DoctorExperience
	var r1 error
	if rf, ok := ret.Get(0).(func(doctor.DoctorExperience) (*doctor.DoctorExperience, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(doctor.DoctorExperience) *doctor.DoctorExperience); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.DoctorExperience)
		}
	}

	if rf, ok := ret.Get(1).(func(doctor.DoctorExperience) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertExpertise provides a mock function with given fields: newData
func (_m *DoctorDataInterface) InsertExpertise(newData doctor.DoctorExpertiseRelation) (*doctor.DoctorExpertiseRelation, error) {
	ret := _m.Called(newData)

	var r0 *doctor.DoctorExpertiseRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(doctor.DoctorExpertiseRelation) (*doctor.DoctorExpertiseRelation, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(doctor.DoctorExpertiseRelation) *doctor.DoctorExpertiseRelation); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.DoctorExpertiseRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(doctor.DoctorExpertiseRelation) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertWorkadays provides a mock function with given fields: newData
func (_m *DoctorDataInterface) InsertWorkadays(newData doctor.DoctorWorkdays) (*doctor.DoctorWorkdays, error) {
	ret := _m.Called(newData)

	var r0 *doctor.DoctorWorkdays
	var r1 error
	if rf, ok := ret.Get(0).(func(doctor.DoctorWorkdays) (*doctor.DoctorWorkdays, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(doctor.DoctorWorkdays) *doctor.DoctorWorkdays); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.DoctorWorkdays)
		}
	}

	if rf, ok := ret.Get(1).(func(doctor.DoctorWorkdays) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsLinkUsed provides a mock function with given fields: meetLink
func (_m *DoctorDataInterface) IsLinkUsed(meetLink string) bool {
	ret := _m.Called(meetLink)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(meetLink)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UpdateDoctorDatapokok provides a mock function with given fields: id, newData
func (_m *DoctorDataInterface) UpdateDoctorDatapokok(id int, newData doctor.DoctorDatapokokUpdate) (bool, error) {
	ret := _m.Called(id, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, doctor.DoctorDatapokokUpdate) (bool, error)); ok {
		return rf(id, newData)
	}
	if rf, ok := ret.Get(0).(func(int, doctor.DoctorDatapokokUpdate) bool); ok {
		r0 = rf(id, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, doctor.DoctorDatapokokUpdate) error); ok {
		r1 = rf(id, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDoctorEducation provides a mock function with given fields: id, doctorID, newData
func (_m *DoctorDataInterface) UpdateDoctorEducation(id int, doctorID int, newData doctor.DoctorEducation) (bool, error) {
	ret := _m.Called(id, doctorID, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorEducation) (bool, error)); ok {
		return rf(id, doctorID, newData)
	}
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorEducation) bool); ok {
		r0 = rf(id, doctorID, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, int, doctor.DoctorEducation) error); ok {
		r1 = rf(id, doctorID, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDoctorExperience provides a mock function with given fields: id, doctorID, newData
func (_m *DoctorDataInterface) UpdateDoctorExperience(id int, doctorID int, newData doctor.DoctorExperience) (bool, error) {
	ret := _m.Called(id, doctorID, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorExperience) (bool, error)); ok {
		return rf(id, doctorID, newData)
	}
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorExperience) bool); ok {
		r0 = rf(id, doctorID, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, int, doctor.DoctorExperience) error); ok {
		r1 = rf(id, doctorID, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDoctorRating provides a mock function with given fields: id, pateintID, newData
func (_m *DoctorDataInterface) UpdateDoctorRating(id int, pateintID int, newData doctor.DoctorRating) (bool, error) {
	ret := _m.Called(id, pateintID, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorRating) (bool, error)); ok {
		return rf(id, pateintID, newData)
	}
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorRating) bool); ok {
		r0 = rf(id, pateintID, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, int, doctor.DoctorRating) error); ok {
		r1 = rf(id, pateintID, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDoctorWorkdays provides a mock function with given fields: id, doctorID, newData
func (_m *DoctorDataInterface) UpdateDoctorWorkdays(id int, doctorID int, newData doctor.DoctorWorkdays) (bool, error) {
	ret := _m.Called(id, doctorID, newData)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorWorkdays) (bool, error)); ok {
		return rf(id, doctorID, newData)
	}
	if rf, ok := ret.Get(0).(func(int, int, doctor.DoctorWorkdays) bool); ok {
		r0 = rf(id, doctorID, newData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, int, doctor.DoctorWorkdays) error); ok {
		r1 = rf(id, doctorID, newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDoctorDataInterface creates a new instance of DoctorDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDoctorDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DoctorDataInterface {
	mock := &DoctorDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}