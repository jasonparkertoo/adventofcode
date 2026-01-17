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

// DataSet represents whether input data should come from an example or challenge file.
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

// ReadLines reads the input file for the specified year and dataset.
// It returns a slice of strings, one for each line of the relevant data section.
// If the dataset is Example, lines are read until the data delimiter.
// If the dataset is Challenge, lines after the data delimiter are returned.
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

// Data holds the input lines read from a data file.
type Data struct {
	lines []string
}

// NewData creates a Data instance for the given dataset and year by reading the appropriate file.
func NewData(ds DataSet, y string) *Data {
	if lines, err := ReadLines(y, ds); err != nil {
		panic(err)
	} else {
		return &Data{lines}
	}
}

// Lines returns all input lines as a slice of strings.
func (d Data) Lines() []string {
	return d.lines
}

// Line returns the requested line (1‑based) or an error if the index is out of range.
func (d Data) Line(n int) (string, error) {
	if n < len(d.lines) || n < 1 {
		return "", fmt.Errorf("invalid line number requested: %d", n)
	}
	return d.lines[n-1], nil
}

// TransformData applies the supplied function to the underlying line slice and returns its result.
func (d Data) TransformData(fn func([]string) any) any {
	return fn(d.lines)
}

// AsGrid converts the data into a two‑dimensional slice of single‑character strings.
func (d Data) AsGrid() (out [][]string) {
	for _, row := range d.lines {
		out = append(out, strings.Split(row, ""))
	}
	return out
}
