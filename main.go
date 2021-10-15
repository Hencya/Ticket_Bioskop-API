package main

import (
	"TiBO_API/app/routes"
	"TiBO_API/businesses/usersEntity"
	"TiBO_API/controllers/usersController"
	"TiBO_API/helpers"

	ConfigJWT "TiBO_API/app/config/auth"
	configDB "TiBO_API/app/config/databases"
	_middleware "TiBO_API/app/middleware/logger"
	_domainFactory "TiBO_API/repository"

	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	var (
		db = configDB.SetupDatabaseConnection()
		//cld = configCLD.SetupCloudinaryConnection()
		jwt = ConfigJWT.SetupJwt()
	)
	timeoutDur, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	timeoutContext := time.Duration(timeoutDur) * time.Millisecond

	echoApp := echo.New()

	//middleware
	echoApp.Validator = &helpers.CustomValidator{Validator: validator.New()}
	echoApp.Use(middleware.CORS())
	echoApp.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))

	//users
	userRepo := _domainFactory.NewUserRepository(db)
	userService := usersEntity.NewUserServices(userRepo, &jwt, timeoutContext)
	userCtrl := usersController.NewUserController(userService, &jwt)

	//routes
	routesInit := routes.ControlerList{
		JWTMiddleware:   jwt.Init(),
		UsersController: *userCtrl,
	}
	routesInit.RouteRegister(echoApp)

	log.Fatal(echoApp.Start(":8000"))
}
