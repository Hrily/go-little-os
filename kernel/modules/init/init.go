package init

import (
	"kernel/modules/memory/gdt"
)

func Init() {
	gdt.Init()
}
