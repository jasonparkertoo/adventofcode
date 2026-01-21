package day01

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDayOnePartOneA(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	expected := 3
	actual, _ := DoorPassword(ex2025)
	assert.Equal(t, expected, actual)
}

func TestDayOnePartOneB(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	expected := 1195
	actual, _ := DoorPassword(ch2025)
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoA(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	expected := 6
	_, actual := DoorPassword(ex2025)
	assert.Equal(t, expected, actual)
}

func TestDayOnePartTwoB(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	expected := 6770
	_, actual := DoorPassword(ch2025)
	assert.Equal(t, expected, actual)
}
