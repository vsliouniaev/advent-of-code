package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2021/6/input")) // 362740
	fmt.Printf("Part 2: %d\n", part2("2021/6/input")) // 1644874076764
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
	for _, d := range util.ReadCSVLine(file) {
		fish[d]++
	}
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
