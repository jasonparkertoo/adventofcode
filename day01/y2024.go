package day01

import (
	"slices"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// LocationIds stores two sorted slices of integer identifiers: one
// for the left hand and one for the right hand. The slices are
// sorted in ascending order to allow efficient lookup and
// comparison operations.
//
// The fields are unexported because callers should use the
// constructor function ToLocationIds to create an instance.
//
// Example usage:
//
//	ids := ToLocationIds(lines)
//	fmt.Println(ids.left, ids.right)
type LocationIds struct {
	left  []int
	right []int
}

// ToLocationIds parses a slice of raw input lines into a LocationIds value.
// Each line is expected to contain two integers separated by three spaces.
// The function ignores any parsing errors and proceeds with the remaining lines.
// It returns a struct whose left and right slices are sorted in ascending order.
func ToLocationIds(lines []string) LocationIds {
	var left, right []int
	for _, line := range lines {
		tok := strings.Split(line, "   ")
		l, _ := strconv.Atoi(tok[0])
		r, _ := strconv.Atoi(tok[1])
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	return LocationIds{
		left:  left,
		right: right,
	}
}

// TotalDistance calculates the sum of absolute differences between
// corresponding left and right identifiers. It assumes that ToLocationIds
// has produced equal‑length slices; if they differ the function will
// panic with an index out of range error. This matches the assumptions
// of the Advent of Code input.
func TotalDistance(d *utils.Data) int {
	ids := ToLocationIds(d.Lines())

	t := 0
	for i := range ids.left {
		abs := ids.left[i] - ids.right[i]
		if abs < 0 {
			abs *= -1
		}
		t += abs
	}
	return t
}

// SimilarityScore computes a weighted count of matching identifiers.
// For each left identifier, the function looks up how many times that
// value appears in the right identifiers and multiplies that count
// by the identifier's value. The sum of these products is returned.
func SimilarityScore(d *utils.Data) int {
	ids := ToLocationIds(d.Lines())

	om := func() map[int]int {
		m := make(map[int]int)
		for i := range ids.right {
			m[ids.right[i]]++
		}
		return m
	}()

	score := 0
	for i := range ids.left {
		score += om[ids.left[i]] * ids.left[i]
	}

	return score
}
