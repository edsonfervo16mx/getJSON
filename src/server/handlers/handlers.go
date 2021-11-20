package handlers

import (
	"fmt"
	"getJSON/src/mydemopackage"
	"getJSON/src/myfilepackage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Demo(c echo.Context) error {
	fmt.Println("Demo...")
	return c.String(http.StatusOK, "on routes")
}

func Hello(c echo.Context) error {
	res := mydemopackage.Test()
	fmt.Println(res)
	return c.String(http.StatusOK, "Hello, World!")
}

/**/
func ReadFile(c echo.Context) error {
	fmt.Println("Read File...")
	var myFile myfilepackage.File
	res := myFile.GetMetadata()
	fmt.Println(res)
	return c.String(http.StatusOK, "Read finish")

}

/**/
