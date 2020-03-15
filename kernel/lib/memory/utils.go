package memory

import "unsafe"

// GetPointerB returns byte pointer to given address
func GetPointerB(address uint32) *byte {
	addr := uintptr(address)
	ptr := unsafe.Pointer(addr)
	return (*byte)(ptr)
}

// PutB puts byte b at given address
func PutB(address uint32, b byte) {
	p := GetPointerB(address)
	*p = b
}

// MoveData moves length bytes from sourceAddress to destAddress
// This does not copy before moving
func MoveData(destAddress, sourceAddress, length uint32) {
	for offset := uint32(0); offset < length; offset++ {
		source := GetPointerB(sourceAddress + offset)
		dest := GetPointerB(destAddress + offset)
		*dest = *source
	}
}
