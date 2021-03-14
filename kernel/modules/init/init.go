package init

import (
	"kernel/lib/logger"
	"kernel/models"
	"kernel/modules/io/idt"
	"kernel/modules/io/interrupts"
	"kernel/modules/memory/gdt"
	"kernel/modules/memory/paging"
	"kernel/modules/memory/tls"
	"kernel/modules/process/processinfo"
)

func Init(p models.KernelParams) {
	logger.COM().Info("Setting up GDT")
	gdt.Init()
	logger.COM().Info("Setting up IDT")
	idt.Init()
	logger.COM().Info("Setting up Interrupts")
	interrupts.Load()
	logger.COM().Info("Setting up Kernel Process Info")
	processinfo.SetupKernelInfo(p)
	logger.COM().Info("Setting up Paging")
	paging.LoadKernelPDT(
		p.KernelPStartAddr, p.KernelVStartAddr,
		p.KernelVEndAddr-p.KernelVStartAddr,
	)
	tls.Setup()
	logger.COM().Info("Initialized")
}
