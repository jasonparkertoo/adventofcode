package day01

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2025 = utils.NewData(utils.Example, utils.Year2025)
var ch2025 = utils.NewData(utils.Challenge, utils.Year2025)

func TestDayOnePartOneA(t *testing.T) {
	expected := 3
	actual, _ := DoorPassword(ex2025)
	assert.Equal(t, expected, actual)
}

func TestDayOnePartOneB(t *testing.T) {
	expected := 1195
	actual, _ := DoorPassword(ch2025)
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoA(t *testing.T) {
	 expected := 6
	 _, actual := DoorPassword(ex2025)
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoB(t *testing.T) {
	expected := 6770
	_, actual := DoorPassword(ch2025)
	assert.Equal(t, expected, actual)
}
