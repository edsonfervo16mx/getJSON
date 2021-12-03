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
func (myFile File) GetMetadata(c echo.Context) (res bool, name string, extension string, size float32) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return false, "", "", 0
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return false, "", "", 0
	}
	filename := strings.Split(file.Filename, ".")
	defer src.Close()

	return true, filename[0], filename[1], float32(file.Size)
}

func (myFile File) ValidateSize() bool {
	if myFile.Size <= 5000000 {
		return true
	}
	return false
}
