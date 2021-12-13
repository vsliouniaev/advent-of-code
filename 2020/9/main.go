package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"math"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(util.RelativeFile("input"))) // 22477624
	fmt.Printf("Part 2: %d\n", part2(util.RelativeFile("input"))) // 2980044
}

func part1(file string) int {
	size := 25
	lines := util.ReadLinesInts(file)
	ints := make([]int, size)
	copy(ints, lines)
	for i := size; i < len(lines); i++ {
		if !canSum(lines[i], ints) {
			return lines[i]
		}
		ints = append(ints[1:], lines[i])
	}
	return -1
}

func part2(file string) int {
	target := part1(file)
	lines := util.ReadLinesInts(file)

	// Contiguous sub-array:
	//  If there is a contiguous array move right pointer until we exceed the target, then move left pointer until we're not.
	//  If we found the answer, we're done. If not start moving right pointer again
	sum := lines[0]
	for a, z := 0, 1; z < len(lines); z++ {
		for ; a < z && sum > target; a++ {
			sum -= lines[a]
		}
		if sum == target {
			return minMaxInRangeSum(a, z-1, lines)
		}
		sum += lines[z]
	}

	return -1
}

func canSum(target int, ints []int) bool {
	h := make(map[int]struct{})
	for _, i := range ints {
		if _, ok := h[target-i]; ok {
			return true
		}
		h[i] = struct{}{}
	}
	return false
}

func minMaxInRangeSum(a, z int, lines []int) int {
	min := math.MaxInt64
	max := math.MinInt64
	for ; a <= z; a++ {
		if min > lines[a] {
			min = lines[a]
		}
		if max < lines[a] {
			max = lines[a]
		}
	}
	return min + max
}
