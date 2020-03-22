package gdt

import (
	"kernel/utils/integer"
)

// ToUInt converts flags to it's bit structure
func (f *Flags) ToUInt() uint32 {
	// Flag structure
	// | 3 | 2 | 1 | 0 |
	// | G | S | 0 | 0 |
	var flag uint32 = 0
	flag |= integer.BoolToUInt32(bool(f.Granularity)) << 3
	flag |= uint32(f.AddressMode << 2)
	return flag
}

// ToUInt converts access byte to it's bit structure
func (a *Access) ToUInt() uint32 {
	var access uint32 = 0
	// Access structure
	// | 7  | 6 | 5 | 4 | 3  | 2   | 1  | 0  |
	// | Ps | Priv  | T | Ex | D/C | RW | Ac |
	access |= integer.BoolToUInt32(a.IsPresentInMemory) << 7
	access |= uint32(a.PrivilageLevel << 5)
	access |= integer.BoolToUInt32(bool(a.DescriptorType)) << 4
	access |= integer.BoolToUInt32(a.IsExecutable) << 3
	access |= integer.BoolToUInt32(
		a.IsExpandingDown || a.IsConforming,
	) << 2
	access |= integer.BoolToUInt32(
		a.IsReadable || a.IsWritable,
	) << 1
	access |= integer.BoolToUInt32(a.IsAccessed)
	return access
}

// ToDescriptorRecord converts GDT to it's physical structure
func (d *Descriptor) ToDescriptorRecord() DescriptorRecord {
	if d == nil {
		return DescriptorRecord(0)
	}
	var gdt uint64 = 0
	flag := uint64(d.Flags.ToUInt())
	access := uint64(d.Access.ToUInt())

	// Descriptor structure
	// | 31 ... 16 | 15 ..... 0 |
	// | Base 0:15 | Limit 0:15 |
	// | 63 ... 56  | 55 . 52 | 51 ..... 48 | 47 . 40 | 39 .... 32 |
	// | Base 24:31 | Flags   | Limit 16:19 | Access  | Base 16:23 |

	// Create the high 32 bit segment
	gdt = uint64(d.Limit) & 0x000F0000
	gdt |= (flag << 20) & 0x00F00000
	gdt |= (access << 8) & 0x0000FF00
	gdt |= uint64(d.Base>>16) & 0x000000FF
	gdt |= uint64(d.Base) & 0xFF000000

	// Shift by 32 to allow for low part of segment
	gdt <<= 32

	// Create the low 32 bit segment
	gdt |= uint64(d.Base) << 16
	gdt |= uint64(d.Limit) & 0x0000FFFF

	return DescriptorRecord(gdt)
}
