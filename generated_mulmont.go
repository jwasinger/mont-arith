package mont_arith

import (
	"errors"
	"math/big"
	"math/bits"
	"unsafe"
)

func MulModMont64(out, x, y []uint, ctx *MontArithContext) error {
	mod := ctx.Modulus

	var product [2]uint
	var c uint

	product[1], product[0] = bits.Mul(x[0], y[0])
	m := product[0] * ctx.MontParamInterleaved
	c, product[0] = madd1(m, mod[0], product[0])
	out[0] = c + product[1]

	if out[0] > mod[0] {
		out[0] = c - mod[0]
	}

	return nil
}

var Zero1Limbs []uint = make([]uint, 1, 1)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont64(z, x, y []uint, ctx *MontArithContext) error {
	var t [1]uint
	var c [3]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
	m := c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 1; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero2Limbs []uint = make([]uint, 2, 2)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont128(z, x, y []uint, ctx *MontArithContext) error {
	var t [2]uint
	var c [3]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
	m := c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd1(v, y[1], c[1])
	t[1], t[0] = madd3(m, mod[1], c[0], c[2], c[1])
	// round 1
	v = x[1]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	z[1], z[0] = madd3(m, mod[1], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 2; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero3Limbs []uint = make([]uint, 3, 3)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont192(z, x, y []uint, ctx *MontArithContext) error {
	var t [3]uint
	var c [3]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
	m := c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd1(v, y[1], c[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd1(v, y[2], c[1])
	t[2], t[1] = madd3(m, mod[2], c[0], c[2], c[1])
	// round 1
	v = x[1]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	t[2], t[1] = madd3(m, mod[2], c[0], c[2], c[1])
	// round 2
	v = x[2]
	c[1], c[0] = madd1(v, y[0], t[0])
	m = c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], z[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	z[2], z[1] = madd3(m, mod[2], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 3; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero4Limbs []uint = make([]uint, 4, 4)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont256(z, x, y []uint, ctx *MontArithContext) error {
	var t [4]uint
	var c [4]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
	m := c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd1(v, y[1], c[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd1(v, y[2], c[1])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd1(v, y[3], c[1])
	t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
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
	c[1], c[0] = madd2(v, y[1], c[1], t[1])
	c[2], z[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd2(v, y[2], c[1], t[2])
	c[2], z[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd2(v, y[3], c[1], t[3])
	z[3], z[2] = madd3(m, mod[3], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 4; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero5Limbs []uint = make([]uint, 5, 5)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont320(z, x, y []uint, ctx *MontArithContext) error {
	var t [5]uint
	var c [5]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
	m := c[0] * modinv
	c[2] = madd0(m, mod[0], c[0])
	c[1], c[0] = madd1(v, y[1], c[1])
	c[2], t[0] = madd2(m, mod[1], c[2], c[0])
	c[1], c[0] = madd1(v, y[2], c[1])
	c[2], t[1] = madd2(m, mod[2], c[2], c[0])
	c[1], c[0] = madd1(v, y[3], c[1])
	c[2], t[2] = madd2(m, mod[3], c[2], c[0])
	c[1], c[0] = madd1(v, y[4], c[1])
	t[4], t[3] = madd3(m, mod[4], c[0], c[2], c[1])
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
	t[4], t[3] = madd3(m, mod[4], c[0], c[2], c[1])
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
	t[4], t[3] = madd3(m, mod[4], c[0], c[2], c[1])
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
	t[4], t[3] = madd3(m, mod[4], c[0], c[2], c[1])
	// round 4
	v = x[4]
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
	z[4], z[3] = madd3(m, mod[4], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 5; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero6Limbs []uint = make([]uint, 6, 6)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont384(z, x, y []uint, ctx *MontArithContext) error {
	var t [6]uint
	var c [6]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
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
	t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
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
	t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
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
	t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
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
	t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
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
	t[5], t[4] = madd3(m, mod[5], c[0], c[2], c[1])
	// round 5
	v = x[5]
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
	z[5], z[4] = madd3(m, mod[5], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 6; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero7Limbs []uint = make([]uint, 7, 7)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont448(z, x, y []uint, ctx *MontArithContext) error {
	var t [7]uint
	var c [7]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
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
	t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
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
	t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
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
	t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
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
	t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
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
	t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
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
	t[6], t[5] = madd3(m, mod[6], c[0], c[2], c[1])
	// round 6
	v = x[6]
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
	z[6], z[5] = madd3(m, mod[6], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 7; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero8Limbs []uint = make([]uint, 8, 8)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont512(z, x, y []uint, ctx *MontArithContext) error {
	var t [8]uint
	var c [8]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
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
	t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
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
	t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
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
	t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
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
	t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
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
	t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
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
	t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
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
	t[7], t[6] = madd3(m, mod[7], c[0], c[2], c[1])
	// round 7
	v = x[7]
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
	z[7], z[6] = madd3(m, mod[7], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 8; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero9Limbs []uint = make([]uint, 9, 9)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont576(z, x, y []uint, ctx *MontArithContext) error {
	var t [9]uint
	var c [9]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
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
	t[8], t[7] = madd3(m, mod[8], c[0], c[2], c[1])
	// round 8
	v = x[8]
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
	z[8], z[7] = madd3(m, mod[8], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 9; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero10Limbs []uint = make([]uint, 10, 10)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont640(z, x, y []uint, ctx *MontArithContext) error {
	var t [10]uint
	var c [10]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
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

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 10; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

var Zero11Limbs []uint = make([]uint, 11, 11)

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

// NOTE: assumes x < mod and y < mod
func MulModMont704(z, x, y []uint, ctx *MontArithContext) error {
	var t [11]uint
	var c [11]uint
	modinv := ctx.MontParamInterleaved
	mod := ctx.Modulus
	// round 0
	v := x[0]
	c[1], c[0] = bits.Mul(v, y[0])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd1(v, y[10], c[1])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
	// round 9
	v = x[9]
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	t[10], t[9] = madd3(m, mod[10], c[0], c[2], c[1])
	// round 10
	v = x[10]
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
	c[2], z[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd2(v, y[10], c[1], t[10])
	z[10], z[9] = madd3(m, mod[10], c[0], c[2], c[1])

	// final subtraction, overwriting z if z > mod
	c[0] = 0
	for i := 0; i < 11; i++ {
		t[i], c[0] = bits.Sub(z[i], mod[i], c[0])
	}

	if c[0] == 0 {
		copy(z, t[:])
	}

	return nil
}

// NOTE: this assumes that x and y are in Montgomery form and can produce unexpected results when they are not
func MulModMontNonInterleaved(out_bytes, x_bytes, y_bytes []byte, m *MontArithContext) error {
	// length x == y assumed

	product := new(big.Int)
	x := LEBytesToInt(x_bytes)
	y := LEBytesToInt(y_bytes)

	if x.Cmp(m.ModulusNonInterleaved) > 0 || y.Cmp(m.ModulusNonInterleaved) > 0 {
		return errors.New("x/y >= modulus")
	}

	// m <- ((x*y mod R)N`) mod R
	product.Mul(x, y)
	x.And(product, m.mask)
	x.Mul(x, m.MontParamNonInterleaved)
	x.And(x, m.mask)

	// t <- (T + mN) / R
	x.Mul(x, m.ModulusNonInterleaved)
	x.Add(x, product)
	x.Rsh(x, m.NumLimbs*64)

	if x.Cmp(m.ModulusNonInterleaved) >= 0 {
		x.Sub(x, m.ModulusNonInterleaved)
	}

	copy(out_bytes, LimbsToLEBytes(IntToLimbs(x, m.NumLimbs)))

	return nil
}
