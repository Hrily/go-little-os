package processinfo

import (
	"kernel/models"
)

func SetupKernelInfo(params models.KernelParams) {
	KernelSectionDataText.StartAddr = params.KernelVStartAddr
	KernelSectionDataText.Size = params.KernelVEndAddr - params.KernelVStartAddr
	KernelSectionDataText.Capacity = params.KernelVEndAddr - params.KernelVStartAddr

	KernelSectionHeap.StartAddr = params.KernelVEndAddr
	KernelSectionHeap.Size = 0
	KernelSectionHeap.Capacity = 0
}
