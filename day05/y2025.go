package day05

import (
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

type db struct {
	ranges [][]int64
	ids    []int64
}

func day42025Transformer(lines []string) any {
	db := &db{
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
			db.ids = append(db.ids, n)
		} else {
			tokens := strings.Split(l, "-")
			left, _ := strconv.ParseInt(tokens[0], 10, 64)
			right, _ := strconv.ParseInt(tokens[1], 10, 64)
			ranges := []int64{left, right}
			db.ranges = append(db.ranges, ranges)
		}
	}
	return db
}

func CountFreshIds(d *utils.Data) int {
	db := d.TransformData(day42025Transformer).(*db)

	count := 0
	for _, id := range db.ids {
		for _, r := range db.ranges {
			if id < r[0] || id > r[1] {
				continue
			}
			count++
			break
		}
	}
	return count
}
