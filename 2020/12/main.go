package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/navigation"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/12/input")) // 582
	fmt.Printf("Part 2: %d\n", part2("2020/12/input")) // 52069
}

func part1(file string) int {
	loc := navigation.Point{}
	dir := navigation.Orientation{}

	for _, line := range u.ReadLinesStrings(file) {
		m, err := strconv.Atoi(line[1:])
		u.Check(err)
		var vec navigation.Vector
		switch line[0] {
		case 'N':
			vec = navigation.GetVector(navigation.North).Magnitude(m)
		case 'S':
			vec = navigation.GetVector(navigation.South).Magnitude(m)
		case 'E':
			vec = navigation.GetVector(navigation.East).Magnitude(m)
		case 'W':
			vec = navigation.GetVector(navigation.West).Magnitude(m)
		case 'F':
			vec = dir.Forward().Magnitude(m)
		case 'R':
			dir.Turn((navigation.Right).Degrees(m))
		case 'L':
			dir.Turn((navigation.Left).Degrees(m))
		default:
			panic(line[0])
		}
		loc.Move(vec)
	}
	return navigation.Point{}.Manhattan(loc)
}

func part2(file string) int {
	ship := navigation.Point{}
	waypoint := navigation.Point{X: 10, Y: 1}
	for _, line := range u.ReadLinesStrings(file) {
		m, err := strconv.Atoi(line[1:])
		u.Check(err)
		switch line[0] {
		case 'N':
			waypoint.Move(navigation.GetVector(navigation.North).Magnitude(m))
		case 'S':
			waypoint.Move(navigation.GetVector(navigation.South).Magnitude(m))
		case 'E':
			waypoint.Move(navigation.GetVector(navigation.East).Magnitude(m))
		case 'W':
			waypoint.Move(navigation.GetVector(navigation.West).Magnitude(m))
		case 'F':
			ship.Move(waypoint.Vector().Magnitude(m))
		case 'R':
			waypoint.Rotate((navigation.Right).Degrees(m))
		case 'L':
			waypoint.Rotate((navigation.Left).Degrees(m))
		}
	}
	return navigation.Point{}.Manhattan(ship)
}
