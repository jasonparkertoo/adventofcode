package day12

type PlotPoint struct {
	x, y int
}

func CalculateTotalPrice(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
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