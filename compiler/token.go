package compiler

import (
	"log"

	"github.com/thee-engineer/mu0/builtin"
)

type token struct {
	t   tokenType
	arg string
}

type tokenType int

var tokenTypeToOPC = map[tokenType]builtin.Word{
	tokenTypeLDA: builtin.OpLDA,
	tokenTypeSTA: builtin.OpSTA,
	tokenTypeADD: builtin.OpADD,
	tokenTypeSUB: builtin.OpSUB,
	tokenTypeJMP: builtin.OpJMP,
	tokenTypeJGE: builtin.OpJGE,
	tokenTypeJNE: builtin.OpJNE,
	tokenTypeBRK: builtin.OpBRK,
	tokenTypeSLP: builtin.OpSLP,
	tokenTypeSTP: builtin.OpSTP,
}

const (
	_tokenTypeInstruction tokenType = 10
	tokenTypeLDA          tokenType = 11
	tokenTypeSTA          tokenType = 12
	tokenTypeADD          tokenType = 13
	tokenTypeSUB          tokenType = 14
	tokenTypeJMP          tokenType = 15
	tokenTypeJGE          tokenType = 16
	tokenTypeJNE          tokenType = 17

	_tokenTypeCDirective tokenType = 100
	tokenTypeEQU         tokenType = 101
	tokenTypeDEF         tokenType = 102
	tokenTypeINC         tokenType = 103

	_tokenTypeVMDirective tokenType = 1000
	tokenTypeBRK          tokenType = 1001
	tokenTypeSLP          tokenType = 1002
	tokenTypeSTP          tokenType = 1003
)

var tokenTypeMap = map[string]tokenType{
	// Instructions
	"LDA":  tokenTypeLDA,
	"STA":  tokenTypeSTA,
	"ADD":  tokenTypeADD,
	"SUB":  tokenTypeSUB,
	"JMP":  tokenTypeJMP,
	"B":    tokenTypeJMP,
	"JGE":  tokenTypeJGE,
	"BGE":  tokenTypeJGE,
	"JNE":  tokenTypeJNE,
	"BNE":  tokenTypeJNE,
	"STP":  tokenTypeSTP,
	"STOP": tokenTypeSTP,

	// Compiler directives
	"DEF":     tokenTypeDEF,
	"DEFW":    tokenTypeDEF,
	"EQU":     tokenTypeEQU,
	"CONST":   tokenTypeEQU,
	"INCLUDE": tokenTypeINC,
	"INC":     tokenTypeINC,

	// VM directives
	"BRK":   tokenTypeBRK,
	"BREAK": tokenTypeBRK,
	"SLP":   tokenTypeSLP,
	"SLEEP": tokenTypeSLP,
}

func newToken(parts []string) *token {
	// Ignore empty tokens
	if len(parts) == 0 {
		return nil
	}
	if len(parts) > 3 {
		log.Fatalln("invalid token length,", parts)
	}

	// Create new token
	tkn := new(token)

	// If token is a label
	if _, ok := tokenTypeMap[parts[0]]; !ok {
		// Store in memory map
		labels[parts[0]] = address

		// Token is label only
		if len(parts) == 1 {
			return nil
		}

		// Label has token attached
		tkn.t, ok = tokenTypeMap[parts[1]]
		if !ok {
			log.Fatalf("new token: unexpected instructions, %s at line %d\n",
				parts[1], _line)
		}
		tkn.arg = parts[2]
	} else {
		if len(parts) != 2 {
			log.Fatalf("new token: not enough parts %s at line %d\n",
				parts, _line)
		}

		// Normal token
		tkn.t = tokenTypeMap[parts[0]]
		tkn.arg = parts[1]
	}

	address++
	return tkn
}
