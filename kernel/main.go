package kernel

import (
	"kernel/lib/memory/gdt"
)

// Main is the first function which is called
func Main() {
	gdt.Init()
}
