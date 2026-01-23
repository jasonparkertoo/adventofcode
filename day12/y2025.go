package day12

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type Present struct {
	Shape [][]byte
	Index int
}

type Region struct {
	Width    int
	Height   int
	Presents []int
}

type presentItem struct {
	rotations [][][]byte
	size      int
}

// CountValidRegions counts the number of valid regions that can fit all presents
func CountValidRegions(d *utils.Data) int {
	lines := d.Lines()

	presents := parsePresents(lines, 0, len(lines))
	regions := parseRegions(lines, 0, len(lines))

	// Pre-compute all rotations for each present
	presentRotations := make([][][][]byte, len(presents))
	for i, present := range presents {
		presentRotations[i] = getRotations(present.Shape)
	}

	count := 0
	for _, region := range regions {
		if canFitAllPresents(region, presents, presentRotations) {
			count++
		}
	}

	return count
}

// parsePresents parses presents from the input lines
func parsePresents(lines []string, start, end int) []Present {
	var presents []Present
	i := start

	for i < end && (lines[i] == "" || strings.TrimSpace(lines[i]) == "") {
		i++
	}

	for i < end {
		line := lines[i]

		if strings.Contains(line, ":") && !strings.Contains(line, "x") {
			before, _, _ := strings.Cut(line, ":")
			presentIndex, err := strconv.Atoi(strings.TrimSpace(before))
			if err != nil {
				i++
				continue
			}

			var shape [][]byte
			i++
			for i < end && lines[i] != "" && !strings.Contains(lines[i], ":") {
				row := []byte(lines[i])
				shape = append(shape, row)
				i++
			}
			if len(shape) > 0 {
				presents = append(presents, Present{Shape: shape, Index: presentIndex})
			}
		} else {
			i++
		}
	}

	return presents
}

// parseRegions parses region definitions from the input lines
// Each region definition has the format "widthxheight: count1 count2 ..."
// where width and height are positive integers, and counts represent
// how many of each present type can fit in that region
func parseRegions(lines []string, start, end int) []Region {
	var regions []Region
	i := start

	for i < end {
		if strings.Contains(lines[i], "x") && strings.Contains(lines[i], ":") {
			break
		}
		i++
	}

	for i < end {
		line := lines[i]
		if line == "" || strings.TrimSpace(line) == "" {
			i++
			continue
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			i++
			continue
		}

		dimensions := strings.Split(parts[0], "x")
		if len(dimensions) != 2 {
			i++
			continue
		}

		width, err1 := strconv.Atoi(strings.TrimSpace(dimensions[0]))
		height, err2 := strconv.Atoi(strings.TrimSpace(dimensions[1]))
		if err1 != nil || err2 != nil {
			i++
			continue
		}

		counts := strings.Fields(parts[1])
		var presentCounts []int
		for _, count := range counts {
			num, err := strconv.Atoi(strings.TrimSpace(count))
			if err == nil {
				presentCounts = append(presentCounts, num)
			}
		}

		if len(presentCounts) > 0 {
			regions = append(regions, Region{Width: width, Height: height, Presents: presentCounts})
		}
		i++
	}

	return regions
}

