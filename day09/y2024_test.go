package day09

import (
	"testing"

	"adventofcode.dev/utils"
	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func TestDay9A(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 1928
	actual := ChecksumHarddrive(data, CompactNormal)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay9B(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 6262891638328
	actual := ChecksumHarddrive(data, CompactNormal)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay9C(t *testing.T) {
	data := utils.NewData(utils.Example, utils.Year2024)

	expected := 2858
	actual := ChecksumHarddrive(data, CompactLeft)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}

func TestDay9D(t *testing.T) {
	data := utils.NewData(utils.Challenge, utils.Year2024)

	expected := 6287317016845
	actual := ChecksumHarddrive(data, CompactLeft)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}
