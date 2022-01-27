package oortframework

import (
	"github.com/mrratatosk/oort-framework/memory"
	"github.com/mrratatosk/oort-framework/processor"
	"github.com/mrratatosk/oort-framework/tools"
)

type IEmulator interface {
	Start()
}

type Emulator[A tools.TSize, D tools.TSize] struct {
	Memory *memory.Memory[A, D]
	Units  []processor.ITicker
	Bios   string
	IEmulator
}
