package init

import (
	"kernel/lib/logger"
	"kernel/modules/io/idt"
	"kernel/modules/io/interrupts"
	"kernel/modules/memory/gdt"
)

func Init() {
	gdt.Init()
	idt.Init()
	interrupts.Load()
	logger.COM().Info("Initialized")
}
