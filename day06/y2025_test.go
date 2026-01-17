package day06

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2025 = utils.NewData(utils.Example, utils.Year2025)
var ch2025 = utils.NewData(utils.Challenge, utils.Year2025)


func Test2025Day6A(t *testing.T) {
	assert.Equal(t, 4277556, CalculateGrandTotal(ex2025))
}

func Test2025Day6B(t *testing.T) {
	assert.Equal(t, 5595593539811, CalculateGrandTotal(ch2025))
}

func Test2025Day6C(t *testing.T) {
	assert.Equal(t, 3263827, CalculateGrandTotal2(ex2025))
}

func Test2025Day6D(t *testing.T) {
	assert.Equal(t, 10153315705125, CalculateGrandTotal2(ch2025))
}
