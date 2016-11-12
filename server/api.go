package main

import (
	"github.com/labstack/echo"
	"net/http"
	"ui"
)

// Example of how you can set an API endpoint
func Hi(c echo.Context) error {
	return c.JSON(http.StatusOK, &struct {
		Hi string
	}{"there"})
}

func API() *echo.Echo {
	api := echo.New()
	api.Use(ui.HTTPHandler)

	/* Add your API endpoints below.
	I recommend you to create a folder under ./src and locate your business logic there
	to keep here separate from your API business logic. */
	api.Get("/api/hi", Hi)

	return api
}
