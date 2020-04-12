package interrupts

import (
	"kernel/modules/io/interrupts/exceptions"
	"kernel/modules/io/interrupts/pic"
)

func Enable()
func Disable()

func Load() {
	exceptions.Load()
	pic.Load()
	Enable()
}
