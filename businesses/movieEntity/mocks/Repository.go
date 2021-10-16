// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	moviesEntity "TiBO_API/businesses/movieEntity"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, slug
func (_m *Repository) Delete(ctx context.Context, slug string) (string, error) {
	ret := _m.Called(ctx, slug)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, slug)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMovieId provides a mock function with given fields: ctx, id
func (_m *Repository) GetByMovieId(ctx context.Context, id uint) (moviesEntity.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 moviesEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, uint) moviesEntity.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(moviesEntity.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySlug provides a mock function with given fields: ctx, slug
func (_m *Repository) GetBySlug(ctx context.Context, slug string) (moviesEntity.Domain, error) {
	ret := _m.Called(ctx, slug)

	var r0 moviesEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) moviesEntity.Domain); ok {
		r0 = rf(ctx, slug)
	} else {
		r0 = ret.Get(0).(moviesEntity.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *Repository) GetByTitle(ctx context.Context, title string) ([]moviesEntity.Domain, error) {
	ret := _m.Called(ctx, title)

	var r0 []moviesEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []moviesEntity.Domain); ok {
		r0 = rf(ctx, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]moviesEntity.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneByTitle provides a mock function with given fields: ctx, title
func (_m *Repository) GetOneByTitle(ctx context.Context, title string) (moviesEntity.Domain, error) {
	ret := _m.Called(ctx, title)

	var r0 moviesEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) moviesEntity.Domain); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(moviesEntity.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostNewMovie provides a mock function with given fields: ctx, movieDomain
func (_m *Repository) PostNewMovie(ctx context.Context, movieDomain *moviesEntity.Domain) (moviesEntity.Domain, error) {
	ret := _m.Called(ctx, movieDomain)

	var r0 moviesEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *moviesEntity.Domain) moviesEntity.Domain); ok {
		r0 = rf(ctx, movieDomain)
	} else {
		r0 = ret.Get(0).(moviesEntity.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *moviesEntity.Domain) error); ok {
		r1 = rf(ctx, movieDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, slug, movieDomain
func (_m *Repository) Update(ctx context.Context, slug string, movieDomain *moviesEntity.Domain) (moviesEntity.Domain, error) {
	ret := _m.Called(ctx, slug, movieDomain)

	var r0 moviesEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, *moviesEntity.Domain) moviesEntity.Domain); ok {
		r0 = rf(ctx, slug, movieDomain)
	} else {
		r0 = ret.Get(0).(moviesEntity.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *moviesEntity.Domain) error); ok {
		r1 = rf(ctx, slug, movieDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadPoster provides a mock function with given fields: ctx, slug, data
func (_m *Repository) UploadPoster(ctx context.Context, slug string, data *moviesEntity.Domain) (*moviesEntity.Domain, error) {
	ret := _m.Called(ctx, slug, data)

	var r0 *moviesEntity.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, *moviesEntity.Domain) *moviesEntity.Domain); ok {
		r0 = rf(ctx, slug, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*moviesEntity.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *moviesEntity.Domain) error); ok {
		r1 = rf(ctx, slug, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
