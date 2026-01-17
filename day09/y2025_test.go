package day09

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var exampleData = utils.NewData(utils.Example, utils.Year2025)
var challengeData = utils.NewData(utils.Challenge, utils.Year2025)

func Test2025A(t *testing.T) {
	assert.Equal(t, 50, findLargestRectangle(exampleData))
}

func Test2025B(t *testing.T) {
	assert.Equal(t, 4756718172, findLargestRectangle(challengeData))
}

func Test2025C(t *testing.T) {
	assert.Equal(t, 24, findLargestRectangleOfAny(exampleData))
}

func Test2025D(t *testing.T) {
	assert.Equal(t, 1665679194, findLargestRectangleOfAny(challengeData))
}
