package interrupts

import (
	"kernel/modules/io/interrupts/exceptions"
	"kernel/modules/io/interrupts/pic"
	"kernel/modules/io/interrupts/systemcall"
)

func Enable()
func Disable()

func Load() {
	exceptions.Load()
	pic.Load()
	systemcall.Load()
	Enable()
}
