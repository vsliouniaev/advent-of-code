package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 383
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input"))) // 272
}

func part1(file string) int {
	corrects := 0
	for _, line := range u.ReadLinesStrings(file) {
		s := strings.Split(line, ": ")
		lower, upper, letter := parseCriteria(s[0])
		password := s[1]
		matches := 0
		for _, c := range password {
			if c == letter {
				matches++
			}
		}
		if matches >= lower && matches <= upper {
			corrects++
		}
	}
	return corrects
}

func part2(file string) int {
	corrects := 0
	for _, line := range u.ReadLinesStrings(file) {
		s := strings.Split(line, ": ")
		lower, upper, letter := parseCriteria(s[0])
		lower--
		upper--
		password := s[1]
		l := rune(password[lower])
		u := rune(password[upper])
		if l != u && (l == letter || u == letter) {
			corrects++
		}
	}
	return corrects
}

func parseCriteria(s string) (lower int, upper int, letter rune) {
	split := strings.Split(s, " ")
	limits := strings.Split(split[0], "-")
	l, err := strconv.ParseInt(limits[0], 10, 64)
	lower = int(l)
	u.Check(err)
	up, err := strconv.ParseInt(limits[1], 10, 64)
	upper = int(up)
	u.Check(err)
	letter = rune(split[1][0])
	return
}
