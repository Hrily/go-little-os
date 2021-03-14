package gdt

import "unsafe"

const (
	// ByteGranularity represents granularity of 1B blocks
	ByteGranularity Granularity = false
	// PageGranularity represents granularity of 4KiB blocks
	PageGranularity Granularity = true

	// SegmentAddressMode16b represents 16 bit protected mode addressing of the
	// segment
	SegmentAddressMode16b SegmentAddressMode = 0
	// SegmentAddressMode32b represents 32 bit protected mode addressing of the
	// segment
	SegmentAddressMode32b SegmentAddressMode = 1

	// Ring0 is the privilage level 0
	Ring0 PrivilageLevel = 0
	// Ring1 is the privilage level 1
	Ring1 PrivilageLevel = 1
	// Ring2 is the privilage level 2
	Ring2 PrivilageLevel = 2
	// Ring3 is the privilage level 3
	Ring3 PrivilageLevel = 3

	// SystemDescriptor represents that the descriptor points to system segment
	SystemDescriptor DescriptorType = false
	// CodeOrDataDescriptor represents that the descriptor points to code or data
	// descriptor
	CodeOrDataDescriptor DescriptorType = true
)

// KernelSegmentFlags are flags for kernel segment
var KernelSegmentFlags = Flags{
	Granularity: PageGranularity,
	AddressMode: SegmentAddressMode32b,
}

// KernelCodeSegmentAccess are access values for kernel Code segment
var KernelCodeSegmentAccess = Access{
	IsPresentInMemory: true,
	PrivilageLevel:    Ring0,
	DescriptorType:    CodeOrDataDescriptor,
	IsExecutable:      true,
	IsConforming:      false,
	IsReadable:        true,
	IsAccessed:        false,
}

// KernelCodeSegment is the kernel's code segment descriptor
var KernelCodeSegment = Descriptor{
	Base:   0,
	Limit:  0xffffffff,
	Flags:  &KernelSegmentFlags,
	Access: &KernelCodeSegmentAccess,
}

// KernelDataSegmentAccess are access values for kernel Data segment
var KernelDataSegmentAccess = Access{
	IsPresentInMemory: true,
	PrivilageLevel:    Ring0,
	DescriptorType:    CodeOrDataDescriptor,
	IsExecutable:      false,
	IsExpandingDown:   false,
	IsWritable:        true,
	IsAccessed:        false,
}

// KernelDataSegment is the kernel's code segment descriptor
var KernelDataSegment = Descriptor{
	Base:   0,
	Limit:  0xffffffff,
	Flags:  &KernelSegmentFlags,
	Access: &KernelDataSegmentAccess,
}

const tlsStorageSize = 1024

var tlsStorage = [tlsStorageSize]byte{}

// KernelTSSSegmentAccess are access values for kernel Data segment
var KernelTSSSegmentAccess = Access{
	IsPresentInMemory: true,
	PrivilageLevel:    Ring3,
	DescriptorType:    CodeOrDataDescriptor,
	IsExecutable:      false,
	IsExpandingDown:   false,
	IsWritable:        true,
	IsAccessed:        false,
}

// KernelTLSSegment is the kernel's thread level storage segment
var KernelTLSSegment = Descriptor{
	Base:   uint32(uintptr(unsafe.Pointer(&_gdtRecord))),
	Limit:  uint32(uintptr(unsafe.Pointer(&_gdtRecord))) + tlsStorageSize,
	Flags:  &KernelSegmentFlags,
	Access: &KernelTSSSegmentAccess,
}
