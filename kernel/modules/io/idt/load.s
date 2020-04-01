global kernel_modules_io_idt.LoadIDT

idtr DW 0 ; For limit storage
     DD 0 ; For base storage

; kernel_modules_io_idt.LoadIDT
; Loads idt
; stack: [esp + 4] address to idt struct
;        [esp + 8] size to idt struct
kernel_modules_io_idt.LoadIDT:
	cli
	mov  eax, [esp + 4]
	mov  [idtr + 2], eax
	mov  ax, [esp + 8]
	mov  [idtr], ax
	lidt [idtr]
	sti
	ret
