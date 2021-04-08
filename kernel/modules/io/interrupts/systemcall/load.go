package systemcall

import (
	"kernel/modules/io/interrupts/utils"
)

// Handler for system call
func Handler()

// VHandler for v system call i.e. __kernel_vsyscall
func VHandler()

// SetSyscallHandler at gs:0x10
func SetSyscallHandler(handlerAddr uint32)

// Load System call interrupt handler
func Load() {
	utils.LoadIntHandler(0x80, Handler)
	var ptr uint32 = utils.GetFuncAddr(VHandler)
	SetSyscallHandler(ptr)
}
