package main

import (
	"fmt"
	"getJSON/src/server/routes"
	_ "net/http"

	_ "github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Start getJSON API...")

	routes.ConfigureRoutes()
}
