package vm

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/thee-engineer/mu0-vm/mu0"
)

// VM attempts to simulate the components found on the UoM MU0 boards
type VM struct {
	running bool // State of the VM

	ACC      mu0.Word         // Accumulator (main register)
	PC       mu0.Word         // Program Counter
	Memory   [0xFFFF]mu0.Word // Physical memory space
	StopCode mu0.Word         // Exit code / Stop code
}

// New create a virtual machine
func New() VM {
	return VM{
		ACC:      0,
		PC:       0,
		Memory:   [0xFFFF]mu0.Word{},
		StopCode: 0,
	}
}

// Load a compiled program into memory
func (vm *VM) Load(data []byte, start int) {
	for index := start; index < len(data) && index/2 < cap(vm.Memory); index += 2 {
		vm.Memory[index/2] = mu0.Word(data[index])<<8 | mu0.Word(data[index+1])
	}
}

// LoadFile takes a file path and loads all binary data from it
func (vm *VM) LoadFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm.Load(data, 0)
}

// Stop VM execution
func (vm *VM) Stop(code mu0.Word) {
	vm.StopCode = code
	vm.running = false
}

// Run starts OP execution from the ORIG address
func (vm *VM) Run() {
	vm.running = true

	var instruction mu0.Word // Current instruction
	var opc mu0.Word         // Operation code
	var arg mu0.Word         // Operation arg

	for vm.running {
		// Check PC in memory range
		if int(vm.PC) > len(vm.Memory)-1 {
			log.Println("VM: PC out of memory address space")
			vm.Stop(400)
			return
		}

		instruction = vm.Memory[vm.PC]  // Load instruction from memory (PC)
		opc = instruction & mu0.OpcMask // Extract operation code
		arg = instruction & mu0.ArgMask // Extract operation arg

		vm.PC++ // Increment PC

		// Check which instruction to execute and how
		switch opc {
		case mu0.OpLDA:
			vm.ACC = vm.Memory[arg]
			break
		case mu0.OpSTA:
			vm.Memory[arg] = vm.ACC
			break
		case mu0.OpADD:
			vm.ACC += vm.Memory[arg]
			break
		case mu0.OpSUB:
			vm.ACC -= vm.Memory[arg]
			break
		case mu0.OpJMP:
			vm.PC = arg
			break
		case mu0.OpJGE:
			if vm.ACC >= 0 {
				vm.PC = arg
			}
			break
		case mu0.OpJNE:
			if vm.ACC != 0 {
				vm.PC = arg
			}
			break
		case mu0.OpSTP:
			vm.Stop(arg)
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
