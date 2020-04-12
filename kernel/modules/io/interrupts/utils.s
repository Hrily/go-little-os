; kernel_modules_io_interrupts.Enable
; enables interrupts
global kernel_modules_io_interrupts.Enable
kernel_modules_io_interrupts.Enable:
	sti
	ret

; kernel_modules_io_interrupts.Disable
; disables interrupts
global kernel_modules_io_interrupts.Disable
kernel_modules_io_interrupts.Disable:
	cli
	ret
