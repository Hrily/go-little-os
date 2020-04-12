package paging

import (
	"unsafe"

	"kernel/lib/logger"
)

var kernelPDT = PDT{}

// LoadPDT loads pdt. Defined in load.s
// pdtAddr is the physical address of pdt.
func LoadPDT(pdtAddr uint32)

// InvalidateTLB invalidates  Translation Lookaside Buffer (TLB). Defined in
// load.s
func InvalidateTLB()

// LoadKernelPDT loads page directory table for kernel
func LoadKernelPDT(pAddr, vAddr, size uint32) {
	// We'll load kernel in 4MB pages
	var _4mb uint32 = 4 * 1024 * 1024
	nPages := uint32(size / _4mb)
	if size%_4mb > 0 {
		nPages++
	}
	logger.COM().LogUint(logger.Debug, "# Kernel Pages", uint64(nPages))

	// check if physical address is 4mb aligned
	_4mbMask := _4mb - 1
	if pAddr&_4mbMask != 0 {
		// pAddr is not 4mb aligned
		alignedPAddr := pAddr & ^_4mbMask
		vAddr -= pAddr - alignedPAddr
		pAddr = alignedPAddr
	}

	for i := uint32(0); i < nPages; i++ {
		e := Entry{
			VirtualAddress: vAddr + i*_4mb,
			PageAddress:    pAddr + i*_4mb,
			Size:           Size4MB,
			IsReadWrite:    true,
			IsPresent:      true,
		}
		logger.COM().LogUint(logger.Debug, "Mapping vAddr", uint64(e.VirtualAddress))
		logger.COM().LogUint(logger.Debug, "     To pAddr", uint64(e.PageAddress))
		kernelPDT.Load(e)
	}
	pdtVAddr := uint32(uintptr(unsafe.Pointer(&kernelPDT)))
	logger.COM().LogUint(logger.Debug, "kernelPDT", uint64(pdtVAddr))
	// convert pdt address to physical address
	pdtPAddr := pAddr + (pdtVAddr - vAddr)
	LoadPDT(pdtPAddr)
	logger.COM().Info("Kernel PDT loaded succesfully")
}
