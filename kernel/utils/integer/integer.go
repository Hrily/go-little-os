package integer

func UInt16GetHighByte(i uint16) uint8 {
	return uint8((i >> 8) & 0xff)
}

func UInt16GetLowByte(i uint16) uint8 {
	return uint8(i & 0xff)
}
