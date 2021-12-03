package _go

import (
	"fmt"
	"strconv"
	"strings"
)

type opcodefunc = func([]int, int) int

type Day2 struct{}

var ops = map[int]opcodefunc{
	1: func(buf []int, i int) int {
		buf[buf[i+3]] = buf[buf[i+1]] + buf[buf[i+2]]
		return i + 4
	},
	2: func(buf []int, i int) int {
		buf[buf[i+3]] = buf[buf[i+1]] * buf[buf[i+2]]
		return i + 4
	},
	99: func(buf []int, i int) int {
		return -1
	},
}

func (d Day2) part1(data []int, n, v int) int {
	i := 0
	buf := make([]int, len(data))
	copy(buf, data)
	buf[1] = n
	buf[2] = v
	for i != -1 {
		i = ops[buf[i]](buf, i)
	}
	return buf[0]
}

func (d Day2) part2(data []int) int {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			if d.part1(data, n, v) == 19690720 {
				return (100 * n) + v
			}
		}
	}
	return -1
}

func (d Day2) parse(file string) (ops []int) {
	lines := ParseFile(file)
	for _, opstr := range strings.Split(lines[0], ",") {
		op, _ := strconv.Atoi(opstr)
		ops = append(ops, op)
	}
	return
}

func (d Day2) Go() {
	buf := d.parse("in/day2")
	fmt.Println(d.part1(buf, 12, 2))
	fmt.Println(d.part2(buf))
}
