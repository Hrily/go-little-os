package processinfo

// Section in process memory
type Section int

const (
	// SectionDataText
	SectionDataText Section = 1 + iota
	// SectionHeap
	SectionHeap
	// SectionStack
	SectionStack
)

// SectionInfo ...
type SectionInfo struct {
	StartAddr, Size, Capacity uint32
}

func GetSectionInfo(info *Info, section Section) *SectionInfo {
	return info.Sections[section]
}
