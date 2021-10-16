package cinemasController_test

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/businesses"
	"TiBO_API/businesses/cinemasEntity"
	_mockCinema "TiBO_API/businesses/cinemasEntity/mocks"
	"TiBO_API/controllers/cinemasController"
	"TiBO_API/helpers"
	//"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var(
	mockCinemasService _mockCinema.Service
	cinemasCtrl *cinemasController.CinemasController
	cinemaReq string
	cinemaReqInvalidBind string
	cinemaReqInvalidStruct string
	cinemaDomain cinemasEntity.Domain
	claims *auth.JwtCustomClaims
)
func TestMain(m *testing.M){
	cinemasCtrl = cinemasController.NewCinemaController(&mockCinemasService)
	cinemaReq = `{
		"name": "cinema uji coba",
		"phone_number": "(021)1234568",
		"is_open": true,
		"addresses": {
			"street": "Jl. Boulevard Raya, RT.13/RW.17",
			"City": "Jakarta Selatan",
			"Provinci": "DKi Jakarta"
		}
	}`
	cinemaReqInvalidBind = `{
		"name": "cinema uji coba",
		"phone_number": "(021)1234568",
		"is_open": true,
		"addresses": {
			"street": "Jl. Boulevard Raya, RT.13/RW.17",
			"City": "Jakarta Selatan",
			"Provinci": "DKi Jakarta"
		}
	}`
	cinemaReqInvalidStruct = `{
		"name": "cinema uji coba",
		"is_open": 24234,
	}`
	cinemaDomain = cinemasEntity.Domain{
		ID:          1,
		Slug:        "cinema-uji-coba",
		Name:        "cinema uji coba",
		PhoneNumber: "(021)1234568",
		IsOpen:      true,
		AddressID:   1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	claims = &auth.JwtCustomClaims{
		Uuid: "admin",
		Role: "admin",
	}
	m.Run()
}

func TestCinemasController_CreateCinema(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/cinema", strings.NewReader(cinemaReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req,rec)
		c.Set("admin", &jwt.Token{Claims: claims})

		mockCinemasService.On("CreateCinema", mock.Anything, mock.Anything,mock.Anything).Return(&cinemaDomain, nil).Once()

		if assert.NoError(t, cinemasCtrl.CreateCinema(c)){
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
	t.Run("Invalid Test | fail validate", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/cinema", strings.NewReader(cinemaReqInvalidStruct))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		if assert.NoError(t, cinemasCtrl.CreateCinema(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("fail store", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(cinemaReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockCinemasService.On("CreateCinema",  mock.Anything,mock.Anything,mock.Anything).Return(&cinemasEntity.Domain{}, assert.AnError).Once()

		if assert.NoError(t, cinemasCtrl.CreateCinema(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestCinemasController_FindCinemaByIP(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/cinema/find-ip", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindByIP", mock.Anything,mock.AnythingOfType("string")).Return([]cinemasEntity.Domain{cinemaDomain}, nil).Once()

		if assert.NoError(t, cinemasCtrl.FindCinemaByIP(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test | cinema Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/cinema/find-ip", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)

		mockCinemasService.On("FindByIP", mock.Anything, mock.AnythingOfType("string")).Return([]cinemasEntity.Domain{}, businesses.ErrNearestCinemaNotFound).Once()

		if assert.NoError(t, cinemasCtrl.FindCinemaByIP(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestCinemasController_FindCinemaByName(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/cinema/find-name/Test", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindByName",mock.Anything, mock.AnythingOfType("string")).Return([]cinemasEntity.Domain{cinemaDomain}, nil).Once()

		if assert.NoError(t, cinemasCtrl.FindCinemaByName(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test | Cinema Not Found", func(t *testing.T){
		req := httptest.NewRequest(http.MethodGet, "/cinema/find-name/Test", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindByName",mock.Anything, mock.AnythingOfType("string")).Return([]cinemasEntity.Domain{}, businesses.ErrCinemaNotFound).Once()

		if assert.NoError(t, cinemasCtrl.FindCinemaByName(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestCinemasController_DeleteCinemaBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodDelete, "/laundro/cinema-uji-coba", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindBySlug", mock.Anything,mock.AnythingOfType("string")).Return(cinemaDomain, nil).Once()
		mockCinemasService.On("DeleteBySlug", mock.Anything,mock.AnythingOfType("string")).Return("Successfully Deleted a cinemas", nil).Once()

		if assert.NoError(t, cinemasCtrl.DeleteCinemaBySlug(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodDelete, "/cinema/cinema-uji-coba", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindBySlug", mock.Anything,mock.AnythingOfType("string")).Return(cinemaDomain, nil).Once()
		mockCinemasService.On("DeleteBySlug", mock.Anything,mock.AnythingOfType("string")).Return("", businesses.ErrCinemaNotFound).Once()


		if assert.NoError(t, cinemasCtrl.DeleteCinemaBySlug(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
	t.Run("Invalid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodDelete, "/cinema/cinema-uji-coba", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindBySlug", mock.Anything,mock.AnythingOfType("string")).Return(cinemasEntity.Domain{}, assert.AnError).Once()
		mockCinemasService.On("DeleteBySlug", mock.Anything,mock.AnythingOfType("string")).Return("", businesses.ErrCinemaNotFound).Once()


		if assert.NoError(t, cinemasCtrl.DeleteCinemaBySlug(c)){
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}

func TestCinemasController_UpdateCinemaBySlug(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/cinema/edit/cinema-uji-coba", strings.NewReader(cinemaReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindBySlug", mock.Anything,mock.AnythingOfType("string")).Return(cinemaDomain, nil).Once()
		mockCinemasService.On("UpdateCinema",mock.Anything,mock.Anything,mock.Anything, mock.AnythingOfType("string")).Return(&cinemaDomain, nil).Once()

		if assert.NoError(t, cinemasCtrl.UpdateCinemaBySlug(c)){
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPut, "/cinema/edit/cinema-uji-coba", strings.NewReader(cinemaReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req,rec)

		mockCinemasService.On("FindBySlug", mock.Anything,mock.AnythingOfType("string")).Return(cinemaDomain, nil).Once()
		mockCinemasService.On("UpdateCinema",mock.Anything,mock.Anything,mock.Anything, mock.AnythingOfType("string")).Return(&cinemasEntity.Domain{}, assert.AnError).Once()


		if assert.NoError(t, cinemasCtrl.UpdateCinemaBySlug(c)){
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

