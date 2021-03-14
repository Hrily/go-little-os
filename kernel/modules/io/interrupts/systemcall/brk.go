package systemcall

import (
	"kernel/lib/logger"
	"kernel/modules/process/processinfo"
)

// brk systemcall
// sets the program break to a new given addr
// returns new addresss on success, otherwise returns the given addr
func brk(addr uint32) uint32 {
	logger.COM().LogUint(logger.Debug, "brk addr:", uint64(addr))

	// TODO: Remove hard coding
	processBrk := processinfo.KernelSectionHeap.StartAddr + processinfo.KernelSectionHeap.Size

	// Check if caller is requesting program break
	if addr == 0 {
		return processBrk
	}

	// Validate
	if addr <= processBrk {
		return addr
	}

	// If we have capacity to increase brk point
	if addr <= processinfo.KernelSectionHeap.StartAddr+processinfo.KernelSectionHeap.Capacity {
		return addr
	}

	return 0
}
