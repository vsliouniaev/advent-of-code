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

func BenchmarkWithoutFileLoad(t *testing.B) {
	d := u.ReadLinesStrings(u.RelativeFile("input"))[0]
	for i := 0; i < t.N; i++ {
		evalExpression(d)
	}
}

func TestDay16Samples(t *testing.T) {

	tests := []struct {
		input string
		want  int64
	}{
		{"D2FE28", 2021},
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
		{"8A004A801A8002F478", 15},
		{"620080001611562C8802118E34", 46},
		{"C0015000016115A2E0802F182340", 46},
		{"A0016C880162017C3686B18A3D4780", 54},
		{u.ReadLinesStrings(u.RelativeFile("input"))[0], 13476220616073},
	}

	for _, tc := range tests {
		actual := evalExpression(tc.input)
		if tc.want != actual {
			t.Errorf("%s: want: %d but got %d", tc.input, tc.want, actual)
		}
	}
}
