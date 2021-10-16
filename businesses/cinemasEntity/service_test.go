package cinemasEntity_test

import (
	"TiBO_API/businesses/addressesEntity"
	mockAdresses "TiBO_API/businesses/addressesEntity/mocks"
	"TiBO_API/businesses/cinemasEntity"
	mockCinema "TiBO_API/businesses/cinemasEntity/mocks"
	"TiBO_API/businesses/geolocationEntity"
	mockGeo "TiBO_API/businesses/geolocationEntity/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var (
	mockCinemaRepository      mockCinema.Repository
	mockAddressesRepository   mockAdresses.Repository
	mockGeolocationRepository mockGeo.Repository
	cinemaService             cinemasEntity.Service
	cinemaDomain              cinemasEntity.Domain
	newCinemaDomain           cinemasEntity.Domain
	newAddCinemaDomain        cinemasEntity.Domain
	addressesDomain           addressesEntity.Domain
	addressesNewDomain        addressesEntity.Domain
	geoDomain                 geolocationEntity.Domain
)

func TestMain(m *testing.M) {
	cinemaService = cinemasEntity.NewCinemaServices(&mockCinemaRepository, &mockAddressesRepository,
		&mockGeolocationRepository, time.Second*2)
	cinemaDomain = cinemasEntity.Domain{
		ID:          1,
		Slug:        "asd;sasjdlkajd",
		Name:        "Test Cinema",
		PhoneNumber: "(021)123123",
		IsOpen:      true,
		AddressID:   2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newCinemaDomain = cinemasEntity.Domain{

		ID:          1,
		Slug:        "asd;sasjdlkajd",
		Name:        "Test New Cinema",
		PhoneNumber: "(021)123123",
		IsOpen:      false,
		AddressID:   2,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newAddCinemaDomain = cinemasEntity.Domain{
		ID:          1,
		Slug:        "asd;sasjdlkajd",
		Name:        "Test New Cinema",
		AddressID:   2,
		PhoneNumber: "(021)123123",
		IsOpen:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	addressesDomain = addressesEntity.Domain{
		ID:       1,
		Street:   "Test Street",
		City:     "Test City",
		Province: "Test Province",
	}
	addressesNewDomain = addressesEntity.Domain{
		ID:       2,
		Street:   "Test Street 2",
		City:     "Test City 2",
		Province: "Test Province 2",
	}
	geoDomain = geolocationEntity.Domain{
		IP:   "0.0.0.0",
		City: "Test City",
	}
	m.Run()
}

func TestCinemaServices_CreateCinema(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockAddressesRepository.On("Insert", mock.Anything).Return(addressesDomain, nil).Once()
		mockCinemaRepository.On("PostNewCinema", mock.Anything, mock.Anything).Return(cinemaDomain, nil).Once()

		input := cinemasEntity.Domain{
			Name:   "Test Cinema",
			IsOpen: true,
		}

		resp, _ := cinemaService.CreateCinema(context.Background(), &input, &addressesDomain)
		assert.Equal(t, *resp, cinemaDomain)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockAddressesRepository.On("Insert", mock.Anything).Return(addressesDomain, nil).Once()
		mockCinemaRepository.On("PostNewCinema", mock.Anything, mock.Anything).Return(cinemasEntity.Domain{}, assert.AnError).Once()

		input := cinemasEntity.Domain{
			Name:   "Test Cinema",
			IsOpen: true,
		}

		resp, err := cinemaService.CreateCinema(context.Background(), &input, &addressesDomain)

		assert.NotNil(t, err)
		assert.Equal(t, *resp, cinemasEntity.Domain{})
	})
}

func TestCinemaServices_FindByIP(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockGeolocationRepository.On("GetLocationByIP", mock.Anything,mock.AnythingOfType("string")).Return(geoDomain, nil).Once()
		mockAddressesRepository.On("FindByCity", mock.AnythingOfType("string")).Return([]addressesEntity.Domain{addressesDomain}, nil).Once()
		mockCinemaRepository.On("GetByAddress",mock.Anything, mock.AnythingOfType("[]uint")).Return([]cinemasEntity.Domain{cinemaDomain}, nil).Once()

		resp, err := cinemaService.FindByIP(context.Background(),"0.0.0.0")

		assert.Nil(t, err)
		assert.Contains(t, resp, cinemaDomain)
	})
	t.Run("Invalid Test | Location is Not Found", func(t *testing.T){
		mockGeolocationRepository.On("GetLocationByIP", mock.Anything,mock.AnythingOfType("string")).Return(geolocationEntity.Domain{}, assert.AnError).Once()

		resp, err := cinemaService.FindByIP(context.Background(),"0.0.0.0")

		assert.NotNil(t, err)
		assert.NotContains(t, resp, cinemaDomain)
	})
	t.Run("Invalid Test | Address is Not Found", func(t *testing.T){
		mockGeolocationRepository.On("GetLocationByIP", mock.Anything,mock.AnythingOfType("string")).Return(geoDomain, nil).Once()
		mockAddressesRepository.On("FindByCity", mock.AnythingOfType("string")).Return([]addressesEntity.Domain{}, assert.AnError).Once()

		resp, err := cinemaService.FindByIP(context.Background(),"0.0.0.0")

		assert.NotNil(t, err)
		assert.NotContains(t, resp, cinemaDomain)
	})
	t.Run("Invalid Test | Cinema is Not Found", func(t *testing.T){
		mockGeolocationRepository.On("GetLocationByIP",mock.Anything, mock.AnythingOfType("string")).Return(geoDomain, nil).Once()
		mockAddressesRepository.On("FindByCity", mock.AnythingOfType("string")).Return([]addressesEntity.Domain{addressesDomain}, nil).Once()
		mockCinemaRepository.On("GetByAddress",mock.Anything, mock.AnythingOfType("[]uint")).Return([]cinemasEntity.Domain{}, assert.AnError).Once()

		resp, err := cinemaService.FindByIP(context.Background(),"0.0.0.0")

		assert.NotNil(t, err)
		assert.NotContains(t, resp, cinemaDomain)
	})
}

