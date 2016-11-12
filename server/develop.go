package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"ui"
)

func Develop() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case _ = <-watcher.Event:
				ui.BuildUI()
			}
		}
	}()

	err = watcher.Watch("./ui")
	if err != nil {
		log.Fatal(err)
	}

	<-done
	watcher.Close()
}
