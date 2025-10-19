package tests

import (
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

const (
	Day string = "day10"
)

func Test_Day10_A(t *testing.T) {
	lines, _ := ReadLines(Example, Day)
	lavaTrails := c.NewLavaTrails(lines)
	
	expected := 36
	actual := c.TotalTrailheadScore(lavaTrails)
	
	assert.Equal(t, expected, actual, "expected %d, have %d", expected, actual)	
}

func Test_Day10_B(t *testing.T) {
	lines, _ := ReadLines(Challenge, Day)
	lavaTrails := c.NewLavaTrails(lines)
	
	expected := 776
	actual := c.TotalTrailheadScore(lavaTrails)
	
	assert.Equal(t, expected, actual, "expected %d, have %d", expected, actual)	
}

func Test_Day10_C(t *testing.T) {
	lines, _ := ReadLines(Example, Day)
	lavaTrails := c.NewLavaTrails(lines)
	
	expected := 81
	actual := c.TotalTrailheadRating(lavaTrails)
	
	assert.Equal(t, expected, actual, "expected %d, have %d", expected, actual)	
}

func Test_Day10_D(t *testing.T) {
	lines, _ := ReadLines(Challenge, Day)
	lavaTrails := c.NewLavaTrails(lines)
	
	expected := 1657
	actual := c.TotalTrailheadRating(lavaTrails)
	
	assert.Equal(t, expected, actual, "expected %d, have %d", expected, actual)	
}
