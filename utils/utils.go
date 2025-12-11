package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

const (
	MsgPanic string = "unable to find input data"
	MsgExpected string = "expected %d, got %d"
)

type Part string

const (
	PartA Part = "example"
	PartB Part = "challenge"
)

type Day int

const (
	Day1 Day = iota
	Day2
	Day3
	Day4
	Day5
	Day6
	Day7
	Day8
	Day9
	Day10
	Day11
	Day12
	Day13
	Day14
	Day15
	Day16
	Day17
	Day18
	Day19
	Day20
	Day21
	Day22
	Day23
	Day24
	Day25
	Day26
	Day27
	Day28
	Day29
	Day30
	Day31
)

func (d Day) String() string {
	return [...]string{
		"day1",
		"day2",
		"day3",
		"day4",
		"day5",
		"day6",
		"day7",
		"day8",
		"day9",
		"day10",
		"day11",
		"day12",
		"day13",
		"day14",
		"day15",
		"day16",
		"day17",
		"day18",
		"day19",
		"day20",
		"day21",
		"day22",
		"day23",
		"day24",
		"day25",
		"day26",
		"day27",
		"day28",
		"day29",
		"day30",
		"day31",
	}[d]
}

const dataDelimiter = "-BREAK-"

func ReadLines(d Day, p Part) ([]string, error) {
	cwd, _ := os.Getwd()
	file := filepath.Join(cwd, "data", d.String())

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
		
		if p == PartA {
			if line == dataDelimiter {
				break
			}	
			lines = append(lines, line)		
		}
		
	 	if p == PartB {
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
