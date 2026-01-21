package day02

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay2A(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	expected := 1227775554
	actual, _ := SumInvalidIds(ex2025)
	assert.Equal(t, expected, actual)
}

func TestDay2B(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	expected := 15873079081
	actual, _ := SumInvalidIds(ch2025)
	assert.Equal(t, expected, actual)
}

func TestDay2C(t *testing.T) {
	ex2025, err := utils.NewData(utils.Example, utils.Year2025)
	assert.NoError(t, err)
	expected := 4174379265
	_, actual := SumInvalidIds(ex2025)
	assert.Equal(t, expected, actual)
}

func TestDay2D(t *testing.T) {
	ch2025, err := utils.NewData(utils.Challenge, utils.Year2025)
	assert.NoError(t, err)
	expected := 22617871034
	_, actual := SumInvalidIds(ch2025)
	assert.Equal(t, expected, actual)
}