func TestCinemaServices_FindBySlug(t *testing.T){
	t.Run("Valid Test", func(t *testing.T){
		mockCinemaRepository.On("GetBySlug", mock.Anything,mock.AnythingOfType("string")).Return(cinemaDomain, nil).Once()

		resp, err := cinemaService.FindBySlug(context.Background(),"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, cinemaDomain, resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockCinemaRepository.On("GetBySlug",  mock.Anything,mock.AnythingOfType("string")).Return(cinemasEntity.Domain{}, assert.AnError).Once()

		resp, err := cinemaService.FindBySlug(context.Background(),"asd;sasjdlkajd")

		assert.NotNil(t, err)
		assert.Equal(t, cinemasEntity.Domain{}, resp)
	})
}

func TestCinemaServices_DeleteBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockCinemaRepository.On("Delete", mock.Anything,mock.AnythingOfType("string")).Return("Cinemas was Deleted", nil).Once()

		resp, err := cinemaService.DeleteBySlug(context.Background(),"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, "Cinemas was Deleted", resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockCinemaRepository.On("Delete", mock.Anything,mock.AnythingOfType("string")).Return("", assert.AnError).Once()

		resp, err := cinemaService.DeleteBySlug(context.Background(),"asd;sasjdlkajd")

		assert.NotNil(t, err)
		assert.Equal(t, "", resp)
	})
}

func TestCinemaServices_FindByName(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		mockCinemaRepository.On("GetByName",mock.Anything, mock.AnythingOfType("string")).Return([]cinemasEntity.Domain{cinemaDomain}, nil).Once()

		resp, err := cinemaService.FindByName(context.Background(),"Test New Cinema")

		assert.Nil(t, err)
		assert.Contains(t, resp, cinemaDomain)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockCinemaRepository.On("GetByName",mock.Anything, mock.AnythingOfType("string")).Return([]cinemasEntity.Domain{}, assert.AnError).Once()

		resp, err := cinemaService.FindByName(context.Background(),"Test New Cinema")

		assert.NotNil(t, err)
		assert.NotContains(t, resp, cinemaDomain)
	})
}

func TestCinemaServices_UpdateCinema(t *testing.T) {
	t.Run("Valid Test | Address Unchanged", func(t *testing.T){
		mockAddressesRepository.On("Insert", mock.Anything).Return(addressesDomain, nil).Once()
		mockCinemaRepository.On("Update",  mock.Anything,mock.AnythingOfType("string"), mock.Anything).Return(cinemaDomain, nil).Once()

		input := cinemasEntity.Domain{
			Name:   "Test Cinema",
			IsOpen: true,
		}

		resp, err := cinemaService.UpdateCinema(context.Background(), &input, &addressesDomain,"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, cinemaDomain, *resp)
	})
	t.Run("Valid Test | New Address", func(t *testing.T){
		mockAddressesRepository.On("Insert", mock.Anything).Return(addressesNewDomain, nil).Once()
		mockCinemaRepository.On("Update", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(newAddCinemaDomain, nil).Once()

		input := cinemasEntity.Domain{
			Name:   "Test Cinema",
			IsOpen: true,
		}

		resp, err := cinemaService.UpdateCinema(context.Background(), &input, &addressesDomain,"asd;sasjdlkajd")

		assert.Nil(t, err)
		assert.Equal(t, newCinemaDomain, *resp)
	})
	t.Run("Invalid Test", func(t *testing.T){
		mockAddressesRepository.On("Insert", mock.Anything).Return(addressesDomain, nil).Once()
		mockCinemaRepository.On("Update",  mock.Anything,mock.AnythingOfType("string"), mock.Anything).Return(cinemaDomain, assert.AnError).Once()

		input := cinemasEntity.Domain{
			Name:   "Test Cinema",
			IsOpen: true,
		}

		resp, err := cinemaService.UpdateCinema(context.Background(), &input, &addressesDomain,"asd;sasjdlkajd")

		assert.NotNil(t, err)
		assert.Equal(t, cinemasEntity.Domain{}, *resp)
	})
}
