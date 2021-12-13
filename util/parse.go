package util

import (
	"bufio"
	"github.com/vsliouniaev/aoc/util/nav"
	"os"
	"path/filepath"
	"runtime"
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

func ReadLinesRuneGrid(path string) nav.Grid {
	//file, err := os.Open(path)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//var lines [][]rune
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	lines = append(lines, []rune(scanner.Text()))
	//}
	//err = scanner.Err()
	//if err != nil {
	//	panic(err)
	//}
	//return lines
	lines := ReadLinesStrings(path)
	out := make([][]interface{}, len(lines))
	for y, line := range lines {
		out[y] = make([]interface{}, len(line))
		for x, c := range line {
			out[y][x] = c
		}
	}
	return nav.NewGrid(out)
}

func ReadLinesIntGrid(path string) nav.Grid {
	lines := ReadLinesStrings(path)
	out := make([][]interface{}, len(lines))
	for y, line := range lines {
		out[y] = make([]interface{}, len(line))
		for x, c := range line {
			ch := int(c - '0')
			out[y][x] = ch
		}
	}
	return nav.NewGrid(out)
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

func RelativeFile(file string) string {
	_, b, _, _ := runtime.Caller(1)
	path, err := os.Getwd()
	Check(err)
	trimPath := strings.TrimPrefix(filepath.Dir(b), path)
	if trimPath == "" {
		return filepath.Join(path, file)
	} else {
		return filepath.Join(trimPath[1:], file)
	}
}
