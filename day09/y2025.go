package day09

import (
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"adventofcode.dev/utils"
)

// point represents a tile’s coordinates on the grid.
// `row` is the zero‑based vertical index, `col` is the
// zero‑based horizontal index.
type point struct {
	row int
	col int
}

// dataTransformer turns each input line in the form
// "col,row" into a point struct. The function returns
// a slice of points that can be type‑asserted to
// []point by callers.
func dataTransformer(lines []string) any {
	pts := make([]point, 0, len(lines))
	for _, l := range lines {
		p := strings.Split(l, ",")
		col, _ := strconv.Atoi(p[0])
		row, _ := strconv.Atoi(p[1])
		pts = append(pts, point{row: row, col: col})
	}
	return pts
}

// abs returns the absolute value of an integer.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// maxInt returns the greater of two ints.
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// findLargestRectangle examines every pair of red tiles
// and returns the area of the largest rectangle that
// can be formed with those tiles as opposite corners.
// The rectangle’s width and height include both corner
// tiles (hence the +1 in each dimension).
func findLargestRectangle(d *utils.Data) int {
	pts := d.TransformData(dataTransformer).([]point)

	maxArea := 0
	for i := range pts {
		for j := i + 1; j < len(pts); j++ {
			x1, y1 := pts[i].col, pts[i].row
			x2, y2 := pts[j].col, pts[j].row

			width := abs(x2-x1) + 1
			height := abs(y2-y1) + 1
			area := width * height

			maxArea = maxInt(maxArea, area)
		}
	}
	return maxArea
}

// boundingBox returns the minimum and maximum row and column
// values that enclose all provided points.
func boundingBox(pts []point) (minR, maxR, minC, maxC int) {
	minR, maxR = pts[0].row, pts[0].row
	minC, maxC = pts[0].col, pts[0].col

	for _, p := range pts {
		minR = min(minR, p.row)
		maxR = max(maxR, p.row)
		minC = min(minC, p.col)
		maxC = max(maxC, p.col)
	}
	return
}

// markCrossings sets entries in the valid slice to true for each
// column index that appears in crossings, offset by minCol.
func markCrossings(valid []bool, crossings []int, minCol int) {
	for _, x := range crossings {
		idx := x - minCol
		if idx >= 0 && idx < len(valid) {
			valid[idx] = true
		}
	}
}

// buildRowRanges constructs a map from row indices to a slice of
// inclusive [start, end] column intervals that are inside the polygon.
func buildRowRanges(poly []point, minRow, maxRow, minCol, maxCol int) map[int][][2]int {
	rows := make(map[int][][2]int)

	for row := minRow; row <= maxRow; row++ {
		crossings := collectCrossings(poly, row)

		if len(crossings) == 0 {
			if pointInPolygon((minCol+maxCol)/2, row, poly) {
				rows[row] = [][2]int{{minCol, maxCol}}
			}
			continue
		}

		sort.Ints(crossings)
		crossings = uniqueInts(crossings)

		valid := make([]bool, maxCol-minCol+1)

		markCrossings(valid, crossings, minCol)
		fillInterior(valid, crossings, row, minCol, maxCol, poly)

		if intervals := extractIntervals(valid, minCol); len(intervals) > 0 {
			rows[row] = intervals
		}
	}

	return rows
}

// collectCrossings returns a slice of column indices where a horizontal
// or vertical edge of the polygon intersects the given row.
func collectCrossings(poly []point, row int) []int {
	var xs []int

	for i := range poly {
		p1 := poly[i]
		p2 := poly[(i+1)%len(poly)]

		// Horizontal edge
		if p1.row == row && p2.row == row {
			a, b := p1.col, p2.col
			if a > b {
				a, b = b, a
			}
			for x := a; x <= b; x++ {
				xs = append(xs, x)
			}
			continue
		}

		// Vertical edge
		if p1.col == p2.col {
			y1, y2 := p1.row, p2.row
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			if y1 <= row && row <= y2 {
				xs = append(xs, p1.col)
			}
		}
	}

	return xs
}

// pointInPolygon returns true if the point (x,y) is inside or on
// the boundary of the polygon.
func pointInPolygon(x, y int, polygon []point) bool {
	n := len(polygon)
	inside := false

	for i := range n {
		p1 := polygon[i]
		p2 := polygon[(i+1)%n]

		// On vertical edge
		if p1.col == p2.col && p1.col == x {
			y1, y2 := p1.row, p2.row
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			if y >= y1 && y <= y2 {
				return true
			}
		}

		// On horizontal edge
		if p1.row == p2.row && p1.row == y {
			x1, x2 := p1.col, p2.col
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if x >= x1 && x <= x2 {
				return true
			}
		}

		// Ray-casting test
		if (p1.row > y) != (p2.row > y) {
			xIntersect := p1.col +
				(y-p1.row)*(p2.col-p1.col)/(p2.row-p1.row)
			if xIntersect < x {
				inside = !inside
			}
		}
	}

	return inside
}

