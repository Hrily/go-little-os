package paging

const (
	// Size4KB represents 4KBpage size
	Size4KB PageSize = false
	// Size4MB represents 4KBpage size
	Size4MB PageSize = true

	// WriteBack is write hit policy where data is written to cache and to memory
	// later
	WriteBack WritePolicy = false
	// WriteThrough is write hit policy where data is written to cache as well as
	// in memory
	WriteThrough WritePolicy = true

	// _4kb = 4KB
	_4kb uint32 = 4 * 1024
	// _4mb = 4MB
	_4mb uint32 = _4kb * 1024
)
