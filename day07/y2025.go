package day07

import (
	"adventofcode.dev/utils"
)

// CountBeamSplits calculates how many times a tachyon beam is split within the manifold.
// A split is counted once per incoming beam that hits a splitter (^).
func CountBeamSplits(d *utils.Data) int {
	grid := d.AsGrid()

	start, ok := findStart(grid)
	if !ok {
		return 0
	}

	// beams present in current row
	current := map[int]bool{start.c: true}
	splitCount := 0

	for row := start.r + 1; row < len(grid) && len(current) > 0; row++ {
		next := make(map[int]bool)

		for col := range current {
			if col < 0 || col >= len(grid[row]) {
				continue
			}

			switch grid[row][col] {
			case "^":
				// splitter activates once
				splitCount++
				next[col-1] = true
				next[col+1] = true

			default:
				next[col] = true
			}
		}

		current = next
	}

	return splitCount
}

type pos struct{ r, c int }

func findStart(grid [][]string) (pos, bool) {
	for r, row := range grid {
		for c, cell := range row {
			if cell == "S" {
				return pos{r, c}, true
			}
		}
	}
	return pos{}, false
}
