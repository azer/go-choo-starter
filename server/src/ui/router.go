package ui

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Render(c echo.Context, path string, state *State) error {
	body, err := runtime.Render(path, state)
	if err != nil {
		log.Error("Can not render. Error: %v", err)
		return err
	}

	html := fmt.Sprintf(indexhtml, body)
	cache.Set(path, html)

	return c.HTML(http.StatusOK, html)
}

func HTTPHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path

		if html, ok := cache.Get(path); ok {
			return c.HTML(http.StatusOK, html)
		}

		if err := runtime.SyncRoutes(); err != nil {
			return err
		}

		match := runtime.Routes.Match(path)

		if match == nil {
			return next(c)
		}

		log.Info("%s is being automatically routed to UI component", path)

		state := &State{}
		state.Location.Pathname = path
		state.Params = match.Params

		return Render(c, path, state)
	}
}
