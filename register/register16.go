package register

import "github.com/mrratatosk/oort-framework/tools"

func newRegister16() *Register[uint16] {
	return &Register[uint16]{Value: 0x0}
}

func (rl *RegisterList) AddRegister16(name string) {
	rl.r16[name] = newRegister16()
}

func (rl *RegisterList) Get16(name string) *Register[uint16] {
	return rl.r16[name]
}

func (rl *RegisterList) Set16(name string, value uint16) {
	rl.r16[name].Value = value
}

func (rl *RegisterList) SplitRL16(name16 string) (uint8, uint8) {
	return tools.Split8(rl.r16[name16].Value)
}

func (rl *RegisterList) SplitRL16ToRL8(name16 string, name1 string, name2 string) {
	rl.SplitR16ToRL8(rl.r16[name16], name1, name2)
}

func (rl *RegisterList) SplitR16ToRL8(reg *Register[uint16], name1 string, name2 string) {
	v1, v2 := tools.Split8(reg.Value)
	rl.r8[name1].Value = v1
	rl.r8[name2].Value = v2
}

func (rl *RegisterList) AddR16Val(name16 string, value uint16) (bool, bool, bool) {
	r, hc, c, z := tools.Add16(rl.r16[name16].Value, value)

	rl.r16[name16].Value = r
	return hc, c, z
}
