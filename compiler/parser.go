package compiler

import (
	"fmt"
	"log"
	"strings"

	"github.com/thee-engineer/mu0/builtin"
)

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

func parseArg(tkn *token, tree []*token) builtin.Word {
	// Labels and constants
	if addr, ok := labels[tkn.arg]; ok {
		// Expand labels to constants
		if tree[addr].t == tokenTypeEQU {
			return valueToWord(tree[addr].arg)
		}

		// Expand labels to addresses
		return builtin.Word(addr)
	}

	// Parse values
	return valueToWord(tkn.arg)
}

func valueToWord(val string) builtin.Word {
	var valInt int

	if strings.HasPrefix(val, "0B") {
		// Parse binary
		_, err := fmt.Sscanf(val, "0B%b", &valInt)
		if err != nil {
			log.Fatalln("failed to parse value", val)
		}
		return builtin.Word(valInt)
	} else if strings.HasPrefix(val, "&") {
		// Parse hex
		_, err := fmt.Sscanf(val, "&%X", &valInt)
		if err != nil {
			log.Fatalln("failed to parse value", val)
		}
		return builtin.Word(valInt)
	} else if strings.HasPrefix(val, "0X") {
		// Parse hex alternative
		_, err := fmt.Sscanf(val, "0X%X", &valInt)
		if err != nil {
			log.Fatalln("failed to parse value", val)
		}
		return builtin.Word(valInt)
	}

	// Parse decimal
	_, err := fmt.Sscanf(val, "%d", &valInt)
	if err != nil {
		log.Fatalln("failed to parse value", val)
	}
	return builtin.Word(valInt)
}
