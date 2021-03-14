package systemcall

import (
	"kernel/modules/io/interrupts/utils"
)

// Handler for system call
func Handler()

// Load System call interrupt handler
func Load() {
	utils.LoadIntHandler(0x80, Handler)
}
