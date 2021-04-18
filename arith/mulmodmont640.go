package arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

var Zero10Limbs []uint64 = make([]uint64, 10, 10)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

func MulModMont640(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	var t [10]uint64
	var c [10]uint64
	var sub_val []uint64 = mod
	modinv := ctx.MontParamInterleaved

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul64(v, y[0])
	m := c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd1(v, y[1], c[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd1(v, y[2], c[1])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd1(v, y[3], c[1])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd1(v, y[4], c[1])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd1(v, y[5], c[1])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd1(v, y[6], c[1])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd1(v, y[7], c[1])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd1(v, y[8], c[1])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd1(v, y[9], c[1])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 1
	v = x[1]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 2
	v = x[2]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 3
	v = x[3]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 4
	v = x[4]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 5
	v = x[5]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 6
	v = x[6]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 7
	v = x[7]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 8
	v = x[8]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], t[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], t[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], t[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], t[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], t[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	t[9], t[8] = madd3(m, mod[9], c[0], c[2], c[1])
	// round 9
	v = x[9]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], z[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], z[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	c[2], z[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd2(v, y[4], c[1], t[4])
	c[2], z[3] = madd2(m, mod[4], c[2], c[0])
	c[1], c[0] = madd2(v, y[5], c[1], t[5])
	c[2], z[4] = madd2(m, mod[5], c[2], c[0])
	c[1], c[0] = madd2(v, y[6], c[1], t[6])
	c[2], z[5] = madd2(m, mod[6], c[2], c[0])
	c[1], c[0] = madd2(v, y[7], c[1], t[7])
	c[2], z[6] = madd2(m, mod[7], c[2], c[0])
	c[1], c[0] = madd2(v, y[8], c[1], t[8])
	c[2], z[7] = madd2(m, mod[8], c[2], c[0])
	c[1], c[0] = madd2(v, y[9], c[1], t[9])
	z[9], z[8] = madd3(m, mod[9], c[0], c[2], c[1])

	_, c[1] = bits.Sub64(z[0], mod[0], 0)
	_, c[1] = bits.Sub64(z[1], mod[1], 0)
	_, c[1] = bits.Sub64(z[2], mod[2], 0)
	_, c[1] = bits.Sub64(z[3], mod[3], 0)
	_, c[1] = bits.Sub64(z[4], mod[4], 0)
	_, c[1] = bits.Sub64(z[5], mod[5], 0)
	_, c[1] = bits.Sub64(z[6], mod[6], 0)
	_, c[1] = bits.Sub64(z[7], mod[7], 0)
	_, c[1] = bits.Sub64(z[8], mod[8], 0)
	_, c[1] = bits.Sub64(z[9], mod[9], 0)

	if c[1] != 0 { // unnecessary sub
		sub_val = Zero10Limbs
	}

	_, c[1] = bits.Sub64(z[0], sub_val[0], 0)
	_, c[1] = bits.Sub64(z[1], sub_val[1], 0)
	_, c[1] = bits.Sub64(z[2], sub_val[2], 0)
	_, c[1] = bits.Sub64(z[3], sub_val[3], 0)
	_, c[1] = bits.Sub64(z[4], sub_val[4], 0)
	_, c[1] = bits.Sub64(z[5], sub_val[5], 0)
	_, c[1] = bits.Sub64(z[6], sub_val[6], 0)
	_, c[1] = bits.Sub64(z[7], sub_val[7], 0)
	_, c[1] = bits.Sub64(z[8], sub_val[8], 0)
	_, c[1] = bits.Sub64(z[9], sub_val[9], 0)

	return nil
}
