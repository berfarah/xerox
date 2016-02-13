package filelocker_test

import (
	"testing"

	"github.com/berfarah/xerox/filelocker"
)

type filelockerCase struct {
	in, expected string
}

var lockTests = []filelockerCase{
	{"/tmp/foo/bar.lock", "/tmp/foo/bar.lock"},
	{"/tmp/foo/bar.jpg.lock", "/tmp/foo/bar.jpg.lock"},
	{"/tmp/foo/bar", "/tmp/foo/bar.lock"},
	{"/tmp/foo/bar.jpg", "/tmp/foo/bar.jpg.lock"},
}

func TestLock(t *testing.T) {
	for _, c := range lockTests {
		if out := filelocker.Lock(c.in); out != c.expected {
			t.Errorf("Lock(%v) = %v, expected %v", c.in, c.expected, out)
		}
	}
}

var unlockTests = []filelockerCase{
	{"/tmp/foo/bar.lock", "/tmp/foo/bar"},
	{"/tmp/foo/bar.jpg.lock", "/tmp/foo/bar.jpg"},
	{"/tmp/foo/bar", "/tmp/foo/bar"},
	{"/tmp/foo/bar.jpg", "/tmp/foo/bar.jpg"},
}

func TestUnlock(t *testing.T) {
	for _, c := range unlockTests {
		if out := filelocker.Unlock(c.in); out != c.expected {
			t.Errorf("Unlock(%v) = %v, expected %v", c.in, c.expected, out)
		}
	}
}
