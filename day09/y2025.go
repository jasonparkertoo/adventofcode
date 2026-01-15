package day09

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type point struct{ row, col int }

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

// abs returns the absolute value of an integer
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findLargestRectangle(d *utils.Data) int {
	pts := d.TransformData(dataTransformer).([]point)

	maxArea := 0
	for i := range pts {
		for j := i + 1; j < len(pts); j++ {
			x1, y1 := pts[i].col, pts[i].row
			x2, y2 := pts[j].col, pts[j].row

			// Calculate width and height of the rectangle
			width := abs(x2-x1) + 1
			height := abs(y2-y1) + 1

			// Calculate area and update maximum if needed
			area := width * height
			
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}
