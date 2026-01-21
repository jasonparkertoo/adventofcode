// Package utils provides utility functions for reading and processing input data
// for Advent of Code solutions.
package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	// MsgPanic is the error message used when input data cannot be found
	MsgPanic string = "unable to find input data"
	// MsgExpected is the error message format used for expected vs actual value mismatches
	MsgExpected string = "expected %d, got %d"
)

// DataSet represents whether input data should come from an example or challenge file.
type DataSet string

const (
	// Example represents the example input data
	Example DataSet = "example"
	// Challenge represents the challenge input data
	Challenge DataSet = "challenge"
)

const (
	// Year2024 represents the 2024 Advent of Code year
	Year2024 string = "2024"
	// Year2025 represents the 2025 Advent of Code year
	Year2025 string = "2025"
)

const dataDelimiter = "-BREAK-"

// ReadLines reads the input file for the specified year and dataset.
// It returns a slice of strings, one for each line of the relevant data section.
// If the dataset is Example, lines are read until the data delimiter.
// If the dataset is Challenge, lines after the data delimiter are returned.
func ReadLines(y string, ds DataSet) ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}
	file := filepath.Join(cwd, y)

	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", file, err)
	}

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

	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", file, err)
	}

	err = f.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close file %s: %w", file, err)
	}

	return lines, nil
}

// Data holds the input lines read from a data file.
type Data struct {
	lines []string
}

// NewData creates a Data instance for the given dataset and year by reading the appropriate file.
func NewData(ds DataSet, y string) (*Data, error) {
	lines, err := ReadLines(y, ds)
	if err != nil {
		return nil, fmt.Errorf("failed to read lines: %w", err)
	}
	return &Data{lines}, nil
}

// Lines returns all input lines as a slice of strings.
func (d Data) Lines() []string {
	return d.lines
}

// Line returns the requested line (1-based) or an error if the index is out of range.
func (d Data) Line(n int) (string, error) {
	if n > len(d.lines) || n < 1 {
		return "", fmt.Errorf("invalid line number requested: %d, expected 1 to %d", n, len(d.lines))
	}
	return d.lines[n-1], nil
}

// TransformData applies the supplied function to the underlying line slice and returns its result.
func (d Data) TransformData(fn func([]string) any) any {
	return fn(d.lines)
}

// AsGrid converts the data into a two-dimensional slice of single-character strings.
func (d Data) AsGrid() (out [][]string) {
	for _, row := range d.lines {
		out = append(out, strings.Split(row, ""))
	}
	return out
}
