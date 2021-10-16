package moviesController_test

import (
	"TiBO_API/businesses/cinemasEntity"
	moviesEntity "TiBO_API/businesses/movieEntity"
	"TiBO_API/businesses/movieEntity/mocks"
	"TiBO_API/controllers/moviesController"
	"TiBO_API/controllers/moviesController/response"
	"TiBO_API/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	mockMovieService      mocks.Service
	movieCtrl             *moviesController.MoviesController
	movieReq              string
	movieReqInvalidBind   string
	movieReqInvalidStruct string
	movieResp             response.Movies
	movieDomain           moviesEntity.Domain
	cinemaDomain          cinemasEntity.Domain
)

func TestMain(m *testing.M) {
	movieCtrl = moviesController.NewMoviesController(&mockMovieService)
	movieReq = `{
		"title": "Testing movie",
		"trailer_url": "www.testing.com",
		"Synopsis": "ini adalah testing movie",
		"genre": "Testing",
		"duration": "60 menit",
		"language": "Testing Indo",
		"director": "Faiz",
		"censor_rating": "17+",
		"subtitle": "Thai",
		"schedule_date": "01-10-2021",
		"schedule_time": "18:00",
		"status_coming_soon": false,
		"ticket": 100,
		"price": 10000
	}`
	movieReqInvalidBind = `{
		"title": "Testing movie",
		"trailer_url": "www.testing.com",
		"Synopsis": "ini adalah testing movie",
		"genre": "Testing",
		"duration": "60 menit",
		"language": "Testing Indo",
		"director": "Faiz",
		"censor_rating": "17+",
		"subtitle": "Thai",
		"schedule_date": "01-10-2021",
		"schedule_time": "18:00",
		"status_coming_soon": false,
		"ticket": 100,
		"price": 10000
	}`
	movieReqInvalidStruct = `{
		"title": "Testing movie",
		"trailer_url": "www.testing.com",
		"Synopsis": "ini adalah testing movie",
		"genre": "Testing",
		"duration": "60 menit",
		"language": "Testing Indo",
		"director": "Faiz",
		"censor_rating": "17+",
		"subtitle": "Thai",
		"ticket": 100,
		"price": 10000
		}`
	movieResp = response.Movies{
		Slug:             "testing-movie",
		Title:            "Testing movie",
		TrailerUrl:       "www.testing.com",
		Poster:           "images/poster/poster.png",
		Synopsis:         "ini adalah testing movie",
		Genre:            "Testing",
		Duration:         "60 menit",
		Language:         "Testing Indo",
		Director:         "Faiz",
		CensorRating:     "17+",
		Subtitle:         "Thai",
		ScheduleDate:     "01-10-2021",
		ScheduleTime:     "18:00",
		StatusComingSoon: false,
		Ticket:           100,
		Price:            10000,
		AdminFee:         1000,
		CinemaID:         3,
		CinemaName:       "Testing Cinema",
		CinemaAddress:    "Jl Testing Cinema",
	}
	movieDomain = moviesEntity.Domain{
		ID:               1,
		Slug:             "testing-movie",
		Title:            "Testing movie",
		TrailerUrl:       "www.testing.com",
		Poster:           "images/poster/poster.png",
		Synopsis:         "ini adalah testing movie",
		Genre:            "Testing",
		Duration:         "60 menit",
		Language:         "Testing Indo",
		Director:         "Faiz",
		CensorRating:     "17+",
		Subtitle:         "Thai",
		ScheduleDate:     "01-10-2021",
		ScheduleTime:     "18:00",
		StatusComingSoon: false,
		Ticket:           100,
		Price:            10000,
		AdminFee:         1000,
		CinemaID:         3,
		CinemaName:       "Testing Cinema",
		CinemaAddress:    "Jl Testing Cinema",
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	cinemaDomain = cinemasEntity.Domain{
		ID:          3,
		Slug:        "testing-cinema",
		Name:        "Testing Cinema",
		PhoneNumber: "(021)123134",
		IsOpen:      true,
		AddressID:   3,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	m.Run()
}

func TestMoviesController_CreateMovie(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/movie", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("CreateMovie", mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(&movieDomain, nil).Once()

		if assert.NoError(t, movieCtrl.CreateMovie(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
	t.Run("Invalid Test | Insert Failed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/movie", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("CreateMovie", mock.Anything, mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(&moviesEntity.Domain{}, assert.AnError).Once()

		if assert.NoError(t, movieCtrl.CreateMovie(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestMoviesController_FindMovieBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/movie/find-slug/testing-movie", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(movieDomain, nil).Once()

		if assert.NoError(t, movieCtrl.FindMovieBySlug(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test | Cinema Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/movie/find-slug/testing-movie", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(moviesEntity.Domain{}, assert.AnError).Once()

		if assert.NoError(t, movieCtrl.FindMovieBySlug(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestMoviesController_FindMoviesByTitle(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/movie/find-title/Testing", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockMovieService.On("FindByTitle", mock.Anything, mock.AnythingOfType("string")).Return([]moviesEntity.Domain{movieDomain}, nil).Once()

		if assert.NoError(t, movieCtrl.FindMoviesByTitle(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test | Cinema Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/movie/find-title/Testing", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)

		mockMovieService.On("FindByTitle", mock.Anything, mock.AnythingOfType("string")).Return([]moviesEntity.Domain{}, assert.AnError).Once()

		if assert.NoError(t, movieCtrl.FindMoviesByTitle(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestMoviesController_DeleteMovieBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/movie", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(movieDomain, nil).Once()
		mockMovieService.On("DeleteBySlug", mock.Anything, mock.AnythingOfType("string")).Return("Successfully Deleted a movie", nil).Once()

		if assert.NoError(t, movieCtrl.DeleteMovieBySlug(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/movie", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(movieDomain, nil).Once()
		mockMovieService.On("DeleteBySlug", mock.Anything, mock.AnythingOfType("string")).Return("", assert.AnError).Once()

		if assert.NoError(t, movieCtrl.DeleteMovieBySlug(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
	t.Run("Invalid Test | Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/movie", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(moviesEntity.Domain{}, assert.AnError).Once()
		mockMovieService.On("DeleteBySlug", mock.Anything, mock.AnythingOfType("string")).Return("", assert.AnError).Once()

		if assert.NoError(t, movieCtrl.DeleteMovieBySlug(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestMoviesController_UpdateMovieBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/movie/edit", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(movieDomain, nil).Once()
		mockMovieService.On("UpdateMovie", mock.Anything, mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&movieDomain, nil).Once()

		if assert.NoError(t, movieCtrl.UpdateMovieBySlug(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/movie/edit", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(movieDomain, nil).Once()
		mockMovieService.On("UpdateMovie", mock.Anything, mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(&moviesEntity.Domain{}, assert.AnError).Once()

		if assert.NoError(t, movieCtrl.UpdateMovieBySlug(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
	t.Run("Invalid Test | Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/movie/edit", strings.NewReader(movieReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockMovieService.On("FindBySlug", mock.Anything, mock.AnythingOfType("string")).Return(moviesEntity.Domain{}, assert.AnError).Once()
		mockMovieService.On("UpdateMovie", mock.Anything, mock.Anything, mock.Anything, mock.AnythingOfType("string")).Return(&moviesEntity.Domain{}, assert.AnError).Once()

		if assert.NoError(t, movieCtrl.UpdateMovieBySlug(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}
