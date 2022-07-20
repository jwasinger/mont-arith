package mont_arith

import (
	"fmt"
	"math/big"
	"testing"
)

func testMulModMont(t *testing.T, preset *ArithPreset, limbCount uint) {
	mod := GenTestModulus(limbCount)

	montCtx := NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}
	x := LimbsToInt(mod)
	x = x.Sub(x, big.NewInt(10))
	y := LimbsToInt(mod)
	y = y.Sub(y, big.NewInt(15))

	// convert to montgomery form
	x.Mul(x, montCtx.RVal())
	x.Mod(x, LimbsToInt(mod))

	y.Mul(y, montCtx.RVal())
	y.Mod(y, LimbsToInt(mod))

	expected := new(big.Int)
	expected.Mul(x, y)
	expected.Mul(expected, montCtx.RInv())
	expected.Mod(expected, LimbsToInt(mod))

	outLimbs := make(nat, montCtx.NumLimbs)
	xLimbs := IntToLimbs(x, montCtx.NumLimbs)
	yLimbs := IntToLimbs(y, montCtx.NumLimbs)

	montCtx.MulModMont(outLimbs, xLimbs, yLimbs)

	result := LimbsToInt(outLimbs)
	if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
	}
}

func benchmarkMulModMont(b *testing.B, preset *ArithPreset, limbCount uint) {
	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(2)
	x.Lsh(x, (limbCount*64)-10)

	y := big.NewInt(2)
	y.Lsh(y, (limbCount*64)-10)

	// convert x/y to montgomery

	outLimbs := make(nat, montCtx.NumLimbs)
	xLimbs := IntToLimbs(x, limbCount)
	yLimbs := IntToLimbs(y, limbCount)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulModMont(outLimbs, xLimbs, yLimbs)
	}
}

func TestMulModMont(t *testing.T) {
	test := func(t *testing.T, preset *ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("%s/%d-bit", name, i*64), func(t *testing.T) {
				testMulModMont(t, preset, uint(i))
			})
		}
	}

	test(t, DefaultPreset(), "non-unrolled", 1, 12)
}

func BenchmarkMulModMont(b *testing.B) {
	bench := func(b *testing.B, preset *ArithPreset, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			b.Run(fmt.Sprintf("%d-bit", i*64), func(b *testing.B) {
				benchmarkMulModMont(b, preset, uint(i))
			})
		}
	}

	bench(b, DefaultPreset(), 1, 12)
}
