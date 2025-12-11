package y2024

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay6A(t *testing.T) {
	records, err := ReadLines(Day6, PartA)
	if err != nil {
		panic(MsgPanic)
	}
	m, _ := Generate(records)
	path, _ := Explore(m)

	expected := 41
	actual := len(path)

	assert.Equal(t, expected, actual)
}

func TestDay6B(t *testing.T) {
	records, err := ReadLines(Day6, PartB)
	if err != nil {
		panic(MsgPanic)
	}
	m, _ := Generate(records)
	path, _ := Explore(m)

	expected := 5551
	actual := len(path)

	assert.Equal(t, expected, actual)
}

func TestDay6C(t *testing.T) {
	records, err := ReadLines(Day6, PartA)
	if err != nil {
		panic("unable to find input data")
	}
	m, _ := Generate(records)

	expected := 6
	actual, _ := CountLoopPositions(m)

	assert.Equal(t, expected, actual)
}

func TestDay6D(t *testing.T) {
	records, err := ReadLines(Day6, PartB)
	if err != nil {
		panic("unable to find input data")
	}
	m, _ := Generate(records)

	expected := 1939
	actual, _ := CountLoopPositions(m)

	assert.Equal(t, expected, actual)
}
