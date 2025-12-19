package day10

type Point struct{ R, C int }

type LavaTrails struct {
	grid      [][]int
	rows,cols int
}

func NewLavaTrails(lines []string) *LavaTrails {
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]int, rows)
	for r := range grid {
		grid[r] = make([]int, cols)
		for c := range grid[r] {
			grid[r][c] = int(lines[r][c] - '0')
		}
	}
	return &LavaTrails{
		grid: grid,
		rows: rows,
		cols: cols,
	}
}

func cloneSet(src map[Point]struct{}) map[Point]struct{} {
	dst := make(map[Point]struct{}, len(src))
	for k := range src {
		dst[k] = struct{}{}
	}
	return dst
}

func height(t *LavaTrails, p Point) int {
	return t.grid[p.R][p.C]
}

func neighbors(t *LavaTrails, p Point) []Point {
	r, c := p.R, p.C
	cands := [...]Point{
		{r - 1, c}, {r + 1, c},
		{r, c - 1}, {r, c + 1},
	}
	out := make([]Point, 0, 4)
	for _, x := range cands {
		if x.R >= 0 && x.C >= 0 && x.R < t.rows && x.C < t.cols {
			out = append(out, x)
		}
	}
	return out
}

func reachableNines(t *LavaTrails, p Point, memo map[Point]map[Point]struct{}) map[Point]struct{} {
	// READ: return clone so caller can't mutate memo
	if got, ok := memo[p]; ok {
		return cloneSet(got)
	}

	h := height(t, p)
	var res map[Point]struct{}
	if h == 9 {
		res = map[Point]struct{}{p: {}}
	} else {
		res = make(map[Point]struct{})
		for _, n := range neighbors(t, p) {
			if height(t, n) == h+1 {
				child := reachableNines(t, n, memo)
				for q := range child {
					res[q] = struct{}{}
				}
			}
		}
	}

	// WRITE: freeze a clone to memo
	frozen := cloneSet(res)
	memo[p] = frozen
	// RETURN: return a fresh clone so caller can't mutate our stored one
	return cloneSet(frozen)
}

func countPaths(t *LavaTrails, p Point, memo map[Point]int) int {
    if val, exists := memo[p]; exists {
        return val
    }

    h := height(t, p)
    var count int
    if h == 9 {
        count = 1
    } else {
        count = 0
        for _, n := range neighbors(t, p) {
            if height(t, n) == h+1 {
                count += countPaths(t, n, memo)
            }
        }
    }

    memo[p] = count
    return count
}

func TotalTrailheadRating(t *LavaTrails) int {
    sum := 0
    for r := 0; r < t.rows; r++ {
        for c := 0; c < t.cols; c++ {
            p := Point{r, c}
            if height(t, p) == 0 {
                memo := make(map[Point]int)
                sum += countPaths(t, p, memo)
            }
        }
    }
    return sum
}

func TotalTrailheadScore(t *LavaTrails) int {
	memo := make(map[Point]map[Point]struct{})
	sum := 0
	for r := 0; r < t.rows; r++ {
		for c := 0; c < t.cols; c++ {
			p := Point{r, c}
			if height(t, p) == 0 {
				sum += len(reachableNines(t, p, memo))
			}
		}
	}
	return sum
}
