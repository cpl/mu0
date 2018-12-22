package compiler

import (
	"fmt"

	"github.com/thee-engineer/mu0/mu0"
)

func decompileInstruction(w mu0.Word) string {
	var ins string

	switch w & mu0.OpcMask {
	case mu0.OpLDA:
		ins = "LDA"
		break
	case mu0.OpSTA:
		ins = "STA"
		break
	case mu0.OpADD:
		ins = "ADD"
		break
	case mu0.OpSUB:
		ins = "SUB"
		break
	case mu0.OpJMP:
		ins = "JMP"
		break
	case mu0.OpJGE:
		ins = "JGE"
		break
	case mu0.OpJNE:
		ins = "JNE"
		break
	case mu0.OpBRK:
		ins = "BRK"
		break
	case mu0.OpSTP:
		ins = "STP"
		break
	case mu0.OpSLP:
		ins = "SLP"
		break
	default:
		ins = "UNDEFINED"
	}

	return fmt.Sprintf("%s %X", ins, w&mu0.ArgMask)
}
