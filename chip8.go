package main

type Chip8 struct {
	Memory     [4096]uint8
	Stack      [16]uint16
	Registers  [16]uint8 // 16 general purpose 8-bit registers, V0 to VF.
	I          uint16    // Register for storing the memory address.
	DelayTimer uint8     // DelayTimer.
	SoundTimer uint8     // SoundTimer.
	PC         uint16    // Is used to store the currently executing address.
	SP         uint8     // It is used to point to the topmost level of the stack.
}
