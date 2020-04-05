package display

import (
	"kernel/lib/ascii"
	"kernel/lib/io"
	"kernel/utils/integer"
)

// moveCursor moves cursor to given position
func moveCursor(pos uint16) {
	io.OutB(_frameBufferCommandPort, _frameBufferHighByteCommand)
	io.OutB(_frameBufferDataPort, integer.UInt16GetHighByte(pos))
	io.OutB(_frameBufferCommandPort, _frameBufferLowByteCommand)
	io.OutB(_frameBufferDataPort, integer.UInt16GetLowByte(pos))
}

// getColorByte returns frame buffer color byte
func getColorByte(fg, bg uint8) uint8 {
	return ((bg & 0x0F) << 4) | (fg & 0x0F)
}

// getFrameBufferAddress return the address corresponding to given position
func getFrameBufferAddress(pos uint32) uint32 {
	return _frameBufferAddress + (2 * pos)
}

// isPositionModifier returns true if given char is a position modifier
// charachter
func isPositionModifier(char byte) bool {
	switch char {
	case ascii.CR, ascii.LF, ascii.BS, ascii.TAB:
		return true
	}
	return false
}
