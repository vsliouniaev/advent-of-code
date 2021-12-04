package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("3/input"))
	fmt.Printf("Part 2: %d\n", part2("3/input"))
}

func part1(file string) int {
	lines := ReadLinesStrings(file)
	width := len(lines[0])
	ones := make([]int, width)
	zero := make([]int, width)
	for _, l := range ReadLinesStrings(file) {
		for i := 0; i < width; i++ {
			switch l[i] {
			case '0':
				zero[i]++
			case '1':
				ones[i]++
			default:
				panic("unknown num")
			}
		}
	}

	var gam int
	var eps int

	for i := 0; i < width; i++ {
		if ones[width-i-1] > zero[width-i-1] {
			gam = gam | (1 << i)
		} else {
			eps = eps | (1 << i)
		}
	}

	return eps * gam
}

func part2(file string) int64 {
	lines := ReadLinesStrings(file)
	width := len(lines[0])

	oxy := make([]string, len(lines))
	co2 := make([]string, len(lines))
	copy(oxy, lines)
	copy(co2, lines)
	for i := 0; i < width; i++ {
		oxy = filter(oxy, i, mostCommon(oxy, i))
		co2 = filter(co2, i, leastCommon(co2, i))
	}

	o, _ := strconv.ParseInt(oxy[0], 2, 64)
	d, _ :=strconv.ParseInt(co2[0], 2, 64)
	return o * d
}

func countAtPos(lines []string, pos int) (one int, zer int) {
	for _, line := range lines {
		if line[pos] == '0' {
			zer++
		} else {
			one++
		}
	}

	return
}

func mostCommon(lines []string, pos int) byte {
	one, zer := countAtPos(lines, pos)
	if one >= zer {
		return '1'
	} else {
		return '0'
	}
}

func leastCommon(lines []string, pos int) byte {
	one, zer := countAtPos(lines, pos)
	if one >= zer {
		return '0'
	} else {
		return '1'
	}
}

func filter(lines []string, pos int, r byte) []string {
	if len(lines) == 1 {
		return lines
	}
	var out []string
	for _, line := range lines {
		if line[pos] == r {
			out = append(out, line)
		}
	}

	return out
}
