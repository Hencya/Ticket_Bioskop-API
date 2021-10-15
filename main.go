package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	echoApp := echo.New()
	echoApp.Use(middleware.CORS())

	log.Fatal(echoApp.Start(":8080"))
}
