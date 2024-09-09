package emulator

import (
	"errors"
)

const (
	MEMORY_SIZE    = 4096
	STACK_SIZE     = 16
	REGISTERS_SIZE = 16

	DISPLAY_WIDTH  = 64
	DISPLAY_HEIGHT = 32
)

var sprites = []uint8{
	0xF0, 0x90, 0x90, 0x90, 0xF0, //0
	0x20, 0x60, 0x20, 0x20, 0x70, //1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, //2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, //3
	0x90, 0x90, 0xF0, 0x10, 0x10, //4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, //5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, //6
	0xF0, 0x10, 0x20, 0x40, 0x40, //7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, //8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, //9
	0xF0, 0x90, 0xF0, 0x90, 0x90, //A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, //B
	0xF0, 0x80, 0x80, 0x80, 0xF0, //C
	0xE0, 0x90, 0x90, 0x90, 0xE0, //D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, //E
	0xF0, 0x80, 0xF0, 0x80, 0x80, //F
}

type Chip8 struct {
	memory [MEMORY_SIZE]uint8    // Chip-8 memory.
	stack  [STACK_SIZE]uint16    // Chip-8 stack.
	v      [REGISTERS_SIZE]uint8 // 16 general purpose 8-bit registers, goes from V0 to VF.

	// Special Registers
	i  uint16 // Register for storing the memory address.
	dt uint8  // Delay timer register.
	st uint8  // Sound time register.
	pc uint16 // Program Counter. Is used to store the currently executing address.
	sp uint8  // Stack Pointer. It is used to point to the topmost level of the stack.

	// Display buffer
	display [DISPLAY_WIDTH * DISPLAY_HEIGHT]uint8

	shouldDraw bool
	beeper     func()
	rom        string

	// Util struct to do some conversions
	conv Converter
}

func NewChip8() *Chip8 {
	chip8 := Chip8{
		pc:         0x200,
		shouldDraw: true,
		beeper:     func() {},
	}

	for i, sprite := range sprites {
		chip8.memory[i] = sprite
	}

	return &chip8
}

func (c *Chip8) AddBeep(fn func()) {
	c.beeper = fn
}

// Load ROM.
func (c *Chip8) LoadROM(rom string) error {
	if len(rom) == 0 {
		return errors.New("rom doesn't exist")
	}
	c.rom = rom
	return nil
}

// Run the emulator.
func (c Chip8) Run() {}

