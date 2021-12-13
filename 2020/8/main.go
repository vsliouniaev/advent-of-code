package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(util.RelativeFile("input"))) // 1586
	fmt.Printf("Part 2: %d\n", part2(util.RelativeFile("input"))) // 703
}

func part1(file string) int {
	lines := util.ReadLinesStrings(file)
	visited := make([]bool, len(lines))
	acc := 0
	i := 0
	for {
		if visited[i] == true {
			return acc
		}
		visited[i] = true
		line := strings.Split(lines[i], " ")
		command := line[0]
		arg, err := strconv.Atoi(line[1])
		util.Check(err)
		switch command {
		case "acc":
			acc += arg
			i++
		case "nop":
			i++
		case "jmp":
			i += arg
		default:
			panic(command)
		}
	}
}

func part2(file string) int {
	lines := util.ReadLinesStrings(file)
	visited := make([]bool, len(lines))
	changec := 0
	acc := 0
	i := 0
	cmd := 0
	for i < len(lines) {
		if visited[i] == true {
			visited = make([]bool, len(lines))
			changec++
			acc = 0
			i = 0
			cmd = 0
		}
		visited[i] = true
		line := strings.Split(lines[i], " ")
		command := line[0]
		arg, err := strconv.Atoi(line[1])
		util.Check(err)
		switch command {
		case "acc":
			acc += arg
			i++
		case "nop":
			if cmd == changec {
				i += arg
			} else {
				i++
			}
			cmd++
		case "jmp":
			if cmd == changec {
				i++
			} else {
				i += arg
			}
			cmd++
		default:
			panic(command)
		}
	}

	return acc
}
