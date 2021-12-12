package nav

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) Manhattan(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

func (p *Point) Move(vec Vector) {
	p.X += vec.X
	p.Y += vec.Y
}

func (p *Point) Vector() Vector {
	return Vector(*p)
}

// Rotate changes the location of the Point about 0,0
func (p *Point) Rotate(turn Turn) *Point {
	c := math.Cos(float64(turn) * (math.Pi / 2))
	s := math.Sin(float64(turn) * (math.Pi / 2))
	x := float64(p.X)
	y := float64(p.Y)
	p.X, p.Y = int(math.Round(x*c-y*s)), int(math.Round(x*s+y*c))
	return p
}

func (p *Point) AllNeighbours(grid Grid) []*Point {
	return filterOutOfRange([]*Point{
		{X: p.X + 1, Y: p.Y - 0}, // e
		{X: p.X + 1, Y: p.Y - 1}, // se
		{X: p.X + 0, Y: p.Y - 1}, // s
		{X: p.X - 1, Y: p.Y - 1}, // swX
		{X: p.X - 1, Y: p.Y + 0}, // w
		{X: p.X - 1, Y: p.Y + 1}, // nw
		{X: p.X - 0, Y: p.Y + 1}, // n
		{X: p.X + 1, Y: p.Y + 1}, // ne
	}, grid)
}

func (p *Point) CompassNeighbours(grid Grid) []*Point {
	return filterOutOfRange([]*Point{
		{X: p.X + 1, Y: p.Y - 0}, // e
		{X: p.X + 0, Y: p.Y - 1}, // s
		{X: p.X - 1, Y: p.Y + 0}, // w
		{X: p.X - 0, Y: p.Y + 1}, // n
	}, grid)
}

func (p *Point) Clone() *Point {
	return &Point{
		X: p.X,
		Y: p.Y,
	}
}

func filterOutOfRange(all []*Point, g Grid) []*Point {
	var pts []*Point
	for _, a := range all {
		if a.X >= 0 && a.Y >= 0 && a.X <= g.Maxx() && a.Y <= g.Maxy() {
			if a.Y == 91 {
				panic("wat")
			}
			pts = append(pts, a)
		}
	}
	return pts
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}
