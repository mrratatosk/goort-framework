package register

import "github.com/mrratatosk/oort-framework/tools"

func newRegister8() *Register[uint8] {
	return &Register[uint8]{Value: 0x0}
}

func (rl *RegisterList) AddRegister8(name string) {
	rl.r8[name] = newRegister8()
}

func (rl *RegisterList) Get8(name string) *Register[uint8] {
	return rl.r8[name]
}

func (rl *RegisterList) Set8(name string, value uint8) {
	rl.r8[name].Value = value
}

func (rl *RegisterList) Combine8(name1 string, name2 string) *Register[uint16] {
	r16 := newRegister16()

	r16.Value = tools.Combine8(rl.r8[name1].Value, rl.r8[name2].Value)

	return r16
}

func (rl *RegisterList) AddR8Val(name string, value uint8) (bool, bool, bool) {
	r, hc, c, z := tools.Add8(rl.r8[name].Value, value)

	rl.r8[name].Value = r
	return hc, c, z
}

func (rl *RegisterList) AddR8R8(name1 string, name2 string) (bool, bool, bool) {
	return rl.AddR8Val(name1, rl.r8[name2].Value)
}

func (rl *RegisterList) SubR8Val(name string, value uint8) (bool, bool, bool) {
	r, hc, c, z := tools.Sub8(rl.r8[name].Value, value)

	rl.r8[name].Value = r
	return hc, c, z
}

func (rl *RegisterList) SubR8R8(name1 string, name2 string) (bool, bool, bool) {
	return rl.SubR8Val(name1, rl.r8[name2].Value)
}

func (rl *RegisterList) AndR8Val(name string, value uint8) (bool, bool, bool) {
	r, z := tools.And8(rl.r8[name].Value, value)

	rl.r8[name].Value = r
	return true, false, z
}

func (rl *RegisterList) AndR8R8(name1 string, name2 string) (bool, bool, bool) {
	return rl.AndR8Val(name1, rl.r8[name2].Value)
}

func (rl *RegisterList) OrR8Val(name string, value uint8) (bool, bool, bool) {
	r, z := tools.Or8(rl.r8[name].Value, value)

	rl.r8[name].Value = r
	return false, false, z
}

func (rl *RegisterList) OrR8R8(name1 string, name2 string) (bool, bool, bool) {
	return rl.OrR8Val(name1, rl.r8[name2].Value)
}

func (rl *RegisterList) XorR8Val(name string, value uint8) (bool, bool, bool) {
	r, z := tools.Xor8(rl.r8[name].Value, value)

	rl.r8[name].Value = r
	return false, false, z
}

func (rl *RegisterList) XorR8R8(name1 string, name2 string) (bool, bool, bool) {
	return rl.XorR8Val(name1, rl.r8[name2].Value)
}

func (rl *RegisterList) NotR8(name string) {
	rl.r8[name].Value = tools.Not8(rl.r8[name].Value)
}

func (rl *RegisterList) CpR8Val(name string, value uint8) (bool, bool, bool) {
	_, hc, c, z := tools.Sub8(rl.r8[name].Value, value)
	return hc, c, z
}

func (rl *RegisterList) CpR8R8(name1 string, name2 string) (bool, bool, bool) {
	return rl.SubR8Val(name1, rl.r8[name2].Value)
}

func (rl *RegisterList) SwapR8(name string) bool {
	r := tools.Swap8(rl.r8[name].Value)
	rl.r8[name].Value = r
	return r == 0
}

func (rl *RegisterList) ShiftLR8(name string) (bool, bool) {
	r, msb := tools.ShiftL8(rl.r8[name].Value, 1)
	rl.r8[name].Value = r
	return msb, r == 0
}

func (rl *RegisterList) ShiftRR8(name string) (bool, bool) {
	r, msb := tools.ShiftR8(rl.r8[name].Value, 1)
	rl.r8[name].Value = r
	return msb, r == 0
}

func (rl *RegisterList) RotateLR8(name string) (bool, bool) {
	r, c := tools.RotateL8(rl.r8[name].Value, 1)
	rl.r8[name].Value = r
	return c, r == 0
}

func (rl *RegisterList) RotateRR8(name string) (bool, bool) {
	r, c := tools.RotateR8(rl.r8[name].Value, 1)
	rl.r8[name].Value = r
	return c, r == 0
}
