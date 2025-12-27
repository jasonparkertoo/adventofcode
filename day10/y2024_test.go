package day10

import (
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay10A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 36
	actual := TotalTrailheadScore(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay10B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 776
	actual := TotalTrailheadScore(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay10C(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 81
	actual := TotalTrailheadRating(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay10D(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 1657
	actual := TotalTrailheadRating(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
