package day10

import "adventofcode.dev/utils"

// Point represents a coordinate in the grid with row R and column C.
type Point struct{ R, C int }

// LavaTrails holds the parsed elevation grid and its dimensions.
type LavaTrails struct {
	grid       [][]int
	rows, cols int
}

// ToLavaTrails parses the raw input data into a LavaTrails structure.
func ToLavaTrails(d *utils.Data) *LavaTrails {
	lines := d.Lines()

	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]int, rows)
	for r := range grid {
		grid[r] = make([]int, cols)
		for c := range grid[r] {
			grid[r][c] = int(lines[r][c] - '0')
		}
	}

	return &LavaTrails{
		grid: grid,
		rows: rows,
		cols: cols,
	}
}

// cloneSet creates a shallow copy of a set of Points.
func cloneSet(src map[Point]struct{}) map[Point]struct{} {
	dst := make(map[Point]struct{}, len(src))
	for k := range src {
		dst[k] = struct{}{}
	}
	return dst
}

// height returns the elevation at a given point in the grid.
func height(t *LavaTrails, p Point) int {
	return t.grid[p.R][p.C]
}

// neighbors returns the orthogonal neighbors of a point that are within bounds.
func neighbors(t *LavaTrails, p Point) []Point {
	r, c := p.R, p.C
	cands := [...]Point{
		{r - 1, c}, {r + 1, c},
		{r, c - 1}, {r, c + 1},
	}
	out := make([]Point, 0, 4)
	for _, x := range cands {
		if x.R >= 0 && x.C >= 0 && x.R < t.rows && x.C < t.cols {
			out = append(out, x)
		}
	}
	return out
}

// reachableNines returns the set of points at height 9 that can be reached from p
// following strictly increasing elevation steps. Memoization is used to avoid
// recomputation. A fresh clone of the result is returned to prevent callers
// from mutating the memoized data.
func reachableNines(t *LavaTrails, p Point, memo map[Point]map[Point]struct{}) map[Point]struct{} {
	// READ: return clone so caller can't mutate memo
	if got, ok := memo[p]; ok {
		return cloneSet(got)
	}

	h := height(t, p)
	var res map[Point]struct{}
	if h == 9 {
		res = map[Point]struct{}{p: {}}
	} else {
		res = make(map[Point]struct{})
		for _, n := range neighbors(t, p) {
			if height(t, n) == h+1 {
				child := reachableNines(t, n, memo)
				for q := range child {
					res[q] = struct{}{}
				}
			}
		}
	}

	// WRITE: freeze a clone to memo
	frozen := cloneSet(res)
	memo[p] = frozen
	// RETURN: return a fresh clone so caller can't mutate our stored one
	return cloneSet(frozen)
}

// countPaths returns the number of strictly increasing paths from p to any height-9 cell.
// Memoization is employed to cache results for already visited points.
func countPaths(t *LavaTrails, p Point, memo map[Point]int) int {
	if val, exists := memo[p]; exists {
		return val
	}

	h := height(t, p)
	var count int
	if h == 9 {
		count = 1
	} else {
		count = 0
		for _, n := range neighbors(t, p) {
			if height(t, n) == h+1 {
				count += countPaths(t, n, memo)
			}
		}
	}

	memo[p] = count
	return count
}

// TotalTrailheadRating computes the sum of path counts for all trailheads (height 0 cells).
func TotalTrailheadRating(d *utils.Data) int {
	t := ToLavaTrails(d)

	sum := 0
	for r := 0; r < t.rows; r++ {
		for c := 0; c < t.cols; c++ {
			p := Point{r, c}
			if height(t, p) == 0 {
				memo := make(map[Point]int)
				sum += countPaths(t, p, memo)
			}
		}
	}
	return sum
}

// TotalTrailheadScore computes the total number of distinct height‑9 cells reachable
// from all trailheads (height 0 cells).
func TotalTrailheadScore(d *utils.Data) int {
	t := ToLavaTrails(d)

	memo := make(map[Point]map[Point]struct{})
	sum := 0
	for r := 0; r < t.rows; r++ {
		for c := 0; c < t.cols; c++ {
			p := Point{r, c}
			if height(t, p) == 0 {
				sum += len(reachableNines(t, p, memo))
			}
		}
	}
	return sum
}
