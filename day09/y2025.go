package day09

import (
	"sort"
	"strconv"
	"strings"

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
// "col,row" into a point struct. The function
// returns a slice of points that can be type‑asserted
// to []point by callers.
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

// findLargestRectangleOfAny computes the largest rectangle
// that can be formed with red tiles as opposite corners,
// but the rectangle may only contain tiles that are
// red or green. The green tiles are the straight lines
// between consecutive red tiles and all interior tiles
// inside that loop.
func findLargestRectangleOfAny(d *utils.Data) int {
	pts := d.TransformData(dataTransformer).([]point)
	if len(pts) < 2 {
		return 0
	}

	n := len(pts)
	
	// Find bounding box
	minRow, maxRow := pts[0].row, pts[0].row
	minCol, maxCol := pts[0].col, pts[0].col
	for _, p := range pts {
		if p.row < minRow { minRow = p.row }
		if p.row > maxRow { maxRow = p.row }
		if p.col < minCol { minCol = p.col }
		if p.col > maxCol { maxCol = p.col }
	}

	// For each row, determine which columns are valid (red or green)
	rowRanges := make(map[int][][2]int)
	
	for row := minRow; row <= maxRow; row++ {
		// Collect x-coordinates where polygon edges cross this row
		crossings := make([]int, 0)
		
		for i := range n {
			p1 := pts[i]
			p2 := pts[(i+1)%n]
			
			// Horizontal edge on this row
			if p1.row == row && p2.row == row {
				start, end := p1.col, p2.col
				if start > end {
					start, end = end, start
				}
				for x := start; x <= end; x++ {
					crossings = append(crossings, x)
				}
				
			} else if p1.col == p2.col { // Vertical edge crossing this row (not at endpoints for interior)
				y1, y2 := p1.row, p2.row
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				if y1 < row && row < y2 {
					crossings = append(crossings, p1.col)
				} else if row == y1 || row == y2 {
					// On vertical edge endpoint
					crossings = append(crossings, p1.col)
				}
			}
		}
		
		if len(crossings) == 0 {
			// Check if row is entirely inside polygon
			midX := (minCol + maxCol) / 2
			if pointInPolygon(midX, row, pts) {
				rowRanges[row] = [][2]int{{minCol, maxCol}}
			}
			continue
		}
		
		// Sort and deduplicate crossings
		sort.Ints(crossings)
		unique := crossings[:1]
		for i := 1; i < len(crossings); i++ {
			if crossings[i] != crossings[i-1] {
				unique = append(unique, crossings[i])
			}
		}
		crossings = unique
		
		// Determine valid intervals
		validCols := make([]bool, maxCol-minCol+1)
		
		// Mark crossing points
		for _, x := range crossings {
			if x >= minCol && x <= maxCol {
				validCols[x-minCol] = true
			}
		}
		
		// Check intervals between crossings
		// Before first crossing
		if crossings[0] > minCol {
			midX := (minCol + crossings[0] - 1) / 2
			if pointInPolygon(midX, row, pts) {
				for x := minCol; x < crossings[0]; x++ {
					validCols[x-minCol] = true
				}
			}
		}
		
		// Between crossings
		for i := 0; i < len(crossings)-1; i++ {
			if crossings[i+1]-crossings[i] > 1 {
				midX := (crossings[i] + crossings[i+1]) / 2
				if pointInPolygon(midX, row, pts) {
					for x := crossings[i] + 1; x < crossings[i+1]; x++ {
						validCols[x-minCol] = true
					}
				}
			}
		}
		
		// After last crossing
		last := crossings[len(crossings)-1]
		if last < maxCol {
			midX := (last + 1 + maxCol) / 2
			if pointInPolygon(midX, row, pts) {
				for x := last + 1; x <= maxCol; x++ {
					validCols[x-minCol] = true
				}
			}
		}
		
		// Convert to intervals
		var intervals [][2]int
		start := -1
		for i, valid := range validCols {
			col := i + minCol
			if valid && start == -1 {
				start = col
			} else if !valid && start != -1 {
				intervals = append(intervals, [2]int{start, col - 1})
				start = -1
			}
		}
		if start != -1 {
			intervals = append(intervals, [2]int{start, maxCol})
		}
		
		if len(intervals) > 0 {
			rowRanges[row] = intervals
		}
	}

	// Helper to check if a rectangle is valid
	isRectangleValid := func(r1, r2, c1, c2 int) bool {
		top, bottom := r1, r2
		left, right := c1, c2
		if top > bottom {
			top, bottom = bottom, top
		}
		if left > right {
			left, right = right, left
		}
		
		for row := top; row <= bottom; row++ {
			intervals, ok := rowRanges[row]
			if !ok {
				return false
			}
			
			// Binary search for interval containing left
			idx := sort.Search(len(intervals), func(i int) bool {
				return intervals[i][1] >= left
			})
			
			if idx >= len(intervals) || intervals[idx][0] > left || intervals[idx][1] < right {
				return false
			}
		}
		return true
	}

	// Find maximum area rectangle
	maxArea := 0
	
	// Sort points for better cache locality
	sortedPts := make([]point, len(pts))
	copy(sortedPts, pts)
	sort.Slice(sortedPts, func(i, j int) bool {
		if sortedPts[i].row == sortedPts[j].row {
			return sortedPts[i].col < sortedPts[j].col
		}
		return sortedPts[i].row < sortedPts[j].row
	})
	
	// Check all pairs of red points
	for i := 0; i < len(sortedPts); i++ {
		p1 := sortedPts[i]
		
		for j := i + 1; j < len(sortedPts); j++ {
			p2 := sortedPts[j]
			
			height := abs(p2.row-p1.row) + 1
			width := abs(p2.col-p1.col) + 1
			area := width * height
			
			if area <= maxArea {
				continue
			}
			
			if isRectangleValid(p1.row, p2.row, p1.col, p2.col) {
				maxArea = area
			}
		}
	}
	
	return maxArea
}

// pointInPolygon returns true if point (x,y) is inside the polygon
func pointInPolygon(x, y int, polygon []point) bool {
	n := len(polygon)
	inside := false
	
	for i := 0; i < n; i++ {
		p1 := polygon[i]
		p2 := polygon[(i+1)%n]
		
		// Check if point is on edge
		if p1.col == p2.col && p1.col == x {
			// Vertical edge
			y1, y2 := p1.row, p2.row
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			if y >= y1 && y <= y2 {
				return true
			}
		} else if p1.row == p2.row && p1.row == y {
			// Horizontal edge
			x1, x2 := p1.col, p2.col
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if x >= x1 && x <= x2 {
				return true
			}
		}
		
		// Ray casting intersection test
		if (p1.row > y) != (p2.row > y) {
			// Edge crosses horizontal line at y
			xintersect := p1.col + (y-p1.row)*(p2.col-p1.col)/(p2.row-p1.row)
			
			if xintersect < x {
				inside = !inside
			}
		}
	}
	
	return inside
}