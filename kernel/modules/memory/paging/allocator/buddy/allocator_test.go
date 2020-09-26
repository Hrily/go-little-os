package buddy

import (
	"testing"
	"unsafe"

	"kernel/modules/memory/paging/models"

	"github.com/stretchr/testify/assert"
)

const (
	_testNodesSize = 32
)

var (
	_testNodes [_testNodesSize]node
)

// setup common
func setup(ba *buddyAllocator) {
	ba.buddies = make([]freeBuddiesList, _maxOrder+1)

	ba.buddies[_maxOrder].freeMap.maps = make([]uint32, nMaps(_nBigPages))
	for i := uint32(0); i < nMaps(_nBigPages); i++ {
		ba.buddies[_maxOrder].freeMap.maps[i] = _allSet
	}

	for i := 0; i < _maxOrder; i++ {
		var nBuddies uint32 = _nBigPages * (1 << uint32(_maxOrder-i-1))
		ba.buddies[i].freeMap.maps = make([]uint32, nMaps(nBuddies))
		for j := uint32(0); j < nMaps(nBuddies); j++ {
			ba.buddies[i].freeMap.maps[j] = 0
		}
	}

	// test setup of node pool
	_nodePool.freeMap.maps = make([]uint32, nMaps(_testNodesSize))
	for i := uint32(0); i < nMaps(_testNodesSize); i++ {
		_nodePool.freeMap.maps[i] = _allSet
	}
	_nodePool.startAddr = uint32(uintptr(unsafe.Pointer(&_testNodes)))
}

