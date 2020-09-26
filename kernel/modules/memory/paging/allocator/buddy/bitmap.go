package buddy

const (
	_mapSize = 32
	_allSet  = uint32((1 << _mapSize) - 1)
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

func (bm *bitmap) Len() uint32 {
	return uint32(len(bm.maps) * _mapSize)
}

// IsSet tells if bit at index is set
func (bm *bitmap) IsSet(index uint32) bool {
	if index > bm.Len() {
		return false
	}
	return (bm.maps[(index/_mapSize)] & (1 << (index % _mapSize))) > 0
}

// Set bit at index
func (bm *bitmap) Set(index uint32) {
	if index > bm.Len() {
		return
	}
	bm.maps[(index / _mapSize)] |= 1 << (index % _mapSize)
}

// Reset clears bit at index
func (bm *bitmap) Reset(index uint32) {
	if index > bm.Len() {
		return
	}
	bm.maps[(index / _mapSize)] &^= (1 << (index % _mapSize))
}

// Toggle bit at index
func (bm *bitmap) Toggle(index uint32) {
	if index > bm.Len() {
		return
	}
	bm.maps[(index / _mapSize)] ^= 1 << (index % _mapSize)
}

// FirstSet returns first reset bit
func (bm *bitmap) FirstSet() (index uint32, ok bool) {
	for i := uint32(0); i < uint32(len(bm.maps)); i++ {
		// Check if all bits are reset
		if bm.maps[i] == 0 {
			continue
		}
		// There is atleast one reset bit
		for j := uint32(0); j < _mapSize; j++ {
			if bm.IsSet((i * _mapSize) + j) {
				return (i * _mapSize) + j, true
			}
		}
	}
	return 0, false
}

// FirstContiguousSet returns index of n contiguous set bits
func (bm *bitmap) FirstContiguousSet(n uint32) (index uint32, ok bool) {
	i := uint32(0)
	for (i + n) < bm.Len() {
		// check if n bit from i are all set
		j := uint32(0)
		for j < n {
			if !bm.IsSet(i + j) {
				break
			}
			j++
		}

		if j == n {
			// all bits are set from [i, i+n)
			return i, true
		}

		// j is first reset bit after i and j < n, so we know there is no point in
		// searching in [i, j) now
		if j > 0 {
			i += j
		} else {
			i++
		}
	}
	return
}
