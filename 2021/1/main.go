package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
	"math"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part2() int {
	ints := ReadLinesInts("1/input")
	count := 0
	for i := 3; i < len(ints); i++ {
		prev := ints[i-3] + ints[i-2] + ints[i-1]
		cur := ints[i-2] + ints[i-1] + ints[i-0]
		if cur > prev {
			count++
		}
	}
	return count
}

func part1() int {
	ints := ReadLinesInts("1/input")
	count := 0
	prev := math.MaxInt64
	for _, i := range ints {
		if i > prev {
			count++
		}
		prev = i
	}

	return count
}
