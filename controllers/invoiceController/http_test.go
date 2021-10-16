package invoiceController_test

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/businesses"
	"TiBO_API/businesses/invoiceEntity"
	"TiBO_API/businesses/invoiceEntity/mocks"
	"TiBO_API/controllers/invoiceController"
	"TiBO_API/controllers/invoiceController/response"
	"TiBO_API/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"

	//"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

var (
	mockInvoiceService      mocks.Service
	invoiceCtrl             *invoiceController.InvoiceController
	invoiceReq              string
	invoiceReqInvalidBind   string
	invoiceReqInvalidStruct string
	invoiceDomain           invoiceEntity.Domain
	invoiceRes              response.Invoices
	claims                  *auth.JwtCustomClaims
)

func TestMain(m *testing.M) {
	invoiceCtrl = invoiceController.NewInvoiceController(&mockInvoiceService)
	invoiceReq = `{
		"movie_title": "dummy 213",
		"total_ticket": 1
	}`
	invoiceReqInvalidBind = `{
		"movie_title": "dummy 213",
		"total_ticket": 1
	}`
	invoiceReqInvalidStruct = `{
		"total_ticket": 1
	}`
	invoiceDomain = invoiceEntity.Domain{
		ID:               1,
		Date:             time.Now(),
		ShowTime:         "18:00",
		TotalTicket:      2,
		TotalTicketPrice: 20000,
		TotalPrices:      22000,
		AdminFee:         2000,
		MovieID:          4,
		MoviePrice:       10000,
		MovieTitle:       "Testing movie",
		MovieUrl:         "www.testingmovie.com",
		UserID:           3,
		CinemaID:         5,
		CinemaName:       "Testing Cinema",
		CinemaAddress:    "JL Testing",
		CreatedAt:        time.Now(),
	}
	invoiceRes = response.Invoices{
		ID:               1,
		Date:             time.Now(),
		ShowTime:         "18:00",
		TotalTicket:      2,
		TotalTicketPrice: 20000,
		TotalPrices:      22000,
		AdminFee:         2000,
		MovieID:          4,
		MoviePrice:       10000,
		MovieTitle:       "Testing movie",
		MovieUrl:         "www.testingmovie.com",
		UserID:           3,
		CinemaID:         5,
		CinemaName:       "Testing Cinema",
		CinemaAddress:    "JL Testing",
		CreatedAt:        time.Now(),
	}
	claims = &auth.JwtCustomClaims{
		Uuid: "testing-user",
		Role: "admin",
	}
	m.Run()
}

func TestInvoiceController_Create(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/buy", strings.NewReader(invoiceReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockInvoiceService.On("Create", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(&invoiceDomain, nil).Once()

		if assert.NoError(t, invoiceCtrl.Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
	t.Run("Invalid Test | Internal Server Error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/buy", strings.NewReader(invoiceReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockInvoiceService.On("Create", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(&invoiceEntity.Domain{}, assert.AnError).Once()

		if assert.NoError(t, invoiceCtrl.Create(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestInvoiceController_GetByUserID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get-invoice", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockInvoiceService.On("GetByUserID", mock.Anything,mock.AnythingOfType("string")).Return([]invoiceEntity.Domain{invoiceDomain}, nil).Once()

		if assert.NoError(t, invoiceCtrl.GetByUserID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Invalid Test | No Order Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/get-invoice", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: claims})

		mockInvoiceService.On("GetByUserID",mock.Anything, mock.AnythingOfType("string")).Return([]invoiceEntity.Domain{}, businesses.ErrInvoiceNotFound).Once()

		if assert.NoError(t, invoiceCtrl.GetByUserID(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}
