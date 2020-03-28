package kernel

import (
	"kernel/modules/init"
)

var i int

// Main is the first function which is called
func Main() {
	init.Init()
	i = 1 / i
	_ = i + 1
}
