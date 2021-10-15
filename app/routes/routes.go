package routes

import (
	"TiBO_API/app/middleware/auth"
	"TiBO_API/controllers/cinemasController"
	"TiBO_API/controllers/moviesController"
	"TiBO_API/controllers/usersController"
	"TiBO_API/helpers"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControlerList struct {
	UsersController   usersController.UserController
	CinemasController cinemasController.CinemasController
	MoviesController  moviesController.MoviesController
	JWTMiddleware     middleware.JWTConfig
}

func (cl *ControlerList) RouteRegister(echo *echo.Echo) {
	//users
	users := echo.Group("api/v1/user")
	users.POST("/register", cl.UsersController.Registration)
	users.POST("/login", cl.UsersController.LoginUser)
	users.POST("/logout", cl.UsersController.LogoutUser)
	users.GET("/:uuid", cl.UsersController.FindUserByUuid)

	//users with admin role
	users.PUT("/update", cl.UsersController.UpdateUserById, middleware.JWTWithConfig(cl.JWTMiddleware))
	users.POST("/uploadAvatar", cl.UsersController.UploadAvatar, middleware.JWTWithConfig(cl.JWTMiddleware))
	users.DELETE("/delete", cl.UsersController.DeleteUserByUuid, middleware.JWTWithConfig(cl.JWTMiddleware))

	//Cinemas
	cinema := echo.Group("/api/v1/cinema")
	cinema.GET("/find-ip", cl.CinemasController.FindCinemaByIP)
	cinema.GET("/find-name/:name", cl.CinemasController.FindCinemaByName)

	//cinemas with admin role
	cinemaAdmin := cinema
	cinemaAdmin.Use(middleware.JWTWithConfig(cl.JWTMiddleware), AdminValidation())
	cinemaAdmin.POST("", cl.CinemasController.CreateCinema)
	cinemaAdmin.PUT("/edit/:slug", cl.CinemasController.UpdateCinemaBySlug)
	cinemaAdmin.DELETE("/:slug", cl.CinemasController.DeleteCinemaBySlug)

	//movies
	movie := echo.Group("/api/v1/movie")
	movie.GET("/find-slug/:slug", cl.MoviesController.FindMovieBySlug)
	movie.GET("/find-title/:title", cl.MoviesController.FindMoviesByTitle)

	//movies with admin role
	movieAdmin := movie
	movieAdmin.Use(middleware.JWTWithConfig(cl.JWTMiddleware), AdminValidation())
	movieAdmin.POST("/upload-poster/:slug", cl.MoviesController.UploadPosterBySlug)
	movieAdmin.POST("/:slug", cl.MoviesController.CreateMovie)
	movieAdmin.PUT("/edit", cl.MoviesController.UpdateMovieBySlug)
	movieAdmin.DELETE("/:slug", cl.MoviesController.DeleteMovieBySlug)
}

func AdminValidation() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := auth.GetUser(c)

			if user.Role != "admin" {
				return c.JSON(http.StatusForbidden,
					helpers.BuildErrorResponse("you are not a admin",
						errors.New("Pleas Login as admin"), helpers.EmptyObj{}))
			} else {
				return hf(c)
			}

		}
	}
}
