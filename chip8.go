package main

const (
	MEMORY_SIZE    = 4096
	STACK_SIZE     = 16
	REGISTERS_SIZE = 16
)

type Chip8 struct {
	Memory     [MEMORY_SIZE]uint8    // Chip-8 memory.
	Stack      [STACK_SIZE]uint16    // Chip-8 stack.
	Registers  [REGISTERS_SIZE]uint8 // 16 general purpose 8-bit registers, go from V0 to VF.
	I          uint16                // Register for storing the memory address.
	DelayTimer uint8                 // DelayTimer.
	SoundTimer uint8                 // SoundTimer.
	PC         uint16                // Is used to store the currently executing address.
	SP         uint8                 // It is used to point to the topmost level of the stack.
}
