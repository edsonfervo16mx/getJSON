package myfilepackage

import (
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
)

type File struct {
	Name      string
	Extension string
	Size      float32
}

//Obtener datos del archivo
func (myFile File) GetMetadata(c echo.Context) bool {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return false
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return false
	}
	// fmt.Println(file)
	fmt.Println("*****")
	fmt.Println(file.Header)
	fmt.Println("*****")
	fmt.Println(reflect.TypeOf(file.Header))
	fmt.Println(file.Header.Values("Content-Disposition"))
	nameFile := file.Header.Values("Content-Disposition")
	fmt.Println(reflect.TypeOf(nameFile))
	for i := range nameFile {
		fmt.Println(nameFile[i])
	}
	defer src.Close()

	return true
}
