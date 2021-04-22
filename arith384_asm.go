package mont_arith

import (
	"unsafe"
	"github.com/jwasinger/mont-arith/arith384_asm"
)

func MulModMont384_asm(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	arith384_asm.MulMod(z, x, y, mod, ctx.MontParamInterleaved)
}

func AddMod384_asm(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	arith384_asm.AddMod(z, x, y, mod)
}

func SubMod384_asm(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) (error) {
	x := (*[6]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[6]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[6]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[6]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]
	arith384_asm.SubMod(z, x, y, mod)
}
