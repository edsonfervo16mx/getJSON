package main

import (
	"fmt"
	"getJSON/src/mydemopackage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	res := mydemopackage.Test()
	fmt.Println(res)
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "getJSON API")
	})

	e.GET("/test", hello)

	e.Logger.Fatal(e.Start(":3200"))
}
