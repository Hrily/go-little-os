package paging

// LoadPDT loads pdt. Defined in load.s
// pdtAddr is the physical address of pdt.
func LoadPDT(pdtAddr uint32)

// InvalidateTLB invalidates  Translation Lookaside Buffer (TLB). Defined in
// load.s
func InvalidateTLB()
