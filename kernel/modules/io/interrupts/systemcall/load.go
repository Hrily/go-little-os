package systemcall

import (
	"kernel/modules/io/interrupts/utils"
)

// Load System call interrupt handler
func Load() {
	utils.LoadIntHandler(0x80, Int0x80)
}
