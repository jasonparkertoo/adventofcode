package day09

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var exampleData = utils.NewData(utils.Example, utils.Year2025)
var challengeData = utils.NewData(utils.Challenge, utils.Year2025)

func TestDay92025A(t *testing.T) {
	assert.Equal(t, 50, findLargestRectangle(exampleData))
}

func TestDay92025B(t *testing.T) {
	assert.Equal(t, 4756718172, findLargestRectangle(challengeData))
}
