package keyboard

import (
	"kernel/lib/io"
	"kernel/lib/logger"
	"kernel/modules/io/interrupts/models"
)

const (
	keyboardDataPort = 0x60
)

func Handle(r models.Registers) {
	logger.COM().Info("KeyboardInterrupt")
	code := io.InB(keyboardDataPort)
	logger.COM().LogUint(logger.Debug, "Code", uint64(code))
}
