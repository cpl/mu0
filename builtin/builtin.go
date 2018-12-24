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

package builtin

// Word is 2 byte long on the MU0
type Word uint16

// Number of bits in a Word
const wordSizeB = 16

// Operation Code defined constants
const (
	OpcSizeB Word = 4                                 // OP code is 4 bits
	OpcShift Word = 12                                // OP code is bits 12:15
	OpcMask  Word = ((1 << OpcSizeB) - 1) << OpcShift // OP code mask 11110...
)

// Remaining bits are used as "arguments" for the OP
const (
	ArgSize Word = wordSizeB - OpcSizeB //
	ArgMask Word = (1 << ArgSize) - 1   //
)

const (
	OpLDA     Word = ((0 << OpcShift) & OpcMask)  // 0x0000 | LDA v
	OpSTA     Word = ((1 << OpcShift) & OpcMask)  // 0x1000 | STA v
	OpADD     Word = ((2 << OpcShift) & OpcMask)  // 0x2000 | ADD v
	OpSUB     Word = ((3 << OpcShift) & OpcMask)  // 0x3000 | SUB v
	OpJMP     Word = ((4 << OpcShift) & OpcMask)  // 0x4000 | JMP v
	OpJGE     Word = ((5 << OpcShift) & OpcMask)  // 0x5000 | JGE v
	OpJNE     Word = ((6 << OpcShift) & OpcMask)  // 0x6000 | JNE v
	OpSTP     Word = ((7 << OpcShift) & OpcMask)  // 0x7000 | STP
	OpBRK     Word = ((8 << OpcShift) & OpcMask)  // 0x8000 | BRK
	OpSLP     Word = ((9 << OpcShift) & OpcMask)  // 0x9000 | SLP ms
	OpJML     Word = ((10 << OpcShift) & OpcMask) // 0xA000 | JML v
	OpRET     Word = ((11 << OpcShift) & OpcMask) // 0xB000 | RET
	OpUNDEF12 Word = ((12 << OpcShift) & OpcMask) // 0xC000 |
	OpUNDEF13 Word = ((13 << OpcShift) & OpcMask) // 0xD000 |
	OpUNDEF14 Word = ((14 << OpcShift) & OpcMask) // 0xE000 |
	OpUNDEF15 Word = ((15 << OpcShift) & OpcMask) // 0xF000 |
)
