package gdt

// DescriptorRecord represents the physical structure of a descriptor.
// This is the actual bit representation of descriptor.
type DescriptorRecord uint64

// Granularity is the granularity of the blocks
type Granularity bool

// SegmentAddressMode is the addressing mode of the segment
type SegmentAddressMode uint8

// Flags are the segment flags
type Flags struct {
	Granularity Granularity
	AddressMode SegmentAddressMode
}

// PrivilageLevel is the privilage level of the segment
type PrivilageLevel uint8

// DescriptorType represents the type of descriptor
type DescriptorType bool

// Access represents the access specifiers for the segment
type Access struct {
	IsPresentInMemory bool
	PrivilageLevel    PrivilageLevel
	DescriptorType    DescriptorType
	IsExecutable      bool
	IsExpandingDown   bool
	// IsConforming tells if code in segment can be executed with lower privilage
	// level
	IsConforming bool
	IsReadable   bool
	IsWritable   bool
	IsAccessed   bool
}

// Descriptor is a segment descriptor in GDT
type Descriptor struct {
	Base   uint32
	Limit  uint32
	Flags  *Flags
	Access *Access
}
