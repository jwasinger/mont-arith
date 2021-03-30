package evmmax_arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

var Zero12Limbs []uint64 = make([]uint64, 12, 12)

func MulModMont_12Limbs(out_bytes, x_bytes, y_bytes *[]byte, modCtx *ModContext) error {
	x := (*[12]uint64)(unsafe.Pointer(&(*x_bytes)[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&(*y_bytes)[0]))[:]
	out := (*[12]uint64)(unsafe.Pointer(&(*out_bytes)[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&modCtx.Modulus[0]))[:]

	return _MulModMont_12Limbs(out, x, y, mod, modCtx.MontParamInterleaved)
}

// NOTE: len(z) == len(x) == len(y) == len(mod) && x,y < mod assumed
func _MulModMont_12Limbs(z, x, y, mod []uint64, modinv uint64) error {
	var t [12]uint64
	var c [12]uint64
	var sub_val []uint64 = mod

	if x[11] >= mod[11] || y[11] >= mod[11] {
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
	c[2], t[8] = madd2(m, mod[9], c[2], c[0])
	c[1], c[0] = madd1(v, y[10], c[1])
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd1(v, y[11], c[1])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
	// round 10
	v = x[10]
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
	c[2], t[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	t[11], t[10] = madd3(m, mod[11], c[0], c[2], c[1])
	// round 11
	v = x[11]
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
	c[2], z[9] = madd2(m, mod[10], c[2], c[0])
	c[1], c[0] = madd2(v, y[11], c[1], t[11])
	z[11], z[10] = madd3(m, mod[11], c[0], c[2], c[1])

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
	_, c[1] = bits.Sub64(z[10], mod[10], 0)
	_, c[1] = bits.Sub64(z[11], mod[11], 0)

	if c[1] != 0 { // unnecessary sub
		sub_val = Zero12Limbs
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
	_, c[1] = bits.Sub64(z[10], sub_val[10], 0)
	_, c[1] = bits.Sub64(z[11], sub_val[11], 0)

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