// canFitAllPresents determines whether all presents can fit into a given region
// by trying to place them using backtracking with pruning.
// It returns true if a valid arrangement exists, false otherwise.
func canFitAllPresents(region Region, presents []Present, presentRotations [][][][]byte) bool {
	// Early exit: check if total area is sufficient
	totalArea := 0
	for i, count := range region.Presents {
		if i >= len(presents) {
			return false
		}
		totalArea += countShapeCells(presents[i].Shape) * count
	}
	if totalArea > region.Width*region.Height {
		return false
	}

	grid := make([][]byte, region.Height)
	for i := range grid {
		grid[i] = make([]byte, region.Width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// Build list of presents to place, sorted by size (largest first for better pruning)
	var toPlace []presentItem

	for i, count := range region.Presents {
		if i >= len(presents) {
			return false
		}
		for range count {
			toPlace = append(toPlace, presentItem{
				rotations: presentRotations[i],
				size:      countShapeCells(presents[i].Shape),
			})
		}
	}

	// Sort by size descending (place larger pieces first)
	for i := 0; i < len(toPlace)-1; i++ {
		for j := i + 1; j < len(toPlace); j++ {
			if toPlace[j].size > toPlace[i].size {
				toPlace[i], toPlace[j] = toPlace[j], toPlace[i]
			}
		}
	}

	return placePresents(toPlace, grid, 0)
}

// countShapeCells counts the number of cells that are '#' in the given shape
func countShapeCells(shape [][]byte) int {
	count := 0
	for _, row := range shape {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}

// placePresents attempts to place all presents in the given grid using backtracking.
// It tries different rotations and positions for each present, and returns true if
// a valid arrangement exists, false otherwise.
func placePresents(toPlace []presentItem, grid [][]byte, index int) bool {
	if index >= len(toPlace) {
		return true
	}

	rotations := toPlace[index].rotations

	// Try to place in a more strategic order - start from top-left
	for _, shape := range rotations {
		if len(shape) == 0 {
			continue
		}

		maxY := len(grid) - len(shape) + 1
		maxX := len(grid[0]) - len(shape[0]) + 1

		for y := range maxY {
			for x := range maxX {
				if canPlaceShapeFast(shape, grid, x, y) {
					placeShape(shape, grid, x, y)

					if placePresents(toPlace, grid, index+1) {
						return true
					}

					removeShape(shape, grid, x, y)
				}
			}
		}
	}

	return false
}

// canPlaceShapeFast checks if a shape can be placed at the given position in the grid.
// It returns true if the shape fits entirely within the grid boundaries and does not
// overlap with any existing placed presents ('#'), and false otherwise.
func canPlaceShapeFast(shape [][]byte, grid [][]byte, startX, startY int) bool {
	for y, row := range shape {
		gridY := startY + y
		if gridY >= len(grid) {
			return false
		}
		for x, cell := range row {
			if cell == '#' {
				gridX := startX + x
				if gridX >= len(grid[0]) || grid[gridY][gridX] != '.' {
					return false
				}
			}
		}
	}
	return true
}

// placeShape places a shape onto the grid at the specified starting position.
// It assumes the shape can be placed without overlapping existing placed presents.
func placeShape(shape [][]byte, grid [][]byte, startX, startY int) {
	for y, row := range shape {
		gridY := startY + y
		for x, cell := range row {
			if cell == '#' {
				grid[gridY][startX+x] = '#'
			}
		}
	}
}

// removeShape removes a shape from the grid at the specified starting position.
// It assumes the shape was previously placed and only '.' cells are overwritten.
func removeShape(shape [][]byte, grid [][]byte, startX, startY int) {
	for y, row := range shape {
		gridY := startY + y
		for x, cell := range row {
			if cell == '#' {
				grid[gridY][startX+x] = '.'
			}
		}
	}
}
// getRotations returns all unique rotations and reflections of a given shape.
// It includes the original shape, all 90-degree rotations, and all flipped versions
// of those rotations, ensuring no duplicates are returned.
func getRotations(shape [][]byte) [][][]byte {
	var rotations [][][]byte
	seen := make(map[string]bool)

	addUnique := func(s [][]byte) {
		key := shapeKey(s)
		if !seen[key] {
			seen[key] = true
			rotations = append(rotations, s)
		}
	}

	// Original
	addUnique(copyShape(shape))

	// Rotations
	r90 := rotate90(shape)
	addUnique(r90)

	r180 := rotate90(r90)
	addUnique(r180)

	r270 := rotate90(r180)
	addUnique(r270)

	// Flipped
	flipped := flip(shape)
	addUnique(flipped)

	fr90 := rotate90(flipped)
	addUnique(fr90)

	fr180 := rotate90(fr90)
	addUnique(fr180)

	fr270 := rotate90(fr180)
	addUnique(fr270)

	return rotations
}

// shapeKey generates a unique string key for a shape by concatenating all rows
// and separating them with a pipe character.
func shapeKey(shape [][]byte) string {
	var sb strings.Builder
	for _, row := range shape {
		sb.Write(row)
		sb.WriteByte('|')
	}
	return sb.String()
}

// copyShape creates a deep copy of a given shape.
func copyShape(shape [][]byte) [][]byte {
	copied := make([][]byte, len(shape))
	for i, row := range shape {
		copied[i] = make([]byte, len(row))
		copy(copied[i], row)
	}
	return copied
}

// rotate90 rotates a shape 90 degrees clockwise.
// The returned shape has dimensions swapped (rows become columns and vice versa).
func rotate90(shape [][]byte) [][]byte {
	rows := len(shape)
	if rows == 0 {
		return shape
	}
	cols := len(shape[0])
	if cols == 0 {
		return shape
	}

	rotated := make([][]byte, cols)
	for i := range rotated {
		rotated[i] = make([]byte, rows)
	}

	for y := range rows {
		for x := 0; x < len(shape[y]); x++ {
			rotated[x][rows-1-y] = shape[y][x]
		}
	}

	return rotated
}

// flip creates a horizontal reflection of a shape.
// Each row is reversed, effectively flipping the shape horizontally.
func flip(shape [][]byte) [][]byte {
	rows := len(shape)
	if rows == 0 {
		return shape
	}
	cols := len(shape[0])
	if cols == 0 {
		return shape
	}

	flipped := make([][]byte, rows)
	for i := range flipped {
		flipped[i] = make([]byte, cols)
	}

	for y := range rows {
		for x := 0; x < len(shape[y]); x++ {
			flipped[y][cols-1-x] = shape[y][x]
		}
	}

	return flipped
}
