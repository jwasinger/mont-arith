package mont_arith

import (
	"math/bits"
	"unsafe"
	"errors"
)




func AddModNonUnrolled128(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [2]uint64 { 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 2; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 2; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled192(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [3]uint64 { 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 3; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 3; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled256(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [4]uint64 { 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 4; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 4; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled320(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [5]uint64 { 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 5; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 5; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled384(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [6]uint64 { 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 6; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 6; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled448(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [7]uint64 { 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 7; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 7; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled512(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [8]uint64 { 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 8; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 8; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled576(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [9]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 9; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 9; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled640(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [10]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 10; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 10; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled704(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [11]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 11; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 11; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled768(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [12]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 12; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 12; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled832(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[13]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[13]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[13]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[13]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [13]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 13; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 13; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled896(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[14]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[14]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[14]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[14]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [14]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 14; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 14; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled960(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[15]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[15]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[15]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[15]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [15]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 15; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 15; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1024(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[16]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[16]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[16]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[16]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [16]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 16; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 16; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1088(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[17]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[17]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[17]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[17]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [17]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 17; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 17; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1152(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[18]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[18]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[18]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[18]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [18]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 18; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 18; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1216(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[19]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[19]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[19]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[19]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [19]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 19; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 19; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1280(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[20]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[20]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[20]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[20]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [20]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 20; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 20; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1344(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[21]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[21]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[21]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[21]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [21]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 21; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 21; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1408(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[22]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[22]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[22]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[22]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [22]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 22; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 22; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1472(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[23]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[23]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[23]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[23]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [23]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 23; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 23; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1536(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[24]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[24]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[24]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[24]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [24]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 24; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 24; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1600(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[25]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[25]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[25]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[25]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [25]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 25; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 25; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1664(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[26]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[26]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[26]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[26]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [26]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 26; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 26; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1728(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[27]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[27]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[27]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[27]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [27]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 27; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 27; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1792(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[28]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[28]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[28]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[28]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [28]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 28; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 28; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1856(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[29]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[29]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[29]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[29]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [29]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 29; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 29; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1920(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[30]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[30]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[30]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[30]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [30]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 30; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 30; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled1984(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[31]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[31]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[31]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[31]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [31]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 31; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 31; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2048(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[32]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[32]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[32]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[32]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [32]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 32; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 32; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2112(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[33]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[33]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[33]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[33]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [33]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 33; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 33; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2176(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[34]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[34]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[34]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[34]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [34]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 34; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 34; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2240(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[35]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[35]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[35]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[35]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [35]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 35; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 35; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2304(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[36]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[36]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[36]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[36]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [36]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 36; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 36; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2368(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[37]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[37]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[37]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[37]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [37]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 37; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 37; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2432(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[38]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[38]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[38]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[38]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [38]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 38; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 38; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2496(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[39]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[39]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[39]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[39]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [39]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 39; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 39; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2560(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[40]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[40]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[40]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[40]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [40]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 40; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 40; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2624(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[41]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[41]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[41]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[41]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [41]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 41; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 41; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2688(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[42]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[42]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[42]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[42]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [42]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 42; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 42; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2752(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[43]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[43]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[43]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[43]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [43]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 43; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 43; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2816(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[44]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[44]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[44]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[44]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [44]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 44; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 44; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2880(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[45]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[45]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[45]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[45]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [45]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 45; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 45; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled2944(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[46]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[46]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[46]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[46]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [46]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 46; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 46; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3008(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[47]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[47]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[47]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[47]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [47]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 47; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 47; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3072(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[48]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[48]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[48]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[48]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [48]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 48; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 48; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3136(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[49]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[49]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[49]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[49]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [49]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 49; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 49; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3200(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[50]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[50]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[50]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[50]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [50]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 50; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 50; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3264(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[51]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[51]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[51]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[51]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [51]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 51; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 51; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3328(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[52]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[52]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[52]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[52]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [52]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 52; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 52; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3392(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[53]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[53]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[53]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[53]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [53]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 53; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 53; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3456(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[54]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[54]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[54]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[54]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [54]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 54; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 54; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3520(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[55]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[55]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[55]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[55]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [55]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 55; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 55; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3584(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[56]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[56]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[56]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[56]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [56]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 56; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 56; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3648(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[57]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[57]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[57]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[57]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [57]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 57; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 57; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3712(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[58]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[58]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[58]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[58]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [58]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 58; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 58; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3776(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[59]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[59]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[59]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[59]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [59]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 59; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 59; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3840(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[60]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[60]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[60]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[60]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [60]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 60; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 60; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3904(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[61]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[61]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[61]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[61]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [61]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 61; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 61; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled3968(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[62]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[62]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[62]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[62]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [62]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 62; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 62; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled4032(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[63]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[63]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[63]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[63]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [63]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 63; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 63; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}





func AddModNonUnrolled4096(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[64]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[64]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[64]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[64]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64 = 0
	var c1 uint64 = 0

	tmp := [64]uint64 { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}

	for i := 0; i < 64; i++ {
		tmp[i], c = bits.Add64(x[i], y[i], c)
	}

	for i := 0; i < 64; i++ {
		z[i], c1 = bits.Sub64(tmp[i], mod[i], c1)
	}

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

