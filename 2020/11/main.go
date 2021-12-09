package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
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
	nextLines := grid(util.ReadLinesRunes(file))
	maxy := len(nextLines) - 1
	maxx := len(nextLines[0]) - 1

	for {
		lines := nextLines
		nextLines = makeNextGrid(lines)
		changed := 0
		for y := 0; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				p := point{x: x, y: y}
				cell := lines.val(p)
				if cell == Floor {
					continue
				}
				neighbours := p.neighbours(nextLines)
				occ := 0
				for _, n := range neighbours {
					if lines.val(n) == Occupied {
						occ++
					}
				}

				if occ >= 4 && cell == Occupied {
					nextLines.set(p, Empty)
					changed++
				}

				if occ == 0 && cell == Empty {
					nextLines.set(p, Occupied)
					changed++
				}
			}
		}

		if changed == 0 {
			break
		}
	}

	return countOccupied(nextLines)
}

func countOccupied(lines [][]rune) int {
	cnt := 0
	for _, row := range lines {
		for _, c := range row {
			if c == Occupied {
				cnt++
			}
		}
	}
	return cnt
}

func makeNextGrid(in [][]rune) grid {
	out := make([][]rune, len(in))
	for y, row := range in {
		out[y] = make([]rune, len(row))
		for x, c := range row {
			out[y][x] = c
		}
	}

	return out
}

type point struct {
	x int
	y int
}

func (p point) neighbours(g grid) []point {
	maxy := len(g) - 1
	maxx := len(g[0]) - 1
	all := []point{
		{x: p.x + 1, y: p.y - 0}, // e
		{x: p.x + 1, y: p.y - 1}, // se
		{x: p.x + 0, y: p.y - 1}, // s
		{x: p.x - 1, y: p.y - 1}, // sw

		{x: p.x - 1, y: p.y + 0}, // w
		{x: p.x - 1, y: p.y + 1}, // nw
		{x: p.x - 0, y: p.y + 1}, // n
		{x: p.x + 1, y: p.y + 1}, // ne
	}
	// Remove out of range
	var pts []point
	for _, a := range all {
		if a.x >= 0 && a.y >= 0 && a.x <= maxx && a.y <= maxy {
			pts = append(pts, a)
		}
	}

	return pts
}

func occupied(p point, g grid) int {
	maxy := len(g) - 1
	maxx := len(g[0]) - 1
	getResult := func(addx, addy int) int {
		x := p.x + addx
		y := p.y + addy
		for x >= 0 && y >= 0 && x <= maxx && y <= maxy {
			switch g[y][x] {
			case Floor:
				x += addx
				y += addy
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

type grid [][]rune

func (g grid) val(p point) rune {
	return g[p.y][p.x]
}

func (g grid) set(p point, r rune) {
	g[p.y][p.x] = r
}

func part2(file string) int {
	nextLines := grid(util.ReadLinesRunes(file))
	maxy := len(nextLines) - 1
	maxx := len(nextLines[0]) - 1

	for {
		lines := nextLines
		nextLines = makeNextGrid(lines)
		changed := 0
		for y := 0; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				p := point{x: x, y: y}
				cell := lines.val(p)
				if cell == Floor {
					continue
				}

				occ := occupied(p, lines)
				if occ >= 5 && cell == Occupied {
					nextLines.set(p, Empty)
					changed++
				}

				if occ == 0 && cell == Empty {
					nextLines.set(p, Occupied)
					changed++
				}
			}
		}
		if changed == 0 {
			break
		}
	}

	return countOccupied(nextLines)
}
