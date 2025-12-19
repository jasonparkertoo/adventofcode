package day01

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDayOnePartOneA(t *testing.T) {
	input, _ := ReadLines(Year2025, Example)
	doc := NewDocument(input)
	result := doc.DoorPassword()

	expected := 3
	actual := result[0]
	
	assert.Equal(t, expected, actual)
}

func TestDayOnePartOneB(t *testing.T) {
	input, _ := ReadLines(Year2025, Challenge)
	doc := NewDocument(input)
	result := doc.DoorPassword()

	expected := 1195
	actual := result[0]
	
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoA(t *testing.T) {
	input, _ := ReadLines(Year2025, Example)
	doc := NewDocument(input)
	result := doc.DoorPassword()

	expected := 6
	actual := result[1]
	
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoB(t *testing.T) {
	input, _ := ReadLines(Year2025, Challenge)
	doc := NewDocument(input)
	result := doc.DoorPassword()

	expected := 6770
	actual := result[1]
	
	assert.Equal(t, expected, actual)
}
