package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("sample")))
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("sample")))
}

func part1(file string) int {
	fmt.Println(u.RelativeFile("sample"))
	return 0
}

func part2(file string) int64 {
	return 0
}
