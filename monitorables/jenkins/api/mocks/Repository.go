// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/monitorables/jenkins/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetJob provides a mock function with given fields: jobName, branch
func (_m *Repository) GetJob(jobName string, branch string) (*models.Job, error) {
	ret := _m.Called(jobName, branch)

	var r0 *models.Job
	if rf, ok := ret.Get(0).(func(string, string) *models.Job); ok {
		r0 = rf(jobName, branch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(jobName, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBuildStatus provides a mock function with given fields: job
func (_m *Repository) GetLastBuildStatus(job *models.Job) (*models.Build, error) {
	ret := _m.Called(job)

	var r0 *models.Build
	if rf, ok := ret.Get(0).(func(*models.Job) *models.Build); ok {
		r0 = rf(job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Job) error); ok {
		r1 = rf(job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
