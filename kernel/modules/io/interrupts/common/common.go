package common

import (
	"kernel/lib/logger"
	"kernel/modules/io/interrupts/exceptions"
	"kernel/modules/io/interrupts/models"
)

func InterruptHandler(r models.Registers) {
	logger.COM().LogUint(logger.Error, "Interrupt occured with Number", uint64(r.IntNumber))
	switch {
	case 0 <= r.IntNumber && r.IntNumber <= 31:
		exceptions.Handle(r)
	default:
		logger.COM().Error("Unknown Interrupt")
	}
	// Halting
	logger.COM().Error("Halting")
	for true {
	}
}
