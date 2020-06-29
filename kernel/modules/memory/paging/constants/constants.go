package constants

import (
	"kernel/modules/memory/paging/models"
)

const (
	// Size4KB represents 4KBpage size
	Size4KB models.PageSize = false
	// Size4MB represents 4KBpage size
	Size4MB models.PageSize = true

	// WriteBack is write hit policy where data is written to cache and to memory
	// later
	WriteBack models.WritePolicy = false
	// WriteThrough is write hit policy where data is written to cache as well as
	// in memory
	WriteThrough models.WritePolicy = true
)
