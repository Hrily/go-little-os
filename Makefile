SOURCE_DIRECTORY = .
BUILD_DIRECTORY  = build/

uniq=$(shell echo $(1) | tr " " "\n" | cat -n | sort -uk2 | sort -nk1| cut -f2- | tr "\n" " ")

GO_SOURCES     = $(shell find $(SOURCE_DIRECTORY) -name '*.go')
GO_SOURCE_DIRS = $(call uniq,$(dir $(GO_SOURCES)))
GO_OBJECTS     = $(GO_SOURCE_DIRS:/=.o)

AS_SOURCES     = $(shell find $(SOURCE_DIRECTORY) -name '*.s')
AS_OBJECTS     = $(AS_SOURCES:.s=.s.o)

OBJECTS = $(GO_OBJECTS) $(AS_OBJECTS)

DEPS = $(addsuffix deps.mk,$(GO_SOURCE_DIRS))

GCCGO := gccgo
GCCGOFLAGS = -fno-split-stack -fno-go-check-divide-zero

LD := ld
LDFLAGS = -T link.ld -m elf_i386 -lgo -L/usr/lib/gcc/i686-linux-gnu/8/

AS := nasm
ASFLAGS = -f elf

QEMU = qemu-system-i386
QEMUFLAGS = -serial file:serial.out

-include $(DEPS)

all: os.iso

kernel.elf: $(OBJECTS)
	$(LD) $(LDFLAGS) $(addprefix $(BUILD_DIRECTORY),$^) -o $@

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

.SECONDEXPANSION:
REQ_AS_SOURCES = $(shell ls $(dir $(1))/*.s)
$(AS_OBJECTS): %.s.o : $$(call REQ_AS_SOURCES,%.s.o)
	$(eval CURRENT_DIR := $(dir $@))
	$(eval AS_SOURCES  := $(shell ls $(CURRENT_DIR)*.s))
	mkdir -p $(BUILD_DIRECTORY)$(CURRENT_DIR)
	$(AS) $(ASFLAGS) $(AS_SOURCES) -o $(BUILD_DIRECTORY)$@

.SECONDEXPANSION:
REQ_GO_SOURCES = $(shell ls $(1)/*.go)
$(GO_OBJECTS): %.o : $$(call REQ_GO_SOURCES,%)
	$(eval CURRENT_DIR := $(subst .o,/,$@))
	$(eval GO_SOURCES  := $(shell ls $(CURRENT_DIR)*.go))
	mkdir -p $(BUILD_DIRECTORY)$(CURRENT_DIR)
	$(GCCGO) $(GCCGOFLAGS) -fgo-pkgpath=$(CURRENT_DIR:/=) -c $(GO_SOURCES) \
		-o $(BUILD_DIRECTORY)$@ -I$(BUILD_DIRECTORY) -L$(BUILD_DIRECTORY)

.SECONDEXPANSION:
$(DEPS): %.mk: $$(dir %)
	$(eval CURRENT_DIR := $(dir $@))
	./scripts/generate_go_deps.sh $(CURRENT_DIR)

run-bochs:
	bochs -f bochsrc.txt -q

run-qemu:
	$(QEMU) -cdrom os.iso $(QEMUFLAGS)

build-docker:
	docker build . -t littleosbook; \
	docker create --name littleosbook-container littleosbook; \
	docker cp littleosbook-container:/littleosbook/kernel.elf .; \
	docker cp littleosbook-container:/littleosbook/os.iso .; \
	docker rm littleosbook-container

clean:
	rm -rf $(AS_OBJECTS) $(BUILD_DIRECTORY) kernel.elf os.iso

clean-deps:
	find . -name deps.mk | xargs -n1 rm
