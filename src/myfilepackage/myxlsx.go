package myfilepackage

import (
	"github.com/labstack/echo/v4"
)

type Xlsx struct {
	Book  string
	Sheet string
}

func (myXlsx Xlsx) GetData(c echo.Context) (res bool) {

	return true
}
