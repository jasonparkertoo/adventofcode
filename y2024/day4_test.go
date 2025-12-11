package y2024

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay4PartA(t *testing.T) {
	records, err := ReadLines(Day4, PartA)
	if err != nil {
		panic(MsgPanic)
	}
	puzzle := NewPuzzle(records)

	expected := 18
	actual := Count("xmas", &puzzle)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay4PartB(t *testing.T) {
	records, err := ReadLines(Day4, PartB)
	if err != nil {
		panic(MsgPanic)
	}
	puzzle := NewPuzzle(records)

	expected := 2504
	actual := Count("xmas", &puzzle)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay4PartC(t *testing.T) {
	records, err := ReadLines(Day4, PartA)
	if err != nil {
		panic("unable to find input data")
	}
	puzzle := NewPuzzle(records)

	expected := 9
	actual := CountPattern("mas", &puzzle, SearchXPattern)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay4PartD(t *testing.T) {
	records, err := ReadLines(Day4, PartB)
	if err != nil {
		panic("unable to find input data")
	}
	puzzle := NewPuzzle(records)

	expected := 1923
	actual := CountPattern("mas", &puzzle, SearchXPattern)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
