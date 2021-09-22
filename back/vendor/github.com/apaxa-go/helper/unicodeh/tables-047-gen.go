package unicodeh

import "unicode"

// Unicode property "Hex_Digit" (known as "Hex", "Hex_Digit").
// Kind of property: "Binary".
// Based on file "PropList.txt".
var (
	HexDigitNo  = hexDigitNo  // Value "No" (known as "N", "No", "F", "False").
	HexDigitYes = hexDigitYes // Value "Yes" (known as "Y", "Yes", "T", "True").
)

var (
	hexDigitNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x2f, 0x1}, {0x3a, 0x40, 0x1}, {0x47, 0x60, 0x1}, {0x67, 0xff0f, 0x1}, {0xff1a, 0xff20, 0x1}, {0xff27, 0xff40, 0x1}, {0xff47, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x10ffff, 0x1}}, 3}
	hexDigitYes = &unicode.RangeTable{[]unicode.Range16{{0x30, 0x39, 0x1}, {0x41, 0x46, 0x1}, {0x61, 0x66, 0x1}, {0xff10, 0xff19, 0x1}, {0xff21, 0xff26, 0x1}, {0xff41, 0xff46, 0x1}}, nil, 3}
)