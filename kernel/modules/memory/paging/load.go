package paging

import (
	"kernel/lib/logger"
	"kernel/modules/memory/paging/models"
	"kernel/utils/pointer"
)

const (
	_4kb uint32 = 4 * 1024
	_4mb uint32 = _4kb * 1024
)

// LoadKernelPDT loads page directory table for kernel
func LoadKernelPDT(pAddr, vAddr, size uint32) {
	getPAddr := func(_vAddr uint32) uint32 {
		return pAddr + (_vAddr - vAddr)
	}

	// check if physical address is 4mb aligned
	_4mbMask := _4mb - 1
	if pAddr&_4mbMask != 0 {
		// pAddr is not 4mb aligned
		alignedPAddr := pAddr & ^_4mbMask
		vAddr -= pAddr - alignedPAddr
		size += pAddr - alignedPAddr
		pAddr = alignedPAddr
	}

	// We'll load kernel in 4MB pages
	nPages := uint32(size / _4mb)
	if size%_4mb > 0 {
		nPages++
	}
	logger.COM().LogUint(logger.Debug, "# Kernel Pages", uint64(nPages))

	// start kernel PDT at the end of kernel space
	kernelPDTAddr := vAddr + nPages*_4mb
	kernelPDT := (*models.PDT)(pointer.Get(kernelPDTAddr))
	logger.COM().LogUint(logger.Debug, "kernelPDTAddr", uint64(kernelPDTAddr))

	for i := uint32(0); i < nPages; i++ {
		e := &models.Frame{
			VirtualAddress: vAddr + i*_4mb,
			PageAddress:    pAddr + i*_4mb,
			Size:           models.Size4MB,
			IsReadWrite:    true,
			IsPresent:      true,
		}
		logger.COM().LogUint(logger.Debug, "Mapping vAddr", uint64(e.VirtualAddress))
		logger.COM().LogUint(logger.Debug, "     To pAddr", uint64(e.PageAddress))
		kernelPDT.Load(e)
	}

	// map kernelPDT in kernelPDT
	ptForKernelPDTAddr := vAddr + nPages*_4mb + _4kb
	ptForKernelPDT := (*models.PT)(pointer.Get(ptForKernelPDTAddr))
	logger.COM().LogUint(logger.Debug, "ptForKernelPDTAddr", uint64(ptForKernelPDTAddr))
	ptForKernelPDT.Load(&models.Frame{
		VirtualAddress: kernelPDTAddr,
		PageAddress:    getPAddr(kernelPDTAddr),
		Size:           models.Size4KB,
		IsPresent:      true,
	})
	kernelPDT.Load(&models.Frame{
		VirtualAddress: ptForKernelPDTAddr,
		PageAddress:    getPAddr(ptForKernelPDTAddr),
		Size:           models.Size4KB,
		IsPresent:      true,
	})

	LoadPDT(getPAddr(kernelPDTAddr))
	logger.COM().Info("Kernel PDT loaded succesfully")
}
