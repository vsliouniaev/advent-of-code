package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	math2 "github.com/vsliouniaev/aoc/util/math"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2020/13/input")) // 410
	fmt.Printf("Part 2: %d\n", part2("2020/13/input")) // 600691418730595
}

func part1(file string) int {
	lines := u.ReadLinesStrings(file)
	target, err := strconv.Atoi(lines[0])
	u.Check(err)

	min := math.MaxInt64
	sel := 0
	for _, b := range strings.Split(lines[1], ",") {
		busInterval, err := strconv.Atoi(b)
		if err != nil {
			continue
		}

		numBusesBefore := target / busInterval
		firstBusAfter := numBusesBefore + 1
		offset := (busInterval * firstBusAfter) - target

		if min > offset {
			min = offset
			sel = busInterval * (offset)
		}
	}

	return sel
}

func part2(file string) int64 {
	lines := u.ReadLinesStrings(file)
	var buses math2.ChineseRemainder
	for i, n := range strings.Split(lines[1], ",") {
		if b, err := strconv.Atoi(n); err == nil {
			buses = append(buses, &math2.ChineseRemainderArg{
				Remainder: big.NewInt(int64(b - i)),
				Modulus:   big.NewInt(int64(b)),
			})
		}
	}

	return buses.Solve().Int64()
}