func Test_buddyAllocator_Allocate(t *testing.T) {
	tests := []struct {
		name   string
		order  uint32
		setup  func(ba *buddyAllocator)
		assert func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool)
	}{
		{
			name:  "Multiple big pages, available",
			order: 4 * _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// keep 4 pages free at index 12
				// 0b 1111 0111 0110 0101
				ba.buddies[_maxOrder].freeMap.maps[0] = 0xf765
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, true, ok)
				assert.Equal(t, 12*models.Size4MB.ToBytes(), addr)
				assert.Equal(t, uint32(0x0765), ba.buddies[_maxOrder].freeMap.maps[0])
			},
		},
		{
			name:  "Multiple big pages, not available",
			order: 4 * _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// Mark all as used
				for i := uint32(0); i < nMaps(_nBigPages); i++ {
					ba.buddies[_maxOrder].freeMap.maps[i] = 0
				}
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, false, ok)
				assert.Equal(t, uint32(0), addr)
			},
		},
		{
			name:  "Single big page, available",
			order: _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// Mark 12th page as free
				ba.buddies[_maxOrder].freeMap.maps[0] = 0x1000
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, true, ok)
				assert.Equal(t, 12*models.Size4MB.ToBytes(), addr)
				assert.Equal(t, uint32(0), ba.buddies[_maxOrder].freeMap.maps[0])
			},
		},
		{
			name:  "Single big page, not available",
			order: _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// Mark all as used
				for i := uint32(0); i < nMaps(_nBigPages); i++ {
					ba.buddies[_maxOrder].freeMap.maps[i] = 0
				}
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, false, ok)
				assert.Equal(t, uint32(0), addr)
			},
		},
		{
			name:  "Single small page, available",
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// Mark 12th page as free
				ba.buddies[0].freeMap.maps[0] = 0x1000
				ba.buddies[0].freeList.Append(
					(2*12 + 1) * models.Size4KB.ToBytes(),
				)
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, true, ok)
				assert.Equal(t, (2*12+1)*models.Size4KB.ToBytes(), addr)
				assert.Equal(t, uint32(0), ba.buddies[0].freeMap.maps[0])
				assert.True(t, ba.buddies[0].freeList.IsEmpty())
			},
		},
		{
			name:  "Single small page, not available",
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// Mark all pages as used
				for i := uint32(0); i < nMaps(_nBigPages); i++ {
					ba.buddies[_maxOrder].freeMap.maps[i] = 0
				}
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, false, ok)
				assert.Equal(t, uint32(0), addr)
			},
		},
		{
			name:  "Single small page, break from a big page",
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// keep 1 page free at index 12
				ba.buddies[_maxOrder].freeMap.maps[0] = 0x1000
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, true, ok)
				assert.Equal(
					t,
					// we get second page
					((12 * models.Size4MB.ToBytes()) + models.Size4KB.ToBytes()),
					addr,
				)
				assert.Equal(t, uint32(0), ba.buddies[_maxOrder].freeMap.maps[0])
				for i := _maxOrder - 1; i >= 0; i-- {
					index := uint32(12 * (1 << (_maxOrder - i - 1)))
					assert.Equal(
						t, true,
						ba.buddies[i].freeMap.IsSet(index),
						"buddy not marked used",
						"order", i, "index", index,
					)
					// freelist should have first page in buddy
					assert.False(t, ba.buddies[0].freeList.IsEmpty())
					assert.Equal(
						t,
						(2 * index * (1 << i) * models.Size4KB.ToBytes()),
						ba.buddies[i].freeList.Head().Value(),
						"incorrect address added to freelist",
						"order", i, "index", index,
					)
				}
			},
		},
		{
			name:  "Single small page, break from a higher order page",
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// keep 1 page free at index 6, i.e 12th big page
				ba.buddies[_maxOrder-1].freeMap.maps[0] = 0x20
				ba.buddies[_maxOrder-1].freeList.Append(
					12 * models.Size4MB.ToBytes(),
				)
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, true, ok)
				assert.Equal(
					t,
					// we get second page
					((12 * models.Size4MB.ToBytes()) + models.Size4KB.ToBytes()),
					addr,
				)
				assert.Equal(t, false, ba.buddies[_maxOrder-1].freeMap.IsSet(6))
				for i := _maxOrder - 2; i >= 0; i-- {
					index := uint32(12 * (1 << (_maxOrder - i - 1)))
					assert.Equal(
						t, true,
						ba.buddies[i].freeMap.IsSet(index),
						"buddy not marked used",
						"order", i, "index", index,
					)
					// freelist should have first page in buddy
					assert.False(t, ba.buddies[0].freeList.IsEmpty())
					assert.Equal(
						t,
						(2 * index * (1 << i) * models.Size4KB.ToBytes()),
						ba.buddies[i].freeList.Head().Value(),
						"incorrect address added to freelist",
						"order", i, "index", index,
					)
				}
			},
		},
		{
			name:  "Node Pool Empty",
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// keep 1 page free at index 12
				ba.buddies[_maxOrder].freeMap.maps[0] = 0x1000
				// make node pool full
				for i := uint32(0); i < nMaps(_testNodesSize); i++ {
					_nodePool.freeMap.maps[i] = 0
				}
			},
			assert: func(t *testing.T, ba *buddyAllocator, addr uint32, ok bool) {
				assert.Equal(t, false, ok)
				assert.Equal(t, uint32(0), addr)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ba := &buddyAllocator{}
			test.setup(ba)
			addr, ok := ba.Allocate(test.order)
			test.assert(t, ba, addr, ok)
		})
	}
}

