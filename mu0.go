package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thee-engineer/mu0/compiler"
	"github.com/thee-engineer/mu0/vm"
)

const usage = "mu0 <run <source.o> | build <source.s> [source.o]>"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		if len(os.Args) == 4 {
			compiler.Compile(os.Args[2], os.Args[3])
		}

		os.Exit(0)
	case "run":
	default:
		fmt.Println(usage)
		os.Exit(1)
	}

	myVM := vm.New()
	myVM.LoadFile("./source.out")
	myVM.Run()

	log.Println("ACC", myVM.ACC, "PC", myVM.PC, "SC", myVM.StopCode)
}
