package emulator

type Converter struct{}

func (c Converter) iToNnn(i uint16) uint16 {
	return i & 0x0FFF
}

func (c Converter) iToVx(i uint16) uint8 {
	return uint8(i & 0x0F00)
}

func (c Converter) iToVy(i uint16) uint8 {
	return uint8(i & 0x00F0)
}

func (c Converter) iToKk(i uint16) uint8 {
	return uint8(i & 0x00FF)
}

func (c Converter) iToN(i uint16) uint8 {
	return uint8(i & 0x000F)
}
