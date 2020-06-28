ENTRY(loader)                /* the name of the entry label */

SECTIONS {

		. = 0xC0200000;          /* the code should be loaded at 3GB + 2 MB */

		/* these labels get exported to the code files */
		kernel_virtual_start  = .;
		kernel_physical_start = . - 0xC0000000;

		.multiboot ALIGN (0x1000) : AT(ADDR(.multiboot)-0xC0000000)
		{
				*(.multiboot)
		}


		.text ALIGN (0x1000) : AT(ADDR(.text)-0xC0000000)
		{
				*(.text)             /* all text sections from all files */
		}

		.rodata ALIGN (0x1000) : AT(ADDR(.rodata)-0xC0000000)
		{
				*(.rodata*)          /* all read-only data sections from all files */
		}

		.data ALIGN (0x1000) : AT(ADDR(.data)-0xC0000000)
		{
				*(.data)             /* all data sections from all files */
		}

		.bss ALIGN (0x1000) : AT(ADDR(.bss)-0xC0000000)
		{
				*(COMMON)            /* all COMMON sections from all files */
				*(.bss)              /* all bss sections from all files */
		}

		.preinit_array ALIGN (0x1000) : AT(ADDR(.preinit_array)-0xC0000000)
		{
			PROVIDE_HIDDEN (__preinit_array_start = .);
			KEEP (*(.preinit_array))
			PROVIDE_HIDDEN (__preinit_array_end = .);
		}
		.init_array ALIGN (0x1000) : AT(ADDR(.init_array)-0xC0000000)
		{
			PROVIDE_HIDDEN (__init_array_start = .);
			KEEP (*(SORT_BY_INIT_PRIORITY(.init_array.*) SORT_BY_INIT_PRIORITY(.ctors.*)))
			KEEP (*(.init_array EXCLUDE_FILE (*crtbegin.o *crtbegin?.o *crtend.o *crtend?.o ) .ctors))
			PROVIDE_HIDDEN (__init_array_end = .);
		}
		.fini_array ALIGN (0x1000) : AT(ADDR(.fini_array)-0xC0000000)
		{
			PROVIDE_HIDDEN (__fini_array_start = .);
			KEEP (*(SORT_BY_INIT_PRIORITY(.fini_array.*) SORT_BY_INIT_PRIORITY(.dtors.*)))
			KEEP (*(.fini_array EXCLUDE_FILE (*crtbegin.o *crtbegin?.o *crtend.o *crtend?.o ) .dtors))
			PROVIDE_HIDDEN (__fini_array_end = .);
		}

		kernel_virtual_end  = .;
		kernel_physical_end = . - 0xC0000000;
}
