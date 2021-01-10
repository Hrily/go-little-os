package buddy

import (
	"unsafe"

	"kernel/modules/memory/paging/allocator/pageframe"
	"kernel/modules/memory/paging/models"
	"kernel/utils/pointer"
)

const (
	// MaxOrder is the maximum order of 2 pages we can allocate.
	// for now, we assume this to be 4MB = 1024 * 4KB pages.
	MaxOrder = 10

	// _nBigPages is total number of 4mb pages. We assume initial memory of 4GB.
	_nBigPages = 1024
)

// freeBuddiesList stores free pages of a specific order. It uses one bit per
// buddy pair of pages.
// If a bit for a buddy pair is:
//   - 1 : one of the pages in buddy is free
//   - 0 : both pages are either free or allocated
type freeBuddiesList struct {
	freeMap  bitmap
	freeList list
}

// buddyAllocator uses buddy list to allocate page frames.
// It implements Allocator interface.
// Resources will explain a lot better than I ever will be able to, so just
// links, no comments.
//   - OS Dev Wiki: https://wiki.osdev.org/Page_Frame_Allocation#Buddy_Allocation_System
//   - Kernel Docs: https://www.kernel.org/doc/gorman/html/understand/understand009.html
type buddyAllocator struct {
	// buddies is collection of free buddies list
	buddies []freeBuddiesList
}

// _buddyAllocator is the instance of buddyAllocator
var _buddyAllocator buddyAllocator

// AllocatorSize returns total size which will be used by buddyAllocator.
// This is useful while kernel page allocation
// To know how this is computed, check Init function
func AllocatorSize() (size uint32) {
	// size of freeBuddiesList
	size += MaxOrder * uint32(unsafe.Sizeof(freeBuddiesList{}))

	// size of bigPagesBitmap
	size += nMaps(_nBigPages) * 4

	// size of individual freeBuddiesList
	for i := 0; i < MaxOrder; i++ {
		// maxOrder - 1 pages of order i, further divide by 2 since we use 1 bit
		// for buddy pair.
		var nBuddies uint32 = _nBigPages * (1 << uint32(MaxOrder-i-1))
		size += nMaps(nBuddies) * 4
	}

	// size of node pool
	size += nodePoolSize()

	return
}

// InitAllocator the buddyAllocator at given address
func InitAllocator(addr uint32) {
	// Warning: Lot of unsafe pointer usage and manipulation ahead.
	// This function basically initializes memory at given address, so that
	// _buddyAllocator, and it's members, point to this address.

	// create freeBuddiesList
	var _freeBuddiesList *[MaxOrder + 1]freeBuddiesList
	_freeBuddiesList = (*[MaxOrder + 1]freeBuddiesList)(pointer.Get(addr))
	addr += (MaxOrder + 1) * uint32(unsafe.Sizeof(freeBuddiesList{}))

	// set freeBuddiesList in buddyAllocator
	_buddyAllocator.buddies = (*_freeBuddiesList)[:MaxOrder+1]

	// create bigPagesBitmap
	// _nBigPages size for array is way more than needed. this is done since
	// array sizes need to be constant.
	var _bigPagesBitmap *[_nBigPages]uint32
	_bigPagesBitmap = (*[_nBigPages]uint32)(pointer.Get(addr))
	addr += nMaps(_nBigPages) * 4

	// set bigPagesBitmap in buddyAllocator
	_buddyAllocator.buddies[MaxOrder].freeMap.maps = (*_bigPagesBitmap)[:nMaps(_nBigPages)]

	// mark all big pages as free
	for i := uint32(0); i < nMaps(_nBigPages); i++ {
		_buddyAllocator.buddies[MaxOrder].freeMap.maps[i] = _allSet
	}

	// create the individual freeBuddiesList
	for i := 0; i < MaxOrder; i++ {
		// maxOrder - 1 pages of order i, further divide by 2 since we use 1 bit
		// for buddy pair.
		var nBuddies uint32 = _nBigPages * (1 << uint32(MaxOrder-i-1))
		nMaps := nMaps(nBuddies)

		// set address for this freeBuddiesList
		// we are addressing way more memory than needed, because array sizes need
		// to be constant. this will be more than enough for largest freeBuddiesList
		var maps *[_nBigPages * _nBigPages]uint32
		maps = (*[_nBigPages * _nBigPages]uint32)(pointer.Get(addr))
		addr += nMaps * 4

		// set the freeBuddiesList
		_buddyAllocator.buddies[i].freeMap.maps = (*maps)[:nMaps]

		// zero out the freeBuddiesList
		for j := uint32(0); j < nMaps; j++ {
			_buddyAllocator.buddies[i].freeMap.maps[j] = 0
		}
	}

	initNodePool(addr)
}

// Allocator returns buddy allocator
func Allocator() pageframe.Allocator {
	return &_buddyAllocator
}

// getAddress returns address given index and order. In case of buddy pair,
// first will determine if returned address is of first page or second.
func (ba *buddyAllocator) getAddress(
	index, order uint32, first bool,
) (addr uint32) {
	// If it is bigPage
	if order >= MaxOrder {
		return index * models.Size4MB.ToBytes()
	}
	// not a bigPage, is part of buddy
	pageIndex := 2 * index
	if !first {
		pageIndex++
	}
	return pageIndex * (1 << order) * models.Size4KB.ToBytes()
}

