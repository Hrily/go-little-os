%ifndef KERNEL_MODULES_IO_INTERRYPTS_MACROS_M
%define KERNEL_MODULES_IO_INTERRYPTS_MACROS_M

%macro no_error_code_interrupt_handler 2
extern kernel_modules_io_interrupts_common.interruptHandler
global kernel_modules_io_interrupts_%1.Int%2
kernel_modules_io_interrupts_%1.Int%2:
	push    dword 0                     ; push 0 as error code
	push    dword %2                    ; push the interrupt number
	jmp     kernel_modules_io_interrupts_common.interruptHandler
%endmacro

%macro error_code_interrupt_handler 2
extern kernel_modules_io_interrupts_common.interruptHandler
global kernel_modules_io_interrupts_%1.Int%2
kernel_modules_io_interrupts_%1.Int%2:
	push    dword %2                    ; push the interrupt number
	jmp     kernel_modules_io_interrupts_common.interruptHandler
%endmacro

%endif
