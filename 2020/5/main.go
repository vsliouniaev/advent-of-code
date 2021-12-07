package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
	"sort"
)

func main() {

	fmt.Printf("Part 1: %d\n", part1("2020/5/input")) // 890
	fmt.Printf("Part 2: %d\n", part2("2020/5/input"))
}

func part1(file string) int {
	h := 0
	for _, l := range ReadLinesStrings(file) {
		rs := 0
		re := 127
		cs := 0
		ce := 7
		for i := range l {
			midrow := (re - rs) / 2
			midcol := (ce - cs) / 2
			switch l[i] {
			case 'F':
				re = rs + midrow
			case 'B':
				rs = re - midrow
			case 'L':
				ce = cs + midcol
			case 'R':
				cs = ce - midcol
			default:
				panic(l[i])
			}
		}
		// 4 5
		if cs != ce {
			panic("cs!=ce")
		}
		if re != rs {
			panic("re!=rs")
		}

		id := (rs)*8 + cs
		if id > h {
			h = id
		}
	}

	return h
}

func part2(file string) int {
	var ints []int
	for _, l := range ReadLinesStrings(file) {
		rs := 0
		re := 127
		cs := 0
		ce := 7
		for i := range l {
			midrow := (re - rs) / 2
			midcol := (ce - cs) / 2
			switch l[i] {
			case 'F':
				re = rs + midrow
			case 'B':
				rs = re - midrow
			case 'L':
				ce = cs + midcol
			case 'R':
				cs = ce - midcol
			default:
				panic(l[i])
			}
		}

		id := rs*8 + cs
		ints = append(ints, id)
	}

	sort.Ints(ints)
	for i := 1; i < len(ints); i++ {
		if ints[i] != 1+ints[i-1] {
			return ints[i] - 1
		}
	}
	return 0
}
