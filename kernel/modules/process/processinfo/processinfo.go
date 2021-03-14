package processinfo

// Info of a process
type Info struct {
	PID      uint32
	Sections map[Section]*SectionInfo
}

var info map[uint32]*Info = map[uint32]*Info{
	KernelPID: &KernelInfo,
}

func GetProcessInfo(pid uint32) *Info {
	return info[pid]
}