// getIndex returns index given address and order. This is inverse function of
// getAddress. first will be true if this is first page in buddy of that order
func (ba *buddyAllocator) getIndex(
	addr, order uint32,
) (index uint32, first bool) {
	if order >= MaxOrder {
		return addr / models.Size4MB.ToBytes(), false
	}
	pageIndex := addr / models.Size4KB.ToBytes() / (1 << order)
	first = (pageIndex % 2) == 0
	index = pageIndex / 2
	return
}

// getBuddyAddress returns address of buddy of given page
func (ba *buddyAllocator) getBuddyAddress(
	index, order uint32, first bool,
) (addr uint32) {
	return ba.getAddress(index, order, !first)
}

// Allocate 2^order number of pages and return physical address of the first
// page frame. Will return ok as false if it failed to allocate.
func (ba *buddyAllocator) Allocate(order uint32) (addr uint32, ok bool) {
	if order >= MaxOrder {
		// we need order / _maxOrder pages
		n := order / MaxOrder
		index, ok := ba.buddies[MaxOrder].freeMap.FirstContiguousSet(n)
		if !ok {
			return 0, false
		}
		// Mark n pages from index as used
		for i := uint32(0); i < n; i++ {
			ba.buddies[MaxOrder].freeMap.Reset(index + i)
		}
		return ba.getAddress(index, order, false), true
	}

	// Check if freeBuddiesList of that order has a free page
	if !ba.buddies[order].freeList.IsEmpty() {
		addr = ba.buddies[order].freeList.Head().Value()
		// remove this addr from the list
		ba.buddies[order].freeList.Delete(addr)
		// map page as used
		index, _ := ba.getIndex(addr, order)
		ba.buddies[order].freeMap.Toggle(index)
		return addr, true
	}

	// The hard part, check higher order for free page
	higherOrder := order + 1
	higherOrderIndex := uint32(0)
	for higherOrder < MaxOrder {
		// Check if higherOrder freeBuddiesList has free page
		if !ba.buddies[higherOrder].freeList.IsEmpty() {
			higherOrderAddr := ba.buddies[higherOrder].freeList.Head().Value()
			// remove this addr from the list
			ba.buddies[higherOrder].freeList.Delete(higherOrderAddr)
			// map page as used
			higherOrderIndex, _ = ba.getIndex(higherOrderAddr, higherOrder)
			break
		}
		higherOrder++
	}

	if higherOrder == MaxOrder {
		higherOrderIndex, ok = ba.buddies[MaxOrder].freeMap.FirstSet()
		if !ok {
			return 0, false
		}
		// mark higherOrderIndex as used
		ba.buddies[MaxOrder].freeMap.Reset(higherOrderIndex)
	}

	// till we reach desired order
	for higherOrder > order {
		// go to previous order
		higherOrder--
		// previous order index is x2, unless we are on _maxOrder - 1
		if higherOrder < MaxOrder-1 {
			higherOrderIndex <<= 1
		}

		// Add first buddy page to free list
		if ok := ba.buddies[higherOrder].freeList.Append(
			ba.getAddress(higherOrderIndex, higherOrder, true),
		); !ok {
			// we were not able to append the page to free list
			// this should not happen. TODO: handle gracefully
			return 0, false
		}

		// Mark buddy
		ba.buddies[higherOrder].freeMap.Toggle(higherOrderIndex)
	}

	// now higherOrder == order, return addr of second page in buddy
	return ba.getAddress(higherOrderIndex, higherOrder, false), true
}

// Release marks 2^order number of pages starting from addr as free and unused.
func (ba *buddyAllocator) Release(addr, order uint32) {
	index, first := ba.getIndex(addr, order)

	if order >= MaxOrder {
		// we need to free order / _maxOrder pages
		n := order / MaxOrder
		// Mark n pages from index as free
		for i := uint32(0); i < n; i++ {
			ba.buddies[MaxOrder].freeMap.Set(index + i)
		}
		return
	}

	// go up until we get a buddy with one page allocated
	for order <= MaxOrder {
		// check if other buddy is still allocated
		if !ba.buddies[order].freeMap.IsSet(index) {
			// mark buddy as free
			ba.buddies[order].freeMap.Toggle(index)
			// add this page to free list
			ba.buddies[order].freeList.Append(addr)
			// no need to go further
			break
		}
		// buddy is also free, combine buddies
		buddyAddress := ba.getBuddyAddress(index, order, first)
		ba.buddies[order].freeList.Delete(buddyAddress)
		// now, mark both pages are free
		ba.buddies[order].freeMap.Toggle(index)
		// go to next order
		order++
		if buddyAddress < addr {
			addr = buddyAddress
		}
		index, first = ba.getIndex(addr, order)
	}
	return
}

// Mark 2^order number of pages starting from addr as used.
func (ba *buddyAllocator) Mark(addr, order uint32) {
	index, first := ba.getIndex(addr, order)

	if order >= MaxOrder {
		// we need to mark order / _maxOrder pages
		n := order / MaxOrder
		// Mark n pages from index as allocated
		for i := uint32(0); i < n; i++ {
			ba.buddies[MaxOrder].freeMap.Reset(index + i)
		}
		return
	}

	// Mark index as allocated
	ba.buddies[order].freeMap.Toggle(index)
	// If buddy is free add it to free list
	if ba.buddies[order].freeMap.IsSet(index) {
		buddyAddress := ba.getBuddyAddress(index, order, first)
		ba.buddies[order].freeList.Append(buddyAddress)
	}
	return
}
