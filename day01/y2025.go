package day01

import (
	"strconv"
)

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

func (d Document) DoorPassword() []int {
	pos := 50
	landed, visited := 0, 0

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
				visited++
			}
		}
		if pos == 0 {
			landed++
		}
	}

	r := []int{landed, visited}
	return r
}
