package day01

import (
	"slices"
	"strconv"
	"strings"
)

type LocationIds struct {
	left  []int
	right []int
}

func NewLocationIds(lines []string) LocationIds {
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

func TotalDistance(d LocationIds) int {
	t := 0
	for i := range d.left {
		abs := d.left[i] - d.right[i]
		if abs < 0 {
			abs *= -1
		}
		t += abs
	}
	return t
}

func SimilarityScore(d LocationIds) int {
	om := func() map[int]int {
		m := make(map[int]int)
		for i := range d.right {
			m[d.right[i]]++
		}
		return m
	}()

	score := 0
	for i := range d.left {
		score += om[d.left[i]] * d.left[i]
	}

	return score
}
