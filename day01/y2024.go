package day01

import (
	"slices"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type LocationIds struct {
	left  []int
	right []int
}

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
