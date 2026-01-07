package day07

import (
	"strings"

	"adventofcode.dev/utils"
)

// DataTransformer converts each line of the input into a slice of single-character strings.
func DataTransformer(lines []string) any {
	out := make([][]string, len(lines))
	for i, row := range lines {
		out[i] = strings.Split(row, "")
	}
	return out
}

// CountBeamSplits calculates how many times a tachyon beam is split within the manifold.
// It performs an iterative depth‑first traversal starting from 'S', propagating beams downward and splitting at '^'.
// Each splitter is counted only once regardless of how many beams reach it.
func CountBeamSplits(d *utils.Data) int {
	// Convert raw input into a 2D grid of characters.
	grid := d.AsGrid()

	// Locate the start position marked by 'S'.
	var sr, sc int
	found := false
	for r, row := range grid {
		for c, cell := range row {
			if cell == "S" {
				sr, sc = r, c
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		return 0 // No starting point; nothing to count.
	}

	type pos struct{ r, c int }
	stack := []pos{{sr, sc}}      // Stack for DFS traversal.
	visited := make(map[pos]bool) // Track visited cells to avoid re‑processing.
	splitCount := 0

	for len(stack) > 0 {
		// Pop a position from the stack.
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[p] { // Skip if we've already processed this cell.
			continue
		}
		visited[p] = true

		// Bounds check – skip positions outside the grid.
		if p.r < 0 || p.r >= len(grid) || p.c < 0 || p.c >= len(grid[p.r]) {
			continue
		}

		cell := grid[p.r][p.c]
		if cell == "^" { // Splitter: count once and emit two new beams.
			splitCount++
			stack = append(stack, pos{p.r + 1, p.c - 1})
			stack = append(stack, pos{p.r + 1, p.c + 1})
		} else { // Empty space: continue downward.
			stack = append(stack, pos{p.r + 1, p.c})
		}
	}

	return splitCount
}
