package ui

import (
	"github.com/azer/logger"
	"gopkg.in/olebedev/go-duktape.v2"
	"strings"
)

func EvalJS(code string) (string, error) {
	ctx := duktape.New()

	if err := ctx.PevalString(code); err != nil {
		derr := err.(*duktape.Error)
		lines := strings.Split(code, "\n")

		log.Error("Failed to evaluate JavaScript.", logger.Attrs{
			"error": derr.Message,
			"code":  lines[derr.LineNumber-1],
		})

		return "", err
	}

	result := ctx.GetString(-1)
	ctx.DestroyHeap()

	return result, nil
}
