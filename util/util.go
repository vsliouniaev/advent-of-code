package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLinesStrings(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
	return lines
}

func ReadLinesInts(path string) []int {
	var ints []int
	for _, s := range ReadLinesStrings(path) {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func ReadCSVLine(path string) []int {
	lines := ReadLinesStrings(path)
	var ints []int
	for _, str := range strings.Split(lines[0], ",") {
		i, err := strconv.Atoi(str)
		Check(err)
		ints = append(ints, i)
	}
	return ints
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
