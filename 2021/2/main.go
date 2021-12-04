package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2/input"))
	fmt.Printf("Part 2: %d\n", part2("2/input"))
}

func part1(file string) int {
	lines := ReadLinesStrings(file)
	x, d := 0, 0
	for _, l := range lines {
		inp := strings.Split(l, " ")
		m, err := strconv.Atoi(inp[1])
		if err != nil {
			panic(err)
		}
		switch inp[0] {
		case "forward":
			x += m
		case "down":
			d += m
		case "up":
			d -= m
		}
	}
	return x * d
}

func part2(file string) int {
	lines := ReadLinesStrings(file)
	x, a, d := 0, 0, 0
	for _, l := range lines {
		inp := strings.Split(l, " ")
		m, err := strconv.Atoi(inp[1])
		if err != nil {
			panic(err)
		}
		switch inp[0] {
		case "forward":
			d += a * m
			x += m
		case "down":
			a += m
		case "up":
			a -= m
		}
	}
	return x * d
}
