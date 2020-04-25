package init

import (
	"kernel/lib/logger"
	"kernel/models"
	"kernel/modules/io/idt"
	"kernel/modules/io/interrupts"
	"kernel/modules/memory/gdt"
	"kernel/modules/memory/paging"
)

func Init(p models.KernelParams) {
	gdt.Init()
	idt.Init()
	interrupts.Load()
	paging.LoadKernelPDT(
		p.KernelPStartAddr, p.KernelVStartAddr,
		p.KernelVEndAddr-p.KernelVStartAddr,
	)
	logger.COM().Info("Initialized")
}
