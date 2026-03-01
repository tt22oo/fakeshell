package proc

import "sync"

type Process struct {
	User string   `json:"user"`
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

type ProcessTable struct {
	Mu    sync.Mutex
	Table map[int]*Process
}

type SystemInfo struct {
	CpuInfo string
	MemInfo string
	Version string
}

type Proc struct {
	System SystemInfo
	Table  ProcessTable
}
