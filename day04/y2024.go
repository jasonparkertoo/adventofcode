package day04

import (
	"slices"
	"strings"

	"adventofcode.dev/utils"
)

type Puzzle struct {
	letters [][]string
}

// ToPuzzle converts a utils.Data into a Puzzle by splitting each line into individual letters.
// It ignores empty lines.
func ToPuzzle(data *utils.Data) *Puzzle {
	lines := data.Lines()

	var letters [][]string

	for _, line := range lines {
		if line == "" {
			continue
		}
		letters = append(letters, strings.Split(line, ""))
	}

	return &Puzzle{
		letters: letters,
	}
}

type SearchFunction func(string, *Puzzle) int

// compare returns 1 if curr and target are equal or reverse of each other, otherwise 0.
func compare(curr []string, target []string) int {
	l := make([]string, len(curr))
	r := make([]string, len(curr))

	copy(l, curr)
	copy(r, target)

	out := 0
	if slices.Equal(l, r) {
		out = 1
	}
	slices.Reverse(r)
	if slices.Equal(l, r) {
		out = 1
	}
	return out
}

// toSlice converts a string to a slice of its uppercase characters.
func toSlice(s string) []string {
	return strings.Split(strings.ToUpper(s), "")
}

// SearchHorizontal counts the number of occurrences of w horizontally (left-to-right or right-to-left) in p.
func SearchHorizontal(w string, p *Puzzle) int {
	count := 0
	target := toSlice(w)
	for row := range len(p.letters) {
		for col := range len(p.letters[row]) - len(target) + 1 {
			curr := p.letters[row][col : col+len(target)]
			count += compare(curr, target)
		}
	}
	return count
}

// SearchVertical counts the number of occurrences of w vertically (top-to-bottom or bottom-to-top) in p.
func SearchVertical(w string, p *Puzzle) int {
	count := 0
	target := toSlice(w)
	for col := range len(p.letters[0]) {
		for row := range len(p.letters) - (len(target) - 1) {
			var curr []string
			for k := range len(target) {
				curr = append(curr, p.letters[row+k][col])
			}
			count += compare(curr, target)
		}
	}
	return count
}

// SearchLeftToRightDiagonal counts diagonal occurrences from top-left to bottom-right.
func SearchLeftToRightDiagonal(w string, p *Puzzle) int {
	count := 0
	letters := toSlice(w)
	for row := range len(p.letters) - 3 {
		for col := range len(p.letters) - 3 {
			var curr []string
			for k := range len(letters) {
				curr = append(curr, p.letters[row+k][col+k])
			}
			count += compare(curr, letters)
		}
	}
	return count
}

// SearchRightToLeftDiagonal counts diagonal occurrences from top-right to bottom-left.
func SearchRightToLeftDiagonal(w string, p *Puzzle) int {
	count := 0
	letters := toSlice(w)
	for row := range len(p.letters) - 3 {
		for col := 3; col < len(p.letters); col++ {
			var curr []string
			for k := range len(letters) {
				curr = append(curr, p.letters[row+k][col-k])
			}
			count += compare(curr, letters)
		}
	}
	return count
}

// SearchXPattern counts X-shaped pattern occurrences of w in p.
func SearchXPattern(w string, p *Puzzle) int {
	letters := toSlice(w)

	count := 0
	for i := range len(p.letters) - 2 {
		for j := range len(p.letters[i]) - 2 {
			l := make([]string, len(letters))
			r := make([]string, len(letters))

			l[0] = p.letters[i][j]
			l[1] = p.letters[i+1][j+1]
			l[2] = p.letters[i+2][j+2]

			r[0] = p.letters[i][j+2]
			r[1] = p.letters[i+1][j+1]
			r[2] = p.letters[i+2][j]

			if compare(l, letters)+compare(r, letters) == 2 {
				count++
			}
		}
	}
	return count
}

// Count counts horizontal, vertical, and both diagonal occurrences of w in the puzzle derived from d.
func Count(w string, d *utils.Data) int {
	p := ToPuzzle(d)

	count := 0
	count += SearchHorizontal(w, p)
	count += SearchVertical(w, p)
	count += SearchLeftToRightDiagonal(w, p)
	count += SearchRightToLeftDiagonal(w, p)
	return count
}

// CountPattern counts occurrences using the supplied search function pt.
func CountPattern(w string, d *utils.Data, pt func(string, *Puzzle) int) int {
	p := ToPuzzle(d)
	return pt(w, p)
}
