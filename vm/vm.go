package vm

import (
	"fmt"
	"log"
)

// A word is 2 byte long on the MU0
type word uint16

// Number of bits in a word
const wordSizeB = 16

// Operation Code defined constants
const (
	opcSizeB word = 4                                 // OP code is 4 bits
	opcShift word = 12                                // OP code is bits 12:15
	opcMask  word = ((1 << opcSizeB) - 1) << opcShift // OP code mask 11110...
)

// Remaining bits are used as "arguments" for the OP
const (
	argSize word = wordSizeB - opcSizeB //
	argMask word = (1 << argSize) - 1   //
)

const (
	opLDA word = ((0 << opcShift) & opcMask) // LDA s
	opSTA word = ((1 << opcShift) & opcMask) // STA s
	opADD word = ((2 << opcShift) & opcMask) // ADD s
	opSUB word = ((3 << opcShift) & opcMask) // SUB s
	opJMP word = ((4 << opcShift) & opcMask) // JMP s
	opJGE word = ((5 << opcShift) & opcMask) // JGE s
	opJNE word = ((6 << opcShift) & opcMask) // JNE s
	opSTP word = ((7 << opcShift) & opcMask) // STP
)

// VM attempts to simulate the components found on the UoM MU0 boards
type VM struct {
	running bool // State of the VM

	ACC    word         // Accumulator (main register)
	PC     word         // Program Counter
	Memory [0xFFFF]word // Physical memory space
}

// New create a virtual machine
func New() VM {
	return VM{
		ACC:    0,
		PC:     0,
		Memory: [0xFFFF]word{0},
	}
}

// Load a compiled program into memory
func (vm *VM) Load(data []byte, start int) {
	for index := start; index+2 < len(data) && index/2 < cap(vm.Memory); index += 2 {
		vm.Memory[index/2] = word(data[index])<<8 | word(data[index+1])
	}
}

// Stop VM execution
func (vm *VM) Stop() {
	vm.running = false
}

// Run starts OP execution from the ORIG address
func (vm *VM) Run() {
	vm.running = true

	var instruction word // Current instruction
	var opc word         // Operation code
	var arg word         // Operation arg

	for vm.running {
		instruction = vm.Memory[vm.PC] // Load instruction from memory (PC)
		opc = instruction & opcMask    // Extract operation code
		arg = instruction & argMask    // Extract operation arg

		vm.PC++ // Increment PC

		// Check which instruction to execute and how
		switch opc {
		case opLDA:
			vm.ACC = vm.Memory[arg]
			break
		case opSTA:
			vm.Memory[arg] = vm.ACC
			break
		case opADD:
			vm.ACC += vm.Memory[arg]
			break
		case opSUB:
			vm.ACC -= vm.Memory[arg]
			break
		case opJMP:
			vm.PC = arg
			break
		case opJGE:
			if vm.ACC >= 0 {
				vm.PC = arg
			}
			break
		case opJNE:
			if vm.ACC != 0 {
				vm.PC = arg
			}
			break
		case opSTP:
			vm.Stop()
			break
		default:
			log.Fatalf("%04x %04x\n", opc, arg)
		}
	}
}

// MemoryDump writes the memory contents to stdout
func (vm *VM) MemoryDump() {
	for index := 0; index+8 < cap(vm.Memory); index += 8 {
		fmt.Printf(
			"%04x : %04x %04x %04x %04x %04x %04x %04x %04x\n", index,
			vm.Memory[index],
			vm.Memory[index+1],
			vm.Memory[index+2],
			vm.Memory[index+3],
			vm.Memory[index+4],
			vm.Memory[index+5],
			vm.Memory[index+6],
			vm.Memory[index+7])
	}
}
