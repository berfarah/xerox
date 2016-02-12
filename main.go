package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/berfarah/xerox/filewatcher"
	"github.com/berfarah/xerox/locker"
)

func main() {
	fromDir := "/Users/bfarah1/Desktop/example/1"
	toDir := "/Users/bfarah1/Desktop/example/2"

	w := filewatcher.New(fromDir, func(f string) string {
		toLocation := path.Join(toDir, path.Base(f))
		os.Rename(f, toLocation)
		unlocked := locker.Unlock(toLocation)
		time.Sleep(1 * time.Second)
		return "Moved " + unlocked
	})

	w.Start()

	log.Println(<-w.Pulse)
}
