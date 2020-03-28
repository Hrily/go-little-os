%include "kernel/modules/io/interrupts/common/macro.m"

; Interrupt Handlers for Exceptions

; Exception # 	Description 	                          Error Code?
; 0 	          Division By Zero Exception 	            No
no_error_code_interrupt_handler exceptions,0x00

; 1 	          Debug Exception 	                      No
no_error_code_interrupt_handler exceptions,0x01

; 2 	          Non Maskable Interrupt Exception 	      No
no_error_code_interrupt_handler exceptions,0x02

; 3 	          Breakpoint Exception 	                  No
no_error_code_interrupt_handler exceptions,0x03

; 4 	          Into Detected Overflow Exception 	      No
no_error_code_interrupt_handler exceptions,0x04

; 5 	          Out of Bounds Exception 	              No
no_error_code_interrupt_handler exceptions,0x05

; 6 	          Invalid Opcode Exception 	              No
no_error_code_interrupt_handler exceptions,0x06

; 7 	          No Coprocessor Exception 	              No
no_error_code_interrupt_handler exceptions,0x07

; 8 	          Double Fault Exception 	                Yes
error_code_interrupt_handler    exceptions,0x08

; 9 	          Coprocessor Segment Overrun Exception 	No
no_error_code_interrupt_handler exceptions,0x09

; 10 	          Bad TSS Exception 	                    Yes
error_code_interrupt_handler    exceptions,0x0a

; 11 	          Segment Not Present Exception 	        Yes
error_code_interrupt_handler    exceptions,0x0b

; 12 	          Stack Fault Exception 	                Yes
error_code_interrupt_handler    exceptions,0x0c

; 13 	          General Protection Fault Exception 	    Yes
error_code_interrupt_handler    exceptions,0x0d

; 14 	          Page Fault Exception 	                  Yes
error_code_interrupt_handler    exceptions,0x0e

; 15 	          Unknown Interrupt Exception 	          No
no_error_code_interrupt_handler exceptions,0x0f

; 16 	          Coprocessor Fault Exception 	          No
no_error_code_interrupt_handler exceptions,0x10

; 17 	          Alignment Check Exception (486+) 	      No
no_error_code_interrupt_handler exceptions,0x11

; 18 	          Machine Check Exception (Pentium/586+) 	No
no_error_code_interrupt_handler exceptions,0x12

; 19 to 31 	    Reserved Exceptions 	                  No
no_error_code_interrupt_handler exceptions,0x13
no_error_code_interrupt_handler exceptions,0x14
no_error_code_interrupt_handler exceptions,0x15
no_error_code_interrupt_handler exceptions,0x16
no_error_code_interrupt_handler exceptions,0x17
no_error_code_interrupt_handler exceptions,0x18
no_error_code_interrupt_handler exceptions,0x19
no_error_code_interrupt_handler exceptions,0x1a
no_error_code_interrupt_handler exceptions,0x1b
no_error_code_interrupt_handler exceptions,0x1c
no_error_code_interrupt_handler exceptions,0x1d
no_error_code_interrupt_handler exceptions,0x1e
no_error_code_interrupt_handler exceptions,0x1f
