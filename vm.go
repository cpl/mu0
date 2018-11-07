package main

import (
	"fmt"
)

type word uint16

const (
	opcSize  word = 4
	opcShift word = 12
	opcMask  word = ((1 << opcSize) - 1) << opcShift
)

const (
	argSize word = 12
	argMask word = (1 << argSize) - 1
)

const (
	opLDA word = ((0 << opcShift) & opcMask)
	opSTA word = ((1 << opcShift) & opcMask)
	opADD word = ((2 << opcShift) & opcMask)
	opSUB word = ((3 << opcShift) & opcMask)
	opJMP word = ((4 << opcShift) & opcMask)
	opJGE word = ((5 << opcShift) & opcMask)
	opJNE word = ((6 << opcShift) & opcMask)
	opSTP word = ((7 << opcShift) & opcMask)
)

// VM ...
type VM struct {
	running bool

	ACC    word
	PC     word
	Memory [0xFFFF]word
}

// NewVM ...
func NewVM() *VM {
	return &VM{
		ACC:    0,
		PC:     0,
		Memory: [0xFFFF]word{0},
	}
}

// Load ...
func (vm *VM) Load(data []byte) {
	for index := 0; index+2 < len(data) && index/2 < cap(vm.Memory); index += 2 {
		vm.Memory[index/2] = word(data[index])<<8 | word(data[index+1])
	}
}

// Stop ...
func (vm *VM) Stop() {
	vm.running = false
}

// Run ...
func (vm *VM) Run() {
	vm.running = true

	var instruction word
	var opc word
	var arg word

	for vm.running {
		instruction = vm.Memory[vm.PC]
		opc = instruction & opcMask
		arg = instruction & argMask

		vm.PC++

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
			fmt.Printf("%04x %04x\n", opc, arg)
		}
	}
}

// MemoryDump ...
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
