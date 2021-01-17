%include "kernel/modules/io/interrupts/common/macro.m"

; System Call Interrupt Handler
no_error_code_interrupt_handler systemcall,0x80
