package kernel

import (
	"kernel/lib/logger"
)

// Main is the first function which is called
func Main() {
	logger.COM().Debug("Hello World")
}
