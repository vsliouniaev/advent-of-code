package main

import (
	"fmt"
	. "github.com/vsliouniaev/aoc/2021/util"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("4/input")) // 16674
	fmt.Printf("Part 2: %d\n", part2("4/input")) // 7075
}

func part1(file string) int {
	lines := ReadLinesStrings(file)

	calls := parseCalls(lines)
	boards := parseBoards(lines)

	for _, c := range calls {
		for _, b := range boards {
			if b.mark(c) {
				if b.isWinner() {
					return b.sumUnmarked() * c
				}
			}
		}
	}

	return 0
}

func part2(file string) int {
	lines := ReadLinesStrings(file)

	calls := parseCalls(lines)
	boards := parseBoards(lines)

	var lastBoard *bingo
	var lastCall int
	for _, c := range calls {
		for _, board := range boards {
			if !board.won && board.mark(c) {
				if board.isWinner() {
					board.won = true
					lastBoard = board
					lastCall = c
				}
			}
		}
	}

	return lastBoard.sumUnmarked() * lastCall
}

func parseBoards(lines []string) (boards []*bingo) {
	var board [][]int
	for i := 1; i < len(lines); i++ {
		l := lines[i]
		if l == "" {
			if board != nil {
				boards = append(boards, fromSlices(board))
			}
			board = [][]int{}
			continue
		}

		var row []int
		for _, c := range strings.Split(l, " ") {
			if c == "" {
				continue
			}
			i, err := strconv.Atoi(c)
			Check(err)
			row = append(row, i)
		}
		board = append(board, row)
	}
	boards = append(boards, fromSlices(board))
	return
}

func parseCalls(lines []string) (calls []int) {
	for _, c := range strings.Split(lines[0], ",") {
		i, err := strconv.Atoi(c)
		Check(err)
		calls = append(calls, i)
	}

	return
}

type bingo struct {
	board [][]int
	markd [][]bool
	won   bool
}

func fromSlices(slices [][]int) *bingo {
	h := len(slices)
	w := len(slices[0])
	b := &bingo{board: slices}
	for i := 0; i < h; i++ {
		b.markd = append(b.markd, make([]bool, w))
	}

	return b
}

func (b *bingo) mark(i int) bool {
	for row := range b.board {
		for col := range b.board[row] {
			if b.board[row][col] == i {
				b.markd[row][col] = true
				return true
			}
		}
	}

	return false
}

func (b *bingo) isWinner() bool {
	for row := range b.markd {
		rowComplete := true
		for col := range b.markd[row] {
			if !b.markd[row][col] {
				rowComplete = false
				break
			}
		}
		if rowComplete {
			return true
		}
	}

	for col := 0; col < len(b.markd[0]); col++ {
		colComplete := true
		for row := 0; row < len(b.markd); row++ {
			if !b.markd[row][col] {
				colComplete = false
				continue
			}
		}
		if colComplete {
			return true
		}
	}

	// Board x, y has to be identical
	// diagonal \
	diagComplete := true
	var row, col = 0, 0
	for row < len(b.markd) {
		if !b.markd[row][col] {
			diagComplete = false
			break
		}
		row++
		col++
	}
	if diagComplete {
		return true
	}

	// diagonal /
	diagComplete = true
	row, col = len(b.markd)-1, len(b.markd)-1
	for row != 0 {
		if !b.markd[row][col] {
			diagComplete = false
			break
		}
		row--
		col--
	}

	return diagComplete
}

func (b *bingo) sumUnmarked() (sum int) {
	for row := range b.markd {
		for col := range b.markd[row] {
			if !b.markd[row][col] {
				sum += b.board[row][col]
			}
		}
	}
	return
}
