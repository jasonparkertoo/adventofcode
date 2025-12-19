package day06

import (
	"fmt"
)

type Maze struct {
	Layout [][]string
}

const (
	GUARD      = "^"
	OBSTRUCTION = "#"
)

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type GuardState struct {
	X, Y     int
	Dir      Direction
}

type Position struct {
	X, Y int
}

func Generate(lines []string) (*Maze, error) {
	var layout [][]string
	for _, line := range lines {
		row := make([]string, len(line))
		for i, ch := range line {
			row[i] = string(ch)
		}
		layout = append(layout, row)
	}
	return &Maze{Layout: layout}, nil
}

func (m *Maze) FindGuard() (Position, error) {
	for y, row := range m.Layout {
		for x, cell := range row {
			if cell == GUARD {
				return Position{X: x, Y: y}, nil
			}
		}
	}
	return Position{}, fmt.Errorf("guard not found")
}

func Move(state GuardState, m *Maze) GuardState {
	return MoveWithObstruction(state, nil, m)
}

func MoveWithObstruction(state GuardState, obstruction *Position, m *Maze) GuardState {
	width := len(m.Layout[0])
	height := len(m.Layout)

	nx, ny := state.X, state.Y
	dir := state.Dir

	switch dir {
	case NORTH:
		ny--
	case EAST:
		nx++
	case SOUTH:
		ny++
	case WEST:
		nx--
	}

	// Out of bounds -> finished
	if nx < 0 || nx >= width || ny < 0 || ny >= height {
		return GuardState{-1, -1, -1}
	}

	blocked := m.Layout[ny][nx] == OBSTRUCTION ||
		(obstruction != nil && obstruction.X == nx && obstruction.Y == ny)

	if blocked {
		// Turn right
		switch dir {
		case NORTH:
			dir = EAST
		case EAST:
			dir = SOUTH
		case SOUTH:
			dir = WEST
		case WEST:
			dir = NORTH
		}
		return GuardState{X: state.X, Y: state.Y, Dir: dir}
	}

	return GuardState{X: nx, Y: ny, Dir: dir}
}

func Explore(m *Maze) (map[Position]struct{}, error) {
	guardPos, err := m.FindGuard()
	if err != nil {
		return nil, err
	}
	start := GuardState{X: guardPos.X, Y: guardPos.Y, Dir: NORTH}

	visited := make(map[Position]struct{})
	cur := start
	for cur.Dir != -1 {
		visited[Position{X: cur.X, Y: cur.Y}] = struct{}{}
		cur = Move(cur, m)
	}
	return visited, nil
}

func CausesLoop(current Position, m *Maze) bool {
	guardStart, err := m.FindGuard()
	if err != nil {
		return false
	}

	if m.Layout[current.Y][current.X] == OBSTRUCTION ||
		(current.X == guardStart.X && current.Y == guardStart.Y) {
		return false
	}

	visited := make(map[GuardState]struct{})
	cur := GuardState{X: guardStart.X, Y: guardStart.Y, Dir: NORTH}

	for cur.Dir != -1 {
		if _, ok := visited[cur]; ok {
			return true
		}
		visited[cur] = struct{}{}
		cur = MoveWithObstruction(cur, &current, m)
	}
	return false
}

func CountLoopPositions(m *Maze) (int, error) {
	guardPos, err := m.FindGuard()
	if err != nil {
		return 0, err
	}
	explored, err := Explore(m)
	if err != nil {
		return 0, err
	}

	count := 0
	for pos := range explored {
		if pos == guardPos {
			continue
		}
		if CausesLoop(pos, m) {
			count++
		}
	}
	return count, nil
}
