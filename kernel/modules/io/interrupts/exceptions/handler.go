package exceptions

import (
	"kernel/lib/logger"
	"kernel/modules/io/interrupts/models"
)

var messages = [32]string{
	"Division By Zero Exception",
	"Debug Exception",
	"Non Maskable Interrupt Exception",
	"Breakpoint Exception",
	"Into Detected Overflow Exception",
	"Out of Bounds Exception",
	"Invalid Opcode Exception",
	"No Coprocessor Exception",
	"Double Fault Exception",
	"Coprocessor Segment Overrun Exception",
	"Bad TSS Exception",
	"Segment Not Present Exception",
	"Stack Fault Exception",
	"General Protection Fault Exception",
	"Page Fault Exception",
	"Unknown Interrupt Exception",
	"Coprocessor Fault Exception",
	"Alignment Check Exception",
	"Machine Check Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
	"Reserved Exception",
}

func Handle(r models.Registers) {
	if r.IntNumber > 31 {
		return
	}
	logger.COM().Error(messages[r.IntNumber])
	logger.COM().LogUint(logger.Debug, "EIP", uint64(r.EIP))
	// Halting
	logger.COM().Error("Halting")
	for true {
	}
}
