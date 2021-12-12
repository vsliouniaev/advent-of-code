package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/nav"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/11/input")) // 2178
	fmt.Printf("Part 2: %d\n", part2("2020/11/input")) // 1978
}

const (
	Empty    = 'L'
	Occupied = '#'
	Floor    = '.'
)

func part1(file string) int {
	nextLines := util.ReadLinesRuneGrid(file)

	for {
		g := nextLines
		nextLines = g.Clone()
		changed := 0
		for i, p := g.GetIterator(); p != nil; p = i.Next() {
			cell := g.Get(p).(rune)
			if cell == Floor {
				continue
			}
			neighbours := p.AllNeighbours(g)
			occ := 0
			for _, n := range neighbours {
				if g.Get(n).(rune) == Occupied {
					occ++
				}
			}

			if occ >= 4 && cell == Occupied {
				nextLines.Set(p, Empty)
				changed++
			}

			if occ == 0 && cell == Empty {
				nextLines.Set(p, Occupied)
				changed++
			}
		}

		if changed == 0 {
			break
		}
	}

	return countOccupied(nextLines)
}

func countOccupied(lines nav.Grid) int {
	cnt := 0
	for i, p := lines.GetIterator(); p != nil; p = i.Next() {
		if lines.Get(p).(rune) == Occupied {
			cnt++
		}
	}
	return cnt
}

func occupied(p *nav.Point, g nav.Grid) int {
	getResult := func(addx, addy int) int {
		pp := &nav.Point{X: p.X + addx, Y: p.Y + addy}
		for g.Contains(pp) {
			switch g.Get(pp).(rune) {
			case Floor:
				pp.X += addx
				pp.Y += addy
			case Occupied:
				return 1
			case Empty:
				return 0
			}
		}
		return 0
	}

	return getResult(+1, +0) + // e
		getResult(+1, -1) + // se
		getResult(+0, -1) + // s
		getResult(-1, -1) + // sw
		getResult(-1, +0) + // w
		getResult(-1, +1) + // nw
		getResult(-0, +1) + // n
		getResult(+1, +1) // ne

}
func part2(file string) int {
	nextLines := util.ReadLinesRuneGrid(file)

	for {
		lines := nextLines
		nextLines = lines.Clone()
		changed := 0
		for i, p := lines.GetIterator(); p != nil; p = i.Next() {
			cell := lines.Get(p).(rune)
			if cell == Floor {
				continue
			}

			occ := occupied(p, lines)
			if occ >= 5 && cell == Occupied {
				nextLines.Set(p, Empty)
				changed++
			}

			if occ == 0 && cell == Empty {
				nextLines.Set(p, Occupied)
				changed++
			}

		}
		if changed == 0 {
			break
		}
	}

	return countOccupied(nextLines)
}
