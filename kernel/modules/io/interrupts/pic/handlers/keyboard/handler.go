package keyboard

import (
	"kernel/lib/display"
	"kernel/lib/io"
	"kernel/lib/logger"
	"kernel/modules/io/interrupts/models"
)

const (
	keyboardDataPort = 0x60
)

func Handle(r models.Registers) {
	code := uint32(io.InB(keyboardDataPort))
	logger.COM().LogUint(logger.Debug, "Code", uint64(code))
	if isModifierKey(code) {
		state.ToggleModifier()
		return
	}
	isRelease := code >= 0x80
	if isRelease {
		code -= 0x80
	}
	if !isRelease && keys[code] != nil &&
		keys[code].Characters[state.ModifierLevel] != byte(0) {
		display.FB().WriteB(keys[code].Characters[state.ModifierLevel])
	}
	logger.COM().LogUint(logger.Debug, "ModL", uint64(state.ModifierLevel))
}

func isModifierKey(code uint32) bool {
	switch code {
	case 0x2a, 0x36, 0x3a:
		return true
	case 0xaa, 0xb6:
		return true
	}
	return false
}
