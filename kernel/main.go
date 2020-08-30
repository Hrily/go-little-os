package kernel

import (
	"kernel/lib/logger"
	"kernel/models"
	"kernel/modules/init"
)

type I interface {
	F()
}

type S struct{}

func (s *S) F() {
	logger.COM().Info("F()")
}

var _s S = S{}

func f() {
	var i I = &_s
	i.F()
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
	f()
}
