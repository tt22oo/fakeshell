package shell

import (
	"fmt"

	"github.com/tt22oo/fakeshell/loader"
	"github.com/tt22oo/fakeshell/shell/dir"
	"github.com/tt22oo/fakeshell/shell/proc"
)

type Shell struct {
	User string
	Dir  dir.Directory
	Proc proc.Proc
}

type Config struct {
	HomePath string
	DirPath  string
	Proc     *loader.ProcConfig
}

func (cfg *Config) New() (*Shell, error) {
	var s Shell
	err := loader.Dir(cfg.DirPath, &s.Dir)
	if err != nil {
		return nil, fmt.Errorf("Read Dir Error: %s", err.Error())
	}

	err = cfg.Proc.Proc(&s.Proc)
	if err != nil {
		return nil, fmt.Errorf("Read Proc Error: %s", err.Error())
	}

	s.Dir.HomePath = cfg.HomePath
	s.Dir.Entry, err = s.Dir.GetEntry(cfg.HomePath)
	if err != nil {
		return nil, fmt.Errorf("Home Directory Error: %s", err.Error())
	}

	return &s, nil
}
