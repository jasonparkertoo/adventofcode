package day01

import (
	"strconv"

	"adventofcode.dev/utils"
)

// Pair represents a single rotation instruction consisting of a
// direction ("L" for left or "R" for right) and a distance to move.
// The distance is an integer number of steps.
type Pair struct {
	Direction string
	Distance  int
}

// Document holds a slice of Pair values parsed from the input.
// It represents the entire sequence of rotations.
type Document struct {
	Pairs []Pair
}

// ToDocument converts the raw input data into a *Document.
// It reads each line from the utils.Data object, splits the first
// character as the direction and parses the remaining substring as
// an integer distance. The function silently ignores any parsing
// errors (as the Advent of Code input is guaranteed to be valid).
// The returned Document is a shallow copy of the data and can be
// safely used by the caller.
func ToDocument(d *utils.Data) *Document {
	var pairs []Pair

	for _, rotation := range d.Lines() {
		d := rotation[0:1]
		t, _ := strconv.Atoi(rotation[1:])
		pair := Pair{Direction: d, Distance: t}
		pairs = append(pairs, pair)
	}

	return &Document{
		Pairs: pairs,
	}
}

// DoorPassword processes the parsed document and returns two
// integers: the number of times the door lands at position 0 at the
// end of a rotation sequence, and the total number of times the
// door passes position 0 during the entire sequence.
// The calculation uses a circular track of 100 positions. Each
// rotation step moves the current position left or right by one
// step. Position 0 is considered the top of the door. The initial
// position is 50.
func DoorPassword(d *utils.Data) (int, int) {
	doc := ToDocument(d)
	
	landed, visited := 0, 0

	pos := 50
	for _, p := range doc.Pairs {
		dis := p.Distance
		for range dis {
			switch p.Direction {
			case "L":
				pos = (pos - 1 + 100) % 100
			case "R":
				pos = (pos + 1) % 100
			}
			if pos == 0 {
				visited++
			}
		}
		if pos == 0 {
			landed++
		}
	}

	return landed, visited
}
