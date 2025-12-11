package y2025

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDayOnePartOneA(t *testing.T) {
	input, _ := ReadLines(Day1, PartA)
	doc := NewDocument(input)

	expected := 3
	actual := doc.DoorPassword(MethodNone)

	assert.Equal(t, expected, actual)
}

func TestDayOnePartOneB(t *testing.T) {
	input, _ := ReadLines(Day1, PartB)
	doc := NewDocument(input)

	expected := 1195
	actual := doc.DoorPassword(MethodNone)

	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoA(t *testing.T) {
	input, _ := ReadLines(Day1, PartA)
	doc := NewDocument(input)

	expected := 6
	actual := doc.DoorPassword(Method0x434C49434B)

	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoB(t *testing.T) {
	input, _ := ReadLines(Day1, PartB)
	doc := NewDocument(input)

	expected := 6770
	actual := doc.DoorPassword(Method0x434C49434B)

	assert.Equal(t, expected, actual)
}
