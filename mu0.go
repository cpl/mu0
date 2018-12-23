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

package main

import (
	"fmt"
	"os"

	"github.com/thee-engineer/mu0/compiler"
	"github.com/thee-engineer/mu0/module"
	"github.com/thee-engineer/mu0/mu0"
	"github.com/thee-engineer/mu0/vm"
)

const usage = "mu0 <run <source.o> | build <source.s> [source.o]>"

func main() {
	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		if len(os.Args) == 4 {
			compiler.Compile(os.Args[2], os.Args[3])
		}

		compiler.Compile(os.Args[2], os.Args[2]+".o")

		os.Exit(0)
	case "run":
		_vm := vm.New()

		// ! DEBUG
		_vm.AddModule(module.NewDummy([]mu0.Word{0x100, 0x102, 0x104}))

		_vm.LoadFile(os.Args[2])
		_vm.Run()

		// ! DEBUG
		fmt.Println(
			"ACC", _vm.ACC, "PC", _vm.PC,
			"EXIT CODE", _vm.StopCode, "EXEC COUNT", _vm.CountExec)

		// ! DEBUG
		_vm.MemoryDump(0x200)

		os.Exit(0)
	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
