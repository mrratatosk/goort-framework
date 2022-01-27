package tools

func Add16(value1 uint16, value2 uint16) (uint16, bool, bool, bool) {
	r := value1 + value2
	c := value1 > r
	hc := (((value1 & 0xFFF) + (value2 & 0xFFF)) & 0x1000) == 0x1000

	return r, c, hc, (r == 0)
}

func Sub16(value1 uint16, value2 uint16) (uint16, bool, bool, bool) {
	r := value1 - value2
	c := value1 < r
	hc := (((value1 & 0xFFF) - (value2 & 0xFFF)) & 0x1000) == 0x1000

	return r, c, hc, (r == 0)
}
