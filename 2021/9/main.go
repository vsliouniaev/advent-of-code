package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/data"
	"github.com/vsliouniaev/aoc/util/nav"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("input")) // 494
	fmt.Printf("Part 2: %d\n", part2("input")) // 1048128
}

func part1(file string) int {
	g := util.ReadLinesIntGrid(file)
	riskLevels := 0
	for i, p := g.GetIterator(); p != nil; p = i.Next() {
		v := g.Get(p).(int)
		lowest := true
		for _, s := range p.CompassNeighbours(g) {
			if g.Get(s).(int) <= v {
				lowest = false
			}
		}
		if lowest {
			riskLevels += v + 1
		}

	}
	return riskLevels
}

func part2(file string) int {
	g := util.ReadLinesIntGrid(file)
	visited := make(map[string]struct{})
	var basins []int
	for i, p := g.GetIterator(); p != nil; p = i.Next() {
		if _, ok := visited[p.String()]; ok {
			continue
		}
		visited[p.String()] = struct{}{}
		if g.Get(p) == 9 {
			continue
		}
		b := basin(p, g)
		basins = append(basins, len(b))
		for v := range b {
			visited[v] = struct{}{}
		}

	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func basin(p *nav.Point, g nav.Grid) map[string]struct{} {
	q := &data.Queue{}
	visited := map[string]struct{}{p.String(): {}}

	addNeighbours := func(p *nav.Point) {
		for _, n := range p.CompassNeighbours(g) {
			v := g.Get(n).(int)
			if v != 9 {
				q.Push(n)
			}
		}
	}

	addNeighbours(p)

	for q.Len() != 0 {
		s := q.Pop().(*nav.Point)
		// Check if already visited
		if _, ok := visited[s.String()]; ok {
			continue
		}
		// And add to visited hash
		visited[s.String()] = struct{}{}
		addNeighbours(s)
	}

	return visited
}
