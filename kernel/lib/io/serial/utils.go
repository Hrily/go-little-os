package serial

func serialDataPort(base uint16) uint16 {
	return base
}

func serialFIFOCommandPort(base uint16) uint16 {
	return base + 2
}

func serialLineCommandPort(base uint16) uint16 {
	return base + 3
}

func serialModemCommandPort(base uint16) uint16 {
	return base + 4
}

func serialLineStatusPort(base uint16) uint16 {
	return base + 5
}
