package compiler

import (
	"fmt"
)

// Compile ...
func Compile(filePath string) {
	tree := lex(filePath)

	for idx, tkn := range tree {

		// Treat defines as single word
		if tkn.t == tokenTypeDEF {
			fmt.Println(idx, "DEF")
			continue
		}

		fmt.Printf("%4d %4d %4d\n", idx, tkn.t, parseArg(tkn, tree))
	}
}

func parseArg(tkn *token, tree []*token) int {

	// Labels and constants
	if addr, ok := labels[tkn.arg]; ok {
		// Expand labels to constants
		if tree[addr].t == tokenTypeEQU {
			return -10 // TODO, parse tree[addr].arg as value
		}

		// Expand labels to addresses
		return addr
	}

	// Parse values
	return -3 // TODO, parse tkn.arg as value
}
