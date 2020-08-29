package kernel

import (
	"kernel/lib/logger"
	"kernel/models"
	"kernel/modules/init"
	"unsafe"
)

type I interface {
	F()
}

type S struct{}

func (s *S) F() {
	logger.COM().Info("F()")
}

func f() {
	var i I = &S{}
	i.F()
}

type J uint32

type Y struct {
	i uint32
	j J
	k uint32
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

	y := Y{}
	logger.COM().LogUint(
		logger.Info, "&y.i", uint64(uintptr(unsafe.Pointer(&y.i))),
	)
	logger.COM().LogUint(
		logger.Info, "&y.j", uint64(uintptr(unsafe.Pointer(&y.j))),
	)
	logger.COM().LogUint(
		logger.Info, "&y.k", uint64(uintptr(unsafe.Pointer(&y.k))),
	)
}
