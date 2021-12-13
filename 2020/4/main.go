package main

import (
	"fmt"
	//. "github.com/vsliouniaev/aoc/util"
	u "github.com/vsliouniaev/aoc/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(u.RelativeFile("input"))) // 226
	fmt.Printf("Part 2: %d\n", part2(u.RelativeFile("input"))) // 160
}

func part1(file string) int {
	valids := 0
	val := validator()
	for _, line := range u.ReadLinesStrings(file) {
		if line == "" {
			if len(val) == 0 {
				valids++
			}
			val = validator()
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			kind := strings.Split(field, ":")[0]
			delete(val, kind)
		}
	}
	if len(val) == 0 {
		valids++
	}

	return valids
}

func part2(file string) int {
	valids := 0
	passVal := validator()
	for _, line := range u.ReadLinesStrings(file) {
		if line == "" {
			if len(passVal) == 0 {
				valids++
			}
			passVal = validator()
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			if field == "" {
				continue
			}
			kv := strings.Split(field, ":")
			var key, value = kv[0], kv[1]
			if fun, ok := passVal[key]; ok && fun(value) {
				delete(passVal, key)
			}
		}
	}
	if len(passVal) == 0 {
		valids++
	}

	return valids
}

func validator() map[string]func(string) bool {
	return map[string]func(string) bool{
		"byr": func(s string) bool {
			// byr (Birth Year) - four digits; at least 1920 and at most 2002.
			i, err := strconv.ParseInt(s, 10, 64)
			return len(s) == 4 && err == nil && i >= 1920 && i <= 2002
		},
		"iyr": func(s string) bool {
			// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
			i, err := strconv.ParseInt(s, 10, 64)
			return len(s) == 4 && err == nil && i >= 2010 && i <= 2020
		},
		"eyr": func(s string) bool {
			// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
			i, err := strconv.ParseInt(s, 10, 64)
			return len(s) == 4 && err == nil && i >= 2020 && i <= 2030
		},
		"hgt": func(s string) bool {
			//hgt (Height) - a number followed by either cm or in:
			// If cm, the number must be at least 150 and at most 193.
			// If in, the number must be at least 59 and at most 76.
			h := strings.TrimSuffix(strings.TrimSuffix(s, "cm"), "in")
			i, err := strconv.ParseInt(h, 10, 64)
			if err != nil || len(h) == len(s) {
				return false
			}
			if strings.HasSuffix(s, "cm") {
				return i >= 150 && i <= 193
			}
			if strings.HasSuffix(s, "in") {
				return i >= 59 && i <= 76
			}
			panic("Failed: " + s)
		},
		"hcl": func(s string) bool {
			// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
			return hcl.MatchString(s)
		},
		"ecl": func(s string) bool {
			// ecl (Eye Color) - exactly one of:
			_, ok := ecl[s]
			return ok
		},
		"pid": func(s string) bool {
			// pid (Passport ID) - a nine-digit number, including leading zeroes.
			_, err := strconv.ParseInt(s, 10, 64)
			return len(s) == 9 && err == nil
		},
		//"cid": false,
	}
}

var (
	hcl = regexp.MustCompile("#[0-9a-f]{6}")
	ecl = map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}
)
