package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"os"
)

func init() {
	// Import env files
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	api := API()

	// Setup static files
	api.Static("/public", "./public")
	api.Get("/favicon.ico", func(c echo.Context) error {
		c.Redirect(301, "/public/favicon.ico")
		return nil
	})

	// Watch & build front-end stuff if develop mode is enabled
	if len(os.Getenv("DEVELOP")) > 0 {
		go Develop()
	}

	// And finally, run the server. You can edit the ADDR from .env file on the project folder.
	api.Run(standard.New(os.Getenv("ADDR")))
}
