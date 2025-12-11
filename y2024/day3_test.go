package y2024

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay3PartOneA(t *testing.T) {
	records, err := ReadLines(Day3, PartA)
	if err != nil {
		panic(MsgPanic)
	}
	instructions := NewInstructions(records)

	expected := 161
	actual := CalculateUncorrupted(instructions)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay3PartOneB(t *testing.T) {
	records, err := ReadLines(Day3, PartB)
	if err != nil {
		panic("unable to find input data")
	}
	instructions := NewInstructions(records)

	expected := 156388521
	actual := CalculateUncorrupted(instructions)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay3PartTwoB(t *testing.T) {
	records, err := ReadLines(Day3, PartB)
	if err != nil {
		panic("unable to find input data")
	}
	instructions := NewInstructions(records)

	expected := 75920122
	actual := Calculate(instructions)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
