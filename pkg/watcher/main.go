package watcher

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

func Start() {

	watch := watcher.New()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if err := watch.Add(cwd); err != nil {
		log.Fatalln(err)
	}

	watch.IgnoreHiddenFiles(true)

	watch.AddFilterHook(watcher.RegexFilterHook(regexp.MustCompile(".md$"), false))

	go func() {
		for {
			select {
			case event := <-watch.Event:
				log.Println("File Changed: ", event) // Print the event's info.
			case err := <-watch.Error:
				log.Fatalln(err)
			case <-watch.Closed:
				return
			}
		}
	}()

	for path, f := range watch.WatchedFiles() {
		fmt.Printf("Watching: %s: %s\n", path, f.Name())
	}
	// Start the watching process - it'll check for changes every 100ms.
	if err := watch.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}

}
