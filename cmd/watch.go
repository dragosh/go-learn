package main

import (
	"log"

	"github.com/dragosh/go-learn/pkg/watcher"
)

func StartWatcher() {
	watcher.Start(".", func(path string) {
		log.Println(path)
	})
}

func main() {
	StartWatcher()

}
