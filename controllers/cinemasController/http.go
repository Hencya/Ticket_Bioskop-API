package cinemasController

import (
	"TiBO_API/businesses/cinemasEntity"
	"TiBO_API/controllers/cinemasController/request"
	"TiBO_API/controllers/cinemasController/response"
	"TiBO_API/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CinemasController struct {
	cinemasServices cinemasEntity.Service
}

func NewCinemaController(cs cinemasEntity.Service) *CinemasController {
	return &CinemasController{
		cinemasServices: cs,
	}
}

func (ctrl *CinemasController) CreateCinema(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.Cinemas)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while inputing data",
				err, helpers.EmptyObj{}))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	cinemaData, addrData := req.ToDomain()
	res, err := ctrl.cinemasServices.CreateCinema(ctx, cinemaData, addrData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created a cinemas",
			response.FromDomain(*res)))
}

func (ctrl *CinemasController) FindCinemaByIP(c echo.Context) error {
	ip := c.RealIP()
	data, err := ctrl.cinemasServices.FindByIP(c.Request().Context(), ip)
	if len(data) == 0 {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("No one Cinema found within your area",
				err, helpers.EmptyObj{}))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Cinema By IP",
			data))
}

func (ctrl *CinemasController) FindCinemaByName(c echo.Context) error {
	name := c.Param("name")
	data, err := ctrl.cinemasServices.FindByName(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Cinema doesn't exist",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Cinema By Name",
			data))
}

func (ctrl *CinemasController) FindCinemaBySlug(c echo.Context) error {
	slug := c.Param("slug")

	cinema, err := ctrl.cinemasServices.FindBySlug(c.Request().Context(), slug)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Cinema doesn't exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get User By id",
			response.FromDomain(cinema)))

}

func (ctrl *CinemasController) UpdateCinemaBySlug(c echo.Context) error {
	slug := c.Param("slug")
	_, errGet := ctrl.cinemasServices.FindBySlug(c.Request().Context(), slug)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Cinema doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	ctx := c.Request().Context()
	req := new(request.Cinemas)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while inputing data",
				err, helpers.EmptyObj{}))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("An error occurred while validating the request data",
				err, helpers.EmptyObj{}))
	}

	cinemaData, addrData := req.ToDomain()
	res, err := ctrl.cinemasServices.UpdateCinema(ctx, cinemaData, addrData, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully updated a cinemas",
			response.FromDomain(*res)))
}

func (ctrl *CinemasController) DeleteCinemaBySlug(c echo.Context) error {
	slug := c.Param("slug")
	_, errGet := ctrl.cinemasServices.FindBySlug(c.Request().Context(), slug)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Cinema doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.cinemasServices.DeleteBySlug(c.Request().Context(), slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully Deleted a cinemas",
			nil))
}
