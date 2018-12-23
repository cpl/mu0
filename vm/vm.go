package vm

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/thee-engineer/mu0/mu0"
)

// VM attempts to simulate the components found on the UoM MU0 boards
type VM struct {
	isRunning  bool // State of the VM (running)
	isSleeping bool // State of the VM (sleeping)
	isInBreak  bool // State of the VM (break, wait for input)

	ACC      mu0.Word         // Accumulator (main register)
	PC       mu0.Word         // Program Counter
	Memory   [0xFFFF]mu0.Word // Physical memory space
	StopCode mu0.Word         // Exit code / Stop code
}

// New create a virtual machine
func New() *VM {
	return new(VM)
}

// Load a compiled program into memory
func (v *VM) Load(data []byte, start int) {
	for index := start; index < len(data) && index/2 < cap(v.Memory); index += 2 {
		v.Memory[index/2] = mu0.Word(data[index])<<8 | mu0.Word(data[index+1])
	}
}

// LoadFile takes a file path and loads all binary data from it
func (v *VM) LoadFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	v.Load(data, 0)
}

// Stop VM execution
func (v *VM) Stop(code mu0.Word) {
	v.StopCode = code
	v.isRunning = false
}

// Run starts OP execution from the ORIG address
func (v *VM) Run() {
	v.isRunning = true

	var instruction mu0.Word // Current instruction
	var opc mu0.Word         // Operation code
	var arg mu0.Word         // Operation arg

	for v.isRunning {
		// Check PC in memory range
		if int(v.PC) > len(v.Memory)-1 {
			log.Println("VM: PC out of memory address space")
			v.Stop(400)
			return
		}

		instruction = v.Memory[v.PC]    // Load instruction from memory (PC)
		opc = instruction & mu0.OpcMask // Extract operation code
		arg = instruction & mu0.ArgMask // Extract operation arg

		v.PC++ // Increment PC

		// Check which instruction to execute and how
		switch opc {
		case mu0.OpLDA:
			v.ACC = v.Memory[arg]
			break
		case mu0.OpSTA:
			v.Memory[arg] = v.ACC
			break
		case mu0.OpADD:
			v.ACC += v.Memory[arg]
			break
		case mu0.OpSUB:
			v.ACC -= v.Memory[arg]
			break
		case mu0.OpJMP:
			v.PC = arg
			break
		case mu0.OpJGE:
			if v.ACC >= 0 {
				v.PC = arg
			}
			break
		case mu0.OpJNE:
			if v.ACC != 0 {
				v.PC = arg
			}
			break
		case mu0.OpSTP:
			v.Stop(arg)
			break
		case mu0.OpSLP:
			// Convert argument word to duration string then duration
			d, err := time.ParseDuration(fmt.Sprintf("%dms", arg))
			if err != nil {
				log.Fatalln(err)
			}

			// Set sleeping state and sleep
			v.isSleeping = true
			time.Sleep(d)
			v.isSleeping = false

			break
		default:
			log.Fatalf("%04x %04x\n", opc, arg)
		}
	}
}

// MemoryDump writes the memory contents to stdout
func (v *VM) MemoryDump() {
	for index := 0; index+8 < cap(v.Memory); index += 8 {
		fmt.Printf(
			"%04x : %04x %04x %04x %04x %04x %04x %04x %04x\n", index,
			v.Memory[index],
			v.Memory[index+1],
			v.Memory[index+2],
			v.Memory[index+3],
			v.Memory[index+4],
			v.Memory[index+5],
			v.Memory[index+6],
			v.Memory[index+7])
	}
}
