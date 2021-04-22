package mont_arith

import (
	"math/bits"
	"unsafe"
	"errors"
)





var Zero4Limbs []uint64 = make([]uint64, 4, 4)


/* NOTE: addmod/submod/mulmodmont assume:
	len(z) == len(x) == len(y) == len(mod)
*/

func MulModMont256(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
    	var t [4]uint64
	var c [4]uint64
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
					t[3], t[2]  = madd3(m, mod[3], c[0], c[2], c[1])
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
					t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
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
					t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
		// round 3
			v = x[3]
			c[1], c[0] = madd1(v, y[0], t[0])
			m = c[0] * modinv
			c[2] = madd0(m, mod[0], c[0])
				c[1], c[0] = madd2(v, y[1],  c[1], t[1])
					c[2], z[0] = madd2(m, mod[1],  c[2], c[0])
				c[1], c[0] = madd2(v, y[2],  c[1], t[2])
					c[2], z[1] = madd2(m, mod[2],  c[2], c[0])
				c[1], c[0] = madd2(v, y[3],  c[1], t[3])
					z[3], z[2] = madd3(m, mod[3], c[0], c[2], c[1])

	_, c[1] = bits.Sub64(z[0], mod[0], 0)
	    _, c[1] = bits.Sub64(z[1], mod[1], 0)
	    _, c[1] = bits.Sub64(z[2], mod[2], 0)
	    _, c[1] = bits.Sub64(z[3], mod[3], 0)

	if c[1] != 0 { // unnecessary sub
		sub_val = Zero4Limbs
	}

	_, c[1] = bits.Sub64(z[0], sub_val[0], 0)
	    _, c[1] = bits.Sub64(z[1], sub_val[1], 0)
	    _, c[1] = bits.Sub64(z[2], sub_val[2], 0)
	    _, c[1] = bits.Sub64(z[3], sub_val[3], 0)

	return nil
}

