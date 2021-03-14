package kernel

import (
	"kernel/lib/logger"
	"kernel/models"
	"kernel/modules/init"
)

type S struct {
	A uint32
}

func f() *S {
	return &S{A: 32}
}

// Main is the first function which is called
func Main(p models.KernelParams) {
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

	init.Init(p)
	s := f()
	logger.COM().LogUint(
		logger.Info, "S.A ", uint64(s.A),
	)
}
