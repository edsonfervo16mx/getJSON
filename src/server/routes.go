package routes

import (
	_ "fmt"
	_ "getJSON/src/mydemopackage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func demo(c echo.Context) error {
	return c.String(http.StatusOK, "on routes")
}

func ConfigureRoutes() {
	e := echo.New()
	e.GET("/test/package", demo)
}
