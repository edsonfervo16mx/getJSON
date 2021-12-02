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

func ReadFile(c echo.Context) error {
	apikey := c.FormValue("apikey")
	fmt.Println("***APIKEY***")
	fmt.Println(apikey)
	var myFile myfilepackage.File
	// Obtener Metadatos del Archivo
	res, name, extension := myFile.GetMetadata(c)
	fmt.Println("***METADATOS***")
	fmt.Println(res)
	fmt.Println(name)
	fmt.Println(extension)

	return c.String(http.StatusOK, "Read finish")

}
