package arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

func SubModUnrolled128(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [2]uint64{0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 2; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 2; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled192(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [3]uint64{0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 3; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 3; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled256(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [4]uint64{0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 4; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 4; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled320(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [5]uint64{0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 5; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 5; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled384(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [6]uint64{0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 6; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 6; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled448(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [7]uint64{0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 7; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 7; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled512(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [8]uint64{0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 8; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 8; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled576(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [9]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 9; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 9; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled640(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [10]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 10; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 10; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled704(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [11]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 11; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 11; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled768(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [12]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 12; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 12; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled832(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[13]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[13]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[13]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[13]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [13]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 13; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 13; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled896(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[14]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[14]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[14]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[14]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [14]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 14; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 14; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled960(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[15]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[15]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[15]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[15]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [15]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 15; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 15; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled1024(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[16]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[16]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[16]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[16]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [16]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 16; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 16; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled1088(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[17]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[17]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[17]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[17]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [17]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 17; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 17; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled1152(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[18]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[18]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[18]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[18]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [18]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 18; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 18; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled1216(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[19]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[19]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[19]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[19]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [19]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 19; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 19; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func SubModUnrolled1280(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[20]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[20]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[20]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[20]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [20]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 20; i++ {
		tmp[i], c = bits.Sub64(x[i], y[i], c)
	}

	for i := 0; i < 20; i++ {
		z[i], c1 = bits.Add64(tmp[i], mod[i], c1)
	}

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}
