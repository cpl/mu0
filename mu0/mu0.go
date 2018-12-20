package mu0

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
	OpLDA Word = ((0 << OpcShift) & OpcMask) // 0x0000 | LDA v
	OpSTA Word = ((1 << OpcShift) & OpcMask) // 0x1000 | STA v
	OpADD Word = ((2 << OpcShift) & OpcMask) // 0x2000 | ADD v
	OpSUB Word = ((3 << OpcShift) & OpcMask) // 0x3000 | SUB v
	OpJMP Word = ((4 << OpcShift) & OpcMask) // 0x4000 | JMP v
	OpJGE Word = ((5 << OpcShift) & OpcMask) // 0x5000 | JGE v
	OpJNE Word = ((6 << OpcShift) & OpcMask) // 0x6000 | JNE v
	OpSTP Word = ((7 << OpcShift) & OpcMask) // 0x7000 | STP
	OpBRK Word = ((8 << OpcShift) & OpcMask) // 0x8000 | BRK
	OpSLP Word = ((9 << OpcShift) & OpcMask) // 0x9000 | SLP ms
)
