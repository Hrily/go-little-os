kernel/modules/init.o : kernel/lib/logger.o \
	kernel/models.o \
	kernel/modules/io/idt.o \
	kernel/modules/io/interrupts.o \
	kernel/modules/memory/gdt.o \
	kernel/modules/memory/paging.o \
	kernel/modules/memory/tls.o \
	kernel/modules/process/processinfo.o
