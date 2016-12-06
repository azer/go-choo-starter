package ui

import (
	"github.com/azer/logger"
	. "github.com/azer/on-tree-change"
	"os"
	"strings"
	"time"
)

const CHANGE_THRESHOLD_MS = 250

func CheckRuntimeErrors() {
	if err := runtime.CheckForErrors(); err == nil {
		log.Info("🏁  JavaScript runtime is ready for server-side rendering.")
	} else {
		log.Error("JavaScript runtime could not be compiled.", logger.Attrs{
			"error": err,
		})
	}
}

func WatchCodeChanges() {
	go CheckRuntimeErrors()

	var lastchange = 0

	OnTreeChange("ui", filter, func(name string) {
		now := int(time.Now().UnixNano() / 1000000)

		if now-lastchange < CHANGE_THRESHOLD_MS {
			return
		}

		lastchange = now

		BuildUI()

		runtime.CleanCache()
		CheckRuntimeErrors()

		loadIndexHTML()
	})
}

func filter(name string, info os.FileInfo) bool {
	return (info != nil && !info.IsDir()) && !strings.Contains(name, "node_modules/")
}
