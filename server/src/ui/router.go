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

	return c.HTML(http.StatusOK, fmt.Sprintf(indexhtml, body))
}

func HTTPHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := runtime.SyncRoutes(); err != nil {
			return err
		}

		path := c.Request().URL.Path
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
