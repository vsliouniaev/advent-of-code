package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/nav"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/12/input")) // 582
	fmt.Printf("Part 2: %d\n", part2("2020/12/input")) // 52069
}

func part1(file string) int {
	loc := nav.Point{}
	dir := nav.Orientation{}

	for _, line := range u.ReadLinesStrings(file) {
		m, err := strconv.Atoi(line[1:])
		u.Check(err)
		var vec nav.Vector
		switch line[0] {
		case 'N':
			vec = nav.GetVector(nav.North).Magnitude(m)
		case 'S':
			vec = nav.GetVector(nav.South).Magnitude(m)
		case 'E':
			vec = nav.GetVector(nav.East).Magnitude(m)
		case 'W':
			vec = nav.GetVector(nav.West).Magnitude(m)
		case 'F':
			vec = dir.Forward().Magnitude(m)
		case 'R':
			dir.Turn((nav.Right).Degrees(m))
		case 'L':
			dir.Turn((nav.Left).Degrees(m))
		default:
			panic(line[0])
		}
		loc.Move(vec)
	}
	return nav.Point{}.Manhattan(loc)
}

func part2(file string) int {
	ship := nav.Point{}
	waypoint := nav.Point{X: 10, Y: 1}
	for _, line := range u.ReadLinesStrings(file) {
		m, err := strconv.Atoi(line[1:])
		u.Check(err)
		switch line[0] {
		case 'N':
			waypoint.Move(nav.GetVector(nav.North).Magnitude(m))
		case 'S':
			waypoint.Move(nav.GetVector(nav.South).Magnitude(m))
		case 'E':
			waypoint.Move(nav.GetVector(nav.East).Magnitude(m))
		case 'W':
			waypoint.Move(nav.GetVector(nav.West).Magnitude(m))
		case 'F':
			ship.Move(waypoint.Vector().Magnitude(m))
		case 'R':
			waypoint.Rotate((nav.Right).Degrees(m))
		case 'L':
			waypoint.Rotate((nav.Left).Degrees(m))
		}
	}
	return nav.Point{}.Manhattan(ship)
}
