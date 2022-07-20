package mont_arith

import (
	"errors"
	"math/bits"
	"unsafe"
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

func SubMod64(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[1]uint)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint
	var tmp uint = 0

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	z[0], c = bits.Sub(x[0], y[0], 0)
	tmp, _ = bits.Add(z[0], mod[0], 0)

	if c != 0 {
		z[0] = tmp
	}

	return nil
}

func AddMod64(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[1]uint)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint
	var tmp uint = 0

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	tmp, c = bits.Add(x[0], y[0], 0)
	z[0], _ = bits.Sub(tmp, mod[0], 0)

	if c == 0 {
		z[0] = tmp
	}
	return nil
}

func MulModMont64(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[1]uint)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint)(unsafe.Pointer(&y_bytes[0]))[:]
	out := (*[1]uint)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var product [2]uint
	var c uint

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	product[1], product[0] = bits.Mul(x[0], y[0])
	m := product[0] * ctx.MontParamInterleaved
	c, product[0] = madd1(m, mod[0], product[0])
	out[0] = c + product[1]

	if out[0] > mod[0] {
		out[0] = c - mod[0]
	}

	return nil
}
