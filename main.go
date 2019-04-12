package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/igor822/gomicro/controller"
)

func main() {
    e := echo.New()
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))

    e.GET("/health", controller.CheckHealth)

    e.Logger.Fatal(e.Start(":1323"))
}
