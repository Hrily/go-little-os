package pageframe

import (
	"kernel/modules/memory/paging/models"
)

// Allocator is page frame allocator
type Allocator interface {
	// Allocate allocates n continuous pages if given size and returns the
	// physical address of the first page frame.
	// Will return ok as false if it failed to allocate
	Allocate(n uint32, size models.PageSize) (addr uint32, ok bool)
	// Release marks pages as free and unused
	Release(addr, n uint32, size models.PageSize)
	// Mark will mark given pages as used
	Mark(addr, n uint32, size models.PageSize)
}
