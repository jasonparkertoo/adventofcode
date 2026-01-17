package day03

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func Test2024A(t *testing.T) {
	assert.Equal(t, 161, CalculateUncorrupted(ex2024))
}

func Test2024B(t *testing.T) {
	assert.Equal(t, 156388521, CalculateUncorrupted(ch2024))
}

func Test2024C(t *testing.T) {
	assert.Equal(t, 75920122, Calculate(ch2024))
}
