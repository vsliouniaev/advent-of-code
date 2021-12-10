package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 387
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input"))) // 6428
}

func part1(file string) int {
	return do(file, 2020)
}

func part2(file string) int {
	return do(file, 30000000)
}

type num struct {
	prevTurn     int
	prevPrevTurn int
}

func do(file string, limit int) int {
	spoken := make(map[int]*num)
	lastSpoken := -1
	currentTurn := 0
	for _, c := range strings.Split(u.ReadLinesStrings(file)[0], ",") {
		currentTurn++
		i, err := strconv.Atoi(c)
		u.Check(err)
		spoken[i] = &num{
			prevTurn:     currentTurn,
			prevPrevTurn: currentTurn,
		}
		lastSpoken = i
	}

	currentTurn++
	for ; currentTurn <= limit; currentTurn++ {
		numToSpeak := 0
		if l, wasSpoken := spoken[lastSpoken]; wasSpoken {
			numToSpeak = l.prevTurn - l.prevPrevTurn
		}
		if w, wasSpoken := spoken[numToSpeak]; wasSpoken {
			w.prevPrevTurn = w.prevTurn
			w.prevTurn = currentTurn
		} else {
			spoken[numToSpeak] = &num{
				prevTurn:     currentTurn,
				prevPrevTurn: currentTurn,
			}
		}
		lastSpoken = numToSpeak
	}

	return lastSpoken
}
