package processinfo

const (
	// KernelPID
	KernelPID uint32 = 0
)

var (
	KernelSectionDataText = SectionInfo{
		StartAddr: 0,
		Size:      0,
		Capacity:  0,
	}
	KernelSectionHeap = SectionInfo{
		StartAddr: 0,
		Size:      0,
		Capacity:  0,
	}
	KernelSectionStack = SectionInfo{
		StartAddr: 0,
		Size:      0,
		Capacity:  0,
	}

	KernelInfo = Info{
		PID: KernelPID,
		Sections: map[Section]*SectionInfo{
			SectionDataText: &KernelSectionDataText,
			SectionHeap:     &KernelSectionHeap,
			SectionStack:    &KernelSectionStack,
		},
	}
)
