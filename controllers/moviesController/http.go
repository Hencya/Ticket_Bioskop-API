package moviesController

import (
	moviesEntity "TiBO_API/businesses/movieEntity"
	"TiBO_API/controllers/moviesController/request"
	"TiBO_API/controllers/moviesController/response"
	"TiBO_API/helpers"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type MoviesController struct {
	moviesServices moviesEntity.Service
}

func NewMoviesController(ms moviesEntity.Service) *MoviesController {
	return &MoviesController{
		moviesServices: ms,
	}
}

func (ctrl *MoviesController) CreateMovie(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(request.Movies)

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

	moviesData := req.ToDomain()
	slug := c.Param("slug")

	res, err := ctrl.moviesServices.CreateMovie(ctx, moviesData, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created a cinemas",
			response.FromDomain(*res)))
}

func (ctrl *MoviesController) FindMoviesByTitle(c echo.Context) error {
	name := c.Param("title")

	data, err := ctrl.moviesServices.FindByTitle(c.Request().Context(), name)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Movie  doesn't exist",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get Movie By Name",
			data))
}

func (ctrl *MoviesController) FindMovieBySlug(c echo.Context) error {
	slug := c.Param("slug")

	movie, err := ctrl.moviesServices.FindBySlug(c.Request().Context(), slug)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Movie doesn't exist",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Get User By id",
			response.FromDomain(movie)))

}

func (ctrl *MoviesController) UpdateMovieBySlug(c echo.Context) error {
	slugMovie := c.QueryParam("slugMovie")
	slugCinema := c.QueryParam("slugCinema")

	_, errGet := ctrl.moviesServices.FindBySlug(c.Request().Context(), slugMovie)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("movie doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	ctx := c.Request().Context()
	req := new(request.Movies)

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

	movieData := req.ToDomain()
	res, err := ctrl.moviesServices.UpdateMovie(ctx, movieData, slugCinema, slugMovie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully updated a movie",
			response.FromDomain(*res)))
}

func (ctrl *MoviesController) DeleteMovieBySlug(c echo.Context) error {
	slug := c.Param("slug")
	_, errGet := ctrl.moviesServices.FindBySlug(c.Request().Context(), slug)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Movie doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.moviesServices.DeleteBySlug(c.Request().Context(), slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Deleted a movie",
			nil))
}

func (ctrl *MoviesController) UploadPosterBySlug(c echo.Context) error {
	ctx := c.Request().Context()

	slug := c.Param("slug")
	file, err := c.FormFile("poster")
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Failed to Upload poster",
				err, helpers.EmptyObj{}))
	}

	movie, errGet := ctrl.moviesServices.FindBySlug(c.Request().Context(), slug)
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("movie doesn't exist",
				errGet, helpers.EmptyObj{}))
	}

	path := fmt.Sprintf("images/poster/%v-%s", movie.Slug, file.Filename)

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Failed to Upload poster",
				err, helpers.EmptyObj{}))
	}
	defer src.Close()

	destination, err := os.Create(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Failed to Upload poster",
				err, helpers.EmptyObj{}))
	}
	defer destination.Close()

	if _, err = io.Copy(destination, src); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Failed to Upload poster",
				err, helpers.EmptyObj{}))
	}

	res, err := ctrl.moviesServices.UploadPoster(ctx, slug, path)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("Failed to Upload poster",
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully upload avatar",
			response.FromDomain(*res)))
}
