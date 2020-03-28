extern kernel_modules_io_interrupts_common.InterruptHandler
; func common_interrupt_handler
; This is our common ISR stub. It saves the processor state, sets
; up for kernel mode segments, calls the C-level fault handler,
; and finally restores the stack frame.
global kernel_modules_io_interrupts_common.interruptHandler
kernel_modules_io_interrupts_common.interruptHandler:
	push ds
	push es
	push fs
	push gs
	pusha
	mov ax, 0x10   ; Load the Kernel Data Segment descriptor!
	mov ds, ax
	mov es, ax
	mov fs, ax
	mov gs, ax
	mov eax, kernel_modules_io_interrupts_common.InterruptHandler
	call eax       ; A special call, preserves the 'eip' register
	popa
	pop gs
	pop fs
	pop es
	pop ds
	add esp, 8     ; Cleans up the pushed error code and pushed ISR number
	iret           ; pops 5 things at once: CS, EIP, EFLAGS, SS, and ESP!
