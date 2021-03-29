package evmmax_arith

import (
	"math/big"
	"testing"
)

func BenchmarkMulModMontNonInterleaved(b *testing.B) {
	modCtx, err := NewModContext("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab", 16)
	_ = err
	_ = modCtx

	// 5 limb - val = (1<<512) - 1
	x, _ := new(big.Int).SetString("13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084095", 10)
	y, _ := new(big.Int).SetString("13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084095", 10)
	z := new(big.Int)

	z.Mul(x, y)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		z := new(big.Int)
		z.SetBytes(x.Bytes())

		z = new(big.Int)
		z.SetBytes(x.Bytes())

		z = new(big.Int)
		z.SetBytes(x.Bytes())

		z.Mul(x, y)
		z.Mul(x, y)
		z.Mul(x, y)
		z.Mul(x, y)
	}
}
