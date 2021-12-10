package main

import (
	"fmt"
	u "github.com/vsliouniaev/aoc/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 17934269678453
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input"))) // 3440662844064
}

func part1(file string) uint64 {
	var m1, m0 uint64
	mem := make(map[uint64]uint64)
	for _, line := range u.ReadLinesStrings(file) {
		if strings.HasPrefix(line, "mask = ") {
			m1 = uint64(0)
			m0 = uint64(0)
			line = strings.TrimPrefix(line, "mask = ")
			for i := 0; i < len(line); i++ {
				switch line[len(line)-i-1] {
				case '0':
					m0 |= 1 << i
				case '1':
					m1 |= 1 << i
				case 'X':
				default:
					panic(line[i])
				}
			}
		} else {
			loc, bits := parseMem(line)
			bits |= m1  // |  is or
			bits &^= m0 // &^ is the "clear bit operator"
			mem[loc] = bits
		}
	}
	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}

	return sum
}

func part2(file string) uint64 {
	computerMemory := make(map[uint64]uint64)
	// onesMask defines bits that are always 1 in the resulting output
	var onesMask uint64

	// Permutations of "floating" bits
	var floatingPerm []uint64
	// floatingMask defines bits with multiple values in the problem, a.k.a. "floating" bits. These can be 0 or 1 and
	// must be combined with values of floatingPerm to determine which bits to actually set and clear
	var floatingMask uint64

	for _, line := range u.ReadLinesStrings(file) {
		if strings.HasPrefix(line, "mask = ") {
			onesMask = uint64(0)
			floatingMask = uint64(0)
			line = strings.TrimPrefix(line, "mask = ")
			floatingPerm = gen(strings.ReplaceAll(line, "1", "0"), 0)
			for i := 0; i < len(line); i++ {
				switch line[len(line)-i-1] {
				case '0':
					continue
				case '1':
					onesMask |= 1 << i
				case 'X':
					floatingMask |= 1 << i
				default:
					panic(line[i])
				}
			}
		} else {
			loc, bits := parseMem(line)
			for _, floatBits := range floatingPerm {
				setBits := floatBits & floatingMask
				clearBits := ^floatBits & floatingMask
				setBits |= onesMask
				l := loc &^ clearBits
				l |= setBits
				computerMemory[l] = bits
			}
		}
	}
	sum := uint64(0)
	for _, v := range computerMemory {
		sum += v
	}

	return sum
}

func gen(l string, start int) []uint64 {
	var out []uint64
	for i := start; i < len(l); i++ {
		if l[i] == 'X' {
			one := []rune(l)
			one[i] = '1'
			zer := []rune(l)
			zer[i] = '0'
			out = append(out, gen(string(one), i+1)...)
			out = append(out, gen(string(zer), i+1)...)
			return out
		}
	}

	i, err := strconv.ParseUint(l, 2, 64)
	u.Check(err)
	return []uint64{i}
}

func parseMem(line string) (uint64, uint64) {
	r := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	s := r.FindAllStringSubmatch(line, -1)
	loc, err := strconv.Atoi(s[0][1])
	u.Check(err)
	bits, err := strconv.Atoi(s[0][2])
	u.Check(err)
	return uint64(loc), uint64(bits)
}
