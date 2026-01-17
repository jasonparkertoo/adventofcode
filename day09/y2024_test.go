package day09

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func Test2024A(t *testing.T) {
	assert.Equal(t, 1928, ChecksumHarddrive(ex2024, CompactNormal))
}

func Test2024B(t *testing.T) {
	assert.Equal(t, 6262891638328, ChecksumHarddrive(ch2024, CompactNormal))
}

func Test2024C(t *testing.T) {
	assert.Equal(t, 2858, ChecksumHarddrive(ex2024, CompactLeft))
}

func Test2024D(t *testing.T) {
	assert.Equal(t, 6287317016845, ChecksumHarddrive(ch2024, CompactLeft))
}
