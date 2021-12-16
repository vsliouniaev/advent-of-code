package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"math"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 3411
	fmt.Printf("Part 1: %d\n", part2(u.RelativeFile("input"))) // 7477815755570
}

func polymerise(file string, iterations int) uint64 {
	template, rules := parseInput(file)
	letterCounts := map[rune]uint64{template[0]: 1}
	cur := make(map[string]uint64)
	for i := 1; i < len(template); i++ {
		k := string([]rune{template[i-1], template[i]})
		letterCounts[template[i]]++
		cur[k]++
	}

	for i := 0; i < iterations; i++ {
		newMap := make(map[string]uint64)
		for k, v := range cur {
			n := rules[k]
			l := string([]rune{rune(k[0]), n})
			r := string([]rune{n, rune(k[1])})
			newMap[l] += v
			newMap[r] += v
			letterCounts[n] += v
		}
		cur = newMap
	}

	min, max := uint64(math.MaxUint64), uint64(0)
	for _, v := range letterCounts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}

func part1(file string) uint64 {
	return polymerise(file, 10)
}

func part2(file string) uint64 {
	return polymerise(file, 40)
}

func parseInput(file string) ([]rune, map[string]rune) {
	var template []rune
	rules := make(map[string]rune)
	for _, line := range u.ReadLinesStrings(file) {
		if line == "" {
			continue
		}
		if strings.Contains(line, " -> ") {
			spl := strings.Split(line, " -> ")
			rules[spl[0]] = rune(spl[1][0])
		} else {
			template = []rune(line)
		}
	}

	return template, rules
}
