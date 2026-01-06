package day06

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2025Day6A(t *testing.T) {
	d := utils.NewData(utils.Example, utils.Year2025)

	expected := 4277556
	actual := CalculateGrandTotal(d)

	assert.Equal(t, expected, actual)
}

func Test2025Day6B(t *testing.T) {
	d := utils.NewData(utils.Challenge, utils.Year2025)

	expected := 5595593539811
	actual := CalculateGrandTotal(d)

	assert.Equal(t, expected, actual)
}

func Test2025Day6C(t *testing.T) {
	d := utils.NewData(utils.Example, utils.Year2025)

	expected := 3263827
	actual := CalculateGrandTotal2(d)

	assert.Equal(t, expected, actual)
}

func Test2025Day6D(t *testing.T) {
	d := utils.NewData(utils.Challenge, utils.Year2025)

	expected := 10153315705125
	actual := CalculateGrandTotal2(d)

	assert.Equal(t, expected, actual)
}
