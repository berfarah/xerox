package file_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/berfarah/xerox/file"
)

var files = []string{"a", "b", "c", "d", "e"}

func shuffle(a []string) []string {
	dest := make([]string, len(a))
	perm := rand.Perm(len(a))
	for i, v := range perm {
		dest[v] = a[i]
	}
	return dest
}

func TestByName(t *testing.T) {
	shuffled := shuffle(files)
	sorted := file.FromArray(shuffled).ByName().ToArray()

	if !reflect.DeepEqual(sorted, files) {
		t.Errorf("New(%v).ByName() = %v, expected %v", shuffled, sorted, files)
	}
}
