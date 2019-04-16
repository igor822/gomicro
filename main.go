package main

import (
	"github.com/igor822/gomicro/config"
	"github.com/igor822/gomicro/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	g := e.Group("/gomicro/v1")
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g.GET("/health", controller.CheckHealth)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
