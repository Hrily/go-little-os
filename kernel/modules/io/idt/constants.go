package idt

const (
	// Source: https://wiki.osdev.org/Interrupt_Descriptor_Table

	// InterruptGate is used to specify an interrupt service routine
	InterruptGate GateType = 0x8e
	// TrapGate is similar to InterruptGate, but doesn't have interrupts disabled
	// during execution
	TrapGate GateType = 0x8f
)
