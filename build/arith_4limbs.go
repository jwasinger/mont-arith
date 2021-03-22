package evmmax_arith

import (
	"errors"
	"fmt"
	"math/bits"
	"unsafe"
)

var Zero4Limbs [4]uint64 = [4]uint64{0, 0, 0, 0}

func MulModMont(out_bytes, x_bytes, y_bytes *[]byte, modCtx *ModContext) error {
	x := (*[4]uint64)(unsafe.Pointer(&(*x_bytes)[0]))
	y := (*[4]uint64)(unsafe.Pointer(&(*y_bytes)[0]))
	out := (*[4]uint64)(unsafe.Pointer(&(*out_bytes)[0]))
	mod := (*[4]uint64)(unsafe.Pointer(&modCtx.Modulus[0]))

	return _MulModMont(out, x, y, mod, modCtx.NInv)
}

func _MulModMont(z, x, y *[4]uint64, mod *[4]uint64, modinv uint64) error {
	var t [4]uint64
	var c [4]uint64
	var sub_val [4]uint64 = *mod

	if x[3] >= mod[3] || y[3] >= mod[3] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	fmt.Println("hello there")
	fmt.Println(x[0])
	fmt.Println(y[0])
	fmt.Println(modinv)
	fmt.Println(mod)
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

/*
	Modular Addition
*/
/*
func (out *Element) AddMod(x, y, mod *Element) (error) {
	var c uint64
	var tmp Element

	if x[3] > mod[3] || y[3] > mod[3] {
		return errors.New("x/y must be smaller than modulus")
	}

	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)

	out[0], c = bits.Sub64(tmp[0], mod[0], 0)
	out[1], c = bits.Sub64(tmp[1], mod[1], c)
	out[2], c = bits.Sub64(tmp[2], mod[2], c)
	out[3], c = bits.Sub64(tmp[3], mod[3], c)

	if c != 0 { // unnecessary sub
		*out = tmp
	}

	return nil
}
*/

/*
	Modular Subtraction
*/
/*
func (out *Element) SubMod(x, y, mod *Element) (error) {
	var c, c1 uint64
	var tmp Element

	if x[3] > mod[3] || y[3] > mod[3] {
		return errors.New("x/y must be smaller than modulus")
	}

	tmp[0], c1 = bits.Sub64(x[0], y[0], 0)
	tmp[1], c1 = bits.Sub64(x[1], y[1], c1)
	tmp[2], c1 = bits.Sub64(x[2], y[2], c1)
	tmp[3], c1 = bits.Sub64(x[3], y[3], c1)

	out[0], c = bits.Add64(tmp[0], mod[0], 0)
	out[1], c = bits.Add64(tmp[1], mod[1], c)
	out[2], c = bits.Add64(tmp[2], mod[2], c)
	out[3], c = bits.Add64(tmp[3], mod[3], c)

	if c1 == 0 { // unnecessary add
		*out = tmp
	}

	return nil
}
*/
