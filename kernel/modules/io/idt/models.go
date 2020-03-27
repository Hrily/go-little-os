package idt

// DescriptorRecord is the physical representation of IDT record
type DescriptorRecord uint64

// GateType represents gate type of the interrrupt
type GateType uint8

// Descriptor is an entry in IDT
type Descriptor struct {
	Number          uint16
	Offset          uint32
	SegmentSelector uint16
	GateType        GateType
}
