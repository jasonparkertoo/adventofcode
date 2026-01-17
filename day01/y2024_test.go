package day01

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2024 = utils.NewData(utils.Example, utils.Year2024)
var ch2024 = utils.NewData(utils.Challenge, utils.Year2024)

func Test2024A(t *testing.T) {
	assert.Equal(t, 11, TotalDistance(ex2024))
}

func Test2024B(t *testing.T) {
	assert.Equal(t, 2285373, TotalDistance(ch2024))
}

func Test2024C(t *testing.T) {
	assert.Equal(t, 31, SimilarityScore(ex2024))
}

func Test2024D(t *testing.T) {
	assert.Equal(t, 21142653, SimilarityScore(ch2024))
}
