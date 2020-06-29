package pointer

import (
	"unsafe"
)

// Get returns a pointer to given address
func Get(addr uint32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(addr))
}
