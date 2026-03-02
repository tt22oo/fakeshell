package dir

import (
	"fmt"
	"strings"
)

func (d *Dir) GetDir(path string) (*Entry, error) {
	if path == "" {
		return nil, fmt.Errorf("invalid path")
	}
	d.Mu.Lock()
	defer d.Mu.Unlock()

	path = strings.Trim(path, "/")
	if !strings.Contains(path, "/") {
		entry := d.Entries[path]
		if entry == nil {
			return nil, fmt.Errorf("%s is not found", path)
		}
		return entry, nil
	}

	paths := strings.Split(path, "/")
	entry := d.Entries[paths[0]]
	if entry == nil {
		return nil, fmt.Errorf("%s is not found", path)
	}

	for _, p := range paths[1:] {
		entry = entry.Children[p]
		if entry == nil {
			return nil, fmt.Errorf("%s is not found", path)
		}
	}

	return entry, nil
}
