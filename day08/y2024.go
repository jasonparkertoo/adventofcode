package day08

import "adventofcode.dev/utils"

// Pos represents a coordinate on the grid.
type Pos struct {
	X, Y int
}

// Antenna represents a transmitter with a frequency and a location.
type Antenna struct {
	Freq rune
	Pos  Pos
}

// CountUniqueLocations returns the number of distinct antenna nodes
// that can be reached when each pair of antennas on the same
// frequency uses the Part 1 strategy for generating antinodes.
func CountUniqueLocations(d *utils.Data) int {
	return countAntinodes(antinodesPart1, d)
}

// CountUniqueLocationsHarmonics returns the number of distinct antenna nodes
// that can be reached when each pair of antennas on the same
// frequency uses the Part 2 strategy for generating antinodes.
func CountUniqueLocationsHarmonics(d *utils.Data) int {
	return countAntinodes(antinodesPart2, d)
}

// AntinodeStrategy defines the algorithm used to compute the set of
// antinode positions for a pair of antennas.  The function receives
// the positions of the two antennas, the grid width, and the grid height,
// and returns a map whose keys are the positions of the antinodes.
type AntinodeStrategy func(a, b Pos, width, height int) map[Pos]struct{}

// countAntinodes is the core routine that, given a strategy,
// parses the input data, groups antennas by frequency, and
// accumulates all antinode positions produced by the strategy.
// The return value is the number of unique antinode positions.
func countAntinodes(strategy AntinodeStrategy, d *utils.Data) int {
	lines := d.Lines()

	height := len(lines)
	width := len(lines[0])

	antennas := parseAntennas(lines)
	byFreq := make(map[rune][]Antenna)
	for _, ant := range antennas {
		byFreq[ant.Freq] = append(byFreq[ant.Freq], ant)
	}

	antinodes := make(map[Pos]struct{})

	for _, group := range byFreq {
		if len(group) < 2 {
			continue
		}
		for i := range group {
			for j := i + 1; j < len(group); j++ {
				for pos := range strategy(group[i].Pos, group[j].Pos, width, height) {
					antinodes[pos] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

// parseAntennas scans the grid lines and returns a slice of Antenna
// structs for every non‑empty cell.  The frequency is the rune in the
// cell and the position is its (X, Y) coordinates.
func parseAntennas(lines []string) []Antenna {
	var antennas []Antenna
	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				antennas = append(antennas, Antenna{Freq: ch, Pos: Pos{X: x, Y: y}})
			}
		}
	}
	return antennas
}

// antinodesPart1 implements the Part 1 antinode strategy:
// For a pair of antennas, it returns the two positions that lie
// directly beyond each antenna along the line defined by the pair,
// provided those positions are inside the grid bounds.
func antinodesPart1(a, b Pos, width, height int) map[Pos]struct{} {
	result := make(map[Pos]struct{})
	dx := b.X - a.X
	dy := b.Y - a.Y

	for _, p := range []Pos{
		{X: a.X - dx, Y: a.Y - dy},
		{X: b.X + dx, Y: b.Y + dy},
	} {
		if inBounds(p, width, height) {
			result[p] = struct{}{}
		}
	}

	return result
}

// antinodesPart2 implements the Part 2 antinode strategy:
// For a pair of antennas, it walks along the line connecting them
// and records every grid cell that lies on that line and is
// within bounds.  The walk starts at the extreme point opposite
// the pair, steps outward, and continues until it exits the grid.
func antinodesPart2(a, b Pos, width, height int) map[Pos]struct{} {
	result := make(map[Pos]struct{})
	dx := b.X - a.X
	dy := b.Y - a.Y

	g := gcd(abs(dx), abs(dy))
	sx := dx / g
	sy := dy / g

	// Step back to the edge of the grid
	bx, by := a.X, a.Y
	for inBounds(Pos{X: bx - sx, Y: by - sy}, width, height) {
		bx -= sx
		by -= sy
	}

	// Walk forward along the line, collecting positions
	px, py := bx, by
	for inBounds(Pos{X: px, Y: py}, width, height) {
		result[Pos{X: px, Y: py}] = struct{}{}
		px += sx
		py += sy
	}

	return result
}

// inBounds reports whether a position lies within the grid of the
// specified width and height.
func inBounds(p Pos, width, height int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < width && p.Y < height
}

// gcd returns the greatest common divisor of a and b
// using the Euclidean algorithm.  The function accepts
// negative values and returns a non‑negative result.
func gcd(a, b int) int {
	if b == 0 {
		return abs(a)
	}
	return gcd(b, a%b)
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
