package mont_arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

func SubMod64(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64
	var tmp uint64 = 0

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	z[0], c = bits.Sub64(x[0], y[0], 0)
	tmp, _ = bits.Add64(z[0], mod[0], 0)

	if c != 0 {
		z[0] = tmp
	}

	return nil
}

func AddMod64(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64
	var tmp uint64 = 0

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	tmp, c = bits.Add64(x[0], y[0], 0)
	z[0], _ = bits.Sub64(tmp, mod[0], 0)

	if c == 0 {
		z[0] = tmp
	}
	return nil
}

func MulModMont64(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[1]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[1]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	out := (*[1]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[1]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var product [2]uint64
	var c uint64

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	product[1], product[0] = bits.Mul64(x[0], y[0])
	m := product[0] * ctx.MontParamInterleaved
	c, product[0] = madd1(m, mod[0], product[0])
	out[0] = c + product[1]

	if out[0] > mod[0] {
		out[0] = c - mod[0]
	}

	return nil
}