// fillInterior marks the interior of a polygon row by setting
// entries in valid to true for columns between crossings that lie
// within the polygon. It handles segments before the first crossing,
// between consecutive crossings, and after the last crossing.
func fillInterior(valid []bool, crossings []int, row, minCol, maxCol int, poly []point) {
	// Before first
	if crossings[0] > minCol {
		mid := (minCol + crossings[0] - 1) / 2
		if pointInPolygon(mid, row, poly) {
			for x := minCol; x < crossings[0]; x++ {
				valid[x-minCol] = true
			}
		}
	}

	// Between crossings
	for i := 0; i < len(crossings)-1; i++ {
		if crossings[i+1]-crossings[i] > 1 {
			start := crossings[i] + 1
			end := crossings[i+1] - 1
			mid := (start + end) / 2
			if pointInPolygon(mid, row, poly) {
				for x := start; x <= end; x++ {
					valid[x-minCol] = true
				}
			}
		}
	}

	// After last
	last := crossings[len(crossings)-1]
	if last < maxCol {
		mid := (last + 1 + maxCol) / 2
		if pointInPolygon(mid, row, poly) {
			for x := last + 1; x <= maxCol; x++ {
				valid[x-minCol] = true
			}
		}
	}
}

// extractIntervals converts a slice of bools indicating validity
// into a slice of inclusive [start, end] column intervals.
func extractIntervals(valid []bool, minCol int) [][2]int {
	var out [][2]int
	start := -1

	for i, ok := range valid {
		col := i + minCol
		if ok && start == -1 {
			start = col
		} else if !ok && start != -1 {
			out = append(out, [2]int{start, col - 1})
			start = -1
		}
	}

	if start != -1 {
		out = append(out, [2]int{start, minCol + len(valid) - 1})
	}

	return out
}

// rectangleIsValid checks whether the axis‑aligned rectangle defined by
// points p1 and p2 lies entirely within the provided rowRanges map,
// which contains per‑row intervals of valid columns.
func rectangleIsValid(p1, p2 point, rows map[int][][2]int) bool {
	top := min(p1.row, p2.row)
	bot := max(p1.row, p2.row)
	left := min(p1.col, p2.col)
	right := max(p1.col, p2.col)

	for r := top; r <= bot; r++ {
		ints, ok := rows[r]
		if !ok {
			return false
		}

		i := sort.Search(len(ints), func(i int) bool {
			return ints[i][1] >= left
		})
		if i >= len(ints) || ints[i][0] > left || ints[i][1] < right {
			return false
		}
	}
	return true
}

// rectangleArea returns the area of an axis‑aligned rectangle
// defined by points a and b, including both corner tiles.
func rectangleArea(a, b point) int {
	return (abs(a.row-b.row) + 1) * (abs(a.col-b.col) + 1)
}

// uniqueInts removes consecutive duplicate integers from xs.
// It assumes that the slice is sorted.
func uniqueInts(xs []int) []int {
	out := xs[:1]
	for i := 1; i < len(xs); i++ {
		if xs[i] != xs[i-1] {
			out = append(out, xs[i])
		}
	}
	return out
}

// findLargestRectangleOfAny - concurrent version for maximum speed
func findLargestRectangleOfAny(d *utils.Data) int {
	points := d.TransformData(dataTransformer).([]point)
	if len(points) < 2 {
		return 0
	}

	minRow, maxRow, minCol, maxCol := boundingBox(points)
	rowRanges := buildRowRanges(points, minRow, maxRow, minCol, maxCol)

	// Create candidate pairs sorted by area (descending)
	type candidate struct {
		i, j int
		area int
	}
	
	candidates := make([]candidate, 0, len(points)*(len(points)-1)/2)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			area := rectangleArea(points[i], points[j])
			candidates = append(candidates, candidate{i, j, area})
		}
	}
	
	// Sort by area descending
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].area > candidates[j].area
	})
	
	// Process candidates concurrently
	numWorkers := runtime.GOMAXPROCS(0)
	var maxArea atomic.Int64
	var wg sync.WaitGroup
	
	// Work queue with early termination
	chunkSize := 1000 // Process in chunks for better early termination
	
	for start := 0; start < len(candidates); start += chunkSize {
		end := start + chunkSize
		if end > len(candidates) {
			end = len(candidates)
		}
		
		chunk := candidates[start:end]
		
		// Early termination: if best in chunk can't beat current max, skip rest
		if chunk[0].area <= int(maxArea.Load()) {
			break
		}
		
		wg.Add(numWorkers)
		
		for w := 0; w < numWorkers; w++ {
			workerID := w
			go func() {
				defer wg.Done()
				
				// Each worker processes every Nth item in the chunk
				for i := workerID; i < len(chunk); i += numWorkers {
					c := chunk[i]
					
					// Check if we can still improve
					currentMax := int(maxArea.Load())
					if c.area <= currentMax {
						continue
					}
					
					p1, p2 := points[c.i], points[c.j]
					
					if rectangleIsValidFast(p1, p2, rowRanges) {
						// Update max atomically
						for {
							currentMax = int(maxArea.Load())
							if c.area <= currentMax {
								break
							}
							if maxArea.CompareAndSwap(int64(currentMax), int64(c.area)) {
								break
							}
						}
					}
				}
			}()
		}
		
		wg.Wait()
	}

	return int(maxArea.Load())
}

// rectangleIsValidFast - optimized validation (concurrent-safe, read-only)
func rectangleIsValidFast(p1, p2 point, rows map[int][][2]int) bool {
	top := min(p1.row, p2.row)
	bot := max(p1.row, p2.row)
	left := min(p1.col, p2.col)
	right := max(p1.col, p2.col)

	for r := top; r <= bot; r++ {
		ints, ok := rows[r]
		if !ok {
			return false
		}

		// Binary search
		lo, hi := 0, len(ints)
		for lo < hi {
			mid := (lo + hi) >> 1
			if ints[mid][1] < left {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		
		if lo >= len(ints) || ints[lo][0] > left || ints[lo][1] < right {
			return false
		}
	}
	return true
}