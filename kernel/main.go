package kernel

import (
	"kernel/modules/init"
)

// Main is the first function which is called
func Main() {
	init.Init()
}
