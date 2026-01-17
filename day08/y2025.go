package day08

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"adventofcode.dev/utils"
)

// dataTransformer converts a slice of strings like "162,817,812"
// into a 2‑D slice of ints.
func dataTransformer(lines []string) any {
	out := make([][]int, len(lines))
	for i := range lines {
		row := make([]int, 3)
		parts := strings.Split(lines[i], ",")
		for j := range parts {
			n, _ := strconv.Atoi(parts[j])
			row[j] = n
		}
		out[i] = row
	}
	return out
}

// JunctionBox represents a point in 3‑D space.
type JunctionBox struct {
	X, Y, Z int
}

// Pair holds the squared distance between two junction boxes and
// the indices of those boxes in a slice.
type Pair struct {
	dist int64
	i, j int
}

// DSU implements a disjoint‑set‑union data structure.
type DSU struct {
	parent []int
	size   []int
}

// NewDSU creates a DSU for `n` elements, each initially its own set.
func NewDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := range p {
		p[i] = i
		s[i] = 1
	}
	return &DSU{parent: p, size: s}
}

// find returns the root of element x with path compression.
func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

// union merges the sets containing x and y. It returns true if the
// merge actually happened, i.e. the sets were separate.
func (d *DSU) union(x, y int) bool {
	rootX := d.find(x)
	rootY := d.find(y)

	if rootX == rootY {
		return false
	}

	if d.size[rootX] < d.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	d.parent[rootY] = rootX
	d.size[rootX] += d.size[rootY]
	return true
}

// generatePairs computes the squared Euclidean distance between every
// pair of junction boxes and returns a sorted slice of those pairs.
func generatePairs(boxes []*JunctionBox) []*Pair {
	n := len(boxes)
	pairs := make([]*Pair, 0, n*(n-1)/2)
	for i := range boxes {
		for j := i + 1; j < n; j++ {
			dx := int64(boxes[i].X - boxes[j].X)
			dy := int64(boxes[i].Y - boxes[j].Y)
			dz := int64(boxes[i].Z - boxes[j].Z)
			distance := dx*dx + dy*dy + dz*dz
			pairs = append(pairs, &Pair{dist: distance, i: i, j: j})
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].dist < pairs[j].dist })
	return pairs
}

// processDSU builds a DSU by unioning the first `k` pairs in the
// sorted list. It returns the resulting DSU.
func processDSU(pairs []*Pair, k int, n int) *DSU {
	if k > len(pairs) {
		k = len(pairs)
	}
	dsu := NewDSU(n)
	for idx := 0; idx < k; idx++ {
		p := pairs[idx]
		dsu.union(p.i, p.j)
	}
	return dsu
}

// getTopThreeProduct returns the product of the sizes of the three
// largest connected components in the DSU.
func getTopThreeProduct(dsu *DSU, n int) int {
	componentSizes := make(map[int]int)
	for i := range n {
		root := dsu.find(i)
		componentSizes[root] = dsu.size[root]
	}
	sizes := make([]int, 0, len(componentSizes))
	for _, size := range componentSizes {
		sizes = append(sizes, size)
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	result := 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= sizes[i]
	}
	return result
}

// P represents a pair of junction boxes with an associated distance.
type P struct {
	b1 *JunctionBox
	b2 *JunctionBox
	d  int
}

// Distance calculates the Euclidean distance between two junction boxes.
func Distance(p *P) int {
	dx := int64(p.b2.X - p.b1.X)
	dy := int64(p.b2.Y - p.b1.Y)
	dz := int64(p.b2.Z - p.b1.Z)
	distance := math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
	return int(distance)
}

// ProductOfThreeLargestCircuits returns the product of the sizes of the
// three largest connected components after adding the shortest `k` edges.
func ProductOfThreeLargestCircuits(d *utils.Data, k int) int {
	data := d.TransformData(dataTransformer).([][]int)
	boxes := make([]*JunctionBox, 0, len(data))
	for _, row := range data {
		boxes = append(boxes, &JunctionBox{row[0], row[1], row[2]})
	}
	pairs := generatePairs(boxes)
	dsu := processDSU(pairs, k, len(boxes))
	return getTopThreeProduct(dsu, len(boxes))
}

// ProcessUntilConnected iteratively unions pairs until all boxes are
// connected and returns the DSU at that point.
func ProcessUntilConnected(pairs []*Pair, n int) *DSU {
	dsu := NewDSU(n)
	for _, p := range pairs {
		merged := dsu.union(p.i, p.j)

		if merged {
			// Check if all nodes are now connected
			root := dsu.find(0)
			allConnected := true
			for i := 1; i < n; i++ {
				if dsu.find(i) != root {
					allConnected = false
					break
				}
			}

			if allConnected {
				return dsu
			}
		}
	}
	return dsu
}

// FindLastConnectionXProduct returns the product of the X coordinates of
// the pair that first connects the entire graph.
func FindLastConnectionXProduct(boxes []*JunctionBox, pairs []*Pair) int {
	dsu := NewDSU(len(boxes))

	for _, p := range pairs {
		merged := dsu.union(p.i, p.j)

		if merged {
			root := dsu.find(0)
			allConnected := true
			for i := 1; i < len(boxes); i++ {
				if dsu.find(i) != root {
					allConnected = false
					break
				}
			}

			if allConnected {
				return boxes[p.i].X * boxes[p.j].X
			}
		}
	}

	return 0
}

// ProductOfLastConnectionX returns the product of the X coordinates of
// the pair that connects the entire set of boxes for the given input data.
func ProductOfLastConnectionX(d *utils.Data) int {
	data := d.TransformData(dataTransformer).([][]int)
	boxes := make([]*JunctionBox, 0, len(data))
	for _, row := range data {
		boxes = append(boxes, &JunctionBox{row[0], row[1], row[2]})
	}

	pairs := generatePairs(boxes)
	return FindLastConnectionXProduct(boxes, pairs)
}
