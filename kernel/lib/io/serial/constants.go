package serial

const (
	/* The I/O ports */

	/* All the I/O ports are calculated relative to the data port. This is because
	 * all serial ports (COM1, COM2, COM3, COM4) have their ports in the same
	 * order, but they start at different values.
	 */

	SerialCOM1Base = 0x3F8 /* COM1 base port */

	/* The I/O port commands */

	/* SERIAL_LINE_ENABLE_DLAB:
	 * Tells the serial port to expect first the highest 8 bits on the data port,
	 * then the lowest 8 bits will follow
	 */
	_serialLineEnableDLAB = 0x80
)
