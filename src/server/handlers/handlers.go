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
	res, name, extension, size := myFile.GetMetadata(c)
	fmt.Println("***METADATOS***")
	myFile.Name = name
	myFile.Extension = extension
	myFile.Size = size
	fmt.Println(myFile.Name)
	fmt.Println(myFile.Extension)
	fmt.Println(myFile.Size)
	result := myFile.ValidateSize()
	if !res || !result {
		fmt.Println("Error")
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	fmt.Println(result)
	up := myFile.Upload(c)
	fmt.Println(up)

	switch myFile.Extension {
	case "xlsx":
		fmt.Println("---------------")
		var myXlsx myfilepackage.Xlsx
		myXlsx.Book = myFile.Name
		data := myXlsx.GetData(c)
		fmt.Println(data)

	default:
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.String(http.StatusOK, "Read finish")

}
