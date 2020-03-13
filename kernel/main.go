package kernel

import (
	"kernel/display"
)

func Main() {
	display.FrameBufferWriteCell(0, 'A', 2, 8)
}
