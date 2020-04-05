package display

import (
	"kernel/lib/ascii"
	"kernel/lib/memory"
)

// FrameBuffer represents the frame buffer
type FrameBuffer struct {
	pos    uint16
	fg     uint8
	bg     uint8
	isInit bool
}

var fb = FrameBuffer{}

func FB() *FrameBuffer {
	fb.Init()
	return &fb
}

// Init initializes the frame buffer
func (fb *FrameBuffer) Init() {
	if fb.isInit {
		return
	}
	fb.pos = 0
	fb.fg = Cyan
	fb.bg = Black
	fb.isInit = true
	for i := 0; i < _frameBufferColumns*_frameBufferRows; i++ {
		address := getFrameBufferAddress(uint32(i))
		memory.PutB(address, ' ')
		memory.PutB(address+1, getColorByte(fb.fg, fb.bg))
	}
	moveCursor(fb.pos)
}

/**
 * Exported Methods
 */

// Write puts given string on screen
func (fb *FrameBuffer) Write(buffer string) int {
	for i := 0; i < len(buffer); i++ {
		fb.writeB(buffer[i])
	}
	moveCursor(fb.pos)
	return len(buffer)
}

// WriteB puts given string on screen
func (fb *FrameBuffer) WriteB(b byte) int {
	fb.writeB(b)
	moveCursor(fb.pos)
	return 1
}

// SetFg sets foreground color of frame buffer
func (fb *FrameBuffer) SetFg(fg uint8) {
	fb.fg = fg
}

// SetBg sets background color of frame buffer
func (fb *FrameBuffer) SetBg(bg uint8) {
	fb.bg = bg
}

// ScrollUp scrolls the frame buffer by n rows
func (fb *FrameBuffer) ScrollUp(n uint32) {
	memory.MoveData(
		getFrameBufferAddress(0),
		getFrameBufferAddress(n*_frameBufferColumns),
		getFrameBufferAddress(_frameBufferRows*_frameBufferColumns)-
			getFrameBufferAddress(n*_frameBufferColumns),
	)
	// Save position
	posBackup := fb.pos
	// Empty remaining rows
	remaining := uint16(_frameBufferRows - n)
	fb.pos = remaining * _frameBufferColumns
	for pos := remaining * _frameBufferColumns; pos <
		_frameBufferRows*_frameBufferColumns; pos++ {
		fb.putChar(' ')
	}
	// Restore position
	fb.pos = posBackup
}

/**
 * Unexported Methods
 */

func (fb *FrameBuffer) writeB(b byte) {
	if isPositionModifier(b) {
		fb.handlePositionModifier(b)
	} else {
		fb.putChar(b)
	}
	fb.scrollIfNeeded()
}

// writeCell writes given char to given cell
func (fb *FrameBuffer) writeCell(char byte) {
	address := getFrameBufferAddress(uint32(fb.pos))
	memory.PutB(address, char)
	memory.PutB(address+1, getColorByte(fb.fg, fb.bg))
}

// increementPos increements the frame buffer cursor position
func (fb *FrameBuffer) increementPos() {
	fb.pos++
}

// handlePositionModifier handles position modifier charachters
func (fb *FrameBuffer) handlePositionModifier(char byte) {
	switch char {
	case ascii.CR:
		fb.pos -= (fb.pos % _frameBufferColumns)
	case ascii.LF:
		fb.pos -= (fb.pos % _frameBufferColumns)
		fb.pos += _frameBufferColumns
	case ascii.BS:
		if fb.pos > 0 {
			fb.pos--
			fb.writeCell(' ')
		}
	case ascii.TAB:
		fb.pos += _tabWidth
	}
}

// scrollIfNeeded scrolls the frame buffer if needed
func (fb *FrameBuffer) scrollIfNeeded() {
	for fb.pos >= (_frameBufferRows * _frameBufferColumns) {
		// reset position to start of row
		fb.pos -= _frameBufferColumns
		fb.ScrollUp(1)
	}
}

// putChar writes charachter and increements the cursor position
func (fb *FrameBuffer) putChar(char byte) {
	fb.writeCell(char)
	fb.increementPos()
}
