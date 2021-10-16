package usersController

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/businesses/usersEntity"
	"TiBO_API/controllers/usersController/request"
	"TiBO_API/controllers/usersController/response"
	"TiBO_API/helpers"
	"fmt"
	"io"
	"os"

	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usersService usersEntity.Service
	jwtAuth      *auth.ConfigJWT
}

func NewUserController(us usersEntity.Service, jwtauth *auth.ConfigJWT) *UserController {
	return &UserController{
		usersService: us,
		jwtAuth:      jwtauth,
	}
}

func (ctrl *UserController) Registration(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.UserRegistration)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.usersService.Register(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created an account",
			response.FromDomain(res)))
}

func (ctrl *UserController) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.UserLogin)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	token, err := ctrl.usersService.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	expire := time.Now().Add(5 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:    "is-login",
		Value:   token,
		Expires: expire,
	}
	c.SetCookie(&cookie)

	//return helpers.BuildSuccessResponseContext(c, response)
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("successful to login",
			response))
}

func (ctrl *UserController) LogoutUser(c echo.Context) error {
	err := ctrl.usersService.Logout(c)
	if err != nil {
		return c.JSON(http.StatusRequestTimeout,
			helpers.BuildSuccessResponse("you are not logged in",
				nil))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("you have successfully logged out",
			nil))
}

func (ctrl *UserController) FindUserByUuid(c echo.Context) error {
	//checking cookie is the user was login or not
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}

	uuid := c.Param("uuid")

	user, err := ctrl.usersService.FindByUuid(c.Request().Context(), uuid)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("user doesn't exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get User By id",
			response.FromDomain(&user)))
}

func (ctrl *UserController) UpdateUserById(c echo.Context) error {
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}

	ctx := c.Request().Context()
	req := new(request.UserRegistration)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while input the data",
				err, helpers.EmptyObj{}))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	user := auth.GetUser(c)
	id := user.Uuid

	res, err := ctrl.usersService.UpdateById(ctx, req.ToDomain(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully update an account",
			response.FromDomain(res)))
}

func (ctrl *UserController) UploadAvatar(c echo.Context) error {
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}
	//var err error
	ctx := c.Request().Context()
	//file := new(request.UserUploadAvatar)
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	user := auth.GetUser(c)
	userID := user.Uuid

	path := fmt.Sprintf("images/avatar/%v-%s", userID, file.Filename)

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}
	defer src.Close()

	destination, err := os.Create(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}
	defer destination.Close()

	if _, err = io.Copy(destination, src); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.usersService.UploadAvatar(ctx, userID, path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Uploud Avatar Failed",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully upload avatar",
			response.FromDomain(res)))
}

func (ctrl *UserController) DeleteUserByUuid(c echo.Context) error {
	//checking login user
	_, errC := c.Cookie("is-login")
	if errC != nil {
		return c.JSON(http.StatusForbidden,
			helpers.BuildErrorResponse("you are not logged in",
				errC, helpers.EmptyObj{}))
	}

	user := auth.GetUser(c)
	id := user.Uuid

	_, errGet := ctrl.usersService.FindByUuid(c.Request().Context(), id)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("User doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.usersService.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a User",
			nil))
}
