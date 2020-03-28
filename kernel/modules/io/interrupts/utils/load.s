; func GetFuncAddr
; Returns address of functions. This is based on fact that the func address is
; stored at *[esp + 4]
global kernel_modules_io_interrupts_utils.GetFuncAddr
kernel_modules_io_interrupts_utils.GetFuncAddr:
	mov  eax, [esp + 4]
	mov  eax, [eax]
	ret
