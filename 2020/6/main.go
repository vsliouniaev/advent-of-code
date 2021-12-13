package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
)

func main() {

	fmt.Printf("Part 1: %d\n", part1(RelativeFile("input"))) // 7120
	fmt.Printf("Part 2: %d\n", part2(RelativeFile("input")))
}

func part1(file string) int {
	sum := 0
	qs := make(map[rune]struct{})
	for _, line := range ReadLinesStrings(file) {
		for _, l := range line {
			qs[l] = struct{}{}
		}
		if line == "" {
			sum += len(qs)
			qs = make(map[rune]struct{})
		}
	}
	sum += len(qs)

	return sum
}

func part2(file string) int {
	sum := 0
	ppl := 0
	qs := make(map[rune]int)
	for _, line := range ReadLinesStrings(file) {
		if line == "" {
			for _, v := range qs {
				if v == ppl {
					sum++
				}
			}

			qs = make(map[rune]int)
			ppl = 0

		} else {
			ppl++
			for _, l := range line {
				qs[l] = qs[l] + 1
			}
		}
	}

	for _, v := range qs {
		if v == ppl {
			sum++
		}
	}
	return sum

}
