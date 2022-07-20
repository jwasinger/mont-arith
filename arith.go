package mont_arith

import (
	"math/bits"
)

// madd0 hi = a*b + c (discards lo bits)
func madd0(a, b, c uint) (hi uint) {
	var carry, lo uint
	hi, lo = bits.Mul(a, b)
	_, carry = bits.Add(lo, c, 0)
	hi, _ = bits.Add(hi, 0, carry)
	return
}

// madd1 hi, lo = a*b + c
func madd1(a, b, c uint) (hi uint, lo uint) {
	var carry uint
	hi, lo = bits.Mul(a, b)
	lo, carry = bits.Add(lo, c, 0)
	hi, _ = bits.Add(hi, 0, carry)
	return
}

// madd2 hi, lo = a*b + c + d
func madd2(a, b, c, d uint) (hi uint, lo uint) {
	var carry uint
	hi, lo = bits.Mul(a, b)
	c, carry = bits.Add(c, d, 0)
	hi, _ = bits.Add(hi, 0, carry)
	lo, carry = bits.Add(lo, c, 0)
	hi, _ = bits.Add(hi, 0, carry)
	return
}

func madd3(a, b, c, d, e uint) (hi uint, lo uint) {
	var carry uint
	hi, lo = bits.Mul(a, b)
	c, carry = bits.Add(c, d, 0)
	hi, _ = bits.Add(hi, 0, carry)
	lo, carry = bits.Add(lo, c, 0)
	hi, _ = bits.Add(hi, e, carry)
	return
}
