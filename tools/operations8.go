package tools

import "math/bits"

func Combine8(p1 uint8, p2 uint8) uint16 {
	return (uint16(p1) >> 8) | uint16(p2)
}

func Split8(value uint16) (uint8, uint8) {
	return uint8((value & 0xFF00) >> 8), uint8(value & 0xFF)
}

func Bit8(value uint8, pos uint) bool {
	return (value & (1 << pos)) > 0
}

func Set8(value uint8, pos uint) uint8 {
	return value | (1 << pos)
}

func Clear8(value uint8, pos uint) uint8 {
	return value & ^(1 << pos)
}

func Add8(value1 uint8, value2 uint8) (uint8, bool, bool, bool) {
	r := value1 + value2
	c := value1 > r
	hc := (((value1 & 0xF) + (value2 & 0xF)) & 0x10) == 0x10

	return r, c, hc, (r == 0)
}

func Sub8(value1 uint8, value2 uint8) (uint8, bool, bool, bool) {
	r := value1 - value2
	c := value1 < r
	hc := (((value1 & 0xF) - (value2 & 0xF)) & 0x10) == 0x10

	return r, c, hc, (r == 0)
}

func And8(value1 uint8, value2 uint8) (uint8, bool) {
	r := value1 & value2

	return r, (r == 0)
}

func Or8(value1 uint8, value2 uint8) (uint8, bool) {
	r := value1 & value2

	return r, (r == 0)
}

func Xor8(value1 uint8, value2 uint8) (uint8, bool) {
	r := value1 ^ value2

	return r, (r == 0)
}

func Not8(value1 uint8) uint8 {
	r := ^value1

	return r
}

func Swap8(value uint8) uint8 {
	return ((value&0x0F)<<4 | (value&0xF0)>>4)
}

func Nibbles(value uint8) (uint8, uint8) {
	return (value & 0xF0) >> 4, (value & 0x0F)
}

func ShiftL8(value uint8, len uint) (uint8, bool) {
	msb := (value & (1 << 7)) > 0
	return (value << len), msb
}

func ShiftR8(value uint8, len uint) (uint8, bool) {
	lsb := (value & (1 << 0)) > 0
	return (value >> len), lsb
}

func RotateL8(value uint8, len uint) (uint8, bool) {
	msb := (value & (1 << 7)) > 0
	return bits.RotateLeft8(value, int(len)), msb
}

func RotateR8(value uint8, len uint) (uint8, bool) {
	lsb := (value & (1 << 0)) > 0
	return bits.RotateLeft8(value, -int(len)), lsb
}
