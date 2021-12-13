package main

import (
	u "github.com/vsliouniaev/aoc/util"
	"testing"
)

func BenchmarkPart1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		part1(u.RelativeFile("input"))
	}
}

func BenchmarkPart2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		part2(u.RelativeFile("input"))
	}
}
