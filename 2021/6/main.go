package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 362740
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input"))) // 1644874076764
}

func part1(file string) int {
	return afterDays(file, 80)
}

func part2(file string) int {
	return afterDays(file, 256)
}

func afterDays(file string, days int) int {
	var size = 9
	fish := make([]int, size)
	for _, d := range u.ReadCSVLine(file) {
		fish[d]++
	}

	// @eoincampbell:
	// Day 8 is just day 0 - new fish
	// To reset the ones that reproduce, just add them to day 6
	//
	//for d := 0; d < days; d++ {
	//	fish[(d+7)%9] += fish[d%9]
	//}

	for day := 0; day < days; day++ {
		add := fish[0]
		fish = append(fish[1:], add)
		fish[size-3] += add
	}

	count := 0
	for _, f := range fish {
		count += f
	}
	return count
}
