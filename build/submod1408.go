package modext

import (
	"errors"
	"math/bits"
	"unsafe"
)

func SubMod1408(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[22]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[22]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[22]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[22]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := [22]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[21] >= mod[21] || y[21] >= mod[21] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Sub64(x[0], y[0], 0)
	tmp[1], c = bits.Sub64(x[1], y[1], c)
	tmp[2], c = bits.Sub64(x[2], y[2], c)
	tmp[3], c = bits.Sub64(x[3], y[3], c)
	tmp[4], c = bits.Sub64(x[4], y[4], c)
	tmp[5], c = bits.Sub64(x[5], y[5], c)
	tmp[6], c = bits.Sub64(x[6], y[6], c)
	tmp[7], c = bits.Sub64(x[7], y[7], c)
	tmp[8], c = bits.Sub64(x[8], y[8], c)
	tmp[9], c = bits.Sub64(x[9], y[9], c)
	tmp[10], c = bits.Sub64(x[10], y[10], c)
	tmp[11], c = bits.Sub64(x[11], y[11], c)
	tmp[12], c = bits.Sub64(x[12], y[12], c)
	tmp[13], c = bits.Sub64(x[13], y[13], c)
	tmp[14], c = bits.Sub64(x[14], y[14], c)
	tmp[15], c = bits.Sub64(x[15], y[15], c)
	tmp[16], c = bits.Sub64(x[16], y[16], c)
	tmp[17], c = bits.Sub64(x[17], y[17], c)
	tmp[18], c = bits.Sub64(x[18], y[18], c)
	tmp[19], c = bits.Sub64(x[19], y[19], c)
	tmp[20], c = bits.Sub64(x[20], y[20], c)
	tmp[21], c = bits.Sub64(x[21], y[21], c)
	z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Add64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Add64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Add64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Add64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Add64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Add64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Add64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Add64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Add64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Add64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Add64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Add64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Add64(tmp[19], mod[19], c1)
	z[20], c1 = bits.Add64(tmp[20], mod[20], c1)
	z[21], c1 = bits.Add64(tmp[21], mod[21], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}