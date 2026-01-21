package day02

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2024A(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 2, NumberOfSafeReports(false, ex2024))
}

func Test2024B(t *testing.T) {
	ch2024, err := utils.NewData(utils.Challenge, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 606, NumberOfSafeReports(false, ch2024))
}

func Test2024C(t *testing.T) {
	ex2024, err := utils.NewData(utils.Example, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 4, NumberOfSafeReports(true, ex2024))
}

func Test2024D(t *testing.T) {
	ch2024, err := utils.NewData(utils.Challenge, utils.Year2024)
	assert.NoError(t, err)
	assert.Equal(t, 644, NumberOfSafeReports(true, ch2024))
}
