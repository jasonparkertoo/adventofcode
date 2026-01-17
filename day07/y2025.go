package day07

import (
	"adventofcode.dev/utils"
)

// propagateCounts moves beams from the current row to the next one.
// It iterates over every column that currently holds a beam and, depending
// on the cell contents, determines where those beams will appear in the
// next row.
//
// Parameters:
//
//	curr: a map that associates a column index with the number of beams
//	      currently present in that column.
//	row: a slice of strings representing the current row of the grid.
//
// Returns:
//
//	A map that associates column indices with the number of beams that
//	arrive in those columns on the next row.  A splitter ("^") causes a
//	beam to split into two beams, one to the left and one to the right.
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

// propagateSplits moves beams from the current row to the next one
// and counts how many times a splitter is encountered.  A splitter
// causes a beam to split into two, thereby increasing the split count.
//
// Parameters:
//
//	curr: a map of current beam counts per column.
//	row: the current row of the grid.
//
// Returns:
//
//	A map of beam counts for the next row, just like propagateCounts,
//	and an integer indicating how many splitters were encountered in
//	this transition.
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

// CountBeamSplits calculates the total number of times a tachyon beam
// is split within the manifold.  The function simulates the beam
// propagation row by row, counting split events each time a splitter
// is encountered.
//
// Parameters:
//
//	d: a pointer to utils.Data which provides the grid data for the
//	   problem.
//
// Returns:
//
//	The total number of split events that occur as the beam travels
//	from the start cell to the bottom of the grid.  If the start cell
//	cannot be found, the function returns 0.
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

// NumberOfTachyonParticleTimelines returns the total number of distinct
// timelines that reach the bottom of the grid.  Each timeline corresponds
// to a unique sequence of beam positions from the start cell to the last
// row.  The function performs a breadth‑first simulation of beam
// propagation, accumulating the counts for each column in the last row.
//
// Parameters:
//
//	d: a pointer to utils.Data containing the grid layout.
//
// Returns:
//
//	The sum of all beam counts in the last row, which equals the number
//	of distinct timelines.  If the start cell cannot be found, the
//	function returns 0.
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

// findStart locates the start cell ("S") in the grid.
//
// Parameters:
//
//	grid: a 2‑D slice of strings representing the grid.
//
// Returns:
//
//	A pos struct containing the row and column indices of the start cell,
//	and a boolean that is true if the start cell was found, or false
//	otherwise.
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
