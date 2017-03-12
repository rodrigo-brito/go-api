package main

import (
	"./handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handler.Index)
	e.GET("/user/:id", handler.GetUser)
	e.Logger.Fatal(e.Start(":8080"))
}
