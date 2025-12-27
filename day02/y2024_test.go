package day02

import (
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay2PartOneA(t *testing.T) {
	data := utils.NewData(Example, Year2024)

	expected := 2
	actual := NumberOfSafeReports(false, data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay2PartOneB(t *testing.T) {
	data := utils.NewData(Challenge, Year2024)

	expected := 606
	actual := NumberOfSafeReports(false, data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay2PartTwoA(t *testing.T) {
	data := utils.NewData(Example, Year2024)

	expected := 4
	actual := NumberOfSafeReports(true, data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay2PartTwoB(t *testing.T) {
	data := utils.NewData(Challenge, Year2024)

	expected := 644
	actual := NumberOfSafeReports(true, data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
