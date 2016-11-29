package main

import (
	"github.com/labstack/echo"
	"ui"
)

func SetUIRoutes(server *echo.Echo) {
	// Client-side routes will already be compiled and routed by the server automatically.
	// But you may want to do call the UI manually in some cases like passing a custom state.
	// Check this ExampleForm function to see how we accomplish that.
	// You can try commenting out following line and visiting the same URL again.
	// server.Get("/", ExampleForm)

	// If there is a UI route matching the URL path, it'll respond first. Otherwise, server will try to match other endpoints.
	server.Use(ui.HTTPHandler)
}

/*func ExampleForm(c echo.Context) error {
	return ui.Render(c, "/", &struct {
		ExampleTitle string `json:"title"`
	}{"Hello From The Server-side"})
}*/
