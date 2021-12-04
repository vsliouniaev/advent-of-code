package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/3/input")) // 272
	fmt.Printf("Part 2: %d\n", part2("2020/3/input")) // 3898725600

}

func part1(file string) int {
	lines := ReadLinesStrings(file)
	return countTrees(lines, 3, 1)
}

func countTrees(lines []string, right, down int) int {
	col := 0
	trees := 0
	for row := down; row < len(lines); row += down {
		line := lines[row]
		col += right
		if col >= len(line) {
			col -= len(line)
		}
		if line[col] == '#' {
			trees++
		}
	}
	return trees
}

func part2(file string) int {
	lines := ReadLinesStrings(file)
	return countTrees(lines, 1, 1) *
		countTrees(lines, 3, 1) *
		countTrees(lines, 5, 1) *
		countTrees(lines, 7, 1) *
		countTrees(lines, 1, 2)
}
