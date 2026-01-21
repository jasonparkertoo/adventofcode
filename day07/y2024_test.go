package day07

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2024A(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 3749, TotalCalibrationResult(ex2024))
}

func Test2024B(t *testing.T) {
	ch2024, err := utils.NewData(utils.Challenge, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 3598800864292, TotalCalibrationResult(ch2024))
}

func Test2024C(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 11387, TotalCalibrationResultWithConcat(ex2024))
}

func Test2024D(t *testing.T) {
	ch2024, err := utils.NewData(utils.Challenge, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 340362529351427, TotalCalibrationResultWithConcat(ch2024))
}
