package module

import (
	"github.com/thee-engineer/mu0/builtin"
)

// Dummy is a test module
type Dummy struct {
	module
}

// Handle sets the device registers to 0xF
func (m Dummy) Handle(mm *[0xFFFF]builtin.Word) {
	m.locked = true
	for _, deviceRegister := range m.deviceRegisters {
		mm[deviceRegister] = 0xF
	}
	m.locked = false
}

// NewDummy returns a dummy with device registers to value set
func NewDummy(registers []builtin.Word) Module {
	return Dummy{
		module{
			deviceRegisters: registers,
		}}
}
