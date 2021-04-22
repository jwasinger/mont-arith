package mont_arith

import (
	"math/bits"
	"unsafe"
	"errors"
)





var Zero2Limbs []uint64 = make([]uint64, 2, 2)


/* NOTE: addmod/submod/mulmodmont assume:
	len(z) == len(x) == len(y) == len(mod)
*/

func MulModMont128(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [2]uint64
	var c [3]uint64
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
					t[1], t[0]  = madd3(m, mod[1], c[0], c[2], c[1])
		// round 1
			v = x[1]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					z[1], z[0] = madd3(m, mod[1], c[0], c[2], c[1])

	_, c[1] = bits.Sub64(z[0], mod[0], 0)
	    _, c[1] = bits.Sub64(z[1], mod[1], 0)

	if c[1] != 0 { // unnecessary sub
		sub_val = Zero2Limbs
	}

	_, c[1] = bits.Sub64(z[0], sub_val[0], 0)
	    _, c[1] = bits.Sub64(z[1], sub_val[1], 0)

	return nil
}

