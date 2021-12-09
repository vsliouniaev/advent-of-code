package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/12/input")) // 582
	fmt.Printf("Part 2: %d\n", part2("2020/12/input")) // 52069
}

func part1(file string) int {
	loc := u.Point{}
	dir := u.Orientation{}

	for _, line := range u.ReadLinesStrings(file) {
		m, err := strconv.Atoi(line[1:])
		u.Check(err)
		var vec u.Vector
		switch line[0] {
		case 'N':
			vec = u.GetVector(u.North).Magnitude(m)
		case 'S':
			vec = u.GetVector(u.South).Magnitude(m)
		case 'E':
			vec = u.GetVector(u.East).Magnitude(m)
		case 'W':
			vec = u.GetVector(u.West).Magnitude(m)
		case 'F':
			vec = dir.Forward().Magnitude(m)
		case 'R':
			dir.Turn((u.Right).Degrees(m))
		case 'L':
			dir.Turn((u.Left).Degrees(m))
		default:
			panic(line[0])
		}
		loc.Move(vec)
	}
	return u.Point{}.Manhattan(loc)
}

func part2(file string) int {
	ship := u.Point{}
	waypoint := u.Point{X: 10, Y: 1}
	for _, line := range u.ReadLinesStrings(file) {
		m, err := strconv.Atoi(line[1:])
		u.Check(err)
		switch line[0] {
		case 'N':
			waypoint.Move(u.GetVector(u.North).Magnitude(m))
		case 'S':
			waypoint.Move(u.GetVector(u.South).Magnitude(m))
		case 'E':
			waypoint.Move(u.GetVector(u.East).Magnitude(m))
		case 'W':
			waypoint.Move(u.GetVector(u.West).Magnitude(m))
		case 'F':
			ship.Move(waypoint.Vector().Magnitude(m))
		case 'R':
			waypoint.Rotate((u.Right).Degrees(m))
		case 'L':
			waypoint.Rotate((u.Left).Degrees(m))
		}
	}
	return u.Point{}.Manhattan(ship)
}
