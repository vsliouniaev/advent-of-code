package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2021/9/input")) // 494
	fmt.Printf("Part 2: %d\n", part2("2021/9/input")) // 1048128
}

func part1(file string) int {
	g := grid(util.ReadLinesIntGrid(file))
	riskLevels := 0
	for y := range g {
		for x := range g[y] {
			v := g[y][x]
			surround := point{x: x, y: y}.neighbours(g)
			lowest := true
			for _, s := range surround {
				if g[s.y][s.x] <= v {
					lowest = false
				}
			}
			if lowest {
				riskLevels += v + 1
			}
		}
	}
	return riskLevels
}

func part2(file string) int {
	g := grid(util.ReadLinesIntGrid(file))
	visited := make(map[string]struct{})
	var basins []int
	for y := range g {
		for x := range g[y] {
			p := &point{x: x, y: y}
			if _, ok := visited[p.String()]; ok {
				continue
			}
			visited[p.String()] = struct{}{}
			if g.val(p) == 9 {
				continue
			}
			b := basin(p, g)
			basins = append(basins, len(b))
			for v := range b {
				visited[v] = struct{}{}
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func basin(p *point, g grid) map[string]struct{} {
	q := &util.Queue{}
	visited := map[string]struct{}{p.String(): {}}

	addNeighbours := func(p *point) {
		for _, n := range p.neighbours(g) {
			v := g.val(n)
			if v != 9 {
				q.Push(n)
			}
		}
	}

	addNeighbours(p)

	for q.Len() != 0 {
		s := q.Pop().(*point)
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

type grid [][]int

func (g grid) val(p *point) int {
	return g[p.y][p.x]
}

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p point) neighbours(grid [][]int) []*point {
	maxy := len(grid) - 1
	maxx := len(grid[0]) - 1
	all := []*point{
		{x: p.x + 1, y: p.y - 0}, // e
		{x: p.x + 0, y: p.y - 1}, // s
		{x: p.x - 1, y: p.y + 0}, // w
		{x: p.x - 0, y: p.y + 1}, // n
	}
	// Remove out of range
	var pts []*point
	for _, a := range all {
		if a.x >= 0 && a.y >= 0 && a.x <= maxx && a.y <= maxy {
			pts = append(pts, a)
		}
	}

	return pts
}
