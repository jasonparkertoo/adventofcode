package day03

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay3PartOneA(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 161
	actual := CalculateUncorrupted(data)

	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func TestDay3PartOneB(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)
	
	expected := 156388521
	actual := CalculateUncorrupted(data)

	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func TestDay3PartTwoB(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)
	
	expected := 75920122
	actual := Calculate(data)

	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}
