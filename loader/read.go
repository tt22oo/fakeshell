package loader

import (
	"os"

	"github.com/tt22oo/fakeshell/shell/dir"
	"github.com/tt22oo/fakeshell/shell/proc"
)

type ProcConfig struct {
	ProcessPath string
	CpuInfoPath string
	MemInfoPath string
	VersionPath string
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (cfg *ProcConfig) Proc(proc *proc.Proc) error {
	data, err := os.ReadFile(cfg.ProcessPath)
	if err != nil {
		return err
	}

	err = parseProc(data, &proc.Table)
	if err != nil {
		return err
	}

	proc.System.CpuInfo, err = ReadFile(cfg.CpuInfoPath)
	if err != nil {
		return err
	}

	proc.System.MemInfo, err = ReadFile(cfg.MemInfoPath)
	if err != nil {
		return err
	}

	proc.System.Version, err = ReadFile(cfg.VersionPath)
	if err != nil {
		return err
	}

	return nil
}

func Dir(path string, dir *dir.Directory) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return parseDir(data, dir)
}
