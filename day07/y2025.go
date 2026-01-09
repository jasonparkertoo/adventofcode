package day07

import (
	"adventofcode.dev/utils"
)

// propagateCounts moves beams from the current row to the next one.
// For each column that has a beam, if the cell contains a splitter ("^")
// it creates two beams in the adjacent columns; otherwise the beam continues
// straight down. The function returns the mapping for the next row.
func propagateCounts(curr map[int]int, row []string) map[int]int {
	next := make(map[int]int)
	for col, val := range curr {
		if col < 0 || col >= len(row) {
			continue
		}
		switch row[col] {
		case "^":
			next[col-1] += val
			next[col+1] += val
		default:
			next[col] += val
		}
	}
	return next
}

// propagateSplits moves beams from the current row to the next one and counts splits.
func propagateSplits(curr map[int]int, row []string) (map[int]int, int) {
	next := make(map[int]int)
	splitCount := 0
	for col, val := range curr {
		if col < 0 || col >= len(row) {
			continue
		}
		cell := row[col]
		switch cell {
		case "^":
			splitCount++
			next[col-1] += val
			next[col+1] += val
		default:
			next[col] += val
		}
	}
	return next, splitCount
}

// CountBeamSplits calculates how many times a tachyon beam is split within the manifold.
func CountBeamSplits(d *utils.Data) int {
	grid := d.AsGrid()
	start, ok := findStart(grid)
	if !ok {
		return 0
	}
	current := map[int]int{start.c: 1}
	totalSplits := 0
	for rowIdx := start.r + 1; rowIdx < len(grid) && len(current) > 0; rowIdx++ {
		var splits int
		current, splits = propagateSplits(current, grid[rowIdx])
		totalSplits += splits
	}
	return totalSplits
}

// NumberOfTachyonParticleTimelines returns the total number of distinct timelines that reach the bottom.
func NumberOfTachyonParticleTimelines(d *utils.Data) int {
	grid := d.AsGrid()
	start, ok := findStart(grid)
	if !ok {
		return 0
	}
	current := map[int]int{start.c: 1}
	for rowIdx := start.r + 1; rowIdx < len(grid) && len(current) > 0; rowIdx++ {
		current = propagateCounts(current, grid[rowIdx])
	}
	total := 0
	for _, count := range current {
		total += count
	}
	return total
}

type pos struct{ r, c int }

// findStart returns the position of the start cell ("S") and a boolean indicating success.
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
