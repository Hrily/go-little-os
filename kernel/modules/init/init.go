package init

import (
	"kernel/modules/io/idt"
	"kernel/modules/memory/gdt"
)

func Init() {
	gdt.Init()
	idt.Init()
}
