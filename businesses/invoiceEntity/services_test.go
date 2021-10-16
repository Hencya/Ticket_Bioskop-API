package invoiceEntity_test

import (
	"TiBO_API/businesses/cinemasEntity"
	_mockCinema "TiBO_API/businesses/cinemasEntity/mocks"
	"TiBO_API/businesses/invoiceEntity"
	"TiBO_API/businesses/invoiceEntity/mocks"
	moviesEntity "TiBO_API/businesses/movieEntity"
	_mockMovie "TiBO_API/businesses/movieEntity/mocks"
	"TiBO_API/businesses/usersEntity"
	_mockUser "TiBO_API/businesses/usersEntity/mocks"
	"context"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockInvoiceRepository mocks.Repository
	mockMovieRepository   _mockMovie.Repository
	mockCinemaRepository  _mockCinema.Repository
	mockUserRepository    _mockUser.Repository
	invoiceService        invoiceEntity.Service
	invoiceDomain         invoiceEntity.Domain
	movieDomain           moviesEntity.Domain
	cinemaDomain          cinemasEntity.Domain
	userDomain            usersEntity.Domain
	userService           usersEntity.Service
)

func TestMain(m *testing.M) {
	invoiceService = invoiceEntity.NewInvoiceService(&mockInvoiceRepository, &mockMovieRepository, &mockCinemaRepository,
		&mockUserRepository, time.Second*2)

	uuidUserID, _ := uuid.Parse("asdasldalsdn1")
	userDomain = usersEntity.Domain{
		ID:   12,
		Uuid: uuidUserID,
	}

	movieDomain = moviesEntity.Domain{
		ID:       1,
		Title:    "Avenger Endgame",
		MovieUrl: "www.google/avenger-endgame-movies.com",
		Price:    1000,
		Ticket: 100,
		ScheduleTime:  "18:00",
	}

	cinemaDomain = cinemasEntity.Domain{
		ID:     3,
		IsOpen: true,
		Name:   "Cinema XX1",
	}

	invoiceDomain = invoiceEntity.Domain{
		ID:               1,
		Date:             time.Now(),
		ShowTime:         "18:00",
		TotalTicket:      3,
		TotalTicketPrice: 30000,
		TotalPrices:      33000,
		AdminFee:         3000,
		MovieID:          1,
		MoviePrice:       1000,
		MovieTitle:       "Avenger Endgame",
		MovieUrl:         "www.google/avenger-endgame-movies.com",
		UserID:           12,
		CinemaID:         3,
		CinemaName:       "Cinema XX1",
		CinemaAddress:    "JL. Testing",
		CreatedAt:        time.Time{},
	}
	m.Run()
}

