package ui

import (
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

func HTTPHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		routes, err := runtime.Routes()
		if err != nil {
			log.Error("Can not get routes. Error: %v", err)
			return err
		}

		// Check if it's a URL registered by Choo. Continue with next handler if not.
		path := c.Request().URL().Path()
		if _, ok := routes[path]; !ok {
			return next(c)
		}

		body, err := runtime.Render(path)
		if err != nil {
			log.Error("Can not render. Error: %v", err)
			return err
		}

		return c.HTML(http.StatusOK, fmt.Sprintf(indexhtml, body))
	}
}
