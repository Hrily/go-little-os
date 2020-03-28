package exceptions

import (
	"kernel/modules/io/interrupts/utils"
)

func Load() {
	utils.LoadIntHandler(0x00, Int0x00)
	utils.LoadIntHandler(0x01, Int0x01)
	utils.LoadIntHandler(0x02, Int0x02)
	utils.LoadIntHandler(0x03, Int0x03)
	utils.LoadIntHandler(0x04, Int0x04)
	utils.LoadIntHandler(0x05, Int0x05)
	utils.LoadIntHandler(0x06, Int0x06)
	utils.LoadIntHandler(0x07, Int0x07)
	utils.LoadIntHandler(0x08, Int0x08)
	utils.LoadIntHandler(0x09, Int0x09)
	utils.LoadIntHandler(0x0a, Int0x0a)
	utils.LoadIntHandler(0x0b, Int0x0b)
	utils.LoadIntHandler(0x0c, Int0x0c)
	utils.LoadIntHandler(0x0d, Int0x0d)
	utils.LoadIntHandler(0x0e, Int0x0e)
	utils.LoadIntHandler(0x0f, Int0x0f)
	utils.LoadIntHandler(0x10, Int0x10)
	utils.LoadIntHandler(0x11, Int0x11)
	utils.LoadIntHandler(0x12, Int0x12)
	utils.LoadIntHandler(0x13, Int0x13)
	utils.LoadIntHandler(0x14, Int0x14)
	utils.LoadIntHandler(0x15, Int0x15)
	utils.LoadIntHandler(0x16, Int0x16)
	utils.LoadIntHandler(0x17, Int0x17)
	utils.LoadIntHandler(0x18, Int0x18)
	utils.LoadIntHandler(0x19, Int0x19)
	utils.LoadIntHandler(0x1a, Int0x1a)
	utils.LoadIntHandler(0x1b, Int0x1b)
	utils.LoadIntHandler(0x1c, Int0x1c)
	utils.LoadIntHandler(0x1d, Int0x1d)
	utils.LoadIntHandler(0x1e, Int0x1e)
	utils.LoadIntHandler(0x1f, Int0x1f)
}
