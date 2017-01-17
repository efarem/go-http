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

func main() {

	e := echo.New()

	e.POST("/hello/:name", func(c echo.Context) error {

		d := new(Data)

		if err := c.Bind(d); err != nil {
			return err
		}

		c.Response().Header().Set(HeaderXSecret, d.Secret)

		name := c.Param("name")

		return c.String(http.StatusOK, "<h1>Hello "+name+"</h1>\n")
	})

	e.Logger.Fatal(e.Start(":1337"))
}
