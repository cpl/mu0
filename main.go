package main

import "io/ioutil"

func main() {
	data, err := ioutil.ReadFile("source.mu0")
	if err != nil {
		panic(err)
	}

	vm := NewVM()
	vm.Load(data)
	vm.MemoryDump()
}
