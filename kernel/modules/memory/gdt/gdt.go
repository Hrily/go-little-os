package gdt

import (
	"unsafe"

	"kernel/lib/logger"
)

const (
	nDescriptors = 6
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

var _gdt = [nDescriptors]*Descriptor{
	nil, // first entry is nil
	&KernelCodeSegment,
	&KernelDataSegment,
	&KernelTLSSegment,
	nil, // extra
	nil, // extra
}

var _gdtRecord = [nDescriptors]DescriptorRecord{
	0, 0, 0, 0, 0,
}

// LoadGDT loads gdt using lgdt instruction, defined in load.s
func LoadGDT(gdtAddr uint32, gdtSize uint16)

// AddToGDT adds given descriptor at index
func AddToGDT(index int, d *Descriptor) bool {
	if index > len(_gdt) {
		logger.COM().Error("index greater than gdt size")
		return false
	}
	_gdtRecord[index] = d.ToDescriptorRecord()
	return true
}

// Init initializes gdt and loads it
func Init() {
	gdt := GDT(_gdt[:])
	gdtRecord := _gdtRecord[:]
	gdt.PopulateRecords(gdtRecord)

	address := uint32(uintptr(unsafe.Pointer(&_gdtRecord)))
	size := uint16(8 * nDescriptors)

	// gdtAddr := uintptr(unsafe.Pointer(&p))
	LoadGDT(address, size)
}
