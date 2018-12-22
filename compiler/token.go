package compiler

import (
	"log"

	"github.com/thee-engineer/mu0/mu0"
)

type token struct {
	t   tokenType
	arg string
}

type tokenType int

var tokenTypeToOPC = map[tokenType]mu0.Word{
	tokenTypeLDA: mu0.OpLDA,
	tokenTypeSTA: mu0.OpSTA,
	tokenTypeADD: mu0.OpADD,
	tokenTypeSUB: mu0.OpSUB,
	tokenTypeJMP: mu0.OpJMP,
	tokenTypeJGE: mu0.OpJGE,
	tokenTypeJNE: mu0.OpJNE,
	tokenTypeBRK: mu0.OpBRK,
	tokenTypeSLP: mu0.OpSLP,
	tokenTypeSTP: mu0.OpSTP,
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
	"EQU":   tokenTypeEQU,
	"CONST": tokenTypeEQU,
	"DEF":   tokenTypeDEF,
	"DEFW":  tokenTypeDEF,

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