func TestInvoiceServices_Create(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockMovieRepository.On("GetOneByTitle", mock.Anything,
			mock.Anything).Return(movieDomain, nil).Once()
		mockCinemaRepository.On("FindStatusByTitle", mock.Anything,
			mock.Anything).Return(cinemaDomain, nil).Once()
		mockUserRepository.On("GetByUuid", mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		mockInvoiceRepository.On("Create", mock.Anything,mock.Anything).Return(invoiceDomain, nil).Once()

		req := &invoiceEntity.Domain{
			ID:               1,
			Date:             time.Now(),
			ShowTime:         "18:00",
			TotalTicket:      3,
			TotalTicketPrice: 30000,
			TotalPrices:      33000,
			AdminFee:         3000,
			MovieID:          1,
			MoviePrice:       1000,
			MovieTitle:       "Avenger Endgame",
			MovieUrl:         "www.google/avenger-endgame-movies.com",
			UserID:           12,
			CinemaID:         3,
			CinemaName:       "Cinema XX1",
			CinemaAddress:    "JL. Testing",
			CreatedAt:        time.Time{},
		}

		res, err := invoiceService.Create(context.Background(), "asdasldalsdn1", req)
		assert.Nil(t, err)
		assert.Equal(t, *res, invoiceDomain)
	})
	t.Run("Invalid Test || GetOneByTitle", func(t *testing.T) {
		mockMovieRepository.On("GetOneByTitle", mock.Anything,
			mock.Anything).Return(moviesEntity.Domain{},assert.AnError).Once()
		mockCinemaRepository.On("FindStatusByTitle", mock.Anything,
			mock.Anything).Return(cinemaDomain, nil).Once()
		mockUserRepository.On("GetByUuid", mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		mockInvoiceRepository.On("Create", mock.Anything,mock.Anything).Return(invoiceEntity.Domain{}, nil).Once()

		result, _ := invoiceService.Create(context.Background(), "asdasldalsdn1", &invoiceEntity.Domain{})
		assert.Equal(t, invoiceEntity.Domain{}, *result)
	})
	t.Run("Invalid Test || FindStatusByTitle", func(t *testing.T) {
		mockMovieRepository.On("GetOneByTitle", mock.Anything,
			mock.Anything).Return(movieDomain,nil).Once()
		mockCinemaRepository.On("FindStatusByTitle", mock.Anything,
			mock.Anything).Return(cinemasEntity.Domain{}, assert.AnError).Once()
		mockUserRepository.On("GetByUuid", mock.Anything,
			mock.Anything).Return(userDomain, nil).Once()
		mockInvoiceRepository.On("Create", mock.Anything,mock.Anything).Return(invoiceEntity.Domain{}, nil).Once()

		result, err := invoiceService.Create(context.Background(), "asdasldalsdn1", &invoiceEntity.Domain{})
		assert.NotNil(t, err)
		assert.Equal(t, invoiceEntity.Domain{}, *result)
	})
	t.Run("Invalid Test || GetByUuid", func(t *testing.T) {
		mockMovieRepository.On("GetOneByTitle", mock.Anything,
			mock.Anything).Return(movieDomain,nil).Once()
		mockCinemaRepository.On("FindStatusByTitle", mock.Anything,
			mock.Anything).Return(cinemaDomain, nil).Once()
		mockUserRepository.On("GetByUuid", mock.Anything,
			mock.Anything).Return(usersEntity.Domain{}, assert.AnError).Once()
		mockInvoiceRepository.On("Create", mock.Anything,mock.Anything).Return(invoiceEntity.Domain{}, nil).Once()

		result, _ := invoiceService.Create(context.Background(), "asdasldalsdn1", &invoiceEntity.Domain{})
		assert.Equal(t, invoiceEntity.Domain{}, *result)
	})
}

func TestGetByUserID(t *testing.T) {
	//t.Run("Valid Test", func(t *testing.T) {
	//	mockUserRepository.On("GetByUuid", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
	//	mockInvoiceRepository.On("GetID", mock.Anything, mock.Anything).Return([]invoiceEntity.Domain{invoiceDomain}, nil).Once()
	//	result, _ := invoiceService.GetByUserID(context.Background(),"asdasldalsdn1")
	//	assert.Contains(t, result,invoiceDomain)
	//})
	t.Run("Invalid Test || No invoice", func(t *testing.T) {
		mockUserRepository.On("GetByUuid", mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		mockInvoiceRepository.On("GetID", mock.Anything, mock.Anything).Return([]invoiceEntity.Domain{},  assert.AnError).Once()
		result, _  := invoiceService.GetByUserID(context.Background(),"asdasldalsdn1")
		assert.NotContains(t, result, &[]invoiceEntity.Domain{})
	})
	t.Run("Invalid Test || No invoice", func(t *testing.T) {
		mockUserRepository.On("GetByUuid", mock.Anything, mock.Anything).Return(usersEntity.Domain{}, assert.AnError).Once()
		mockInvoiceRepository.On("GetID", mock.Anything, mock.Anything).Return([]invoiceEntity.Domain{},  assert.AnError).Once()
		result, err := invoiceService.GetByUserID(context.Background(),"asdasldalsdn1")
		assert.NotNil(t, err)
		assert.NotContains(t, result, invoiceDomain)
	})
}
