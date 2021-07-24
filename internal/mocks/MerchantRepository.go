// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/ardafirdausr/posjoo-server/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// MerchantRepository is an autogenerated mock type for the MerchantRepository type
type MerchantRepository struct {
	mock.Mock
}

// CreateMerchant provides a mock function with given fields: ctx, param
func (_m *MerchantRepository) CreateMerchant(ctx context.Context, param entity.CreateMerchantParam) (*entity.Merchant, error) {
	ret := _m.Called(ctx, param)

	var r0 *entity.Merchant
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateMerchantParam) *entity.Merchant); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Merchant)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateMerchantParam) error); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
