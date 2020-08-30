package buddy

import (
	"unsafe"

	"kernel/lib/logger"
	"kernel/utils/pointer"
)

const (
	// _maxOrder is the maximum order of 2 pages we can allocate.
	// for now, we assume this to be 4MB = 1024 * 4KB pages.
	_maxOrder = 10

	// _nBigPages is total number of 4mb pages. We assume initial memory of 4GB.
	_nBigPages = 1024
)

// AllocatorSize returns total size which will be used by buddyAllocator.
// This is useful while kernel page allocation
// To know how this is computed, check Init function
func AllocatorSize() (size uint32) {
	// size of bigPagesBitmap
	size += nMaps(_nBigPages) * 4

	// size of freeBuddiesList
	size += _maxOrder * uint32(unsafe.Sizeof(freeBuddiesList{}))

	// size of individual freeBuddiesList
	for i := 0; i < _maxOrder; i++ {
		// maxOrder - 1 pages of order i, further divide by 2 since we use 1 bit
		// for buddy pair.
		var nBuddies uint32 = _nBigPages * (1 << uint32(_maxOrder-i-1))
		size += nMaps(nBuddies) * 4
	}

	return
}

// freeBuddiesList stores free pages of a specific order. It uses one bit per
// buddy pair of pages.
// If a bit for a buddy pair is:
//   - 1 : first page in buddy pair is free
//   - 0 : both pages are either free or allocated
type freeBuddiesList struct {
	bitmap
}

// buddyAllocator uses buddy list to allocate page frames.
// It implements Allocator interface.
// Resources:
//   - OS Dev Wiki: https://wiki.osdev.org/Page_Frame_Allocation#Buddy_Allocation_System
//   - Kernel Docs: https://www.kernel.org/doc/gorman/html/understand/understand009.html
type buddyAllocator struct {
	// bigPagesBitmap is free bitmap for 4mb pages
	bigPagesBitmap bitmap
	// buddies is collection of free buddies list
	buddies []freeBuddiesList
}

// _buddyAllocator is the instance of buddyAllocator
var _buddyAllocator buddyAllocator

// InitAllocator the buddyAllocator at given address
func InitAllocator(addr uint32) {
	/**
	 * Warning: Lot of unsafe pointer usage and manipulation ahead.
	 * This function basically initializes memory at given address, so that
	 * _buddyAllocator, and it's members, point to this address.
	 */
	logger.COM().LogUint(logger.Info, "_buddyAllocator struct at", uint64(uintptr(unsafe.Pointer(&_buddyAllocator))))

	// create bigPagesBitmap
	// _nBigPages size for array is way more than needed. this is done since
	// array sizes need to be constant.
	var _bigPagesBitmap *[_nBigPages]uint32
	_bigPagesBitmap = (*[_nBigPages]uint32)(pointer.Get(addr))
	addr += nMaps(_nBigPages) * 4

	// set bigPagesBitmap in buddyAllocator
	_buddyAllocator.bigPagesBitmap.maps = (*_bigPagesBitmap)[:nMaps(_nBigPages)]

	// zero out bigPagesBitmap
	for i := uint32(0); i < nMaps(_nBigPages); i++ {
		_buddyAllocator.bigPagesBitmap.maps[i] = 0
	}

	// create freeBuddiesList
	var _freeBuddiesList *[_maxOrder]freeBuddiesList
	_freeBuddiesList = (*[_maxOrder]freeBuddiesList)(pointer.Get(addr))
	addr += _maxOrder * uint32(unsafe.Sizeof(freeBuddiesList{}))

	// set freeBuddiesList in buddyAllocator
	_buddyAllocator.buddies = (*_freeBuddiesList)[:_maxOrder]

	// create the individual freeBuddiesList
	for i := 0; i < _maxOrder; i++ {
		// maxOrder - 1 pages of order i, further divide by 2 since we use 1 bit
		// for buddy pair.
		var nBuddies uint32 = _nBigPages * (1 << uint32(_maxOrder-i-1))
		nMaps := nMaps(nBuddies)

		// set address for this freeBuddiesList
		// we are addressing way more memory than needed, because array sizes need
		// to be constant. this will be more than enough for largest freeBuddiesList
		var maps *[_nBigPages * _nBigPages]uint32
		maps = (*[_nBigPages * _nBigPages]uint32)(pointer.Get(addr))
		addr += nMaps * 4

		// set the freeBuddiesList
		_buddyAllocator.buddies[i].maps = (*maps)[:nMaps]

		// zero out the freeBuddiesList
		for j := uint32(0); j < nMaps; j++ {
			_buddyAllocator.buddies[i].maps[j] = 0
		}
	}
}

// Allocate 2^order number of pages and return physical address of the first
// page frame. Will return ok as false if it failed to allocate.
func (ba *buddyAllocator) Allocate(order uint32) (addr uint32, ok bool) {
	return
}

// Release marks 2^order number of pages starting from addr as free and unused.
func (ba *buddyAllocator) Release(addr, order uint32) {
	return
}

// Mark 2^order number of pages starting from addr as used.
func (ba *buddyAllocator) Mark(addr, order uint32) {
	return
}
