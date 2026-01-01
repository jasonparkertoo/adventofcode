package day05

import (
	"testing"

	"adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

func Test2025Day5A(t *testing.T) {
	d := utils.NewData(utils.Example, utils.Year2025)
	
	expected := 3
	actual := CountFreshIds(d)
	
	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func Test2025Day5B(t *testing.T) {
	d := utils.NewData(utils.Challenge, utils.Year2025)
	
	expected := 896
	actual := CountFreshIds(d)
	
	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func Test2025Day5C(t *testing.T) {
	d := utils.NewData(utils.Example, utils.Year2025)
	
	var expected int64 = 14
	actual := NumberOfFreshRangeIds(d)
	
	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}

func Test2025Day5D(t *testing.T) {
	d := utils.NewData(utils.Challenge, utils.Year2025)
	
	var expected int64 = 346240317247002
	actual := NumberOfFreshRangeIds(d)
	
	assert.Equal(t, expected, actual, utils.MsgExpected, expected, actual)
}
