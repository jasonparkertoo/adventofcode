package day04

import (
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay4PartA(t *testing.T) {
	data := utils.NewData(Example, Year2024)

	expected := 18
	actual := Count("xmas", data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay4PartB(t *testing.T) {
	data := utils.NewData(Challenge, Year2024)

	expected := 2504
	actual := Count("xmas", data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay4PartC(t *testing.T) {
	data := utils.NewData(Example, Year2024)

	expected := 9
	actual := CountPattern("mas", data, SearchXPattern)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay4PartD(t *testing.T) {
	data := utils.NewData(Challenge, Year2024)

	expected := 1923
	actual := CountPattern("mas", data, SearchXPattern)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
