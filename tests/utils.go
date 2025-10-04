package tests

import (
	"bufio"
	"fmt"
	"os"
)

type PartName string

const (
	Challenge PartName = "challenge"
	Example   PartName = "example"
)

func ReadLines(part PartName, fileName string) ([]string, error) {
	cwd, _ := os.Getwd()
	file := fmt.Sprintf("%s/../data/%s/%s", cwd, string(part), fileName)
	
	f, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, nil
}
