package nav

import (
	"fmt"
	"strings"
)

type Grid interface {
	Maxx() int
	Maxy() int
	Set(p *Point, i interface{})
	Get(p *Point) interface{}
	GetIterator() (*Gridterator, *Point)
}

// Not sure this is actually more readable than the normal syntax
// the only real benefit is that I don't need to construct a *Point a bunch of times.
type Gridterator struct {
	x    int
	y    int
	maxx int
	maxy int
}

func (g *Gridterator) Next() *Point {
	if g.x >= g.maxx {
		g.x = -1
		g.y++
	}
	if g.y > g.maxy {
		return nil
	}
	g.x++

	return &Point{
		X: g.x,
		Y: g.y,
	}
}

func (g IntGrid) GetIterator() (*Gridterator, *Point) {
	i := &Gridterator{
		x:    -1,
		y:    0,
		maxx: g.Maxx(),
		maxy: g.Maxy(),
	}
	return i, i.Next()
}

type IntGrid [][]int

func (g IntGrid) Maxx() int {
	return len(g) - 1
}

func (g IntGrid) Maxy() int {
	return len(g[0]) - 1
}

func (g IntGrid) Get(p *Point) interface{} {
	return g[p.Y][p.X]
}

func (g IntGrid) Set(p *Point, i interface{}) {
	g[p.Y][p.X] = i.(int)
}

func (g IntGrid) String() string {
	sb := strings.Builder{}
	for y := range g {
		for x := range g[y] {
			sb.WriteString(fmt.Sprint(g[y][x]))
		}
		sb.WriteRune('\n')
	}
	return strings.TrimSuffix(sb.String(), "\n")
}
