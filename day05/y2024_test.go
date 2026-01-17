package day05

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func Test2024A(t *testing.T) {
	assert.Equal(t, 143, SumMiddlePageNumbers(ex2024))
}

func Test2024B(t *testing.T) {
	assert.Equal(t, 5747, SumMiddlePageNumbers(ch2024))
}

func Test2024C(t *testing.T) {
	assert.Equal(t, 123, SumIncorrectMiddlePageNumbers(ex2024))
}

func Test2024D(t *testing.T) {
	assert.Equal(t, 5502, SumIncorrectMiddlePageNumbers(ch2024))
}
