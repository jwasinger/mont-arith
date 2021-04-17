package arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

func SubMod448(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := [7]uint64{0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Sub64(x[0], y[0], 0)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}
