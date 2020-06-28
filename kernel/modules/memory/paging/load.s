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

; kernel_modules_memory_paging.GetTLSPTAddr
; gets address of tls page table
; returns: eax address of pt
global kernel_modules_memory_paging.GetTLSPTAddr
kernel_modules_memory_paging.GetTLSPTAddr:
	mov eax, tlsPT
	ret

; kernel_modules_memory_paging.GetTLSPage1Addr
; gets address of tls page #1
; returns: eax address of page
global kernel_modules_memory_paging.GetTLSPage1Addr
kernel_modules_memory_paging.GetTLSPage1Addr:
	mov eax, tlsPage1
	ret

; kernel_modules_memory_paging.GetTLSPage2Addr
; gets address of tls page #2
; returns: eax address of page
global kernel_modules_memory_paging.GetTLSPage2Addr
kernel_modules_memory_paging.GetTLSPage2Addr:
	mov eax, tlsPage2
	ret

section .bss
; kernel_modules_memory_paging.kernelPDT
; kernel pdt
; this is written here for aligning it at 4KB
align 0x1000
kernelPDT:
	resb 4 * 1024
; kernel_modules_memory_paging.tlsPT
; TLS Page Table
; this is written here for aligning it at 4KB
align 0x1000
tlsPT:
	resb 4 * 1024
; kernel_modules_memory_paging.tlsPage1
; TLS Page #1
; this is written here for aligning it at 4KB
align 0x1000
tlsPage1:
	resb 4 * 1024
; kernel_modules_memory_paging.tlsPage2
; TLS Page #2
; this is written here for aligning it at 4KB
align 0x1000
tlsPage2:
	resb 4 * 1024
