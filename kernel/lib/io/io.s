 global kernel_lib_io.OutB     ; make the label outb visible outside this file

; kernel_lib_io.OutB - send a byte to an I/O port:
; stack: [esp + 8] the data byte
;        [esp + 4] the I/O port
;        [esp    ] return address
kernel_lib_io.OutB:
	mov al, [esp + 8]    ; move the data to be sent into the al register
	mov dx, [esp + 4]    ; move the address of the I/O port into the dx register
	out dx, al           ; send the data to the I/O port
	ret                  ; return to the calling function
