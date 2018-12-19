package compiler

func isSpace(c byte) bool {
	if c == 0x20 || c == 0x9 {
		return true
	}

	return c >= 0xA && c <= 0xD
}

func isTokenChar(c byte) bool {
	if c >= 0x30 && c <= 0x39 {
		return true
	} else if c >= 0x41 && c <= 0x5A {
		return true
	} else if c == 0x26 {
		return true
	}

	return c >= 0x61 && c <= 0x7A
}

func toUpper(c byte) byte {
	if c >= 0x61 && c <= 0x7A {
		return c - 0x20
	}

	return c
}
