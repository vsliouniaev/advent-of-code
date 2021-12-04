package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
	"sort"
)

var target = 2020

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/1/input"))
	fmt.Printf("Part 2: %d\n", part2("2020/1/input"))
}

func part1(file string) int {
	// This is a standard interview question, which you can do in one loop with a hashmap (Provided vals are unique)
	ints := ReadLinesInts(file)
	lookup := make(map[int]struct{})
	for _, n := range ints {
		other := target - n
		if _, ok := lookup[other]; ok {
			return other * n
		} else {
			if _, ok := lookup[n]; ok {
				panic("Non-unique")
			}
			lookup[n] = struct{}{}
		}
	}
	return -1
}

func part2(file string) int {
	// This is called 3-sum problem, which is solvable in O(n^2)
	// First: sort the list, which is nlogn
	// Then in the sorted list we can fix the first number and look at the next one and the end one
	//  If the sum at those indexes is too low we need to look for a higher number - move the middle up
	//  If the sum at those is too high, then we need to move the index of the end down
	ints := ReadLinesInts(file)
	sort.Ints(ints) // Increasing order
	for low := 0; low < len(ints); low++ {
		mid := low + 1
		high := len(ints) - 1
		for mid < high {
			val := ints[low] + ints[mid] + ints[high]
			if val == target {
				return ints[low] * ints[mid] * ints[high]
			}
			if val < target {
				mid++
			} else {
				high--
			}
		}
	}
	return 0
}
