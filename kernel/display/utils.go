package display

import (
	"unsafe"
)

const (
	_frameBufferAddress = 0x000B8000
)

func putChar(index int32, char byte) {
	var addr uintptr = uintptr(_frameBufferAddress + index)
	ptr := unsafe.Pointer(addr)
	p := (*byte)(ptr)
	*p = char
}

func FrameBufferWriteCell(
	index int32, char byte, fg, bg uint8,
) {
	putChar(index, char)
	putChar(index+1, ((fg&0x0F)<<4)|(bg&0x0F))
}
