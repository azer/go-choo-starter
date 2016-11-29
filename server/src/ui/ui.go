package ui

import (
	"github.com/azer/logger"
	"io/ioutil"
)

var log = logger.New("ui")
var runtime *Runtime
var indexhtml string

func init() {
	runtime = &Runtime{}
	loadIndexHTML()
}

func loadIndexHTML() {
	indexbyt, err := ioutil.ReadFile("./ui/index.html")
	if err != nil {
		panic(err)
	}

	indexhtml = string(indexbyt)
}
