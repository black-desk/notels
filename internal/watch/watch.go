package watch

import (
	"github.com/fsnotify/fsnotify"

	"github.com/black-desk/notels/internal/utils/logger"
)

var log = logger.Get("notels.watch")

func Watch(workspacePath string) error {
	log.Debugw("Watch called",
		"workspacePath", workspacePath)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// go func() {
	// for {
	// select {
	// case event, ok := <-watcher.Events:
	// if !ok {
	// return
	// }
	// log.Debugw("fs event arrived", "event", event)
	// if event.Has(fsnotify.Write) {
	// event.Name != html
	// log.Infow("ready run md2html")
	// log.Println("modified file:")

	// }
	// case err, ok := <-watcher.Errors:
	// if !ok {
	// return
	// }
	// // log.Println("error:", err)
	// }
	// }
	// }()

	// err = watcher.Add("/tmp")
	// if err != nil {
	// log.Fatal(err)
	// }

	// <-make(chan struct{})

	return nil
}
