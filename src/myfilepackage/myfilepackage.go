package myfilepackage

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type File struct {
	Name      string
	Extension string
	Size      float32
}

//Obtener datos del archivo
func (myFile File) GetMetadata(c echo.Context) (res bool, name string, extension string) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return false, "", ""
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return false, "", ""
	}
	filename := strings.Split(file.Filename, ".")
	defer src.Close()

	return true, filename[0], filename[1]
}
