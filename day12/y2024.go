package day12

import "adventofcode.dev/utils"

// PlotPoint represents a coordinate in the plot grid.
// The fields are unexported because they are only used
// internally within the bfs algorithm.
type PlotPoint struct {
	x, y int
}

// CalculateTotalPrice returns the sum of the product of the
// area and perimeter for each contiguous region of identical
// plant types in the input grid.
//
// The input is provided as a utils.Data value, where each
// line represents a row of the plot. Each character in a
// line represents a plant type. The function treats the
// grid as a rectangular matrix. If the grid is empty,
// the function returns 0.
//
// The calculation works by performing a breadth‑first
// search for each unvisited cell. For each region the
// area is the number of cells and the perimeter is
// 4×area minus the number of shared edges between
// adjacent cells of the same type. The total price is
// the sum over all regions of area×perimeter.
func CalculateTotalPrice(d *utils.Data) int {
	grid := d.Lines()

	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range rows {
		visited[i] = make([]bool, cols)
	}

	totalPrice := 0

	for i := range rows {
		for j := range cols {
			if !visited[i][j] {
				plantType := grid[i][j]
				area, perimeter := bfs(grid, visited, i, j, plantType)
				totalPrice += area * perimeter
			}
		}
	}

	return totalPrice
}

// bfs performs a breadth‑first search starting from
// (startX, startY) over cells of the same plantType.
// It marks cells as visited and returns the area and
// perimeter of the discovered region.
// The function is unexported because it is an internal
// helper.
func bfs(grid []string, visited [][]bool, startX, startY int, plantType byte) (int, int) {
	rows, cols := len(grid), len(grid[0])
	directions := []PlotPoint{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	queue := []PlotPoint{{startX, startY}}
	visited[startX][startY] = true

	area := 0
	sharedEdges := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		area++

		// Check all four neighbors for shared edges
		for _, dir := range directions {
			nx, ny := current.x+dir.x, current.y+dir.y

			// If neighbor is within bounds and same type, count shared edge
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == plantType {
				sharedEdges++
				if !visited[nx][ny] {
					visited[nx][ny] = true
					queue = append(queue, PlotPoint{nx, ny})
				}
			}
		}
	}
	perimeter := 4*area - sharedEdges
	return area, perimeter
}
