package modext

import (
	"errors"
	"math/bits"
	"unsafe"
)

var Zero9Limbs []uint64 = make([]uint64, 9, 9)

func SubMod576(out_bytes, x_bytes, y_bytes, modulus []byte) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	out := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&modulus[0]))[:]

	submod576(out, x, y, mod)
	return nil
}

func AddMod576(out_bytes, x_bytes, y_bytes, modulus []byte) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	out := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&modulus[0]))[:]

	addmod576(out, x, y, mod)
	return nil
}

func MulModMont576(out_bytes, x_bytes, y_bytes, modulus []byte, modinv uint64) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	out := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&modulus[0]))[:]

	return mulmodMont576(out, x, y, mod, modinv)
}

/* NOTE: addmod/submod/mulmodmont assume:
len(z) == len(x) == len(y) == len(mod)
*/

func mulmodMont576(z, x, y, mod []uint64, modinv uint64) error {
	var t [9]uint64
	var c [9]uint64
	var sub_val []uint64 = mod

	if x[8] >= mod[8] || y[8] >= mod[8] {
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

	_, c[1] = bits.Sub64(z[0], mod[0], 0)
	_, c[1] = bits.Sub64(z[1], mod[1], 0)
	_, c[1] = bits.Sub64(z[2], mod[2], 0)
	_, c[1] = bits.Sub64(z[3], mod[3], 0)
	_, c[1] = bits.Sub64(z[4], mod[4], 0)
	_, c[1] = bits.Sub64(z[5], mod[5], 0)
	_, c[1] = bits.Sub64(z[6], mod[6], 0)
	_, c[1] = bits.Sub64(z[7], mod[7], 0)
	_, c[1] = bits.Sub64(z[8], mod[8], 0)

	if c[1] != 0 { // unnecessary sub
		sub_val = Zero9Limbs
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

	return nil
}

/*
	Modular Addition
*/

func addmod576(z, x, y, mod []uint64) {
	var c uint64
	tmp := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[8] >= mod[8] || y[8] >= mod[8] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	z[0], c = bits.Add64(x[0], y[0], 0)
	z[1], c = bits.Add64(x[1], y[1], c)
	z[2], c = bits.Add64(x[2], y[2], c)
	z[3], c = bits.Add64(x[3], y[3], c)
	z[4], c = bits.Add64(x[4], y[4], c)
	z[5], c = bits.Add64(x[5], y[5], c)
	z[6], c = bits.Add64(x[6], y[6], c)
	z[7], c = bits.Add64(x[7], y[7], c)
	z[8], c = bits.Add64(x[8], y[8], c)
	tmp[0], c = bits.Sub64(tmp[0], mod[0], 0)
	tmp[1], c = bits.Sub64(tmp[1], mod[1], c)
	tmp[2], c = bits.Sub64(tmp[2], mod[2], c)
	tmp[3], c = bits.Sub64(tmp[3], mod[3], c)
	tmp[4], c = bits.Sub64(tmp[4], mod[4], c)
	tmp[5], c = bits.Sub64(tmp[5], mod[5], c)
	tmp[6], c = bits.Sub64(tmp[6], mod[6], c)
	tmp[7], c = bits.Sub64(tmp[7], mod[7], c)
	tmp[8], c = bits.Sub64(tmp[8], mod[8], c)

	if c == 0 {
		copy(z, tmp[:])
	}
}

func submod576(z, x, y, mod []uint64) {
	var c, c1 uint64
	tmp := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[8] >= mod[8] || y[8] >= mod[8] {
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
	z[0], c1 = bits.Add64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Add64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Add64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Add64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Add64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Add64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Add64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Add64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Add64(tmp[8], mod[8], c1)

	if c == 0 {
		copy(z, tmp[:])
	}

}
