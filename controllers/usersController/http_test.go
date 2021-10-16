package usersController_test

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/businesses/usersEntity"
	"TiBO_API/businesses/usersEntity/mocks"
	"TiBO_API/controllers/usersController"
	"TiBO_API/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var(
	mockUserService mocks.Service
	jwtConfig auth.ConfigJWT
	userCtrl	*usersController.UserController
	userReq	string
	userReqInvalidBind string
	userReqInvalidStruct string
	userLoginReq string
	hashedPassword string
	userDomain usersEntity.Domain
)

func TestMain(m *testing.M){
	userCtrl = usersController.NewUserController(&mockUserService,&jwtConfig)
	userReq = `{
		"name": "testing user",
		"username" : "testing",
		"email"  : "testing123@gmail.com",
		"password" : "testingTibo#1234",
		"phone_number": "081234567899"
	}`
	userReqInvalidBind = `{
		"name": "testing user",
		"username" : "testing",
		"email"  : "testing123@gmail.com",
		"password" : "testingTibo#1234",
		"phone_number": "081234567899"
	}`
	userReqInvalidStruct = `{
		"name": "testing user",
		"username" : "testing",
		"email"  : "testing123@gmail.com",
		"phone_number": "081234567899"
		}
	}`
	uuidID, _ := uuid.Parse("c5c838ba-24f8-11ec-9621-0242ac130002")
	userLoginReq = `{"email": "testing123@gmail.com","password": "testingTibo#1234"}`
	hashedPassword, _ = helpers.HashPassword("testingTibo#1234")
	userDomain = usersEntity.Domain{
		ID:          1,
		Uuid:       uuidID,
		Name:        "testing user",
		Username:    "testing",
		Email:       "testing123@gmail.com",
		Password:    "testingTibo#1234",
		PhoneNumber: "081234567899",
		Role:        "user",
		Avatar:      "images/avatar/avatar.png",
		Token: "token",
	}
	m.Run()
}

func TestUserController_LoginUser(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T){
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(userLoginReq))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		e.Validator = &helpers.CustomValidator{Validator: validator.New()}
		c := e.NewContext(req,rec)
		mockUserService.On("Login",mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("token", nil).Once()



		if assert.NoError(t, userCtrl.LoginUser(c)){
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t,"token", userDomain.Token)
		}
	})
}
