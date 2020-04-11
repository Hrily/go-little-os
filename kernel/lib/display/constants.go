package display

const (
	_frameBufferAddress = 0xC00B8000
	_frameBufferRows    = 25
	_frameBufferColumns = 80

	// I/O ports

	_frameBufferCommandPort = 0x3D4
	_frameBufferDataPort    = 0x3D5

	// I/O port commands

	_frameBufferHighByteCommand = 14
	_frameBufferLowByteCommand  = 15

	// Colors

	Black        = 0
	Blue         = 1
	Green        = 2
	Cyan         = 3
	Red          = 4
	Magenta      = 5
	Brown        = 6
	LightGrey    = 7
	DarkGrey     = 8
	LightBlue    = 9
	LightGreen   = 10
	LightCyan    = 11
	LightRed     = 12
	LightMagenta = 13
	LightBrown   = 14
	White        = 15

	_tabWidth = 4
)
