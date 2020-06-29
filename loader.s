MAGIC_NUMBER equ 0x1BADB002     ; define the magic number constant
FLAGS        equ 0x0            ; multiboot flags
CHECKSUM     equ -MAGIC_NUMBER  ; calculate the checksum
			    ; (magic number + checksum + flags should equal 0)

section .multiboot
	align 4                       ; the code must be 4 byte aligned
	dd MAGIC_NUMBER               ; write the magic number to the machine code,
	dd FLAGS                      ; the flags,
	dd CHECKSUM                   ; and the checksum

; This is the virtual base address of kernel space. It must be used to convert virtual
; addresses into physical addresses until paging is enabled. Note that this is not
; the virtual address where the kernel image itself is loaded -- just the amount that must
; be subtracted from a virtual address to get a physical address.
KERNEL_VIRTUAL_BASE equ 0xC0000000                  ; 3GB
KERNEL_PAGE_NUMBER equ (KERNEL_VIRTUAL_BASE >> 22)  ; Page directory index of kernel's 4MB PTE.
KERNEL_STACK_SIZE equ 4 * 1024 * 1024     ; size of stack in bytes

section .data
align 0x1000
boot_page_directory:
	; This page directory entry identity-maps the first 4MB of the 32-bit physical address space.
	; All bits are clear except the following:
	; bit 7: PS The kernel page is 4MB.
	; bit 1: RW The kernel page is read/write.
	; bit 0: P  The kernel page is present.
	; This entry must be here -- otherwise the kernel will crash immediately after paging is
	; enabled because it can't fetch the next instruction! It's ok to unmap this page later.
	dd 0x00000083 + 0*0x400000
	dd 0x00000083 + 1*0x400000
	dd 0x00000083 + 2*0x400000
	dd 0x00000083 + 3*0x400000   ; will contain the actual kernel pdt when loaded
	times (KERNEL_PAGE_NUMBER - 4) dd 0                 ; Pages before kernel space.
	; This page directory entry defines a 4MB page containing the kernel.
	dd 0x00000083 + 0*0x400000
	dd 0x00000083 + 1*0x400000
	dd 0x00000083 + 2*0x400000
	dd 0x00000083 + 3*0x400000
	times (1024 - KERNEL_PAGE_NUMBER - 8) dd 0  ; Pages after the kernel image.

section .text                  ; start of the text (code) section

; setting up entry point for linker
global loader                   ; the entry symbol for ELF
loader equ (_loader - 0xC0000000)
_loader:                         ; the loader label (defined as entry point in linker script)
mov eax, (boot_page_directory - KERNEL_VIRTUAL_BASE)
; eax has the address of the page directory
mov cr3, eax
mov ebx, cr4        ; read current cr4
or  ebx, 0x00000010 ; set PSE
mov cr4, ebx        ; update cr4
mov ebx, cr0        ; read current cr0
or  ebx, 0x80000000 ; set PG
mov cr0, ebx        ; update cr0
; now paging is enabled

; assembly code executing at around 0x00100000
; enable paging for both actual location of kernel
; and its higher-half virtual location

lea ebx, [higher_half] ; load the address of the label in ebx
jmp ebx                ; jump to the label

higher_half:
	; code here executes in the higher half kernel
	; eip is larger than 0xC0000000
	; can continue kernel initialisation, calling C code, etc.

	; Unmap the identity-mapped first 4MB of physical address space. It should not be needed
	; anymore.
	mov dword [boot_page_directory], 0
	mov dword [boot_page_directory + 4], 0
	invlpg [0]
	; NOTE: From now on, paging should be enabled. The first 4MB of physical address space is
	; mapped starting at KERNEL_VIRTUAL_BASE. Everything is linked to this address, so no more
	; position-independent code or funny business with virtual-to-physical address translation
	; should be necessary. We now have a higher-half kernel.
	mov esp, kernel_stack + KERNEL_STACK_SIZE   ; set up the stack

	; The assembly code

	; GRUB variables
	extern kernel_virtual_start
	extern kernel_virtual_end
	extern kernel_physical_start
	extern kernel_physical_end
	push kernel_physical_end
	push kernel_physical_start
	push kernel_virtual_end
	push kernel_virtual_start
	extern kernel.Main     ; the function sum_of_three is defined elsewhere
	call kernel.Main       ; call the function, the result will be in eax
																							; stack (end of memory area)
	.loop:
	jmp .loop                   ; loop forever

; define __go_runtime_error
; TODO: Add code to handle/print error
__go_runtime_error:
	ret

global main.main
main.main:
	ret

global _end
_end:
	ret

global __go_init_main
__go_init_main:
	ret

section .bss
; Kernel Stack
align 4                                   ; align at 4 bytes
kernel_stack:                             ; label points to beginning of memory
	resb KERNEL_STACK_SIZE                  ; reserve stack for the kernel
