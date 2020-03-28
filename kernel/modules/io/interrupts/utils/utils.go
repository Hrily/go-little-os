package utils

import (
	"kernel/modules/io/idt"
)

func GetFuncAddr(f func()) uint32

func LoadIntHandler(n uint16, handler func()) {
	var ptr uint32 = GetFuncAddr(handler)
	idt.AddInterrupt(idt.Descriptor{
		Offset:          ptr,
		Number:          uint16(n),
		SegmentSelector: 0x08,
		GateType:        idt.InterruptGate,
	})
}
