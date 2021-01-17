extern __libc_setup_tls

; Calls __libc_setup_tls
global kernel_modules_memory_tls.LibCSetupTLS
kernel_modules_memory_tls.LibCSetupTLS:
	call __libc_setup_tls
	ret
