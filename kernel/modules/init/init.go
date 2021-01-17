package init

import (
	"kernel/lib/logger"
	"kernel/models"
	"kernel/modules/io/idt"
	"kernel/modules/io/interrupts"
	"kernel/modules/memory/gdt"
	"kernel/modules/memory/paging"
	"kernel/modules/memory/tls"
)

func Init(p models.KernelParams) {
	gdt.Init()
	idt.Init()
	interrupts.Load()
	paging.LoadKernelPDT(
		p.KernelPStartAddr, p.KernelVStartAddr,
		p.KernelVEndAddr-p.KernelVStartAddr,
	)
	tls.Setup()
	logger.COM().Info("Initialized")
}
