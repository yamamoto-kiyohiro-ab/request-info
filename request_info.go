package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/*", handler)
	e.POST("/*", handler)

	e.Logger.Fatal(e.Start(":80"))
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
