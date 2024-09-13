package emulator

import (
	"testing"
)

var conv = Converter{}

func TestConvNnn(t *testing.T) {
	var instruction uint16 = 0xA123
	var expected uint16 = 0x123
	got := conv.iToNnn(instruction)

	if got != expected {
		t.Errorf("iToNnn(%X) = %X; want %X", instruction, got, expected)
	}
}

func TestConvN(t *testing.T) {
	var instruction uint16 = 0xA123
	var expected uint8 = 0x3
	got := conv.iToN(instruction)

	if got != expected {
		t.Errorf("iToN(%X) = %X; want %X", instruction, got, expected)
	}
}

func TestConvVx(t *testing.T) {
	var instruction uint16 = 0xA123
	var expected uint8 = 0x1
	got := conv.iToVx(instruction)

	if got != expected {
		t.Errorf("iToVx(%X) = %X; want %X", instruction, got, expected)
	}
}

func TestConvVy(t *testing.T) {
	var instruction uint16 = 0xA123
	var expected uint8 = 0x2
	got := conv.iToVy(instruction)

	if got != expected {
		t.Errorf("iToVy(%X) = %X; want %X", instruction, got, expected)
	}

}

func TestConvKk(t *testing.T) {
	var instruction uint16 = 0xA123
	var expected uint8 = 0x23
	got := conv.iToKk(instruction)

	if got != expected {
		t.Errorf("iToKk(%X) = %X; want %X", instruction, got, expected)
	}
}
