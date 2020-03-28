package init

import (
	"kernel/modules/io/idt"
	"kernel/modules/io/interrupts"
	"kernel/modules/memory/gdt"
)

func Init() {
	gdt.Init()
	idt.Init()
	interrupts.Load()
}
