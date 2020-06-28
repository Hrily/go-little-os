package paging

import (
	"unsafe"

	"kernel/lib/logger"
	"kernel/lib/memory"
)

// GetKernelPDTAddr returns kernel's pdt address
func GetKernelPDTAddr() uint32

// GetTLSPTAddr returns tls PT address
func GetTLSPTAddr() uint32

// GetTLSPage1Addr returns tls page #1 address
func GetTLSPage1Addr() uint32

// GetTLSPage2Addr returns tls page #2 address
func GetTLSPage2Addr() uint32

// LoadPDT loads pdt. Defined in load.s
// pdtAddr is the physical address of pdt.
func LoadPDT(pdtAddr uint32)

// InvalidateTLB invalidates  Translation Lookaside Buffer (TLB). Defined in
// load.s
func InvalidateTLB()

// LoadKernelPDT loads page directory table for kernel
func LoadKernelPDT(pAddr, vAddr, size uint32) {

	getPAddr := func(_vAddr uint32) uint32 {
		return pAddr + (_vAddr - vAddr)
	}

	// kernelPDT is defined in load.s since it requires 4KB alignment
	kernelPDTAddr := GetKernelPDTAddr()
	kernelPDT := (*PDT)(unsafe.Pointer(uintptr(kernelPDTAddr)))
	logger.COM().LogUint(logger.Debug, "kernelPDTAddr", uint64(kernelPDTAddr))

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

	// Add thread local storage page
	// This area stores the g pointer
	tlsPT := (*PT)(unsafe.Pointer(uintptr(GetTLSPTAddr())))
	// physical address of tlsPT
	tlsPTPAddr := getPAddr(GetTLSPTAddr())
	tlsPage1Addr := getPAddr(GetTLSPage1Addr())
	tlsPage2Addr := getPAddr(GetTLSPage2Addr())
	// for gs:0x30
	tlsPT.Load(Entry{
		VirtualAddress: 0x0,
		PageAddress:    tlsPage1Addr,
		Size:           Size4KB,
		IsReadWrite:    true,
		IsPresent:      true,
	})
	// for gs:0xffffffec
	tlsPT.Load(Entry{
		VirtualAddress: 0xfffff000,
		PageAddress:    tlsPage2Addr,
		Size:           Size4KB,
		IsReadWrite:    true,
		IsPresent:      true,
	})
	// for gs:0x30
	kernelPDT.Load(Entry{
		VirtualAddress: 0x0,
		PageAddress:    tlsPTPAddr,
		Size:           Size4KB,
		IsPresent:      true,
	})
	// for gs:0xffffffec
	kernelPDT.Load(Entry{
		VirtualAddress: 0xfffff000,
		PageAddress:    tlsPTPAddr,
		Size:           Size4KB,
		IsPresent:      true,
	})

	// Copy TLS header to TLS Page 2
	// TLS header starts from end of vaddr
	memory.MoveData(GetTLSPage2Addr()+0x1000-0x14, vAddr+size+0x200000, 0x14)

	pdtVAddr := uint32(uintptr(unsafe.Pointer(kernelPDT)))
	logger.COM().LogUint(logger.Debug, "kernelPDT", uint64(pdtVAddr))
	// convert pdt address to physical address
	pdtPAddr := getPAddr(pdtVAddr)
	LoadPDT(pdtPAddr)
	logger.COM().Info("Kernel PDT loaded succesfully")
}
