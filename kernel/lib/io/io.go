package io

// OutB:
//   Sends the given data to the given I/O port. Defined in io.s
//   @param port The I/O port to send the data to
//   @param data The data to send to the I/O port
func OutB(port uint16, char byte)

// InB:
//   Read a byte from an I/O port.
//   @param  port The address of the I/O port
//   @return      The read byte
func InB(port uint16) byte
