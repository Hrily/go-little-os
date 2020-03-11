GO_SOURCES = $(wildcard kernel/*.go)
GO_OBJECTS = ${GO_SOURCES:.go=.o}

OBJECTS = ${GO_OBJECTS}

GCCGO := gccgo
GCCGOFLAGS = -fno-split-stack

LD := ld
LDFLAGS = -T link.ld -m elf_i386

AS := nasm
ASFLAGS = -f elf

all: os.iso

kernel.elf: $(OBJECTS) loader.o
	$(LD) $(LDFLAGS) $^ -o $@

os.iso: kernel.elf
	cp kernel.elf iso/boot/kernel.elf
	genisoimage -R                          \
		-b boot/grub/stage2_eltorito    \
		-no-emul-boot                   \
		-boot-load-size 4               \
		-A os                           \
		-input-charset utf8             \
		-quiet                          \
		-boot-info-table                \
		-o os.iso                       \
		iso

%.o: %.s
	$(AS) $(ASFLAGS) $< -o $@

%.o: %.go
	$(GCCGO) $(GCCGOFLAGS) -c $< -o $@

run-bochs:
	bochs -f bochsrc.txt -q

build-docker:
	docker build . -t littleosbook; \
	docker create --name littleosbook-container littleosbook; \
	docker cp littleosbook-container:/littleosbook/kernel.elf .; \
	docker cp littleosbook-container:/littleosbook/os.iso .; \
	docker rm littleosbook-container

clean:
	rm -rf *.o */*.o kernel.elf os.iso
