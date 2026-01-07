package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	MsgPanic    string = "unable to find input data"
	MsgExpected string = "expected %d, got %d"
)

type DataSet string

const (
	Example   DataSet = "example"
	Challenge DataSet = "challenge"
)

const (
	Year2024 string = "2024"
	Year2025 string = "2025"
)

const dataDelimiter = "-BREAK-"

func ReadLines(y string, ds DataSet) ([]string, error) {
	cwd, _ := os.Getwd()
	file := filepath.Join(cwd, y)

	f, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	sc := bufio.NewScanner(f)

	var lines []string
	pastDelimiter := false
	for sc.Scan() {
		line := sc.Text()

		if ds == Example {
			if line == dataDelimiter {
				break
			}
			lines = append(lines, line)
		}

		if ds == Challenge {
			if !pastDelimiter {
				if line == dataDelimiter {
					pastDelimiter = true
				}
				continue
			}
			lines = append(lines, line)
		}
	}
	return lines, nil
}

type Data struct {
	lines []string
}

func NewData(ds DataSet, y string) *Data {
	if lines, err := ReadLines(y, ds); err != nil {
		panic(err)
	} else {
		return &Data{lines}
	}
}

func (d Data) Lines() []string {
	return d.lines
}

func (d Data) Line(n int) (string, error) {
	if n < len(d.lines) || n < 1 {
		return "", fmt.Errorf("invalid line number requested: %d", n)
	}
	return d.lines[n-1], nil
}

func (d Data) TransformData(fn func([]string) any) any {
	return fn(d.lines)
}

func (d Data) AsGrid() (out [][]string) {
	for _, row := range d.lines {
		out = append(out, strings.Split(row, ""))
	}
	return out
}
