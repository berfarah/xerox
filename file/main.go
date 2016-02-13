package file

type File struct{ String string }

type Files []File

func FromArray(a []string) Files {
	files := make(Files, len(a))
	for i, e := range a {
		files[i] = File{e}
	}

	return files
}

func (a Files) ToArray() []string {
	out := make([]string, len(a))
	for i, f := range a {
		out[i] = f.String
	}

	return out
}
