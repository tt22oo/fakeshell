package proc

func (pt *ProcessTable) Delete(pid int) {
	pt.Mu.Lock()
	defer pt.Mu.Unlock()

	delete(pt.Table, pid)
}
