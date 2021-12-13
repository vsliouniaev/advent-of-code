package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
	"regexp"
	"strconv"
)

func main() {

	fmt.Printf("Part 1: %d\n", part1(RelativeFile("input"))) // 101
	fmt.Printf("Part 2: %d\n", part2(RelativeFile("input"))) // 108636
}

var (
	nrx    = regexp.MustCompile(`^(\w+\s\w+)\sbags?`)
	crx    = regexp.MustCompile(`(\d)\s(\w+\s\w+)\sbags?`)
	target = "shiny gold"
)

func part1(file string) int {
	search := make(map[string][]string)
	for _, line := range ReadLinesStrings(file) {
		outer := nrx.FindStringSubmatch(line)[1]
		for _, sub := range crx.FindAllStringSubmatch(line, -1) {
			child := sub[2]
			search[child] = append(search[child], outer)
		}
	}

	roots := make(map[string]struct{})
	cur := map[string]struct{}{target: {}}
	for len(cur) > 0 {
		c := any(cur)
		results, ok := search[c]
		roots[c] = struct{}{}
		if ok {
			add(cur, results)
		}
		delete(cur, c)
	}
	delete(roots, target)
	return len(roots)
}

func part2(file string) int {
	var err error
	search := make(map[string][]bagq)
	for _, line := range ReadLinesStrings(file) {
		outer := nrx.FindStringSubmatch(line)[1]
		for _, sub := range crx.FindAllStringSubmatch(line, -1) {
			b := bagq{}
			b.quantity, err = strconv.Atoi(sub[1])
			b.name = sub[2]
			search[outer] = append(search[outer], b)
			Check(err)
		}
	}

	sum := 0
	bags := []bagq{{target, 1}}
	for len(bags) > 0 {
		b := bags[0]
		contents, ok := search[b.name]
		if ok {
			for _, c := range contents {
				sum += c.quantity * b.quantity
				bags = append(bags, bagq{c.name, c.quantity * b.quantity})
			}
		}
		bags = bags[1:]
	}

	return sum
}

type bagq struct {
	name     string
	quantity int
}

func add(m map[string]struct{}, strings []string) {
	for _, s := range strings {
		m[s] = struct{}{}
	}
}

func any(m map[string]struct{}) string {
	for k := range m {
		return k
	}
	return ""
}
