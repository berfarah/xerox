package file

import "time"

func (a Files) filter(fn func(File) bool) Files {
	files := Files{}
	for _, e := range a {
		if fn(e) {
			files = append(files, e)
		}
	}

	return files
}

func (a Files) MtimeBefore(t time.Time) Files {
	return a.filter(func(f File) bool {
		return f.Mtime().Before(t)
	})
}

func (a Files) MtimeAfer(t time.Time) Files {
	return a.filter(func(f File) bool {
		return f.Mtime().After(t)
	})
}

func (a Files) SizeBiggerThan(size int64) Files {
	return a.filter(func(f File) bool {
		return f.Size() > size
	})
}

func (a Files) SizeSmallerThan(size int64) Files {
	return a.filter(func(f File) bool {
		return f.Size() < size
	})
}
