extern __libc_setup_tls
extern runtime.malg
extern runtime.setg

; Calls __libc_setup_tls
global kernel_modules_memory_tls.LibCSetupTLS
kernel_modules_memory_tls.LibCSetupTLS:
	call __libc_setup_tls
	push 0x0
	push 0x0
	call runtime.malg
	call runtime.setg
	ret
