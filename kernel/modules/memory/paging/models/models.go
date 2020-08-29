package models

import (
	"kernel/utils/integer"
)

// PageSize determines the size of page 4KB or 4MB
type PageSize bool

const (
	// Size4KB represents 4KBpage size
	Size4KB PageSize = false
	// Size4MB represents 4KBpage size
	Size4MB PageSize = true
)

// ToBytes converts page size to number of bytes
func (p PageSize) ToBytes() uint32 {
	switch p {
	case Size4KB:
		return 4 * 1024
	case Size4MB:
		return 4 * 1024 * 1024
	}
	return 0
}

// WritePolicy determines the policy used incase of write hit (cache)
type WritePolicy bool

const (
	// WriteBack is write hit policy where data is written to cache and to memory
	// later
	WriteBack WritePolicy = false
	// WriteThrough is write hit policy where data is written to cache as well as
	// in memory
	WriteThrough WritePolicy = true
)

// Frame is an entry in Page Directory or Table
// May represent a page table entry or directly page entry
type Frame struct {
	VirtualAddress  uint32
	PageAddress     uint32
	IsGlobal        bool
	Size            PageSize
	IsDirty         bool
	IsAccessed      bool
	IsCacheDisabled bool
	WritePolicy     WritePolicy
	IsUserPage      bool
	IsReadWrite     bool
	IsPresent       bool
}

// ToRecord converts entry to a record which can be stored in table
func (e *Frame) ToRecord() uint32 {
	var entry uint32
	// | 31 .. 12 | 11 .. 9 | 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1   | 0 |
	// | PageAddr | Future  | G | S | D | A | C | W | U | R/W | P |
	entry = e.PageAddress & 0xfffff000
	entry |= integer.BoolToUInt32(e.IsGlobal) << 8
	entry |= integer.BoolToUInt32(bool(e.Size)) << 7
	entry |= integer.BoolToUInt32(e.IsDirty) << 6
	entry |= integer.BoolToUInt32(e.IsAccessed) << 5
	entry |= integer.BoolToUInt32(e.IsCacheDisabled) << 4
	entry |= integer.BoolToUInt32(bool(e.WritePolicy)) << 3
	entry |= integer.BoolToUInt32(e.IsUserPage) << 2
	entry |= integer.BoolToUInt32(e.IsReadWrite) << 1
	entry |= integer.BoolToUInt32(e.IsPresent) << 0
	return entry
}

// PT is the Page Table
type PT struct {
	Entries [1024]uint32
}

// Load loads entry into page table
func (p *PT) Load(e *Frame) {
	index := e.VirtualAddress >> 12
	index &= 0x3ff
	p.Entries[index] = e.ToRecord()
}

// PDT is the Page Directory
type PDT PT

// Load loads entry into page table
func (p *PDT) Load(e *Frame) {
	index := e.VirtualAddress >> 22
	p.Entries[index] = e.ToRecord()
}
