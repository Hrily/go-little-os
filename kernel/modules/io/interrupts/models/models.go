package models

// Registers contains the state of registers before the interrupt happens
type Registers struct {
	EDI, ESI, EBP, ESP, EBX, EDX, ECX, EAX int32 /* pushed by 'pusha' */
	GS, FS, ES, DS                         int32 /* pushed the segs last */
	IntNumber, ErrCode                     int32 /* our 'push byte #' and ecodes do this */
	EIP, CS, EFlags, UserESP, SS           int32 /* pushed by the processor automatically */
}
