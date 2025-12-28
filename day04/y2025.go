package day04

import (
	"strings"

	"adventofcode.dev/utils"
)

func Dataformatter(s []string) any {
	out := make([][]string, len(s))
	for i, str := range s {
		out[i] = strings.Split(str, "")
	}
	return out
}

const PaperRoll = "@"

func CountAccessible(d *utils.Data) int {
	grid := d.TransformData(Dataformatter).([][]string)

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
