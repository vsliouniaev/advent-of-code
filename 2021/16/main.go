package main

import (
	"encoding/hex"
	"fmt"
	"github.com/vsliouniaev/aoc/util"
	"github.com/vsliouniaev/aoc/util/bits"
	"math"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1(util.RelativeFile("input"))) // 940
	fmt.Printf("Part 2: %d\n", part2(util.RelativeFile("input"))) // 13476220616073
}

func part1(file string) int64 {
	data, err := hex.DecodeString(util.ReadLinesStrings(file)[0])
	util.Check(err)
	p := Parser{data: data}
	p.ReadNumSubPackets(1)
	return p.sumVersionNums
}

func part2(file string) int64 {
	return evalExpression(util.ReadLinesStrings(file)[0])
}

func evalExpression(text string) int64 {
	data, err := hex.DecodeString(text)
	util.Check(err)
	p := Parser{data: data}
	var out []int64
	out, _ = p.ReadNumSubPackets(1)
	return out[0]
}

// Parser holds data for parsing the Buoyancy Interchange Transmission System (BITS) expression.
type Parser struct {
	data           []byte
	pos            int
	sumVersionNums int64
}

func (p *Parser) Chomp(ln int) (int64, bool) {
	i := bits.GetAsInt(p.data, p.pos, ln)
	p.pos += ln
	return i, (len(p.data)*8)-1 > p.pos
}

// ReadVersion reads the 3 bits at the start of the expression. These are only useful for part 1 of day 16.
func (p *Parser) ReadVersion() (int64, bool) {
	return p.Chomp(3)
}

// ReadMode reads the 3 bits of the mode at the head of an expression (after the version).
func (p *Parser) ReadMode() (int64, bool) {
	return p.Chomp(3)
}

// ReadLiteral reads a single literal value.
func (p *Parser) ReadLiteral() (int64, bool) {
	var ok bool
	var read int64
	val := int64(0)
	for {
		i, _ := p.Chomp(1)
		val = val << 4
		read, ok = p.Chomp(4)
		val |= read

		if i == 0 || !ok {
			break
		}
	}
	return val, ok
}

// ReadNested begins to parse a sub-packet chunk.
// If the length type id is 0 - read packets until the total number of bits has been reached
// If the length type id is 1 - read the specified number of packets
func (p *Parser) ReadNested() ([]int64, bool) {
	lengthTypeId, _ := p.Chomp(1)
	if lengthTypeId == 0 {
		// Total number of bits
		totalNumberOfBits, _ := p.Chomp(15)
		return p.ReadSubPacketsOfLength(int(totalNumberOfBits))
	} else {
		numberOfPackets, _ := p.Chomp(11)
		return p.ReadNumSubPackets(numberOfPackets)
	}
}

// ReadNumSubPackets reads the specified number of sub-packets
func (p *Parser) ReadNumSubPackets(num int64) ([]int64, bool) {
	var out []int64
	for i := int64(0); i < num; i++ {
		o, _ := p.ReadSubPacket()
		out = append(out, o)
	}
	return out, true
}

// ReadSubPacketsOfLength reads sub-packets limited to the number of bits
func (p *Parser) ReadSubPacketsOfLength(lim int) ([]int64, bool) {
	var out []int64
	lim = p.pos + lim
	for p.pos < lim {
		o, _ := p.ReadSubPacket()
		out = append(out, o)
	}
	return out, true
}

// ReadSubPacket reads a single sub-packet. However, since a sub-packet can have more sub-packets it is recursive via
// ReadNested calls back to itself.
func (p *Parser) ReadSubPacket() (int64, bool) {
	var out int64
	var ok bool
	v, _ := p.ReadVersion()
	p.sumVersionNums += v
	mode, _ := p.ReadMode()
	var operands []int64
	if mode == 4 {
		return p.ReadLiteral()
	} else {
		operands, ok = p.ReadNested()
	}
	switch mode {
	case 4: // Literal
		// Do nothing
	case 0: // +
		for _, o := range operands {
			out += o
		}
	case 1: // *
		out = 1
		for _, o := range operands {
			out *= o
		}
	case 2: // min
		out = math.MaxInt64
		for _, o := range operands {
			if out > o {
				out = o
			}
		}
	case 3: // max
		out = math.MinInt64
		for _, o := range operands {
			if out < o {
				out = o
			}
		}
	case 5: // >
		if len(operands) != 2 {
			panic("ops > 2")
		}
		if operands[0] > operands[1] {
			return 1, true
		} else {
			return 0, true
		}
	case 6: // <
		if len(operands) != 2 {
			panic("ops > 2")
		}
		if operands[0] < operands[1] {
			return 1, true
		} else {
			return 0, true
		}
	case 7: // =
		if len(operands) != 2 {
			panic("ops > 2")
		}
		if operands[0] == operands[1] {
			return 1, true
		} else {
			return 0, true
		}
	default:
		panic(mode)
	}

	return out, ok
}
