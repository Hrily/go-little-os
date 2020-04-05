package keyboard

var key0x02 = Key{Code: 0x02, Characters: [2]byte{'1', '!'}}
var key0x03 = Key{Code: 0x03, Characters: [2]byte{'2', '@'}}
var key0x04 = Key{Code: 0x04, Characters: [2]byte{'3', '#'}}
var key0x05 = Key{Code: 0x05, Characters: [2]byte{'4', '$'}}
var key0x06 = Key{Code: 0x06, Characters: [2]byte{'5', '%'}}
var key0x07 = Key{Code: 0x07, Characters: [2]byte{'6', '^'}}
var key0x08 = Key{Code: 0x08, Characters: [2]byte{'7', '&'}}
var key0x09 = Key{Code: 0x09, Characters: [2]byte{'8', '*'}}
var key0x0a = Key{Code: 0x0a, Characters: [2]byte{'9', '('}}
var key0x0b = Key{Code: 0x0b, Characters: [2]byte{'0', ')'}}
var key0x0c = Key{Code: 0x0c, Characters: [2]byte{'-', '_'}}
var key0x0d = Key{Code: 0x0d, Characters: [2]byte{'=', '+'}}
var key0x0e = Key{Code: 0x0e, Characters: [2]byte{'\b'}}
var key0x0f = Key{Code: 0x0f, Characters: [2]byte{'\t'}}
var key0x10 = Key{Code: 0x10, Characters: [2]byte{'q', 'Q'}}
var key0x11 = Key{Code: 0x11, Characters: [2]byte{'w', 'W'}}
var key0x12 = Key{Code: 0x12, Characters: [2]byte{'e', 'E'}}
var key0x13 = Key{Code: 0x13, Characters: [2]byte{'r', 'R'}}
var key0x14 = Key{Code: 0x14, Characters: [2]byte{'t', 'T'}}
var key0x15 = Key{Code: 0x15, Characters: [2]byte{'y', 'Y'}}
var key0x16 = Key{Code: 0x16, Characters: [2]byte{'u', 'U'}}
var key0x17 = Key{Code: 0x17, Characters: [2]byte{'i', 'I'}}
var key0x18 = Key{Code: 0x18, Characters: [2]byte{'o', 'O'}}
var key0x19 = Key{Code: 0x19, Characters: [2]byte{'p', 'P'}}
var key0x1a = Key{Code: 0x1a, Characters: [2]byte{'[', '{'}}
var key0x1b = Key{Code: 0x1b, Characters: [2]byte{']', '}'}}
var key0x1c = Key{Code: 0x1c, Characters: [2]byte{'\n'}}
var key0x1e = Key{Code: 0x1e, Characters: [2]byte{'a', 'A'}}
var key0x1f = Key{Code: 0x1f, Characters: [2]byte{'s', 'S'}}
var key0x20 = Key{Code: 0x20, Characters: [2]byte{'d', 'D'}}
var key0x21 = Key{Code: 0x21, Characters: [2]byte{'f', 'F'}}
var key0x22 = Key{Code: 0x22, Characters: [2]byte{'g', 'G'}}
var key0x23 = Key{Code: 0x23, Characters: [2]byte{'h', 'H'}}
var key0x24 = Key{Code: 0x24, Characters: [2]byte{'j', 'J'}}
var key0x25 = Key{Code: 0x25, Characters: [2]byte{'k', 'K'}}
var key0x26 = Key{Code: 0x26, Characters: [2]byte{'l', 'L'}}
var key0x27 = Key{Code: 0x27, Characters: [2]byte{';', ':'}}
var key0x28 = Key{Code: 0x28, Characters: [2]byte{'\'', '"'}}
var key0x29 = Key{Code: 0x29, Characters: [2]byte{'`', '~'}}
var key0x2b = Key{Code: 0x2b, Characters: [2]byte{'\\', '|'}}
var key0x2c = Key{Code: 0x2c, Characters: [2]byte{'z', 'Z'}}
var key0x2d = Key{Code: 0x2d, Characters: [2]byte{'x', 'X'}}
var key0x2e = Key{Code: 0x2e, Characters: [2]byte{'c', 'C'}}
var key0x2f = Key{Code: 0x2f, Characters: [2]byte{'v', 'V'}}
var key0x30 = Key{Code: 0x30, Characters: [2]byte{'b', 'B'}}
var key0x31 = Key{Code: 0x31, Characters: [2]byte{'n', 'N'}}
var key0x32 = Key{Code: 0x32, Characters: [2]byte{'m', 'M'}}
var key0x33 = Key{Code: 0x33, Characters: [2]byte{',', '<'}}
var key0x34 = Key{Code: 0x34, Characters: [2]byte{'.', '>'}}
var key0x35 = Key{Code: 0x35, Characters: [2]byte{'/', '?'}}
var key0x39 = Key{Code: 0x39, Characters: [2]byte{' '}}

var keys = [100]*Key{
	nil, nil,
	&key0x02,
	&key0x03,
	&key0x04,
	&key0x05,
	&key0x06,
	&key0x07,
	&key0x08,
	&key0x09,
	&key0x0a,
	&key0x0b,
	&key0x0c,
	&key0x0d,
	&key0x0e,
	&key0x0f,
	&key0x10,
	&key0x11,
	&key0x12,
	&key0x13,
	&key0x14,
	&key0x15,
	&key0x16,
	&key0x17,
	&key0x18,
	&key0x19,
	&key0x1a,
	&key0x1b,
	&key0x1c,
	nil,
	&key0x1e,
	&key0x1f,
	&key0x20,
	&key0x21,
	&key0x22,
	&key0x23,
	&key0x24,
	&key0x25,
	&key0x26,
	&key0x27,
	&key0x28,
	&key0x29,
	nil,
	&key0x2b,
	&key0x2c,
	&key0x2d,
	&key0x2e,
	&key0x2f,
	&key0x30,
	&key0x31,
	&key0x32,
	&key0x33,
	&key0x34,
	&key0x35,
	nil, nil, nil,
	&key0x39,
}
