package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input")))
	//fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input")))
}

func part1(file string) int {
	valids := make(map[int]struct{})
	pm := 0
	sum := 0
	for _, line := range u.ReadLinesStrings(file) {
		if line == "your ticket:" {
			pm = 1
			continue
		}
		if line == "nearby tickets:" {
			pm = 2
			continue
		}
		switch pm {
		case 0:
			for _, v := range getRanges(line) {
				valids[v] = struct{}{}
			}
		case 1:
			//return 0
		case 2:
			for _, s := range strings.Split(line, ",") {
				i, err := strconv.Atoi(s)
				u.Check(err)
				if _, ok := valids[i]; !ok {
					sum += i
				}
			}

		default:
			panic(pm)
		}

	}

	return sum
}

func part2(file string) int {
	return 0
}

func getRanges(line string) []int {
	var out []int
	for _, rng := range regexp.MustCompile(`\d+-\d+`).FindAllStringSubmatch(line, -1) {
		split := strings.Split(rng[0], "-")
		s, err := strconv.Atoi(split[0])
		u.Check(err)
		e, err := strconv.Atoi(split[1])
		u.Check(err)
		for ; s <= e; s++ {
			out = append(out, s)
		}
	}
	return out
}
