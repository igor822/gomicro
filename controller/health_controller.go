package controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type Health struct {
    Status string `json:"status"`
}

func CheckHealth(c echo.Context) error {
    health := Health{}
    health.Status = "UP"
    return c.JSON(http.StatusOK, health)
}
