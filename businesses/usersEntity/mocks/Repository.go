// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	usersEntity "TiBO_API/businesses/usersEntity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateNewUser provides a mock function with given fields: ctx, data
func (_m *Repository) CreateNewUser(ctx context.Context, data *usersEntity.Domain) (*usersEntity.Domain, error) {
	ret := _m.Called(ctx, data)

	var r0 *usersEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *usersEntity.Domain) *usersEntity.Domain); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersEntity.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *usersEntity.Domain) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserByUuid provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteUserByUuid(ctx context.Context, id string) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) GetByEmail(ctx context.Context, email string) (usersEntity.Domain, error) {
	ret := _m.Called(ctx, email)

	var r0 usersEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) usersEntity.Domain); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(usersEntity.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUuid provides a mock function with given fields: ctx, uuid
func (_m *Repository) GetByUuid(ctx context.Context, uuid string) (usersEntity.Domain, error) {
	ret := _m.Called(ctx, uuid)

	var r0 usersEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) usersEntity.Domain); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(usersEntity.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, id, data
func (_m *Repository) UpdateUser(ctx context.Context, id string, data *usersEntity.Domain) (*usersEntity.Domain, error) {
	ret := _m.Called(ctx, id, data)

	var r0 *usersEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, *usersEntity.Domain) *usersEntity.Domain); ok {
		r0 = rf(ctx, id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersEntity.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *usersEntity.Domain) error); ok {
		r1 = rf(ctx, id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadAvatar provides a mock function with given fields: ctx, id, data
func (_m *Repository) UploadAvatar(ctx context.Context, id string, data *usersEntity.Domain) (*usersEntity.Domain, error) {
	ret := _m.Called(ctx, id, data)

	var r0 *usersEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, *usersEntity.Domain) *usersEntity.Domain); ok {
		r0 = rf(ctx, id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersEntity.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *usersEntity.Domain) error); ok {
		r1 = rf(ctx, id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}