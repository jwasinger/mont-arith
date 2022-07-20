package mont_arith

import (
	"math/bits"
)

// XXX implement the below methods using these types (conversions might make it awkward/slower)
type Word uint
type nat = []Word
const WordSize = 8 // XXX word size in bytes, hardcoded to 64bit limbs


// madd0 hi = a*b + c (discards lo bits)
func madd0(a, b, c Word) (Word) {
	var carry, lo uint
	hi, lo := bits.Mul(uint(a), uint(b))
	_, carry = bits.Add(lo, uint(c), 0)
	hi, _ = bits.Add(hi, 0, carry)
	return Word(hi)
}

// madd1 hi, lo = a*b + c
func madd1(a, b, c Word) (Word, Word) {
	var carry uint
	hi, lo := bits.Mul(uint(a), uint(b))
	lo, carry = bits.Add(uint(lo), uint(c), 0)
	hi, _ = bits.Add(hi, 0, carry)
	return Word(hi), Word(lo)
}

// madd2 hi, lo = a*b + c + d
func madd2(a, b, c, d Word) (Word, Word) {
	var carry uint
    var c_uint uint
	hi, lo := bits.Mul(uint(a), uint(b))
	c_uint, carry = bits.Add(uint(c), uint(d), 0)
    c = Word(c_uint)
	hi, _ = bits.Add(hi, 0, carry)
	lo, carry = bits.Add(lo, uint(c), 0)
	hi, _ = bits.Add(hi, 0, carry)
	return Word(hi), Word(lo)
}

func madd3(a, b, c, d, e Word) (Word, Word) {
	var carry uint
    var c_uint uint
	hi, lo := bits.Mul(uint(a), uint(b))
	c_uint, carry = bits.Add(uint(c), uint(d), 0)
	hi, _ = bits.Add(hi, 0, carry)
	lo, carry = bits.Add(lo, c_uint, 0)
	hi, _ = bits.Add(hi, uint(e), carry)
	return Word(hi), Word(lo)
}
