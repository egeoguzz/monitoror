// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	configmodels "github.com/monitoror/monitoror/api/config/models"
	mock "github.com/stretchr/testify/mock"

	models "github.com/monitoror/monitoror/monitorables/github/api/models"

	monitorormodels "github.com/monitoror/monitoror/models"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Checks provides a mock function with given fields: params
func (_m *Usecase) Checks(params *models.ChecksParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.ChecksParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.ChecksParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Count provides a mock function with given fields: params
func (_m *Usecase) Count(params *models.CountParams) (*monitorormodels.Tile, error) {
	ret := _m.Called(params)

	var r0 *monitorormodels.Tile
	if rf, ok := ret.Get(0).(func(*models.CountParams) *monitorormodels.Tile); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*monitorormodels.Tile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.CountParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PullRequests provides a mock function with given fields: params
func (_m *Usecase) PullRequests(params interface{}) ([]configmodels.DynamicTileResult, error) {
	ret := _m.Called(params)

	var r0 []configmodels.DynamicTileResult
	if rf, ok := ret.Get(0).(func(interface{}) []configmodels.DynamicTileResult); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]configmodels.DynamicTileResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
