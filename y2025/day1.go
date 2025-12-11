package y2025

import "strconv"

type Pair struct {
	Direction string
	Distance  int
}
type Document struct {
	Pairs []Pair
}

func NewDocument(rotations []string) *Document {
	var pairs []Pair

	for _, rotation := range rotations {
		d := rotation[0:1]
		t, _ := strconv.Atoi(rotation[1:])
		pair := Pair{Direction: d, Distance: t}
		pairs = append(pairs, pair)
	}

	return &Document{
		Pairs: pairs,
	}
}

type Method int

const (
	MethodNone Method = iota
	Method0x434C49434B
)

func (d Document) DoorPassword(method Method) int {
	pos := 50
	landed, during := 0, 0

	for _, p := range d.Pairs {
		dis := p.Distance
		for range dis {
			switch p.Direction {
			case "L":
				pos = (pos - 1 + 100) % 100
			case "R":
				pos = (pos + 1) % 100
			}
			if pos == 0 {
				during++
			}
		}
		if pos == 0 {
			landed++
		}
	}

	if method == Method0x434C49434B {
		return during
	}

	return landed
}
