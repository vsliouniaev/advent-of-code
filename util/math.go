package util

import "math/big"

// Chinese Remainder Theorem:
// Stolen from https://github.com/stripedpajamas/cryptopals-go/blob/master/set8/challenge57/57.go

type ChineseRemainderArg struct {
	Remainder *big.Int
	Modulus   *big.Int
}

type ChineseRemainder []*ChineseRemainderArg

func (residues ChineseRemainder) Solve() *big.Int {
	if len(residues) < 2 {
		panic("not enough residues to compute solution")
	}
	acc := residues[0]
	for i := 1; i < len(residues); i++ {
		a1, a2 := acc.Remainder, residues[i].Remainder
		n1, n2 := acc.Modulus, residues[i].Modulus
		m1, m2 := new(big.Int), new(big.Int)

		solution := new(big.Int).GCD(m1, m2, n1, n2)

		left := new(big.Int).Mul(a1, m2)
		left.Mul(left, n2)

		right := new(big.Int).Mul(a2, m1)
		right.Mul(right, n1)

		solution.Add(left, right)
		combined := new(big.Int).Mul(n1, n2)

		solution.Mod(solution, combined)
		acc = &ChineseRemainderArg{Remainder: solution, Modulus: combined}
	}
	return acc.Remainder
}
