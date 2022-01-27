package processor

import (
	"github.com/mrratatosk/oort-framework/memory"
	"github.com/mrratatosk/oort-framework/tools"
)

type Instruction[A tools.TSize, D tools.TSize] struct {
	Cycle    uint
	Params   uint
	Callback func(ProcessorUnit[A, D], memory.Memory[A, D], ...D)
	Name     string
}
