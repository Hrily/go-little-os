; kernel_modules_memory_paging.LoadPDT
; loads pdt
; stack: [esp + 4] physical address of pdt
global kernel_modules_memory_paging.LoadPDT
kernel_modules_memory_paging.LoadPDT:
	mov eax, [esp + 4]
	; eax has the address of the page directory
	mov cr3, eax
	mov ebx, cr4        ; read current cr4
	or  ebx, 0x00000010 ; set PSE
	mov cr4, ebx        ; update cr4
	mov ebx, cr0        ; read current cr0
	or  ebx, 0x80000000 ; set PG
	mov cr0, ebx        ; update cr0
	ret

; kernel_modules_memory_paging.InvalidateTLB
; invalidates Translation Lookaside Buffer (TLB)
global kernel_modules_memory_paging.InvalidateTLB
kernel_modules_memory_paging.InvalidateTLB:
	; invalidate any TLB references to virtual address 0
	invlpg [0]

; kernel_modules_memory_paging.GetKernelPDTAddr
; gets address of kernel's pdt
; returns: eax address of pdt
global kernel_modules_memory_paging.GetKernelPDTAddr
kernel_modules_memory_paging.GetKernelPDTAddr:
	mov eax, kernelPDT
	ret

section .bss
; kernel_modules_memory_paging.kernelPDT
; kernel pdt
; this is written here for aligning it at 4KB
align 0x1000
kernelPDT:
	resb 4 * 1024
