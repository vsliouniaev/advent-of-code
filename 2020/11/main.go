package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/11/input"))
	//fmt.Printf("Part 2: %d\n", part2("2020/11/input"))
}

const (
	Empty    = 'L'
	Occupied = '#'
	Blocked  = '.'
)

func part1(file string) int {
	nextLines := util.ReadLinesRunes(file)
	maxy := len(nextLines) - 1
	maxx := len(nextLines[0]) - 1

	i := 0
	for {
		lines := nextLines
		nextLines = makeNextGrid(lines)
		changed := 0
		for y := 0; y <= maxy; y++ {
			for x := 0; x <= maxx; x++ {
				cell := lines[y][x]
				if cell == Blocked {
					continue
				}
				neighbours := point{x: x, y: y}.neighbours(maxx, maxy)
				occ := 0
				for _, n := range neighbours {
					if lines[n.y][n.x] == Occupied {
						occ++
					}
				}

				if occ >= 4 && lines[y][x] == Occupied {
					nextLines[y][x] = Empty
					changed++
				}

				if occ == 0 && lines[y][x] == Empty {
					nextLines[y][x] = Occupied
					changed++
				}
			}
		}
		i++
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

func makeNextGrid(in [][]rune) [][]rune {
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

func (p point) neighbours(maxx, maxy int) []point {
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

func part2(file string) int {
	return 0
}
