package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2021/5/input")) // 5145
	fmt.Printf("Part 2: %d\n", part2("2021/5/input")) // 16518
}

func part1(file string) int {
	pos := make(map[string]int)
	for _, line := range ReadLinesStrings(file) {
		// Could rewrite this in terms of part2
		seg := parse(line)
		var s int
		var e int
		var f string
		if seg.isHorz() {
			f = "%d," + strconv.Itoa(seg.sy)
			s = seg.sx
			e = seg.ex
		} else if seg.isVert() {
			f = strconv.Itoa(seg.sx) + ",%d"
			s = seg.sy
			e = seg.ey
		} else {
			continue
		}
		if e < s {
			s, e = e, s
		}
		for ; s <= e; s++ {
			p := fmt.Sprintf(f, s)
			count := pos[p]
			pos[p] = count + 1
		}
	}
	count := 0
	for _, v := range pos {
		if v > 1 {
			count++
		}
	}
	return count
}

func part2(file string) int {
	pos := make(map[string]int)
	for _, line := range ReadLinesStrings(file) {
		seg := parse(line)
		for {
			x, y, zero := seg.shrink()
			p := fmt.Sprintf("%d,%d", x, y)
			count := pos[p]
			pos[p] = count + 1
			if zero {
				break
			}
		}
	}
	count := 0
	for _, v := range pos {
		if v > 1 {
			count++
		}
	}
	return count
}

func parse(line string) (s segment) {
	var err error
	se := strings.Split(line, " -> ")
	start := strings.Split(se[0], ",")
	end := strings.Split(se[1], ",")

	s.sx, err = strconv.Atoi(start[0])
	Check(err)
	s.sy, err = strconv.Atoi(start[1])
	Check(err)

	s.ex, err = strconv.Atoi(end[0])
	Check(err)
	s.ey, err = strconv.Atoi(end[1])
	Check(err)
	return
}

type segment struct {
	sx int
	ex int
	sy int
	ey   int
}

func (s *segment) isHorzVert() bool {
	return s.isHorz() || s.isVert()
}

func (s *segment) isHorz() bool {
	return s.sy == s.ey
}

func (s *segment) isVert() bool {
	return s.sx == s.ex
}

// shrink moves the start point of the segment towards the end by 1 step. It returns the old (un-shrunk) x,y coordinates
// of the start point and whether the length of the remaining segment is zero - i.e. if it makes sense to shrink the
// segment again.
func (s *segment) shrink() (x, y int, zero bool) {
	zero = true
	x = s.sx
	y = s.sy
	dx := s.ex - s.sx
	dy := s.ey - s.sy
	if dx > 0 {
		zero = false
		s.sx++
	} else if dx < 0 {
		zero = false
		s.sx--
	}
	if dy > 0 {
		zero = false
		s.sy++
	} else if dy < 0 {
		zero = false
		s.sy--
	}
	return
}
