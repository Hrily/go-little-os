package systemcall

import (
	"kernel/lib/logger"
	"kernel/modules/io/interrupts/models"
)

// Handle a system call
func Handle(r models.Registers) {
	logger.COM().Error("SystemCall occured")
	logger.COM().LogUint(logger.Debug, "EAX", uint64(r.EAX))
	logger.COM().LogUint(logger.Debug, "EBX", uint64(r.EBX))
	logger.COM().LogUint(logger.Debug, "EIP", uint64(r.EIP))

	switch r.EAX {
	case 0x2d:
		brk(uint32(r.EBX))
	default:
		// Halting
		logger.COM().Error("Halting")
		for true {
		}
	}
}
