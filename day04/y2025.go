package day04

import (
	"adventofcode.dev/utils"
)

const PaperRoll = "@"

// CountAccessible counts the number of PaperRoll cells that are not surrounded by at least four other PaperRoll cells.
// It iterates through the grid and uses the helper count to determine neighbor counts.
// If the neighbor count is less than four, the cell is considered accessible.
func CountAccessible(d *utils.Data) int {
	grid := d.AsGrid()

	numAccessible := 0
	for row := range len(grid) {
		for col := range len(grid[row]) {
			if grid[row][col] != PaperRoll {
				continue
			}
			c := count(row, col, grid)
			if c < 4 {
				numAccessible++
			}
		}
	}
	return numAccessible
}

// count returns the number of adjacent PaperRoll cells around the specified row and column.
// It checks all eight neighboring positions and counts those that contain a PaperRoll symbol.
func count(row, column int, grid [][]string) int {
	count := 0

	directions := []struct {
		dRow, dCol int
	}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := row + dir.dRow
		newCol := column + dir.dCol
		// check for out of bounds condition and values that are not PaperRoll
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[row]) || PaperRoll != grid[newRow][newCol] {
			continue
		}
		count++
	}

	return count
}

// CountRemovable removes all PaperRoll cells that are not surrounded by at least four other PaperRoll cells,
// iteratively applying the removal until no more cells qualify.
// It returns the total number of cells removed.
func CountRemovable(d *utils.Data) int {
	grid := d.AsGrid()
	removed := 0

	for {
		var toRemove [][2]int
		for r := range grid {
			for c := range grid[r] {
				if grid[r][c] != PaperRoll {
					continue
				}
				if count(r, c, grid) < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = "."
		}
		removed += len(toRemove)
	}

	return removed
}
