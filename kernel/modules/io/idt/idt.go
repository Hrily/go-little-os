package idt

import (
	"unsafe"
)

var _idt = [256]DescriptorRecord{}

// AddInterrupt adds interrupt to interrupt descriptor table
func AddInterrupt(d Descriptor) {
	_idt[d.Number] = d.ToDescriptorRecord()
}

// LoadIDT loads idt using lidt instruction, defined in load.s
func LoadIDT(address uint32, size uint16)

// Init loads interrupt descriptor table
func Init() {
	address := uint32(uintptr(unsafe.Pointer(&_idt)))
	size := uint16(256 * 4)

	LoadIDT(address, size)
}
