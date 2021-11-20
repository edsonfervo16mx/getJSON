package routes

import (
	"getJSON/src/server/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "getJSON API")
	})

	e.GET("/test", handlers.Hello)
	e.GET("/test/package", handlers.Demo)

	e.GET("/file", handlers.ReadFile)
	e.Logger.Fatal(e.Start(":3200"))
}
