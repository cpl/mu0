/*
MIT License

Copyright (c) 2018-2019 Alexandru-Paul Copil

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package compiler

import (
	"log"
	"os"
	"time"

	"github.com/thee-engineer/mu0/builtin"
)

// Compile ...
func Compile(inFile, outFile string) []builtin.Word {
	log.Printf("Compile: %s -> %s\n", inFile, outFile)
	t := time.Now()

	// Lexical analysis of source code
	tree := lex(inFile)

	// Compiled binary for builtin
	var binary []builtin.Word
	var instruction builtin.Word

	// Iterate lexical tree
	for _, tkn := range tree {
		// Skip EQU
		if tkn.t == tokenTypeEQU {
			continue
		}

		// If define, create word, else parse token
		if tkn.t == tokenTypeDEF {
			// If link to label, else word define
			if addr, ok := labels[tkn.arg]; ok {
				instruction = builtin.Word(addr)
			} else {
				instruction = parseArg(tkn, tree)
			}
		} else {
			// Extract instruction op code
			instruction = tokenTypeToOPC[tkn.t]

			// Extract instruction argument
			instruction |= parseArg(tkn, tree)
		}

		binary = append(binary, instruction)
	}

	// Create output file
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Create byte array
	byteStream := make([]byte, len(binary)*2)

	// Move binary words to byte array
	for idx, w := range binary {
		// ! DEBUG
		// fmt.Printf("%X %s\n", idx, decompileInstruction(w))

		byteStream[idx*2] = byte(w >> 8)
		byteStream[idx*2+1] = byte(w & 0x00FF)
	}

	// Write binary data to file
	_, err = f.Write(byteStream)
	if err != nil {
		log.Fatalln(err)
	}

	// Log details
	log.Printf("Compile: finished OK, in %d ns\n",
		time.Now().Nanosecond()-t.Nanosecond())
	log.Printf("Compile: wrote %d bytes to %s\n", len(byteStream), outFile)

	// Return binary as word array
	return binary
}
