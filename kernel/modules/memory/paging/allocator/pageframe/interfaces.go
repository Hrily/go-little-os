package pageframe

// Allocator is page frame allocator
type Allocator interface {
	// Allocate 2^order number of pages and return physical address of the first
	// page frame. Will return ok as false if it failed to allocate.
	Allocate(order uint32) (addr uint32, ok bool)
	// Release marks 2^order number of pages starting from addr as free and
	// unused.
	Release(addr, order uint32)
	// Mark 2^order number of pages starting from addr as used.
	Mark(addr, order uint32)
}
