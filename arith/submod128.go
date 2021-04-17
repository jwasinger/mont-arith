package arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

func SubMod128(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := [2]uint64{0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Sub64(x[0], y[0], 0)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}
