package compiler

import (
	"io/ioutil"
	"log"
	"time"
)

var _line = 0

type token struct {
	t tokenType
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
	tokenTypeSTP          tokenType = 18

	_tokenTypeCDirective tokenType = 100
	tokenTypeORG         tokenType = 101
	tokenTypeEQU         tokenType = 102
	tokenTypeDEF         tokenType = 103

	_tokenTypeVMDirective tokenType = 1000
	tokenTypeBRK          tokenType = 1001
	tokenTypeSLP          tokenType = 1002
)

var labels map[string]int

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

// Lex opens a source file for reading and creates the token tree
func Lex(filePath string) {
	log.Println("Lex:", filePath)
	t := time.Now()

	// Read contents of source file
	stream, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Iterate file contents
	partCount := 0
	for idx := 0; idx < len(stream); {

		// Parse comments
		if stream[idx] == ';' {
			eatComment(&idx, stream)
			continue
		}

		// Parse spaces
		if isSpace(stream[idx]) {
			if eatSpaces(&idx, stream) {
				partCount = 0
				print("\n")
			}
			continue
		}

		// Parse tokens
		if isTokenChar(stream[idx]) {
			str := eatTokenPart(&idx, stream)
			if tokenTypeMap[str] == 0 {
				if partCount == 0 {
					print("lbl:", str)
				} else {
					print("arg:", str)
				}
			} else {
				print("tkn:", str)
			}
			print(" ")

			partCount++
			continue
		}

		// Skip
		idx++
	}

	log.Printf("Lex: finished OK, parsed %d lines in %d ns\n", _line,
		(time.Now().Nanosecond() - t.Nanosecond()))
}