func Test_buddyAllocator_Release(t *testing.T) {
	tests := []struct {
		name   string
		addr   uint32
		order  uint32
		setup  func(ba *buddyAllocator)
		assert func(t *testing.T, ba *buddyAllocator)
	}{
		{
			name:  "Release multiple big page",
			addr:  12 * models.Size4MB.ToBytes(),
			order: 4 * _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// mark 4 pages at index 12 as used
				// 0b 0000 0111 0110 0101
				ba.buddies[_maxOrder].freeMap.maps[0] = 0x0765
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0xf765), ba.buddies[_maxOrder].freeMap.maps[0])
			},
		},
		{
			name:  "Release single big page",
			addr:  12 * models.Size4MB.ToBytes(),
			order: 1 * _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// mark 4 pages at index 12 as used
				// 0b 0000 0111 0110 0101
				ba.buddies[_maxOrder].freeMap.maps[0] = 0x0765
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0x1765), ba.buddies[_maxOrder].freeMap.maps[0])
			},
		},
		{
			name:  "Release single small page, buddy is allocated",
			addr:  (2*12 + 1) * models.Size4KB.ToBytes(),
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// Mark 12th page buddy as free
				ba.buddies[0].freeMap.maps[0] = 0x0765
				ba.buddies[0].freeList.Append(
					(2*12 + 1) * models.Size4KB.ToBytes(),
				)
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0x1765), ba.buddies[0].freeMap.maps[0])
				if ok := assert.False(t, ba.buddies[0].freeList.IsEmpty()); ok {
					assert.Equal(
						t,
						(2*12+1)*models.Size4KB.ToBytes(),
						ba.buddies[0].freeList.Head().Value(),
					)
				}
			},
		},
		{
			name:  "Release single small page, buddy is free, higher order buddy is allocated",
			addr:  (2*12 + 1) * models.Size4KB.ToBytes(),
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// Mark 12th page buddy as free
				ba.buddies[0].freeMap.maps[0] = 0x1765
				// Mark higher order buddy as allocated
				ba.buddies[1].freeMap.maps[0] = 0x0000
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0x0765), ba.buddies[0].freeMap.maps[0])
				assert.Equal(t, uint32(0x0040), ba.buddies[1].freeMap.maps[0])
				if ok := assert.False(t, ba.buddies[1].freeList.IsEmpty()); ok {
					assert.Equal(
						t,
						(2*12)*models.Size4KB.ToBytes(),
						ba.buddies[1].freeList.Head().Value(),
					)
				}
			},
		},
		{
			name:  "Release single small page, buddy is free, higher order buddies are also free",
			addr:  (2*1024 + 1) * models.Size4KB.ToBytes(),
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				index := uint32(1024)
				for order := 0; order < _maxOrder; order++ {
					ba.buddies[order].freeMap.Set(index >> order)
				}
				// mark big page as used
				ba.buddies[_maxOrder].freeMap.maps[0] = 0
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				index := uint32(1024)
				for order := 0; order < _maxOrder; order++ {
					assert.False(
						t, ba.buddies[order].freeMap.IsSet(index>>order),
						"order", order, "index >> order", index>>order,
					)
				}
				assert.Equal(t, uint32(0x4), ba.buddies[_maxOrder].freeMap.maps[0])
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ba := &buddyAllocator{}
			test.setup(ba)
			ba.Release(test.addr, test.order)
			test.assert(t, ba)
		})
	}
}

func Test_buddyAllocator_Mark(t *testing.T) {
	tests := []struct {
		name   string
		addr   uint32
		order  uint32
		setup  func(ba *buddyAllocator)
		assert func(t *testing.T, ba *buddyAllocator)
	}{
		{
			name:  "Mark multiple big page",
			addr:  12 * models.Size4MB.ToBytes(),
			order: 4 * _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				// mark 4 pages at index 12 as used
				// 0b 1111 0111 0110 0101
				ba.buddies[_maxOrder].freeMap.maps[0] = 0xf765
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0x0765), ba.buddies[_maxOrder].freeMap.maps[0])
			},
		},
		{
			name:  "Release single big page",
			addr:  12 * models.Size4MB.ToBytes(),
			order: 1 * _maxOrder,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				ba.buddies[_maxOrder].freeMap.maps[0] = 0x1765
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0x0765), ba.buddies[_maxOrder].freeMap.maps[0])
			},
		},
		{
			name:  "Release single small page, buddy is allocated",
			addr:  (2*12 + 1) * models.Size4KB.ToBytes(),
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				ba.buddies[0].freeMap.maps[0] = 0x1765
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0x0765), ba.buddies[0].freeMap.maps[0])
				assert.True(t, ba.buddies[0].freeList.IsEmpty())
			},
		},
		{
			name:  "Release single small page, buddy is free",
			addr:  (2*12 + 1) * models.Size4KB.ToBytes(),
			order: 0,
			setup: func(ba *buddyAllocator) {
				setup(ba)
				ba.buddies[0].freeMap.maps[0] = 0x0765
			},
			assert: func(t *testing.T, ba *buddyAllocator) {
				assert.Equal(t, uint32(0x1765), ba.buddies[0].freeMap.maps[0])
				if ok := assert.False(t, ba.buddies[0].freeList.IsEmpty()); ok {
					assert.Equal(
						t,
						(2*12)*models.Size4KB.ToBytes(),
						ba.buddies[0].freeList.Head().Value(),
					)
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ba := &buddyAllocator{}
			test.setup(ba)
			ba.Mark(test.addr, test.order)
			test.assert(t, ba)
		})
	}
}
