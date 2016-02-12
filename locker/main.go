package locker

import (
	"os"
	"strings"
)

const extension = ".lock"

// Lock a file
func Lock(f string) string {
	if !islocked(f) {
		return rename(f, f+extension)
	}
	return f
}

// Unlock a file
func Unlock(f string) string {
	if islocked(f) {
		return rename(f, f[:len(f)-len(extension)])
	}
	return f
}

func islocked(f string) bool {
	return strings.HasSuffix(f, extension)
}

func rename(from, to string) string {
	os.Rename(from, to)
	return to
}
