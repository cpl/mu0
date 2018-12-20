package main

import (
	"log"

	"github.com/thee-engineer/mu0-vm/compiler"
	"github.com/thee-engineer/mu0-vm/vm"
)

func main() {
	compiler.Compile("./source.asm", "./source.out")
	myVM := vm.New()
	myVM.LoadFile("./source.out")
	myVM.Run()

	log.Println("ACC", myVM.ACC, "PC", myVM.PC)
}
