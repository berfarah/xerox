package locker_test

import (
	"testing"

	"github.com/berfarah/xerox/locker"
)

type LockerCase struct {
	in, expected string
}

var lockTests = []LockerCase{
	{"/tmp/foo/bar.lock", "/tmp/foo/bar.lock"},
	{"/tmp/foo/bar.jpg.lock", "/tmp/foo/bar.jpg.lock"},
	{"/tmp/foo/bar", "/tmp/foo/bar.lock"},
	{"/tmp/foo/bar.jpg", "/tmp/foo/bar.jpg.lock"},
}

func TestLock(t *testing.T) {
	for _, c := range lockTests {
		if out := locker.Lock(c.in); out != c.expected {
			t.Errorf("Lock(%v) = %v, expected %v", c.in, c.expected, out)
		}
	}
}

var unlockTests = []LockerCase{
	{"/tmp/foo/bar.lock", "/tmp/foo/bar"},
	{"/tmp/foo/bar.jpg.lock", "/tmp/foo/bar.jpg"},
	{"/tmp/foo/bar", "/tmp/foo/bar"},
	{"/tmp/foo/bar.jpg", "/tmp/foo/bar.jpg"},
}

func TestUnlock(t *testing.T) {
	for _, c := range unlockTests {
		if out := locker.Unlock(c.in); out != c.expected {
			t.Errorf("Unlock(%v) = %v, expected %v", c.in, c.expected, out)
		}
	}
}
