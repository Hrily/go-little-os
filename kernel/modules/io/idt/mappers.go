package idt

// ToDescriptorRecord returns the physical representation of Descriptor
func (d *Descriptor) ToDescriptorRecord() DescriptorRecord {
	var record uint64 = 0
	// Higher 32 bits
	// Bit:     | 31              16 | 15 | 14 13 | 12 | 11 | 10 9 8 | 7 6 5 | 4 3 2 1 0 |
	// Content: | offset high        | P  | DPL   | 0  | D  | 1  1 0 | 0 0 0 | reserved  |
	record |= uint64(d.Offset & 0xffff0000)
	record |= uint64(d.GateType) << 8
	// Lower 32 bits
	// Bit:     | 31              16 | 15              0 |
	// Content: | segment selector   | offset low        |
	record <<= 32
	record |= uint64(d.SegmentSelector) << 16
	record |= uint64(d.Offset & 0xffff)
	return DescriptorRecord(record)
}
