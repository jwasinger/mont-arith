package arith

import (
	"errors"
	"math/bits"
	"unsafe"
)

func AddModUnrolled128(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[2]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[2]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[2]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[2]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 2, 2)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled192(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[3]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[3]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[3]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[3]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 3, 3)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled256(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[4]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[4]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[4]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[4]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 4, 4)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled320(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[5]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[5]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[5]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[5]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 5, 5)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled384(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 6, 6)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled448(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[7]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[7]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[7]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[7]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 7, 7)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled512(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[8]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[8]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[8]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[8]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 8, 8)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled576(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[9]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[9]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[9]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[9]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 9, 9)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled640(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[10]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[10]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[10]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[10]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 10, 10)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled704(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[11]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[11]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[11]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[11]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 11, 11)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled768(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[12]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[12]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[12]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[12]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 12, 12)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled832(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[13]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[13]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[13]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[13]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 13, 13)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled896(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[14]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[14]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[14]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[14]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 14, 14)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled960(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[15]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[15]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[15]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[15]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 15, 15)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled1024(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[16]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[16]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[16]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[16]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 16, 16)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled1088(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[17]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[17]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[17]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[17]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 17, 17)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled1152(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[18]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[18]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[18]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[18]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 18, 18)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled1216(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[19]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[19]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[19]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[19]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 19, 19)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}

func AddModUnrolled1280(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[20]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[20]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[20]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[20]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c, c1 uint64
	tmp := make([]uint64, 20, 20)

	if x[0] >= mod[0] || y[0] >= mod[0] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)
	tmp[4], c = bits.Add64(x[4], y[4], c)
	tmp[5], c = bits.Add64(x[5], y[5], c)
	tmp[6], c = bits.Add64(x[6], y[6], c)
	tmp[7], c = bits.Add64(x[7], y[7], c)
	tmp[8], c = bits.Add64(x[8], y[8], c)
	tmp[9], c = bits.Add64(x[9], y[9], c)
	tmp[10], c = bits.Add64(x[10], y[10], c)
	tmp[11], c = bits.Add64(x[11], y[11], c)
	tmp[12], c = bits.Add64(x[12], y[12], c)
	tmp[13], c = bits.Add64(x[13], y[13], c)
	tmp[14], c = bits.Add64(x[14], y[14], c)
	tmp[15], c = bits.Add64(x[15], y[15], c)
	tmp[16], c = bits.Add64(x[16], y[16], c)
	tmp[17], c = bits.Add64(x[17], y[17], c)
	tmp[18], c = bits.Add64(x[18], y[18], c)
	tmp[19], c = bits.Add64(x[19], y[19], c)
	z[0], c1 = bits.Sub64(tmp[0], mod[0], 0)
	z[1], c1 = bits.Sub64(tmp[1], mod[1], c1)
	z[2], c1 = bits.Sub64(tmp[2], mod[2], c1)
	z[3], c1 = bits.Sub64(tmp[3], mod[3], c1)
	z[4], c1 = bits.Sub64(tmp[4], mod[4], c1)
	z[5], c1 = bits.Sub64(tmp[5], mod[5], c1)
	z[6], c1 = bits.Sub64(tmp[6], mod[6], c1)
	z[7], c1 = bits.Sub64(tmp[7], mod[7], c1)
	z[8], c1 = bits.Sub64(tmp[8], mod[8], c1)
	z[9], c1 = bits.Sub64(tmp[9], mod[9], c1)
	z[10], c1 = bits.Sub64(tmp[10], mod[10], c1)
	z[11], c1 = bits.Sub64(tmp[11], mod[11], c1)
	z[12], c1 = bits.Sub64(tmp[12], mod[12], c1)
	z[13], c1 = bits.Sub64(tmp[13], mod[13], c1)
	z[14], c1 = bits.Sub64(tmp[14], mod[14], c1)
	z[15], c1 = bits.Sub64(tmp[15], mod[15], c1)
	z[16], c1 = bits.Sub64(tmp[16], mod[16], c1)
	z[17], c1 = bits.Sub64(tmp[17], mod[17], c1)
	z[18], c1 = bits.Sub64(tmp[18], mod[18], c1)
	z[19], c1 = bits.Sub64(tmp[19], mod[19], c1)

	if c == 0 || c != 0 && c1 == 0 {
		copy(z, tmp[:])
	}

	return nil
}
