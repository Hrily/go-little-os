package kernel

import (
	"display"
)

func Main() {
	display.FrameBufferWriteCell(0, 'A', 2, 8)
}
