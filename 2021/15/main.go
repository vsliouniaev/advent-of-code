package main

import (
	"container/heap"
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/data"
	"github.com/vsliouniaev/aoc/util/nav"
	"math"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(util.RelativeFile("input"))) // 755
	fmt.Printf("Part 2: %d\n", part2(util.RelativeFile("input"))) // 3016
}

func djikstraPathLen(g nav.Grid) int {
	weights := g.Clone()
	ph := data.NewPointHeap()
	for i, p := g.GetIterator(); p != nil; p = i.Next() {
		weights.Set(p, math.MaxInt64)
		ph.Push(data.PointVal{Point: *p, Val: math.MaxInt64})
	}
	heap.Init(ph)

	dest := &nav.Point{X: g.Maxx(), Y: g.Maxy()}
	cur := data.PointVal{Point: nav.Point{}, Val: 0}
	weights.Set(&cur.Point, 0)
	//curVal := 0
	for ph.Len() != 0 {
		neighbours := cur.Point.CompassNeighbours(g)
		for _, n := range neighbours {
			if ph.ContainsPoint(*n) {
				nVal := g.Get(n).(int)
				uVal := weights.Get(n).(int)
				if cur.Val+nVal < uVal {
					ph.Set(data.PointVal{Point: *n, Val: cur.Val + nVal})
					weights.Set(n, cur.Val+nVal)
				}
			}
		}
		if cur.Point == *dest {
			break
		}
		cur = heap.Pop(ph).(data.PointVal)
		if cur.Val == math.MaxInt64 {
			break
		}
	}

	// Walk path backwards to sum weights
	sum := 0
	for cur := dest; !cur.Equals(&nav.Point{X: 0, Y: 0}); {
		sum += g.Get(cur).(int)
		m := math.MaxInt64
		for _, n := range cur.CompassNeighbours(weights) {
			val := weights.Get(n).(int)
			if val < m {
				m = val
				cur = n
			}
		}
	}

	return sum
}

func part1(file string) int {
	g := util.ReadLinesIntGrid(file)
	return djikstraPathLen(g)
}

func part2(file string) int {
	g := upsize(util.ReadLinesIntGrid(file), 5)
	return djikstraPathLen(g)
}

func upsize(g nav.Grid, sz int) nav.Grid {
	u := make([][]interface{}, (g.Maxx()+1)*sz)
	for y := 0; y < len(u); y++ {
		u[y] = make([]interface{}, (g.Maxy()+1)*sz)
		for x := 0; x < len(u[y]); x++ {
			pg := &nav.Point{X: x % (g.Maxx() + 1), Y: y % (g.Maxy() + 1)}
			m := x/(g.Maxx()+1) + y/(g.Maxy()+1)
			vg := g.Get(pg).(int) + m
			if vg > 9 {
				vg = vg - 9
			}
			u[y][x] = vg
		}
	}

	return nav.NewGrid(u)
}
