extern kernel_modules_io_interrupts_systemcall.SystemCall

; System Call Interrupt Handler
global kernel_modules_io_interrupts_systemcall.Handler
kernel_modules_io_interrupts_systemcall.Handler:
	cli
	push eax
	push ebx
	push ecx
	push edx
	push esi
	push edi
	push ebp
	call kernel_modules_io_interrupts_systemcall.SystemCall
	pop  ebp
	pop  edi
	pop  esi
	pop  edx
	pop  ecx
	pop  ebx
	add  esp, 4    ; pop eax
	sti
	iret           ; pops 5 things at once: CS, EIP, EFLAGS, SS, and ESP!

; V System Call Interrupt Handler
global kernel_modules_io_interrupts_systemcall.VHandler
kernel_modules_io_interrupts_systemcall.VHandler:
	cli
	push eax
	push ebx
	push ecx
	push edx
	push esi
	push edi
	push ebp
	call kernel_modules_io_interrupts_systemcall.SystemCall
	pop  ebp
	pop  edi
	pop  esi
	pop  edx
	pop  ecx
	pop  ebx
	add  esp, 4    ; pop eax
	sti
	ret

; SetSyscallHandler at gs:0x10
global kernel_modules_io_interrupts_systemcall.SetSyscallHandler
kernel_modules_io_interrupts_systemcall.SetSyscallHandler:
	mov [gs:0x10], eax
	ret
