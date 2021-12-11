package nav

import "fmt"

var directions = []Vector{
	{X: +1, Y: +0}, // e
	{X: -0, Y: +1}, // n
	{X: -1, Y: -0}, // w
	{X: +0, Y: -1}, // s
}

type Direction int
type Turn int
type Vector struct {
	X int
	Y int
}

func (v Vector) Magnitude(m int) Vector {
	return Vector{
		X: v.X * m,
		Y: v.Y * m,
	}
}

func (t Turn) Degrees(i int) Turn {
	return Turn((i / 90) * int(t))
}

const (
	East  Direction = 0
	North           = 1
	West            = 2
	South           = 3
	Left  Turn      = 1
	Right Turn      = -1
)

func (e Direction) String() string {
	switch e {
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	case North:
		return "North"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Orientation struct {
	o int
}

// Turn changes the direction
func (o *Orientation) Turn(t Turn) *Orientation {
	o.o = (o.o + int(t)) % len(directions)
	if o.o < 0 {
		o.o += len(directions)
	}
	return o
}

func (o *Orientation) GetDirection() Direction {
	return Direction(o.o)
}

func (o *Orientation) Forward() Vector {
	return directions[o.o]
}

func (o Orientation) Backward() Vector {
	return directions[(o.o+2)%len(directions)]
}

func GetVector(d Direction) Vector {
	return directions[d]
}
