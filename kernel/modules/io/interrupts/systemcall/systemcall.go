package systemcall

import (
	"kernel/lib/logger"
)

// Params for a SystemCall
type Params struct {
	EBP, EDI, ESI, EDX, ECX, EBX, EAX uint32
}

// Handle a system call
func SystemCall(p Params) uint32 {
	logger.COM().Error("SystemCall occured")
	logger.COM().LogUint(logger.Debug, "EAX", uint64(p.EAX))
	logger.COM().LogUint(logger.Debug, "EBX", uint64(p.EBX))

	switch p.EAX {
	case 0x2d:
		return brk(uint32(p.EBX))
	case 0xc7:
		return geteuid()
	case 0xc8:
		return getgid()
	case 0xc9:
		return geteuid()
	case 0xca:
		return getegid()
	case 0xf3:
		return uint32(setThreadArea(uint32(p.EBX)))
	default:
		// Halting
		logger.COM().LogUint(logger.Debug, "Unknown SystemCall: ", uint64(p.EAX))
		logger.COM().Error("Halting")
		for true {
		}
	}
	return 0
}
