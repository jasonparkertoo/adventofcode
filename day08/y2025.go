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

type JunctionBox struct {
	X, Y, Z int
}

type Pair struct {
	dist int64
	i, j int
}

// disjoint-set-union
type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := range p {
		p[i] = i
		s[i] = 1
	}
	return &DSU{parent: p, size: s}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

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

type P struct {
	b1 *JunctionBox
	b2 *JunctionBox
	d  int
}

func Distance(p *P) int {
	dx := int64(p.b2.X - p.b1.X)
	dy := int64(p.b2.Y - p.b1.Y)
	dz := int64(p.b2.Z - p.b1.Z)
	distance := math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
	return int(distance)
}

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

func ProcessUntilConnected(pairs []*Pair, n int) *DSU {
    dsu := NewDSU(n)
    for _, p := range pairs {
        // Union returns true if components were merged
        merged := dsu.union(p.i, p.j)
        
        // If we just reduced the number of components to 1, return
        if merged {
            // Check if all nodes are now connected
            // We can check this efficiently by seeing if all nodes have the same root
            root := dsu.find(0)
            allConnected := true
            for i := 1; i < n; i++ {
                if dsu.find(i) != root {
                    allConnected = false
                    break
                }
            }
            
            if allConnected {
                // We need the pair that caused the final union
                // Return a DSU and the pair index
                return dsu
            }
        }
    }
    return dsu
}

func FindLastConnectionXProduct(boxes []*JunctionBox, pairs []*Pair) int {
    dsu := NewDSU(len(boxes))
    
    for _, p := range pairs {
        merged := dsu.union(p.i, p.j)
        
        if merged {
            // Check if all connected
            root := dsu.find(0)
            allConnected := true
            for i := 1; i < len(boxes); i++ {
                if dsu.find(i) != root {
                    allConnected = false
                    break
                }
            }
            
            if allConnected {
                // Return product of X coordinates of the last connected pair
                return boxes[p.i].X * boxes[p.j].X
            }
        }
    }
    
    return 0
}

func ProductOfLastConnectionX(d *utils.Data) int {
    data := d.TransformData(dataTransformer).([][]int)
    boxes := make([]*JunctionBox, 0, len(data))
    for _, row := range data {
        boxes = append(boxes, &JunctionBox{row[0], row[1], row[2]})
    }
    
    pairs := generatePairs(boxes)
    return FindLastConnectionXProduct(boxes, pairs)
}