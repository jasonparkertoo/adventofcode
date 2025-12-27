package day01

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDayOnePartOneA(t *testing.T) {
	data := NewData(Example, Year2025)

	expected := 3
	actual := DoorPassword(data)[0]
	
	assert.Equal(t, expected, actual)
}

func TestDayOnePartOneB(t *testing.T) {
	data := NewData(Challenge, Year2025)

	expected := 1195
	actual := DoorPassword(data)[0]
	
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoA(t *testing.T) {
	data := NewData(Example, Year2025)

	expected := 6
	actual := DoorPassword(data)[1]
	
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoB(t *testing.T) {
	data := NewData(Challenge, Year2025)

	expected := 6770
	actual := DoorPassword(data)[1]
	
	assert.Equal(t, expected, actual)
}
