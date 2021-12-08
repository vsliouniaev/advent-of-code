package main

import (
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"math"
	"math/bits"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("2021/8/input")) // 421
	fmt.Printf("Part 2: %d\n", part2("2021/8/input")) // 986163
}

func part1(file string) int {
	inputs := util.ReadLinesStrings(file)
	cnt := 0
	for _, inp := range inputs {
		spl := strings.Split(inp, " | ")
		dig := strings.Split(spl[1], " ")
		for _, d := range dig {
			switch len(d) {
			case 2: // 1
				cnt++
			case 3: // 7
				cnt++
			case 4: // 4
				cnt++
			case 7: // 8
				cnt++
			}
		}
	}
	return cnt
}

func part2(file string) int {
	s := 0
	inputs := util.ReadLinesStrings(file)
	for _, inp := range inputs {
		spl := strings.Split(inp, " | ")
		code := strings.Split(spl[0], " ")
		dig := strings.Split(spl[1], " ")

		one := getLen(2, code)[0]
		seven := getLen(3, code)[0]
		four := getLen(4, code)[0]
		eight := getLen(7, code)[0]

		c_f := encode(one)
		a_c_f := encode(seven)
		b_c_d_f := encode(four)
		abcdefg := encode(eight)
		a := a_c_f &^ c_f
		a_e_g := abcdefg &^ b_c_d_f
		e_g := a_e_g &^ a
		b_d := b_c_d_f &^ c_f

		var c, d_g, g byte
		search := getLen(5, code)
		for _, s := range search {
			a_b_d_f_g := encode(s)
			if a_b_d_f_g|b_d == a_b_d_f_g {
				// FIVE
				c = c_f &^ a_b_d_f_g
				a_g := a_b_d_f_g &^ b_c_d_f
				g = a_g &^ a
			}
			a_c_d_f_g := encode(s)
			if bits.OnesCount8(a_c_d_f_g&^a_c_f) == 2 {
				// THREE
				d_g = a_c_d_f_g &^ a_c_f
			}

		}
		d := d_g &^ g
		e := e_g &^ g
		b := b_d &^ d
		f := c_f &^ c
		lookup := map[byte]int{
			c | f:                     1,
			a | c | d | e | g:         2,
			a | c | d | f | g:         3,
			b | c | d | f:             4,
			a | b | d | f | g:         5,
			a | b | d | e | f | g:     6,
			a | c | f:                 7,
			a | b | c | d | e | f | g: 8,
			a | b | c | d | f | g:     9,
			a | b | c | e | f | g:     0,
		}

		for i, d := range dig {
			s += lookup[encode(d)] * int(math.Pow(10, float64(len(dig)-i-1)))
		}
	}

	return s
}

func getLen(l int, code []string) []string {
	var out []string
	for _, c := range code {
		if len(c) == l {
			out = append(out, c)
		}
	}
	return out
}

func encode(str string) byte {
	// g f e d c b a
	var b byte
	for _, r := range str {
		b |= 1 << (r - 'a')
	}
	return b
}
