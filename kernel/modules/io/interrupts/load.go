package interrupts

import (
	"kernel/modules/io/interrupts/exceptions"
	"kernel/modules/io/interrupts/pic"
)

func Load() {
	exceptions.Load()
	pic.Load()
}
