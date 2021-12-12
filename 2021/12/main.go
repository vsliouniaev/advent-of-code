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

func part1(file string) int {
	graph := buildGraph(u.ReadLinesStrings(file))
	paths := search(start, nil, make(map[string]int), graph, "-")
	return len(paths)
}

func part2(file string) int {
	graph := buildGraph(u.ReadLinesStrings(file))

	// Not all possible paths include variations on visiting some lowe-case cave twice, so there will be multiple
	// identical entries that must be filtered out.
	all := make(map[string]struct{})
	for twiceCave := range graph {
		if twiceCave != start && twiceCave != end && strings.ToLower(twiceCave) == twiceCave {
			paths := search(start, nil, make(map[string]int), graph, twiceCave)
			for _, p := range paths {
				all[strings.Join(p, ",")] = struct{}{}
			}
		}
	}
	return len(all)
}

// search will traverse the graph depth first, building up paths according to the rules in the question. Which paths
// are valid for the search is determined by canVisit.
// Returns a slice of slices of locations, which can be subsequently filtered to remove duplicates when variations of
// twiceCave will return identical paths not involving the twiceCave value
func search(cave string, path []string, visited map[string]int, graph map[string][]string, twiceCave string) [][]string {
	path = append(path, cave)
	if cave == end {
		return [][]string{path}
	}

	visited[cave]++

	var paths [][]string
	for _, loc := range graph[cave] {
		if canVisit(loc, visited, twiceCave) {
			paths = append(paths, search(loc, copySlice(path), copyMap(visited), graph, twiceCave)...)
		}
	}

	return paths
}

// canVisit will return whether a location can be visited by search according to the rules in the question
// - upper-case locations may be visited any number of times
// - lower-case locations may be visited once generally
// - lower-case location specified by twiceCave may be visited twice
func canVisit(loc string, visits map[string]int, twiceCave string) bool {
	return strings.ToUpper(loc) == loc ||
		loc == twiceCave && visits[loc] < 2 ||
		visits[loc] == 0
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

const (
	start = "start"
	end   = "end"
)

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
