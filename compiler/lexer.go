package compiler

import (
	"io/ioutil"
	"log"
	"time"
)

var _line = 0
var labels = make(map[string]int)
var address = 0

// lex opens a source file for reading and creates the token tree
func lex(filePath string) (tokenList []*token) {
	log.Println("Lex:", filePath)
	t := time.Now()

	// Read contents of source file
	stream, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Iterate file contents
	partCount := 0
	partsList := make([]string, 3)
	for idx := 0; idx < len(stream); {
		// Parse comments
		if stream[idx] == ';' {
			eatComment(&idx, stream)
			continue
		}

		// Parse spaces
		if isSpace(stream[idx]) {
			// On newline, parse token parts
			if eatSpaces(&idx, stream) {

				// Append new token if not nil
				nt := newToken(partsList[:partCount])
				if nt != nil {
					tokenList = append(tokenList, nt)
				}

				// Reset part counter
				partCount = 0
			}
			continue
		}

		// Parse tokens
		if isTokenChar(stream[idx]) {
			str := eatTokenPart(&idx, stream)
			partsList[partCount] = str
			partCount++
			continue
		}

		// Skip
		idx++
	}

	log.Printf("Lex: finished OK, parsed %d lines in %d ns\n", _line,
		(time.Now().Nanosecond() - t.Nanosecond()))
	log.Printf("Lex: generated %d tokens, using %d bytes\n", len(tokenList), address*2)
	log.Printf("Lex: generated %d labels\n", len(labels))

	return tokenList
}
