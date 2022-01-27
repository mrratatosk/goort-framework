package memory

import "github.com/mrratatosk/oort-framework/tools"

type Memory[A tools.TSize, D tools.TSize] struct {
	data []D
}

func NewMemory[A tools.TSize, D tools.TSize](size uint) *Memory[A, D] {
	return &Memory[A, D]{make([]D, size)}
}

func (m Memory[A, D]) Read(address A) D {
	return m.data[address]
}

func (m Memory[A, D]) ReadRange(address A, size uint) []D {
	return m.data[address:(address + A(size))]
}

func (m Memory[A, D]) Write(address A, value D) {
	m.data[address] = value
}

func (m Memory[A, D]) WriteRange(address A, value []D) {
	for i, d := range value {
		m.data[address+A(i)] = d
	}
}
