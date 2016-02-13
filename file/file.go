package file

import (
	"os"
	"time"
)

func (f File) Mtime() time.Time {
	stat, _ := os.Stat(f.String)
	return stat.ModTime()
}

func (f File) Name() string {
	stat, _ := os.Stat(f.String)
	return stat.Name()
}

func (f File) Size() int64 {
	stat, _ := os.Stat(f.String)
	return stat.Size()
}
