package ui

type State struct {
	Params   map[string]string `json:"params"`
	Location struct {
		Pathname string `json:"pathname"`
	} `json:"location"`
}
