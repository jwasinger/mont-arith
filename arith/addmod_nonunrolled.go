package arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

func AddModNonUnrolled128(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [2]uint64{0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 2; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 2; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled192(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [3]uint64{0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 3; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 3; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled256(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [4]uint64{0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 4; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 4; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled320(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [5]uint64{0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 5; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 5; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled384(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [6]uint64{0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 6; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 6; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled448(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [7]uint64{0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 7; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 7; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled512(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 8; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 8; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled576(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 9; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 9; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled640(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [10]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 10; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 10; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled704(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [11]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 11; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 11; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled768(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [12]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 12; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 12; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled832(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[13]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[13]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[13]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[13]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [13]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 13; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 13; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled896(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[14]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[14]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[14]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[14]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [14]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 14; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 14; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled960(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[15]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[15]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[15]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[15]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [15]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 15; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 15; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1024(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[16]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[16]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[16]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[16]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [16]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 16; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 16; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1088(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[17]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[17]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[17]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[17]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [17]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 17; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 17; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1152(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[18]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[18]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[18]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[18]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [18]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 18; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 18; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1216(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[19]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[19]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[19]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[19]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [19]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 19; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 19; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1280(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[20]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[20]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[20]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[20]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [20]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 20; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 20; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1344(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[21]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[21]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[21]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[21]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [21]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 21; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 21; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1408(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[22]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[22]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[22]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[22]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [22]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 22; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 22; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1472(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[23]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[23]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[23]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[23]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [23]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 23; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 23; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1536(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[24]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[24]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[24]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[24]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [24]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 24; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 24; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1600(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[25]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[25]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[25]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[25]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [25]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 25; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 25; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1664(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[26]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[26]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[26]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[26]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [26]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 26; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 26; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1728(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[27]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[27]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[27]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[27]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [27]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 27; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 27; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1792(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[28]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[28]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[28]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[28]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [28]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 28; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 28; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1856(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[29]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[29]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[29]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[29]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [29]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 29; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 29; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1920(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[30]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[30]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[30]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[30]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [30]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 30; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 30; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled1984(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[31]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[31]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[31]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[31]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [31]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 31; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 31; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2048(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[32]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[32]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[32]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[32]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [32]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 32; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 32; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2112(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[33]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[33]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[33]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[33]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [33]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 33; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 33; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2176(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[34]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[34]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[34]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[34]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [34]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 34; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 34; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2240(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[35]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[35]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[35]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[35]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [35]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 35; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 35; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2304(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[36]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[36]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[36]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[36]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [36]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 36; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 36; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2368(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[37]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[37]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[37]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[37]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [37]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 37; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 37; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2432(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[38]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[38]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[38]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[38]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [38]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 38; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 38; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2496(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[39]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[39]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[39]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[39]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [39]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 39; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 39; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2560(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[40]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[40]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[40]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[40]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [40]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 40; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 40; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2624(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[41]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[41]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[41]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[41]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [41]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 41; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 41; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2688(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[42]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[42]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[42]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[42]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [42]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 42; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 42; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2752(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[43]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[43]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[43]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[43]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [43]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 43; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 43; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2816(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[44]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[44]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[44]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[44]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [44]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 44; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 44; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2880(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[45]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[45]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[45]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[45]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [45]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 45; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 45; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled2944(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[46]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[46]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[46]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[46]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [46]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 46; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 46; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3008(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[47]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[47]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[47]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[47]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [47]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 47; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 47; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3072(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[48]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[48]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[48]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[48]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [48]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 48; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 48; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3136(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[49]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[49]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[49]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[49]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [49]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 49; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 49; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3200(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[50]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[50]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[50]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[50]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [50]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 50; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 50; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3264(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[51]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[51]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[51]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[51]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [51]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 51; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 51; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3328(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[52]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[52]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[52]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[52]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [52]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 52; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 52; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3392(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[53]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[53]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[53]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[53]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [53]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 53; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 53; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3456(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[54]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[54]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[54]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[54]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [54]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 54; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 54; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3520(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[55]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[55]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[55]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[55]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [55]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 55; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 55; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3584(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[56]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[56]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[56]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[56]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [56]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 56; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 56; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3648(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[57]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[57]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[57]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[57]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [57]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 57; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 57; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3712(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[58]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[58]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[58]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[58]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [58]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 58; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 58; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3776(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[59]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[59]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[59]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[59]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [59]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 59; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 59; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3840(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[60]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[60]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[60]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[60]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [60]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 60; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 60; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3904(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[61]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[61]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[61]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[61]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [61]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 61; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 61; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled3968(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[62]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[62]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[62]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[62]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [62]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 62; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 62; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4032(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[63]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[63]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[63]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[63]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [63]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 63; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 63; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4096(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[64]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[64]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[64]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[64]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [64]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 64; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 64; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4160(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[65]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[65]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[65]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[65]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [65]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 65; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 65; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4224(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[66]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[66]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[66]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[66]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [66]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 66; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 66; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4288(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[67]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[67]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[67]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[67]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [67]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 67; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 67; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4352(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[68]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[68]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[68]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[68]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [68]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 68; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 68; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4416(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[69]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[69]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[69]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[69]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [69]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 69; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 69; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4480(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[70]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[70]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[70]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[70]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [70]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 70; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 70; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4544(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[71]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[71]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[71]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[71]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [71]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 71; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 71; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4608(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[72]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[72]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[72]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[72]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [72]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 72; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 72; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4672(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[73]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[73]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[73]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[73]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [73]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 73; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 73; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4736(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[74]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[74]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[74]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[74]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [74]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 74; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 74; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4800(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[75]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[75]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[75]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[75]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [75]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 75; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 75; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4864(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[76]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[76]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[76]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[76]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [76]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 76; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 76; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4928(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[77]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[77]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[77]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[77]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [77]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 77; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 77; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled4992(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[78]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[78]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[78]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[78]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [78]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 78; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 78; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5056(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[79]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[79]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[79]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[79]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [79]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 79; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 79; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5120(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[80]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[80]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[80]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[80]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [80]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 80; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 80; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5184(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[81]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[81]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[81]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[81]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [81]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 81; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 81; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5248(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[82]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[82]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[82]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[82]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [82]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 82; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 82; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5312(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[83]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[83]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[83]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[83]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [83]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 83; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 83; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5376(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[84]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[84]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[84]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[84]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [84]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 84; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 84; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5440(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[85]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[85]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[85]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[85]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [85]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 85; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 85; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5504(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[86]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[86]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[86]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[86]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [86]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 86; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 86; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5568(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[87]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[87]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[87]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[87]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [87]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 87; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 87; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5632(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[88]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[88]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[88]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[88]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [88]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 88; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 88; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5696(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[89]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[89]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[89]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[89]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [89]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 89; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 89; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5760(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[90]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[90]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[90]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[90]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [90]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 90; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 90; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5824(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[91]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[91]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[91]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[91]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [91]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 91; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 91; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5888(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[92]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[92]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[92]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[92]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [92]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 92; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 92; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled5952(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[93]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[93]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[93]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[93]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [93]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 93; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 93; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6016(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[94]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[94]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[94]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[94]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [94]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 94; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 94; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6080(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[95]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[95]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[95]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[95]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [95]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 95; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 95; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6144(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[96]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[96]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[96]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[96]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [96]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 96; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 96; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6208(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[97]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[97]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[97]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[97]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [97]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 97; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 97; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6272(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[98]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[98]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[98]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[98]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [98]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 98; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 98; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6336(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[99]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[99]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[99]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[99]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [99]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 99; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 99; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6400(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[100]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[100]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[100]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[100]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [100]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 100; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 100; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6464(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[101]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[101]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[101]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[101]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [101]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 101; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 101; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6528(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[102]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[102]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[102]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[102]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [102]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 102; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 102; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6592(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[103]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[103]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[103]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[103]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [103]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 103; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 103; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6656(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[104]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[104]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[104]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[104]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [104]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 104; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 104; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6720(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[105]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[105]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[105]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[105]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [105]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 105; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 105; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6784(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[106]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[106]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[106]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[106]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [106]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 106; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 106; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6848(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[107]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[107]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[107]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[107]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [107]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 107; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 107; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6912(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[108]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[108]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[108]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[108]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [108]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 108; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 108; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled6976(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[109]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[109]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[109]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[109]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [109]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 109; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 109; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7040(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[110]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[110]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[110]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[110]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [110]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 110; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 110; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7104(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[111]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[111]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[111]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[111]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [111]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 111; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 111; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7168(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[112]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[112]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[112]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[112]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [112]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 112; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 112; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7232(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[113]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[113]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[113]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[113]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [113]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 113; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 113; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7296(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[114]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[114]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[114]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[114]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [114]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 114; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 114; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7360(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[115]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[115]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[115]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[115]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [115]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 115; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 115; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7424(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[116]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[116]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[116]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[116]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [116]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 116; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 116; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7488(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[117]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[117]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[117]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[117]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [117]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 117; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 117; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7552(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[118]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[118]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[118]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[118]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [118]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 118; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 118; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7616(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[119]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[119]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[119]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[119]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [119]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 119; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 119; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7680(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[120]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[120]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[120]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[120]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [120]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 120; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 120; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7744(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[121]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[121]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[121]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[121]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [121]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 121; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 121; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7808(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[122]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[122]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[122]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[122]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [122]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 122; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 122; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7872(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[123]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[123]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[123]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[123]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [123]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 123; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 123; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled7936(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[124]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[124]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[124]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[124]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [124]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 124; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 124; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled8000(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[125]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[125]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[125]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[125]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [125]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 125; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 125; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled8064(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[126]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[126]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[126]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[126]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [126]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 126; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 126; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModNonUnrolled8128(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[127]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[127]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[127]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[127]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	tmp := [127]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 127; i++ {
		z[i], c = bits.Add64(x[i], y[i], c)
	}

	c = 0

	for i := 0; i < 127; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}