// Executes Chip-8's next instruction.
func (c *Chip8) NextInstruction(instruction uint16) {
	switch instruction & 0xF000 {
	case 0x0000:
		switch instruction & 0x00FF {
		case 0x00E0:
			c.i00E0()
		case 0x00EE:
			c.i00EE()
		}
	case 0x1000:
		c.i1nnn(c.conv.iToNnn(instruction))
	case 0x2000:
		c.i2nnn(c.conv.iToNnn(instruction))
	case 0x3000:
		c.i3xkk(c.conv.iToVx(instruction), c.conv.iToKk(instruction))
	case 0x4000:
		c.i4xkk(c.conv.iToVx(instruction), c.conv.iToKk(instruction))
	case 0x5000:
		c.i5xy0(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
	case 0x6000:
		c.i6xkk(c.conv.iToVx(instruction), c.conv.iToKk(instruction))
	case 0x7000:
		c.i7xkk(c.conv.iToVx(instruction), c.conv.iToKk(instruction))
	case 0x8000:
		switch instruction & 0x000F {
		case 0x0000:
			c.i8xy0(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x0001:
			c.i8xy1(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x0002:
			c.i8xy2(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x0003:
			c.i8xy3(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x0004:
			c.i8xy4(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x0005:
			c.i8xy5(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x0006:
			c.i8xy6(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x0007:
			c.i8xy7(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		case 0x000E:
			c.i8xyE(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
		}
	case 0x9000:
		c.i9xy0(c.conv.iToVx(instruction), c.conv.iToVy(instruction))
	case 0xA000:
		c.iAnnn(c.conv.iToNnn(instruction))
	case 0xB000:
		c.iBnnn(c.conv.iToNnn(instruction))
	case 0xC000:
		c.iCxkk(c.conv.iToVx(instruction), c.conv.iToKk(instruction))
	case 0xD000:
		c.iDxyn(c.conv.iToVx(instruction), c.conv.iToVy(instruction), c.conv.iToN(instruction))
	case 0xE000:
		switch instruction & 0x00FF {
		case 0x009E:

		case 0x00A1:

		}
	case 0xF000:
		switch instruction & 0x0000 {
		}

	}
}

/* STANDARD CHIP-8 INSTRUCTIONS */

// 00E0 - CLS
//
// Clear the display.
func (c *Chip8) i00E0() {
	for i := range c.display {
		c.display[i] = 0x00
	}
}

// 00EE - RET
//
// Return from a subroutine.
func (c *Chip8) i00EE() {
	c.pc = c.stack[c.sp]
	c.sp--
}

// 1nnn - JP addr
//
// Jump to location nnn.
func (c *Chip8) i1nnn(addr uint16) {
	c.pc = addr
}

// 2nnn - CALL addr
//
// Call subroutine at nnn.
func (c *Chip8) i2nnn(addr uint16) {
	c.sp++
	c.stack[c.sp] = c.pc
	c.pc = addr
}

// 3xkk - SE Vx, byte
//
// Skip next instruction if Vx = kk.
func (c *Chip8) i3xkk(Vx uint8, kk uint8) {
	if Vx == kk {
		c.pc += 2
	}
}

// 4xkk - SNE Vx, byte
//
// Skip next instruction if Vx != kk.
func (c *Chip8) i4xkk(Vx uint8, kk uint8) {
	if Vx != kk {
		c.pc += 2
	}
}

// 5xy0 - SE Vx, Vy
//
// Skip next instruction if Vx = Vy.
func (c *Chip8) i5xy0(Vx uint8, Vy uint8) {
	if Vx == Vy {
		c.pc += 2
	}
}

// 6xkk - LD Vx, byte
//
// Set Vx = kk.
func (c *Chip8) i6xkk(Vx uint8, kk uint8) {
	for i, register := range c.v {
		if register == Vx {
			c.v[i] = kk
		}
	}
}

// 7xkk - ADD Vx, byte
//
// Set Vx = Vx + kk.
func (c *Chip8) i7xkk(Vx uint8, kk uint8) {
	for i, register := range c.v {
		if register == Vx {
			c.v[i] = register + kk
		}
	}
}

// 8xy0 - LD Vx, Vy
//
// Set Vx = Vy.
func (c *Chip8) i8xy0(Vx uint8, Vy uint8) {
	for i, register := range c.v {
		if register == Vx {
			c.v[i] = Vy
		}
	}
}

// 8xy1 - OR Vx, Vy
//
// Set Vx = Vx OR Vy.
func (c *Chip8) i8xy1(Vx uint8, Vy uint8) {
	for i, register := range c.v {
		if register == Vx {
			c.v[i] = Vx | Vy
		}
	}
}

// 8xy2 - AND Vx, Vy
//
// Set Vx = Vx AND Vy.
func (c *Chip8) i8xy2(Vx uint8, Vy uint8) {
	for i, register := range c.v {
		if register == Vx {
			c.v[i] = Vx & Vy
		}
	}
}

// 8xy3 - XOR Vx, Vy
//
// Set Vx = Vx XOR Vy.
func (c *Chip8) i8xy3(Vx uint8, Vy uint8) {
	for i, register := range c.v {
		if register == Vx {
			c.v[i] = Vx ^ Vy
		}
	}
}

// 8xy4 - ADD Vx, Vy
//
// Set Vx = Vx + Vy, set VF = carry.
func (c *Chip8) i8xy4(Vx uint8, Vy uint8) {
	for i, register := range c.v {
		if register == Vx {
			sum := uint16(Vx) + uint16(Vy)
			c.v[i] = uint8(sum)
			if sum > 255 {
				c.v[15] = 1
			} else {
				c.v[15] = 0
			}
		}
	}
}

func (c *Chip8) i8xy5(Vx uint8, Vy uint8) {

}

func (c *Chip8) i8xy6(Vx uint8, Vy uint8) {

}

func (c *Chip8) i8xy7(Vx uint8, Vy uint8) {

}

func (c *Chip8) i8xyE(Vx uint8, Vy uint8) {

}

func (c *Chip8) i9xy0(Vx uint8, Vy uint8) {}

func (c *Chip8) iAnnn(addr uint16) {}

func (c *Chip8) iBnnn(addr uint16) {}

func (c *Chip8) iCxkk(Vx uint8, kk uint8) {}

func (c *Chip8) iDxyn(Vx uint8, Vy uint8, n uint8) {}

func (c *Chip8) iEx9E(Vx uint8) {}

func (c *Chip8) iExA1(Vx uint8) {}

func (c *Chip8) iFx07(Vx uint8) {}

func (c *Chip8) iFx0A(Vx uint8) {}

func (c *Chip8) iFx15(Vx uint8) {}

func (c *Chip8) iFx18(Vx uint8) {}

func (c *Chip8) iFx1E(Vx uint8) {}

func (c *Chip8) iFx29(Vx uint8) {}

func (c *Chip8) iFx33(Vx uint8) {}

func (c *Chip8) iFx55(Vx uint8) {}

func (c *Chip8) iFx65(Vx uint8) {}
