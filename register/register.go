package register

import "github.com/mrratatosk/oort-framework/tools"

type IRegister[T tools.TSize] interface {
	Set(value T)
	Get() T
}

type Register[T tools.TSize] struct {
	Value T
	IRegister[T]
}

func (r *Register[T]) Set(value T) {
	r.Value = value
}

func (r *Register[T]) Get() T {
	return r.Value
}

func (r *Register[T]) Inc() *Register[T] {
	r.Value++
	return r
}

func (r *Register[T]) Add(value T) *Register[T] {
	r.Value += value
	return r
}

func (r *Register[T]) BitSet(pos T) *Register[T] {
	r.Value |= (1 << pos)
	return r
}

func (r *Register[T]) Dec() *Register[T] {
	r.Value--
	return r
}

func (r *Register[T]) Sub(value T) *Register[T] {
	r.Value -= value
	return r
}

func (r *Register[T]) BitClear(pos T) *Register[T] {
	r.Value &= ^(1 << pos)
	return r
}

func (r *Register[T]) Bit(pos T) bool {
	return (r.Value & (1 << pos)) > 0
}

type RegisterList struct {
	r8  map[string]*Register[uint8]
	r16 map[string]*Register[uint16]
}

func NewRegisterList() RegisterList {
	return RegisterList{
		r8:  make(map[string]*Register[uint8]),
		r16: make(map[string]*Register[uint16]),
	}
}
