package logger

const (
	Debug   = 0
	Info    = 1
	Warn    = 2
	Error   = 3
	Fatal   = 4
	Unknown = 5
)

func (l *Logger) Debug(message, app string) error {
	return l.write(Debug, message, app)
}

func (l *Logger) Info(message, app string) error {
	return l.write(Info, message, app)
}

func (l *Logger) Warn(message, app string) error {
	return l.write(Warn, message, app)
}

func (l *Logger) Error(message, app string) error {
	return l.write(Error, message, app)
}

func (l *Logger) Fatal(message, app string) error {
	return l.write(Fatal, message, app)
}

func (l *Logger) Unknown(message, app string) error {
	return l.write(Unknown, message, app)
}
