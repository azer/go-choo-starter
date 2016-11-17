package ui

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

var runtime *Runtime
var indexhtml string

func init() {
	runtime = &Runtime{}
	indexbyt, err := ioutil.ReadFile("./ui/index.html")
	if err != nil {
		panic(err)
	}

	indexhtml = string(indexbyt)
}

func Route(c echo.Context, path string, state interface{}) error {
	if ok, err := IsValidUIRoute(path); err != nil {
		return err
	} else if !ok {
		return errors.New(fmt.Sprintf("Invalid UI Route '%s'.", path))
	}

	body, err := runtime.Render(path, state)
	if err != nil {
		log.Error("Can not render. Error: %v", err)
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf(indexhtml, body))
}

func IsValidUIRoute(path string) (bool, error) {
	routes, err := runtime.Routes()
	if err != nil {
		log.Error("Can not get list of UI routes. Error: %v", err)
		return false, err
	}

	// Check if it's a URL registered by Choo. Continue with next handler if not.
	_, ok := routes[path]
	return ok, nil
}

func HTTPHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL().Path()

		if ok, err := IsValidUIRoute(path); err != nil {
			return err
		} else if !ok {
			return next(c)
		}

		log.Info("%s is being automatically routed to UI component", path)

		return Route(c, path, nil)
	}
}
