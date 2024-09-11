package emulator

import (
	"testing"
)

var conv = Converter{}

func testConvNnn(t *testing.T) {
	var instruction uint16
	converted := conv.iToNnn(instruction)

}

func testConvVx(t *testing.T) {
	var instruction uint16
	conv.iToVx(instruction)
}

func testConvVy(t *testing.T) {
	var instruction uint16
	conv.iToVy(instruction)
}

func testConvKk(t *testing.T) {
	var instruction uint16
	conv.iToKk(instruction)
}

func testConvN(t *testing.T) {
	var instruction uint16
	conv.iToN(instruction)
}
