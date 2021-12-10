package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2021/10/input")) // 369105
	fmt.Printf("Part 2: %d\n", part2("2021/10/input")) // 3999363569
}

var matching = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func part1(file string) int {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	sum := 0
	for _, brackets := range u.ReadLinesStrings(file) {
		stk := &u.Stack{}
		for _, b := range brackets {
			m, ok := matching[b]
			if ok {
				stk.Push(m)
			} else {
				if stk.Pop() != b {
					sum += points[b]
				}
			}
		}
	}
	return sum
}

func part2(file string) int {
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	var scores []int
	for _, brackets := range u.ReadLinesStrings(file) {
		stk := &u.Stack{}
		s := 0

		for _, b := range brackets {
			m, ok := matching[b]
			if ok {
				stk.Push(m)
			} else {
				if stk.Pop() != b {
					goto next
				}
			}
		}

		for stk.Len() != 0 {
			s *= 5
			s += points[stk.Pop().(rune)]
		}
		scores = append(scores, s)
	next:
	}
	sort.Ints(scores)
	return scores[(len(scores) / 2)]
}
