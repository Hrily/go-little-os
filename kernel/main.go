package kernel

import (
	"kernel/lib/logger"
	"kernel/models"
	"kernel/modules/init"
	"unsafe"
)

func Start()

type I interface {
	F()
}

type S struct{}

func (s *S) F() {
	logger.COM().Info("F()")
}

func NewI() I {
	return &S{}
}

func f() {
	i := make([]uint32, 4*1024*1024)
	address := uint64(uintptr(unsafe.Pointer(&i)))
	logger.COM().LogUint(logger.Info, "&i", address)
	si := i[4*1024*1024-1]
	logger.COM().LogUint(logger.Info, "si", uint64(si))
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
	Start()
	f()
}
