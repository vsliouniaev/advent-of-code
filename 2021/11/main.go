package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/nav"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 1679
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input"))) // 519
}

func part1(file string) int {
	grid := nav.IntGrid(u.ReadLinesIntGrid(file))

	flashes := 0
	for s := 0; s < 100; s++ {

		// Pass 1 : Increase all vals by 1
		for i, p := grid.GetIterator(); p != nil; p = i.Next() {
			grid.Set(p, grid.Get(p).(int)+1)
		}

		// Pass 2 : Handle flashes
		again := true
		for again {
			again = false
			for i, p := grid.GetIterator(); p != nil; p = i.Next() {
				if grid.Get(p).(int) > 9 {
					again = true
					flashes++
					flash(p, &grid)
				}
			}
		}
	}

	return flashes
}

func part2(file string) int {
	grid := nav.IntGrid(u.ReadLinesIntGrid(file))
	size := (grid.Maxx() + 1) * (grid.Maxy() + 1)
	flashes := 0
	step := 0
	for ; flashes != size; step++ {
		flashes = 0
		// Pass 1 : Increase all vals by 1
		for i, p := grid.GetIterator(); p != nil; p = i.Next() {
			grid.Set(p, grid.Get(p).(int)+1)
		}

		// Pass 2 : Handle flashes
		again := true
		for again {
			again = false
			for i, p := grid.GetIterator(); p != nil; p = i.Next() {
				if grid.Get(p).(int) > 9 {
					again = true
					flashes++
					flash(p, &grid)
				}
			}
		}
	}

	return step
}

func flash(p *nav.Point, g nav.Grid) {
	g.Set(p, 0)
	for _, n := range p.AllNeighbours(g) {
		v := g.Get(n).(int)
		// There was already a +1 pass, so the only 0 encountered have already flashed this step
		if v > 0 {
			g.Set(n, v+1)
		}
	}
}
