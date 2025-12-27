package day05

import (
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay5A(t *testing.T) {
	data := utils.NewData(Example, Year2024)

	expected := 143
	actual := SumMiddlePageNumbers(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay5B(t *testing.T) {
	data := utils.NewData(Challenge, Year2024)

	expected := 5747
	actual := SumMiddlePageNumbers(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay5C(t *testing.T) {
	data := utils.NewData(Example, Year2024)

	expected := 123
	actual := SumIncorrectMiddlePageNumbers(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay5D(t *testing.T) {
	data := utils.NewData(Challenge, Year2024)

	expected := 5502
	actual := SumIncorrectMiddlePageNumbers(data)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}
