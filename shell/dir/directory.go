package dir

import (
	"fmt"
	"strings"
)

func (e *Entry) NewDirectory(name string) {
	e.Children[name] = &Entry{
		Type:     TypeDirectory,
		Children: make(map[string]*Entry),
	}
}

func (dir *Directory) GetEntry(path string) (*Entry, error) {
	if path == "" {
		return nil, fmt.Errorf("invalid path")
	}

	var entry *Entry
	paths := strings.Split(strings.Trim(path, "/"), "/")
	if strings.HasPrefix(path, "/") {
		// /path
		e := dir.Entries[paths[0]]
		if e == nil {
			return nil, fmt.Errorf("%s is not found", path)
		}
		entry = e

		for _, p := range paths[1:] {
			e := entry.Children[p]
			if e == nil {
				return nil, fmt.Errorf("%s is not found", path)
			}
			entry = e
		}
	} else {
		// path1/path2, path1
		e := dir.Entry.Children[paths[0]]
		if e == nil {
			return nil, fmt.Errorf("%s is not found", path)
		}
		entry = e

		for _, p := range paths[1:] {
			e := entry.Children[p]
			if e == nil {
				return nil, fmt.Errorf("%s is not found", path)
			}
			entry = e
		}
	}

	return entry, nil
}
