package day05

import (
	"sort"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type db struct {
	ranges [][]int64
	ids    []int64
}

func DataTransformer(lines []string) any {
	d := &db{
		ranges: [][]int64{},
		ids:    []int64{},
	}

	foundDemarcation := false
	for _, l := range lines {
		if l == "" {
			foundDemarcation = true
			continue
		}
		if foundDemarcation {
			n, _ := strconv.ParseInt(l, 10, 64)
			d.ids = append(d.ids, n)
		} else {
			tokens := strings.Split(l, "-")
			left, _ := strconv.ParseInt(tokens[0], 10, 64)
			right, _ := strconv.ParseInt(tokens[1], 10, 64)
			d.ranges = append(d.ranges, []int64{left, right})
		}
	}
	return d
}

func MergeIntervals(d *db) *db {
	// Sort the ranges by their start value.
	sort.Slice(d.ranges, func(i, j int) bool {
		return d.ranges[i][0] < d.ranges[j][0]
	})

	merged := [][]int64{d.ranges[0]}
	for i := 1; i < len(d.ranges); i++ {
		last := merged[len(merged)-1]
		current := d.ranges[i]

		if current[0] <= last[1]+1 { // overlapping or contiguous
			// Extend the end of the last interval if needed.
			if current[1] > last[1] {
				merged[len(merged)-1][1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	d.ranges = merged
	return d
}

func NumberOfFreshRangeIds(d *utils.Data) int64 {
	data := MergeIntervals(d.TransformData(DataTransformer).(*db))
	total := int64(0)
	for _, r := range data.ranges {
		total += int64((r[1] - r[0]) + 1)
	}
	return total
}

func CountFreshIds(d *utils.Data) int {
	data := d.TransformData(DataTransformer).(*db)
	
	count := 0
	for _, id := range data.ids {
		for _, r := range data.ranges {
			if id >= r[0] && id <= r[1] {
				count++
				break
			}
		}
	}
	return count
}
