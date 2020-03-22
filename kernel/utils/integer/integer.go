package integer

func UInt16GetHighByte(i uint16) uint8 {
	return uint8((i >> 8) & 0xff)
}

func UInt16GetLowByte(i uint16) uint8 {
	return uint8(i & 0xff)
}

func BoolToUInt32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}
