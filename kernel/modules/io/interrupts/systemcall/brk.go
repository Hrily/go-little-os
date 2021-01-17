package systemcall

import (
	"kernel/lib/logger"
)

// brk systemcall
// sets the program break to a new given addr
// returns new addresss on success, otherwise returns the given addr
func brk(addr uint32) uint32 {
	logger.COM().LogUint(logger.Debug, "brk addr:", uint64(addr))
	return addr
}
