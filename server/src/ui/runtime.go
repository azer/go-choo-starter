package ui

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Runtime struct {
	CachedSourceCode string
}

func (runtime *Runtime) Browserify() error {
	code, err := CompileJS()
	if err != nil {
		return err
	}

	runtime.CachedSourceCode = code
	return nil
}

func (runtime *Runtime) CleanCache() {
	runtime.CachedSourceCode = ""
}

func (runtime *Runtime) CheckForErrors() error {
	_, err := runtime.Render("/", nil)
	return err
}

func (runtime *Runtime) Render(route string, state interface{}) (string, error) {
	body, err := runtime.SourceCode()
	if err != nil {
		return "", err
	}

	encodedState, err := EncodeState(state)
	if err != nil {
		return "", err
	}

	html, err := EvalJS(fmt.Sprintf(`
  var app;
  %s

  app.toString("%s", %s)

  function start (_app) {
    app = _app;
  }`,
		body,
		route,
		encodedState,
	))

	if err != nil {
		return "", err
	}

	return html, nil
}

func (runtime *Runtime) Routes() (map[string]bool, error) {
	body, err := runtime.SourceCode()
	if err != nil {
		return nil, err
	}

	routes, err := EvalJS(fmt.Sprintf(`
  var app;
  var routes;
  %s

  Object.keys(routes).join(',')

  function start (_app, _routes) {
    app = _app;
    routes = _routes
  }`,
		body,
	))

	if err != nil {
		return nil, err
	}

	paths := strings.Split(routes, ",")
	result := map[string]bool{}

	for _, path := range paths {
		result[path] = true
	}

	return result, nil
}

func (runtime *Runtime) SourceCode() (string, error) {
	if runtime.CachedSourceCode != "" {
		return runtime.CachedSourceCode, nil
	}

	if err := runtime.Browserify(); err != nil {
		return "", err
	}

	return runtime.CachedSourceCode, nil
}

func EncodeState(state interface{}) (string, error) {
	encoded, err := json.Marshal(state)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}
