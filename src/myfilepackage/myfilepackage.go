package myfilepackage

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type File struct {
	Name      string
	Extension string
	Size      float32
	Path      string
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

func (myFile File) Upload(c echo.Context) (res bool) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return false
	}
	src, err := file.Open()
	if err != nil {
		return false
	}
	defer src.Close()

	dst, err := os.Create(os.Getenv("PATH_FILES") + myFile.Extension + "/" + file.Filename)
	if err != nil {
		return false
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return false
	}

	return true
}
