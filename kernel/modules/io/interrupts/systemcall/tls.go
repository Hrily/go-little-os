package systemcall

import (
	"kernel/lib/logger"
	"kernel/modules/memory/gdt"
	"kernel/utils/pointer"
)

type TLSEntry struct {
	Number   uint32
	BaseAddr uint32
	Limit    uint32
	Flags    uint32
}

const (
	// tlsEntryNumber in GDT
	tlsEntryNumber = 3
)

var (
	tlsEntry *TLSEntry
)

func setThreadArea(entryAddr uint32) int32 {
	logger.COM().Error("setting thread area")

	tlsEntry = (*TLSEntry)(pointer.Get(entryAddr))
	tlsEntry.Number = tlsEntryNumber

	gdt.KernelTLSSegment.Base = tlsEntry.BaseAddr
	gdt.KernelTLSSegment.Limit = tlsEntry.Limit

	logger.COM().LogUint(logger.Debug, "Base", uint64(tlsEntry.BaseAddr))
	logger.COM().LogUint(logger.Debug, "Limit", uint64(tlsEntry.Limit))

	if ok := gdt.AddToGDT(tlsEntryNumber, &gdt.KernelTLSSegment); ok {
		return 0
	}
	return -1
}

func getThreadArea(entryAddr uint32) int32 {
	if tlsEntry == nil {
		return -1
	}
	entry := (*TLSEntry)(pointer.Get(entryAddr))
	entry.Number = tlsEntry.Number
	entry.BaseAddr = tlsEntry.BaseAddr
	entry.Limit = tlsEntry.Limit
	entry.Flags = tlsEntry.Flags
	return 0
}
