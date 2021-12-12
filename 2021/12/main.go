package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 5178
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input"))) // 130094
}

const (
	start = "start"
	end   = "end"
)

func part1(file string) int {
	graph := buildGraph(u.ReadLinesStrings(file))
	paths := rec(start, nil, make(map[string]int), graph, "-")
	return len(paths)
}

func part2(file string) int {
	graph := buildGraph(u.ReadLinesStrings(file))
	all := make(map[string]struct{})
	for k := range graph {
		if k != start && k != end && strings.ToLower(k) == k {
			paths := rec(start, nil, make(map[string]int), graph, k)
			for _, p := range paths {
				all[strings.Join(p, ",")] = struct{}{}
			}
		}
	}
	return len(all)
}

func buildGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)
	for _, l := range lines {
		sp := strings.Split(l, "-")
		graph[sp[0]] = append(graph[sp[0]], sp[1])
		graph[sp[1]] = append(graph[sp[1]], sp[0])
	}
	return graph
}

func rec(pos string, path []string, visited map[string]int, locs map[string][]string, twiceCave string) [][]string {
	path = append(path, pos)
	if pos == end {
		return [][]string{path}
	}

	visited[pos]++

	var paths [][]string
	for _, loc := range locs[pos] {
		if !canVisit(loc, visited, twiceCave) {
			continue
		}

		children := rec(loc, copySlice(path), copyMap(visited), locs, twiceCave)
		for _, child := range children {
			if child[len(child)-1] == end {
				paths = append(paths, child)
			}
		}
	}

	return paths
}

func copyMap(m map[string]int) map[string]int {
	c := make(map[string]int)
	for k, v := range m {
		c[k] = v
	}
	return c
}

func copySlice(s []string) []string {
	c := make([]string, len(s))
	copy(c, s)
	return c
}

func canVisit(loc string, visits map[string]int, twiceCave string) bool {
	if strings.ToUpper(loc) == loc {
		return true
	}

	if loc == twiceCave && visits[loc] < 2 {
		return true
	}

	return visits[loc] == 0
}
