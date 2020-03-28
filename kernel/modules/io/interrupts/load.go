package interrupts

import (
	"kernel/modules/io/interrupts/exceptions"
)

func Load() {
	exceptions.Load()
}
