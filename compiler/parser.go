package compiler

import "log"

func eatComment(idx *int, stream []byte) {
	for stream[*idx] != '\n' && *idx < (len(stream)-1) {
		*idx++
	}
}

func eatSpaces(idx *int, stream []byte) bool {
	hasNewline := false
	for isSpace(stream[*idx]) && *idx < (len(stream)-1) {
		// Check for newlines
		if stream[*idx] == '\n' {
			_line++
			hasNewline = true
		}

		*idx++

		// Check for comments
		if stream[*idx] == ';' {
			eatComment(idx, stream)
			hasNewline = true
		}
	}

	return hasNewline
}

func eatTokenPart(idx *int, stream []byte) string {
	var str string

	// Iterate consecutive runes
	for isTokenChar(stream[*idx]) && *idx < (len(stream)-1) {
		// Convert everything to uppercase
		str += string(toUpper(stream[*idx]))

		*idx++

		// Invalid comment
		if rune(stream[*idx]) == ';' {
			log.Fatalln("remove comment near token, line", _line)
		}
	}

	return str
}
