package day08

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func Test2024A(t *testing.T) {
	assert.Equal(t, 14, CountUniqueLocations(ex2024))
}

func Test2024B(t *testing.T) {
	assert.Equal(t, 261, CountUniqueLocations(ch2024))
}

func Test2024C(t *testing.T) {
	assert.Equal(t, 34, CountUniqueLocationsHarmonics(ex2024))
}

func Test2024D(t *testing.T) {
	assert.Equal(t, 988, CountUniqueLocationsHarmonics(ch2024))
}
