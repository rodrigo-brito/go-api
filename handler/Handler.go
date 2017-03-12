package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Golang!")
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "ID = "+id)
}
