package y2024

import (
	"testing"

	. "adventofcode.dev/utils"
	"github.com/stretchr/testify/assert"
)

var exampleData []string
var challengeData []string

func init() {
	exampleData, _ = ReadLines(Day9, PartA)
	challengeData, _ = ReadLines(Day9, PartB)
}

func TestDay9A(t *testing.T) {
	hd := Harddrive{DataMap: exampleData[0]}

	expected := 1928
	actual := ChecksumHarddrive(hd, CompactNormal)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay9B(t *testing.T) {
	hd := Harddrive{DataMap: challengeData[0]}

	expected := 6262891638328
	actual := ChecksumHarddrive(hd, CompactNormal)

	assert.Equal(t, expected, actual, MsgExpected, expected, actual)
}

func TestDay9C(t *testing.T) {
	hd := Harddrive{DataMap: exampleData[0]}

	expected := 2858
	actual := ChecksumHarddrive(hd, CompactLeft)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}

func TestDay9D(t *testing.T) {
	hd := Harddrive{DataMap: challengeData[0]}

	expected := 6287317016845
	actual := ChecksumHarddrive(hd, CompactLeft)

	assert.Equal(t, expected, actual, "expected %s, got %d", expected, actual)
}
