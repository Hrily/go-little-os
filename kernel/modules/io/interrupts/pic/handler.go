package pic

import (
	"kernel/lib/io"
	"kernel/modules/io/interrupts/models"
	"kernel/modules/io/interrupts/pic/handlers/keyboard"
)

func Handle(r models.Registers) {
	switch r.IntNumber - pic1Offset {
	case 1:
		keyboard.Handle(r)
	}

	// Acknoledge the PIC interrupt
	// Send ACK to slave controller only if interrupt was from it
	if r.IntNumber >= pic2Offset {
		io.OutB(pic2Command, picACK)
	}
	// Send ACK to master controller
	io.OutB(pic1Command, picACK)
}
