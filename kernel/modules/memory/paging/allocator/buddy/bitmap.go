package buddy

const (
	_mapSize = 32
)

// bitmap provides map of bits
type bitmap struct {
	maps []uint32
}

// nMaps returns number of maps required for nBits in bitmap
func nMaps(nBits uint32) uint32 {
	// make nBits positive multiple of _mapSize
	if (nBits % _mapSize) > 0 {
		nBits += _mapSize - (nBits % _mapSize)
	}
	return nBits / _mapSize
}

// SetMaps of bitmap
func (bm *bitmap) SetMaps(maps []uint32) {
	bm.maps = maps
}

func (bm *bitmap) len() uint32 {
	return uint32(len(bm.maps) * _mapSize)
}

// IsSet tells if bit at index is set
func (bm *bitmap) IsSet(index uint32) bool {
	if index > bm.len() {
		return false
	}
	return (bm.maps[(index/_mapSize)] & (1 << (index % _mapSize))) > 0
}

// Set bit at index
func (bm *bitmap) Set(index uint32) {
	if index > bm.len() {
		return
	}
	bm.maps[(index / _mapSize)] |= 1 << (index % _mapSize)
}

// Reset clears bit at index
func (bm *bitmap) Reset(index uint32) {
	if index > bm.len() {
		return
	}
	bm.maps[(index / _mapSize)] &^= (1 << (index % _mapSize))
}

// Toggle bit at index
func (bm *bitmap) Toggle(index uint32) {
	if index > bm.len() {
		return
	}
	bm.maps[(index / _mapSize)] ^= 1 << (index % _mapSize)
}
