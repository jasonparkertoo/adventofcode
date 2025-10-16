package tests

import (
	"os"
	"strings"
	"testing"

	c "adventofcode.dev/challenges"
	"github.com/stretchr/testify/assert"
)

var exampleData []string
var challengeData []string

func init() {
	var err error
	eData, err := os.ReadFile("../data/example/day9")
	exampleData = strings.Split(string(eData), "\n")
	if err != nil {
		panic(eData)
	}

	cData, err := os.ReadFile("../data/challenge/day9")
	challengeData = strings.Split(string(cData), "\n")
	if err != nil {
		panic(err)
	}
}

func Test_Day9_A(t *testing.T) {
	hd := c.Harddrive{DataMap: exampleData[0]}

	expected := 1928
	actual := c.ChecksumHarddrive(hd, c.CompactNormal)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}

func Test_Day9_B(t *testing.T) {
	hd := c.Harddrive{DataMap: challengeData[0]}

	expected := 6262891638328
	actual := c.ChecksumHarddrive(hd, c.CompactNormal)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}

func Test_Day9_C(t *testing.T) {
	hd := c.Harddrive{DataMap: exampleData[0]}

	expected := 2858
	actual := c.ChecksumHarddrive(hd, c.CompactLeft)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}

func Test_Day9_D(t *testing.T) {
	hd := c.Harddrive{DataMap: challengeData[0]}

	expected := 6287317016845
	actual := c.ChecksumHarddrive(hd, c.CompactLeft)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}
