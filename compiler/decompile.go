package compiler

import (
	"fmt"

	"github.com/thee-engineer/mu0/builtin"
)

// DecompileIns takes a word and returns the string representation of the ins
func DecompileIns(w builtin.Word) string {
	var ins string

	switch w & builtin.OpcMask {
	case builtin.OpLDA:
		ins = "LDA"
		break
	case builtin.OpSTA:
		ins = "STA"
		break
	case builtin.OpADD:
		ins = "ADD"
		break
	case builtin.OpSUB:
		ins = "SUB"
		break
	case builtin.OpJMP:
		ins = "JMP"
		break
	case builtin.OpJGE:
		ins = "JGE"
		break
	case builtin.OpJNE:
		ins = "JNE"
		break
	case builtin.OpBRK:
		ins = "BRK"
		break
	case builtin.OpSTP:
		ins = "STP"
		break
	case builtin.OpSLP:
		ins = "SLP"
		break
	default:
		ins = "UNDEFINED"
	}

	return fmt.Sprintf("%s %X", ins, w&builtin.ArgMask)
}
