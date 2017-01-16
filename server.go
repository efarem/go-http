package main

import (
  "net/http"
  "github.com/labstack/echo"
)

type (
  Data struct {
    Secret string `json:"secret"`
  }
)

func main() {

  e := echo.New()

  e.POST("/hello/:name", func(c echo.Context) error {

    d := new(Data)

    if err := c.Bind(d); err != nil {
      return err
    }

    c.Response().Header().Set("X-Secret", d.Secret)

    name := c.Param("name")

    return c.String(http.StatusOK, "<h1>Hello " + name + "</h1>\n")
  })

  e.Logger.Fatal(e.Start(":1337"))
}
