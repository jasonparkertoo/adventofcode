package day07

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay7A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 3749
	actual := TotalCalibrationResult(data)

	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func TestDay7B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 3598800864292
	actual := TotalCalibrationResult(data)

	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func TestDay7C(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 11387
	actual := TotalCalibrationResultWithConcat(data)

	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func TestDay7D(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)
	
	expected := 340362529351427
	actual := TotalCalibrationResultWithConcat(data)

	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}
