package serial

import (
	"kernel/lib/io"
	"kernel/utils/integer"
)

// COM represents the serial com port
type COM struct {
	Port        uint16
	RateDivisor uint16
	isInit      bool
}

// Init
// Initialized the serial com port by configuring:
//   - baud rate
//   - buffers
//   - line
//   - modem
func (c *COM) Init() {
	if !c.isInit {
		c.configureBaudRate()
		c.configureBuffers()
		c.configureLine()
		c.configureModem()
		c.isInit = true
	}
}

// Write writes given string on to the serial com port.
// Expects serial com port to be initialized,
func (c *COM) Write(buffer string) uint32 {
	for i := 0; i < len(buffer); i++ {
		for !c.isTransmitFIFOEmpty() {
		}
		io.OutB(c.Port, buffer[i])
	}
	return uint32(len(buffer))
}

// configureBaudRate
// Sets the speed of the data being sent. The default speed of a serial
// port is 115200 bits/s. The argument is a divisor of that number, hence
// the resulting speed becomes (115200 / divisor) bits/s.
func (c *COM) configureBaudRate() {
	io.OutB(serialLineCommandPort(c.Port), _serialLineEnableDLAB)
	io.OutB(serialDataPort(c.Port), integer.UInt16GetLowByte(c.RateDivisor))
	io.OutB(serialDataPort(c.Port), integer.UInt16GetHighByte(c.RateDivisor))
}

// configureLine
// Configures the line of the given serial port. The port is set to have a
// data length of 8 bits, no parity bits, one stop bit and break control
// disabled.
func (c *COM) configureLine() {
	// Bit:     | 7 | 6 | 5 4 3 | 2 | 1 0 |
	// Content: | d | b | prty  | s | dl  |
	// Value:   | 0 | 0 | 0 0 0 | 0 | 1 1 | = 0x03
	io.OutB(serialLineCommandPort(c.Port), 0x03)
}

// configureBuffers
// Configures the FIFO buffer queue. The port is set to:
//   - Enable FIFO
//   - Clear both receiver and transmission FIFO queues
//   - Use 14 bytes as size of queue
func (c *COM) configureBuffers() {
	// Bit:     | 7 6 | 5  | 4 | 3   | 2   | 1   | 0 |
	// Content: | lvl | bs | r | dma | clt | clr | e |
	// Value:   | 1 1 | 0  | 0 | 0   | 1   | 1   | 1 | = 0xc7
	io.OutB(serialDataPort(c.Port), 0xc7)
}

// configureModem
// Configures the modem. The port is set to be ready for data transmission
func (c *COM) configureModem() {
	// Bit:     | 7 | 6 | 5  | 4  | 3   | 2   | 1   | 0   |
	// Content: | r | r | af | lb | ao2 | ao1 | rts | dtr |
	// Value:   | 0 | 0 | 0  | 0  | 0   | 0   | 1   | 1   | = 0x03
	io.OutB(serialModemCommandPort(c.Port), 0x03)
}

//  isTransmitFIFOEmpty:
//  Checks whether the transmit FIFO queue is empty or not for the given COM
//  port.
//
//  @param  com The COM port
//  @return 0 if the transmit FIFO queue is not empty
//          1 if the transmit FIFO queue is empty
func (c *COM) isTransmitFIFOEmpty() bool {
	// 0x20 = 0010 0000
	return (io.InB(serialLineStatusPort(c.Port)) & 0x20) > 0
}

var com1 = COM{
	Port:        SerialCOM1Base,
	RateDivisor: 3,
}

func COM1() *COM {
	com1.Init()
	return &com1
}
