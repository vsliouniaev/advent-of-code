package _go

import (
	"fmt"
	"strconv"
)

type Day1 struct{}

func (d Day1) fuelRequired(f int) int {
	return f/3 - 2
}

// Fuel required for fuel + the fuel required for that
func (d Day1) fuelRequiredRequired(f int) (sum int) {
	f = d.fuelRequired(f)
	for f > 0 {
		sum += f
		f = d.fuelRequired(f)
	}
	return
}

func (d Day1) part1(componentMasses []int) (sum int) {
	for _, m := range componentMasses {
		sum += d.fuelRequired(m)
	}
	return
}

func (d Day1) part2(componentMasses []int) (sum int) {
	for _, m := range componentMasses {
		sum += d.fuelRequiredRequired(m)
	}
	return
}

func (d Day1) parse(file string) (masses []int) {
	lines := ParseFile(file)
	for _, l := range lines {
		m, _ := strconv.Atoi(l)
		masses = append(masses, m)
	}

	return
}

func (d Day1) Go(){
	masses := d.parse(`in/day1`)
	fmt.Println(d.part1(masses))
	fmt.Println(d.part2(masses))
}
