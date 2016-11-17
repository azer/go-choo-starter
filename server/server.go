package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"os"
	"ui"
)

func init() {
	// Read env variables from .env file in the root directory
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	server := echo.New()
	SetUIRoutes(server)
	SetAPIRoutes(server)

	// Setup static files
	server.Static("/public", "./public")
	server.Get("/favicon.ico", func(c echo.Context) error {
		c.Redirect(301, "/public/favicon.ico")
		return nil
	})

	// Watch & build front-end stuff if develop mode is enabled
	if len(os.Getenv("DEVELOP")) > 0 {
		go ui.WatchCodeChanges()
	}

	// And finally, run the server. You can edit the ADDR from .env file on the project folder.
	server.Run(standard.New(os.Getenv("ADDR")))
}
