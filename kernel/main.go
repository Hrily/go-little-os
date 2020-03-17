package kernel

import (
	"kernel/lib/io/serial"
)

// Main is the first function which is called
func Main() {
	com := serial.COM{
		Port:        serial.SerialCOM1Base,
		RateDivisor: 3,
	}
	com.Init()
	com.Write("Hello World")
}
