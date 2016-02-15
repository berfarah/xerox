package logger

type Entry struct {
	Severity int
	App      string
	Message  string
}

func (e Entry) IsEmpty() bool {
	return e.Message == ""
}

func (e Entry) IsSeverity(s int) bool {
	return e.Severity >= s
}

func (e Entry) IsApp(a string) bool {
	return e.App == a
}
