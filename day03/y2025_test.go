package day03

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2025 = utils.NewData(utils.Example, utils.Year2025)
var ch2025 = utils.NewData(utils.Challenge, utils.Year2025)

func Test2025A(t *testing.T) {
	assert.Equal(t, int64(357), TotalOutputJoltage(ex2025, 2))
}

func Test2025B(t *testing.T) {
	assert.Equal(t, int64(17376), TotalOutputJoltage(ch2025, 2))
}

func Test2025C(t *testing.T) {
	assert.Equal(t, int64(3121910778619), TotalOutputJoltage(ex2025, 12))
}

func Test2025D(t *testing.T) {
	assert.Equal(t, int64(172119830406258), TotalOutputJoltage(ch2025, 12))
}
