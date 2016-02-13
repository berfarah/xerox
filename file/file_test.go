package file_test

import (
	"os"
	"testing"
	"time"

	"github.com/berfarah/xerox/file"
)

var testFileName = "/tmp/foo"

func setup() {
	f, _ := os.Create(testFileName)
	defer f.Close()
}

func writestuff() {
	f, _ := os.OpenFile(testFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	f.WriteString("Asdf")
	defer f.Close()
}

func teardown() {
	os.Remove(testFileName)
}

func TestMtime(t *testing.T) {
	setup()

	testFile := file.File{testFileName}
	initialMtime := testFile.Mtime()

	if initialMtime.After(time.Now()) {
		t.Error("Expected file.Mtime() to have already happened")
	}

	time.Sleep(1 * time.Second)
	writestuff()

	if testFile.Mtime().Equal(initialMtime) {
		t.Errorf("Expected file.Mtime() to change when the file was updated.\nInit:\n%v\nCurrent:\n%v", initialMtime, testFile.Mtime())
	}

	teardown()
}

func TestName(t *testing.T) {
	setup()

	testFile := file.File{testFileName}
	if got := testFile.Name(); got != "foo" {
		t.Errorf("Expected file.Name() to be foo, got %v", got)
	}

	teardown()
}

func TestSize(t *testing.T) {
	setup()

	testFile := file.File{testFileName}
	if got := testFile.Size(); got != 0 {
		t.Errorf("Expected empty file to be 0 bytes, got %v", got)
	}

	// Add 4 bytes
	writestuff()

	if got := testFile.Size(); got != 4 {
		t.Errorf("Expected file with 'Asdf' to be 4 bytes, got %v", got)
	}

	teardown()
}
