package emulator

type Converter struct{}

// Converts an instruction (uint16) to nnn (the lowest 12 bits of the instruction)
func (c Converter) iToNnn(i uint16) uint16 {
	return i & 0x0FFF
}

// Converts an instruction (uint16) to n (the lowest 4 bits of the instruction)
func (c Converter) iToN(i uint16) uint8 {
	return uint8(i & 0x000F)
}

// Converts an instruction (uint16) to x (the lower 4 bits of the high byte of the instruction)
func (c Converter) iToVx(i uint16) uint8 {
	return uint8(i & 0x0F00)
}

// Converts an instruction (uint16) to y (the upper 4 bits of the low byte of the instruction)
func (c Converter) iToVy(i uint16) uint8 {
	return uint8(i & 0x00F0)
}

// Converts an instruction (uint16) to kk (the lowest 8 bits of the instruction)
func (c Converter) iToKk(i uint16) uint8 {
	return uint8(i & 0x00FF)
}
