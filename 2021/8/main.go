package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"math"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2021/7/input"))
	//fmt.Println(sumToN(14))
	fmt.Printf("Part 2: %d\n", part2("2021/7/input"))
}

func part1(file string) int {
	crabs := util.ReadCSVLine(file)
	min := math.MaxInt64
	for _, pos := range crabs {
		tot := 0
		for _, crab := range crabs {
			tot += int(math.Abs(float64(pos - crab)))
		}
		if min > tot {
			min = tot
		}
	}
	return min
}

func part2(file string) int {
	crabs := util.ReadCSVLine(file)
	min := math.MaxInt64
	max := math.MinInt64
	for _, pos := range crabs {
		if min > pos {
			min = pos
		}
		if max < pos {
			max = pos
		}
	}

	minf := math.MaxInt64
	for pos := min; pos <= max; pos++ {
		tot := 0
		for _, crab := range crabs {
			sum := sumToN(int(math.Abs(float64(pos - crab))))
			tot += sum
		}
		if minf > tot {
			minf = tot
		}
	}

	return minf
}

func sumToN(n int) int {
	return n * (n + 1) / 2
}
