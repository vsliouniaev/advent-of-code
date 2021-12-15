package nav

import (
	"fmt"
	"strings"
)

type Grid interface {
	Maxx() int
	Maxy() int
	Size() int
	Set(p *Point, i interface{})
	Get(p *Point) interface{}
	GetIterator() (*Gridterator, *Point)
	Contains(p *Point) bool
	Clone() Grid
}

// Not sure this is actually more readable than the normal syntax
// the only real benefit is that I don't need to construct a *Point a bunch of times.
type Gridterator struct {
	x    int
	y    int
	maxx int
	maxy int
}

func NewGrid(g [][]interface{}) Grid {
	return grid(g)
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

func (g grid) GetIterator() (*Gridterator, *Point) {
	i := &Gridterator{
		x:    -1,
		y:    0,
		maxx: g.Maxx(),
		maxy: g.Maxy(),
	}
	return i, i.Next()
}

type grid [][]interface{}

func (g grid) Contains(p *Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X <= g.Maxx() && p.Y <= g.Maxy()
}

func (g grid) Clone() Grid {
	var out [][]interface{}
	for y := range g {
		var row []interface{}
		for x := range g[y] {
			row = append(row, g[y][x])
		}
		out = append(out, row)
	}
	return grid(out)
}

func (g grid) Maxx() int {
	return len(g[0]) - 1
}

func (g grid) Maxy() int {
	return len(g) - 1
}

func (g grid) Size() int {
	return len(g) * len(g[0])
}

func (g grid) Get(p *Point) interface{} {
	return g[p.Y][p.X]
}

func (g grid) Set(p *Point, i interface{}) {
	g[p.Y][p.X] = i
}

func (g grid) String() string {
	sb := strings.Builder{}
	for y := range g {
		for x := range g[y] {
			sb.WriteString(fmt.Sprintf("%d", g[y][x]))
		}
		sb.WriteRune('\n')
	}
	return strings.TrimSuffix(sb.String(), "\n")
}
