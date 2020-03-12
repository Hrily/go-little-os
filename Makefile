SOURCE_DIRECTORY = .

get_object_file = $(addsuffix .o,$(join $(dir), $(shell basename $(dir))))

GO_SOURCES     = $(shell find $(SOURCE_DIRECTORY) -name *.go)
GO_SOURCE_DIRS = $(dir $(GO_SOURCES))
GO_OBJECTS     = $(foreach dir,$(GO_SOURCE_DIRS),$(get_object_file))
GO_INCLUDES    = $(addprefix -I,$(GO_SOURCE_DIRS))
GO_LINKS       = $(addprefix -L,$(GO_SOURCE_DIRS))

OBJECTS = ${GO_OBJECTS} loader.o

GCCGO := gccgo
GCCGOFLAGS = -fno-split-stack

LD := ld
LDFLAGS = -T link.ld -m elf_i386

AS := nasm
ASFLAGS = -f elf

all: os.iso

kernel.elf: $(OBJECTS)
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

$(GO_OBJECTS):
	$(GCCGO) $(GCCGOFLAGS) -c $(shell ls $(dir $@)*.go) -o $@ $(GO_INCLUDES) $(GO_LINKS)

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
