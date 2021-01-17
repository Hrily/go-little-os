package common

import (
	"kernel/lib/logger"
	"kernel/modules/io/interrupts/exceptions"
	"kernel/modules/io/interrupts/models"
	"kernel/modules/io/interrupts/pic"
	"kernel/modules/io/interrupts/systemcall"
)

func InterruptHandler(r models.Registers) {
	logger.COM().LogUint(logger.Info, "Interrupt occured with Number", uint64(r.IntNumber))
	logger.COM().LogUint(logger.Info, "Interrupt error code", uint64(r.ErrCode))
	switch {
	case 0x00 <= r.IntNumber && r.IntNumber <= 0x1f:
		exceptions.Handle(r)
	case 0x20 <= r.IntNumber && r.IntNumber <= 0x2f:
		pic.Handle(r)
	case r.IntNumber == 0x80:
		systemcall.Handle(r)
	default:
		logger.COM().Error("Unknown Interrupt")
	}
}
