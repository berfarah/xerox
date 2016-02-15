package logger

type Logger struct {
	ch chan Entry
}

func New() *Logger {
	return &Logger{make(chan Entry, 16)}
}

func (l *Logger) write(severity int, message, app string) error {
	l.ch <- Entry{severity, message, app}
	return nil
}

func (l *Logger) Listen(fn func(string), filter func(Entry) bool) {
	for e := range l.ch {
		if filter(e) {
			fn(e.Message)
		}
	}
}

func (l *Logger) Stop() {
	close(l.ch)
}
