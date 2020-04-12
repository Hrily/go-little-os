package kernel

import (
	"kernel/lib/logger"
	"kernel/modules/init"
)

type Params struct {
	KernelVStartAddr int32
	KernelVEndAddr   int32
	KernelPStartAddr int32
	KernelPEndAddr   int32
}

// Main is the first function which is called
func Main(p Params) {
	logger.COM().LogUint(
		logger.Info, "Kernel Virtual  Start", uint64(p.KernelVStartAddr),
	)
	logger.COM().LogUint(
		logger.Info, "Kernel Virtual  End  ", uint64(p.KernelVEndAddr),
	)
	logger.COM().LogUint(
		logger.Info, "Kernel Physical Start", uint64(p.KernelPStartAddr),
	)
	logger.COM().LogUint(
		logger.Info, "Kernel Physical End  ", uint64(p.KernelPEndAddr),
	)
	init.Init()
}
