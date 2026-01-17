package day06

import (
	"fmt"

	"adventofcode.dev/utils"
)

// Maze represents the grid layout of the maze.
type Maze struct {
	Layout [][]string
}

// Symbols used in the maze.
const (
	// GUARD represents the starting position of the guard.
	GUARD = "^"
	// OBSTRUCTION represents a wall that the guard cannot cross.
	OBSTRUCTION = "#"
)

// Direction enumerates cardinal movement directions.
type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

// GuardState holds the guard's current position and facing direction.
type GuardState struct {
	X, Y int
	Dir  Direction
}

// Position represents an (X, Y) coordinate in the maze.
type Position struct {
	X, Y int
}

// Generate creates a Maze from the provided slice of strings, where each string
// represents a row of the layout.
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

// FindGuard searches the maze for the guard symbol and returns its position.
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

// Move advances the guard one step in the current direction, wrapping the
// behaviour around MoveWithObstruction with no obstruction present.
func Move(state GuardState, m *Maze) GuardState {
	return MoveWithObstruction(state, nil, m)
}

// MoveWithObstruction moves the guard from the given state, taking an optional
// obstruction into account. It returns the new state, or a sentinel state
// with coordinates (-1, -1) and Dir = -1 to signal that the guard has moved
// out of bounds.
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

// Explore simulates the guard walking until it exits the maze or reaches
// a sentinel state, returning the set of all visited positions.
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

// CausesLoop determines whether the guard, when avoiding the specified
// obstruction position, will enter a cycle that does not exit the maze.
// It returns true if a loop is detected, otherwise false.
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

// CountLoopPositions counts the number of positions that, if used as an
// obstruction, cause the guard to loop. The function panics if the maze
// cannot be generated.
func CountLoopPositions(d *utils.Data) (int, error) {
	m, err := Generate(d.Lines())
	if err != nil {
		panic(err)
	}

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

// CountDistinctPositions returns the number of distinct positions the guard
// visits while walking through the maze. It panics if the maze cannot be
// generated.
func CountDistinctPositions(d *utils.Data) int {
	m, err := Generate(d.Lines())
	if err != nil {
		panic(err)
	}
	pos, _ := Explore(m)
	return len(pos)
}
