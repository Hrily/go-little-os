package pic

import (
	"kernel/lib/io"
	"kernel/modules/io/interrupts/utils"
)

func remapPicIRQs() {
	// starts the initialization sequence (in cascade mode)
	io.OutB(pic1Command, picInit)
	io.OutB(pic2Command, picInit)

	// Set Offsets
	io.OutB(pic1Data, pic1Offset)
	io.OutB(pic2Data, pic2Offset)

	// ICW3: tell Master PIC that there is a slave PIC at IRQ2 (0000 0100)
	io.OutB(pic1Data, 0x04)
	// ICW3: tell Slave PIC its cascade identity (0000 0010)
	io.OutB(pic2Data, 0x02)

	// Set PIC to 8086/88 (MCS-80/85) mode
	io.OutB(pic1Data, picMode8086)
	io.OutB(pic2Data, picMode8086)

	// Enable Only Keyboard Interrupt
	io.OutB(pic1Data, 0xfd)
	io.OutB(pic2Data, 0xff)
}

func Load() {
	remapPicIRQs()

	utils.LoadIntHandler(0x20, Int0x20)
	utils.LoadIntHandler(0x21, Int0x21)
	utils.LoadIntHandler(0x22, Int0x22)
	utils.LoadIntHandler(0x23, Int0x23)
	utils.LoadIntHandler(0x24, Int0x24)
	utils.LoadIntHandler(0x25, Int0x25)
	utils.LoadIntHandler(0x26, Int0x26)
	utils.LoadIntHandler(0x27, Int0x27)

	utils.LoadIntHandler(0x28, Int0x28)
	utils.LoadIntHandler(0x29, Int0x29)
	utils.LoadIntHandler(0x2a, Int0x2a)
	utils.LoadIntHandler(0x2b, Int0x2b)
	utils.LoadIntHandler(0x2c, Int0x2c)
	utils.LoadIntHandler(0x2d, Int0x2d)
	utils.LoadIntHandler(0x2e, Int0x2e)
	utils.LoadIntHandler(0x2f, Int0x2f)
}
