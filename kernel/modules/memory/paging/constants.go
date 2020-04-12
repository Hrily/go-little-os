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
)
