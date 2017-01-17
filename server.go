package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	Data struct {
		Secret string `json:"secret"`
	}
)

const (
	HeaderXSecret = "X-Secret"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		d := new(Data)

		if err := c.Bind(d); err != nil {
			return err
		}

		c.Response().Header().Set(HeaderXSecret, d.Secret)

		return next(c)
	}
}

func Hello(c echo.Context) error {

	name := c.Param("name")

	return c.String(http.StatusOK, "<h1>Hello "+name+"</h1>\n")
}

func main() {

	e := echo.New()

	e.Use(ServerHeader)

	e.POST("/hello/:name", Hello)

	e.Logger.Fatal(e.Start(":1337"))
}
