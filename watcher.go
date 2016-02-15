package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/berfarah/xerox/file"
	"github.com/berfarah/xerox/filelocker"
	"github.com/berfarah/xerox/logger"
)

// Watcher is a service
type Watcher struct {
	Pulse      chan bool
	queue      chan string
	logger     *logger.Logger
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
		logger:     xeroxLogger,
		waitGroup:  &sync.WaitGroup{},
		maxWorkers: 8,
		fn:         fn,
	}
}

// Start initiates the watcher
func (w *Watcher) Start() {
	w.waitGroup.Add(1)
	go w.logger.Listen(func(s string) {
		fmt.Println(s, "watcher")
	},
		func(l logger.Entry) bool {
			return l.IsSeverity(1) && l.IsApp("watcher")
		})

	w.waitGroup.Add(1)
	go w.producer()

	for i := 0; i < w.maxWorkers; i++ {
		w.waitGroup.Add(1)
		go w.worker()
	}
}

// Stop watching after executing remaining threads
func (w *Watcher) Stop() {
	w.logger.Stop()
	close(w.Pulse)
	w.waitGroup.Wait()
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
		sortedByCtime := file.FromArray(matches).ByMtime().ToArray()
		for i, file := range sortedByCtime {
			// locked := os.Rename(file, path.Base())
			if i < 5 {
				locked := filelocker.Lock(file)
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
		str := w.fn(image)
		w.logger.Info(str, "worker")
	}
}
