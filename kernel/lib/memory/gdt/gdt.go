package gdt

import (
	"unsafe"

	"kernel/lib/logger"
)

// GDT represents the Global Descriptor Table
type GDT []*Descriptor

// PopulateRecords populates gdt records
func (g GDT) PopulateRecords(record []DescriptorRecord) {
	// Check lengths are compatible
	if len(g) > len(record) {
		logger.COM().Error("Invalid size of gdt record")
		return
	}

	for i, descriptor := range []*Descriptor(g) {
		record[i] = descriptor.ToDescriptorRecord()
	}
}

var _gdt = [5]*Descriptor{
	nil, // first entry is nil
	&KernelCodeSegment,
	&KernelDataSegment,
	nil, // extra
	nil, // extra
}

var _gdtRecord = [5]DescriptorRecord{
	0, 0, 0, 0, 0,
}

// LoadGDT loads gdt using lgdt instruction
func LoadGDT(gdtAddr uint32, gdtSize uint16)

// Init initializes gdt and loads it
func Init() {
	gdt := GDT(_gdt[:])
	gdtRecord := _gdtRecord[:]
	gdt.PopulateRecords(gdtRecord)

	address := uint32(uintptr(unsafe.Pointer(&_gdtRecord)))
	size := uint16(8 * 5)

	// gdtAddr := uintptr(unsafe.Pointer(&p))
	LoadGDT(address, size)
}
