package file

import "sort"

func (a Files) Len() int {
	return len(a)
}

func (a Files) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// =============================================================================
// SORT BY MTIME

// ByMtime sorts by MTIME
func (a Files) ByMtime() Files {
	f := filesByMtime{a}
	sort.Sort(f)
	return f.Files
}

// FilesByMtime gives ability to sort by MTIME
type filesByMtime struct{ Files }

// Less is required by "sort"
func (a filesByMtime) Less(i, j int) bool {
	return a.Files[i].Mtime().Before(a.Files[j].Mtime())
}

// =============================================================================
// SORT BY NAME

// ByName sorts by Name
func (a Files) ByName() Files {
	f := filesByName{a}
	sort.Sort(f)
	return f.Files
}

// FilesByName gives ability to sort by Name
type filesByName struct{ Files }

// Less is required by "sort"
func (a filesByName) Less(i, j int) bool {
	return a.Files[i].String < a.Files[j].String
}

// =============================================================================
// SORT BY SIZE

// BySize sorts by Name
func (a Files) BySize() Files {
	f := filesBySize{a}
	sort.Sort(f)
	return f.Files
}

// FilesBySize gives ability to sort by Size
type filesBySize struct{ Files }

// Less is required by "sort"
func (a filesBySize) Less(i, j int) bool {
	return a.Files[i].Size() < a.Files[j].Size()
}
