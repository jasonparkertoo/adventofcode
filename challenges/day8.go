package challenges


type Pos struct {
	X, Y int
}

type Antenna struct {
	Freq rune
	Pos  Pos
}

type City struct {
	Scan []string
}

// --------------------
// Public API
// --------------------
func (c *City) CountUniqueLocations() int {
	return c.countAntinodes(c.antinodesPart1)
}

func (c *City) CountUniqueLocationsHarmonics() int {
	return c.countAntinodes(c.antinodesPart2)
}

// --------------------
// Core shared logic
// --------------------
type AntinodeStrategy func(a, b Pos, width, height int) map[Pos]struct{}

func (c *City) countAntinodes(strategy AntinodeStrategy) int {
	height := len(c.Scan)
	width := len(c.Scan[0])

	antennas := c.parseAntennas()
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

// --------------------
// Parsing
// --------------------
func (c *City) parseAntennas() []Antenna {
	var antennas []Antenna
	for y, line := range c.Scan {
		for x, ch := range line {
			if ch != '.' {
				antennas = append(antennas, Antenna{Freq: ch, Pos: Pos{X: x, Y: y}})
			}
		}
	}
	return antennas
}

// --------------------
// Strategies
// --------------------
func (c *City) antinodesPart1(a, b Pos, width, height int) map[Pos]struct{} {
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

func (c *City) antinodesPart2(a, b Pos, width, height int) map[Pos]struct{} {
	result := make(map[Pos]struct{})
	dx := b.X - a.X
	dy := b.Y - a.Y

	g := gcd(abs(dx), abs(dy))
	sx := dx / g
	sy := dy / g

	// Step back to edge
	bx, by := a.X, a.Y
	for inBounds(Pos{X: bx - sx, Y: by - sy}, width, height) {
		bx -= sx
		by -= sy
	}

	// Walk forward along line
	px, py := bx, by
	for inBounds(Pos{X: px, Y: py}, width, height) {
		result[Pos{X: px, Y: py}] = struct{}{}
		px += sx
		py += sy
	}

	return result
}

// --------------------
// Helpers
// --------------------
func inBounds(p Pos, width, height int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < width && p.Y < height
}

func gcd(a, b int) int {
	if b == 0 {
		return abs(a)
	}
	return gcd(b, a%b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
