global kernel_modules_memory_gdt.LoadGDT

gdtr DW 0 ; For limit storage
     DD 0 ; For base storage

; kernel_modules_memory_gdt.LoadGDT
; Loads gdt
; stack: [esp + 4] address to gdt struct
;        [esp + 8] size to gdt struct
kernel_modules_memory_gdt.LoadGDT:
	mov  eax, [esp + 4]
	mov  [gdtr + 2], eax
	mov  ax, [esp + 8]
	mov  [gdtr], ax
	lgdt [gdtr]
	jmp  reload_segments
reload_segments:
	; Reload cs register containing code selector:
	jmp   0x08:reload_cs ; 0x08 points at the new code selector
reload_cs:
	; Reload data segment registers:
	mov   ax, 0x10 ; 0x10 points at the new data selector
	mov   ds, ax
	mov   es, ax
	mov   fs, ax
	mov   gs, ax
	mov   ss, ax
	ret
