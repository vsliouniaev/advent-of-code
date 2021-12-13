package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/nav"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input")))   // 818
	fmt.Printf("Part 2: \n%s\n", part2(u.RelativeFile("input"))) // LRGPRECB
}

// foldLeftOrUp is the entire answer to this question. The rest is just string parsing and printing
func foldLeftOrUp(x, y *int, points map[nav.Point]struct{}) map[nav.Point]struct{} {
	out := make(map[nav.Point]struct{})
	for p := range points {
		if y != nil && p.Y > *y {
			p.Y = *y - (p.Y - *y)
		}
		if x != nil && p.X > *x {
			p.X = *x - (p.X - *x)
		}
		out[p] = struct{}{}
	}
	return out
}

func part1(file string) int {
	lines := u.ReadLinesStrings(file)
	points := make(map[nav.Point]struct{})
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if !strings.HasPrefix(line, "fold along") {
			points[parsePoint(line)] = struct{}{}
		} else {
			if strings.HasPrefix(line, "fold along ") {
				x, y := parseFold(line)
				points = foldLeftOrUp(x, y, points)
				break
			}
		}
	}

	return len(points)
}

func part2(file string) string {
	lines := u.ReadLinesStrings(file)
	points := make(map[nav.Point]struct{})
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if !strings.HasPrefix(line, "fold along") {
			points[parsePoint(line)] = struct{}{}
		} else {
			x, y := parseFold(line)
			points = foldLeftOrUp(x, y, points)
		}
	}

	return render(points)
}

func parseFold(line string) (x, y *int) {
	line = strings.TrimPrefix(line, "fold along ")
	spl := strings.Split(line, "=")
	coord, err := strconv.Atoi(spl[1])
	u.Check(err)
	switch spl[0] {
	case "x":
		return &coord, nil
	case "y":
		return nil, &coord
	default:
		panic(spl[0])
	}
}

func parsePoint(line string) nav.Point {
	var err error
	spl := strings.Split(line, ",")
	p := nav.Point{}
	p.X, err = strconv.Atoi(spl[0])
	u.Check(err)
	p.Y, err = strconv.Atoi(spl[1])
	u.Check(err)
	return p
}

func render(points map[nav.Point]struct{}) string {
	maxx, maxy := 0, 0
	for p := range points {
		if p.X > maxx {
			maxx = p.X
		}
		if p.Y > maxy {
			maxy = p.Y
		}
	}
	sb := strings.Builder{}
	for y := 0; y <= maxy; y++ {
		for x := 0; x <= maxx; x++ {
			r := ' '
			if _, ok := points[nav.Point{X: x, Y: y}]; ok {
				r = '#'
			}
			sb.WriteRune(r)
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
