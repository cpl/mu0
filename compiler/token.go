package compiler

import (
	"log"
)

type token struct {
	t   tokenType
	arg string
}

type tokenType int

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
	tokenTypeORG         tokenType = 101
	tokenTypeEQU         tokenType = 102
	tokenTypeDEF         tokenType = 103

	_tokenTypeVMDirective tokenType = 1000
	tokenTypeBRK          tokenType = 1001
	tokenTypeSLP          tokenType = 1002
	tokenTypeSTP          tokenType = 1003
)

var tokenTypeMap = map[string]tokenType{
	// Instructions
	"LDA": tokenTypeLDA,
	"STA": tokenTypeSTA,
	"ADD": tokenTypeADD,
	"SUB": tokenTypeSUB,
	"JMP": tokenTypeJMP,
	"JGE": tokenTypeJGE,
	"JNE": tokenTypeJNE,
	"STP": tokenTypeSTP,

	// Compiler directives
	"ORG":   tokenTypeORG,
	"ORIG":  tokenTypeORG,
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
	log.Println(parts)

	// If token is a label
	if _, ok := tokenTypeMap[parts[0]]; !ok {
		// Store in memory map
		labels[parts[0]] = address

		// Token is label only
		if len(parts) == 1 {
			return nil
		}

		// Label has token attached
		tkn.t = tokenTypeMap[parts[1]]
		if tkn.t != tokenTypeSTP {
			tkn.arg = parts[2]
		}
	} else {
		// Normal token
		tkn.t = tokenTypeMap[parts[0]]
		if tkn.t != tokenTypeSTP {
			tkn.arg = parts[1]
		}
	}

	address++
	return tkn
}
