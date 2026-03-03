package loader

import (
	"encoding/json"

	"github.com/tt22oo/fakeshell/shell/dir"
	"github.com/tt22oo/fakeshell/shell/proc"
)

func parseDir(data []byte, dir *dir.Directory) error {
	dir.Mu.Lock()
	defer dir.Mu.Unlock()
	return json.Unmarshal(data, &dir.Entries)
}

func parseProc(data []byte, pt *proc.ProcessTable) error {
	pt.Mu.Lock()
	defer pt.Mu.Unlock()

	err := json.Unmarshal(data, &pt.Table)
	if err != nil {
		return err
	}

	return nil
}
