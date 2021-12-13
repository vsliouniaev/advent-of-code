package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(util.RelativeFile("input"))) // 2432
	fmt.Printf("Part 2: %d\n", part2(util.RelativeFile("input"))) // 453551299002368
}

func part1(file string) int {
	lines := util.ReadLinesInts(file)
	sort.Ints(lines)
	ones := 0
	thrs := 0
	cur := 0
	for _, i := range lines {
		switch i - cur {
		case 1:
			ones++
		case 3:
			thrs++
		}
		cur = i
	}
	thrs++

	return ones * thrs
}

func part2(file string) int {
	adapters := util.ReadLinesInts(file)
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	var memo = make(map[int]int)
	var pathsTo func(target int) int
	// Start at the end, then search for all paths reachable from that, and recurse for all those as well.
	pathsTo = func(target int) int {
		if target == 0 {
			return 1
		}
		if ps, ok := memo[target]; ok {
			return ps
		}
		cnt := 0
		for _, a := range adapters {
			d := target - a
			if d <= 3 && d > 0 {
				ps := pathsTo(a)
				memo[a] = ps
				cnt += ps
			}
		}
		return cnt
	}

	return pathsTo(adapters[len(adapters)-1])
}
