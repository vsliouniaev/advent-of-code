package util

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

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}
