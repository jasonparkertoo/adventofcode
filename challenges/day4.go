package challenges

import (
	"slices"
	"strings"
)

type Puzzle struct {
	letters [][]string
}

func NewPuzzle(lines []string) Puzzle {
	var letters [][]string

	for _, line := range lines {
		if line == "" {
			continue
		}
		letters = append(letters, strings.Split(line, ""))
	}

	return Puzzle{
		letters: letters,
	}
}

type SearchFunction func(string, *Puzzle) int



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

func toSlice(s string) []string {
	return strings.Split(strings.ToUpper(s), "")
}


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

func Count(w string, p *Puzzle) int {
	count := 0
	count += SearchHorizontal(w, p)
	count += SearchVertical(w, p)
	count += SearchLeftToRightDiagonal(w, p)
	count += SearchRightToLeftDiagonal(w, p)
	return count
}


func CountPattern(w string, p *Puzzle, pt func(string, *Puzzle) int) int {
	return pt(w, p)
}
