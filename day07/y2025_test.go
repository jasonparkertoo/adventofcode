package day07

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var ex2025 = utils.NewData(utils.Example, utils.Year2025)
var ch2025 = utils.NewData(utils.Challenge, utils.Year2025)

func Test2025A(t *testing.T) {
	assert.Equal(t, 21, CountBeamSplits(ex2025))
}

func Test2025B(t *testing.T) {
	assert.Equal(t, 1533, CountBeamSplits(ch2025))
}

func Test2025C(t *testing.T) {
	assert.Equal(t, 40, NumberOfTachyonParticleTimelines(ex2025))
}

func Test2025D(t *testing.T) {
	assert.Equal(t, 10733529153890, NumberOfTachyonParticleTimelines(ch2025))
}
