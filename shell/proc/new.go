package proc

import (
	"sort"
)

func newPID(pt *ProcessTable) int {
	var pids []int
	for pid := range pt.Table {
		pids = append(pids, pid)
	}
	sort.Ints(pids)

	if len(pids) < 1 {
		return 1
	}
	return pids[len(pids)-1] + 1
}

func (p *Process) New(pt *ProcessTable) int {
	pt.Mu.Lock()
	defer pt.Mu.Unlock()

	pid := newPID(pt)
	pt.Table[pid] = p

	return pid
}
