package filewatcher

import (
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/berfarah/xerox/filesort"
	"github.com/berfarah/xerox/locker"
)

// Watcher is a service
type Watcher struct {
	Pulse      chan bool
	queue      chan string
	log        chan string
	dir        string
	waitGroup  *sync.WaitGroup
	maxWorkers int
	fn         func(string) string
}

// New creates a new Watcher
func New(dir string, fn func(string) string) *Watcher {
	return &Watcher{
		Pulse:      make(chan bool),
		dir:        dir,
		queue:      make(chan string, 16),
		log:        make(chan string, 16),
		waitGroup:  &sync.WaitGroup{},
		maxWorkers: 8,
		fn:         fn,
	}
}

// Start initiates the watcher
func (w *Watcher) Start() {
	w.waitGroup.Add(1)
	go w.logger()

	w.waitGroup.Add(1)
	go w.producer()

	for i := 0; i < w.maxWorkers; i++ {
		w.waitGroup.Add(1)
		go w.worker()
	}
}

// Stop watching after executing remaining threads
func (w *Watcher) Stop() {
	defer close(w.log)
	close(w.Pulse)
	w.waitGroup.Wait()
}

func (w *Watcher) logger() {
	for message := range w.log {
		log.Println(message)
	}
}

func (w *Watcher) producer() {
	defer w.waitGroup.Done()
	for {
		select {
		case <-w.Pulse:
			return
		default:
		}

		// do dir globbing and send to chan
		matches, _ := filepath.Glob(filepath.Join(w.dir, "/*[^lock]"))
		sortedByCtime := filesort.New(matches).ByCtime()
		for i, file := range sortedByCtime {
			// locked := os.Rename(file, path.Base())
			if i < 5 {
				locked := locker.Lock(file)
				w.queue <- locked
			}
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func (w *Watcher) worker() {
	defer w.waitGroup.Done()
	for image := range w.queue {
		// w.log <- slice(image)

		// TEMPORARY CODE
		// time.Sleep(1 * time.Second)
		w.log <- w.fn(image)
	}
}
