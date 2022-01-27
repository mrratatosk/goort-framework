package processor

import (
	"sync"

	"github.com/mrratatosk/oort-framework/register"
	"github.com/mrratatosk/oort-framework/tools"
)

type ITicker interface {
	Tick(*sync.WaitGroup)
	ClockDivider() uint8
}

type ProcessorUnit[A tools.TSize, D tools.TSize] struct {
	InstructionSet map[uint]Instruction[A, D]
	ExtensionSet   map[uint]map[uint]Instruction[A, D]
	Registers      register.RegisterList
	ITicker
}
