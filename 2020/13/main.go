package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/13/input")) // 410 (too high)
	//fmt.Printf("Part 2: %d\n", part2("2020/1/input"))
}

func part1(file string) int {
	lines := u.ReadLinesStrings(file)
	target, err := strconv.Atoi(lines[0])
	u.Check(err)

	min := math.MaxInt64
	sel := 0
	for _, b := range strings.Split(lines[1], ",") {
		busInterval, err := strconv.Atoi(b)
		if err != nil {
			continue
		}

		numBusesBefore := target / busInterval
		firstBusAfter := numBusesBefore + 1
		offset := (busInterval * firstBusAfter) - target

		if min > offset {
			min = offset
			sel = busInterval * (offset)
		}
	}

	return sel
}

func part2(file string) int {

	return 0
}
