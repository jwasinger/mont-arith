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
	x := LimbsToInt(mod).Sub(x, big.NewInt(10))
	y := LimbsToInt(mod).Sub(y, big.NewInt(15))

	// convert to montgomery form
	x.Mul(x, montCtx.RVal())
	x.Mod(x, LimbsToInt(mod))

	y.Mul(y, montCtx.RVal())
	y.Mod(y, LimbsToInt(mod))

	expected := new(big.Int)
	expected.Mul(x, y)
	expected.Mul(expected, montCtx.RInv())
	expected.Mod(expected, LimbsToInt(mod))

	outLimbs := make([]uint, montCtx.NumLimbs)
    xLimbs := IntToLimbs(x)
    yLimbs := IntToLimbs(y)

	montCtx.MulModMont(out, x, y)

    result = LimbsToInt(outLimbs)
	if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", out, expected)
	}
}

func testAddMod(t *testing.T, preset *ArithPreset, limbCount uint) {
	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic(err)
	}

	x := LimbsToInt(mod).Sub(x, big.NewInt(10))
	y := LimbsToInt(mod).Sub(y, big.NewInt(15))

	expected := new(big.Int)
	expected.Add(x, y)
	expected.Mod(expected, LimbsToInt(mod))

    xLimbs := IntToLimbs(x)
    yLimbs := IntToLimbs(y)
	outLimbs := make([]uint, montCtx.NumLimbs)
	montCtx.AddMod(outLimbs, xLimbs, yLimbs)

	result := LimbsToInt(outLimbs)
	if out.Cmp(expected) != 0 {
		t.Fatalf("%x != %x", out, expected)
	}
}

func testSubMod(t *testing.T, preset *ArithPreset, limbCount uint) {
	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := LimbsToInt(mod).Sub(x, big.NewInt(10))
	y := LimbsToInt(mod).Sub(y, big.NewInt(16))

	// convert x/y to montgomery

	expected := new(big.Int)
	expected.Sub(x, y)
	expected.Mod(expected, LimbsToInt(mod))

	outLimbs := make([]byte, montCtx.NumLimbs*8)
    xLimbs := IntToLimbs(x)
    yLimbs := IntToLimbs(y)

	montCtx.SubMod(outLimbs, xLimbs, yLimbs)

	result := LimbsToInt(outLimbs)

	if result.Cmp(expected) != 0 {
		t.Fatal()
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

	outLimbs := make([]uint, montCtx.NumLimbs)
	xLimbs := LimbsToLEBytes(IntToLimbs(x, limbCount))
	yLimbs := LimbsToLEBytes(IntToLimbs(y, limbCount))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulModMont(outLimbs, xLimbs, yLimbs)
	}
}

func benchmarkAddMod(b *testing.B, preset *ArithPreset, limbCount uint) {
	mod := MaxModulus(limbCount)
	modBytes := LimbsToLEBytes(mod)
	montCtx := NewMontArithContext(preset)

	err := montCtx.SetMod(modBytes)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(3)
	y := big.NewInt(4)

	outLimbs := make([]uint, montCtx.NumLimbs)
	xLimbs := LimbsToLEBytes(IntToLimbs(x, limbCount))
	yLimbs := LimbsToLEBytes(IntToLimbs(y, limbCount))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.AddMod(outLimbs, xLimbs, yLimbs)
	}
}

func TestAddMod(t *testing.T) {
	test := func(t *testing.T, preset *ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i < maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("/%s/%d-bit", name, i*64), func(t *testing.T) {
				testAddMod(t, preset, uint(i))
			})
		}
	}

	test(t, Asm384Preset(), "asm384", 1, 128)
	test(t, DefaultPreset(), "non-unrolled", 1, 128)
}

func TestSubMod(t *testing.T) {
	test := func(t *testing.T, preset *ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i < maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("/%s/%d-bit", name, i*64), func(t *testing.T) {
				testSubMod(t, preset, uint(i))
			})
		}
	}

	test(t, Asm384Preset(), "asm384", 1, 128)
	test(t, DefaultPreset(), "non-unrolled", 1, 128)
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

	test(t, DefaultPreset(), "non-unrolled", 1, 128)
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

	bench(b, DefaultPreset(), 1, 128)
}

func BenchmarkAddMod(b *testing.B) {
	bench := func(b *testing.B, preset *ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			b.Run(fmt.Sprintf("%s/%d-bit", name, i*64), func(b *testing.B) {
				benchmarkAddMod(b, preset, uint(i))
			})
		}
	}

	bench(b, DefaultPreset(), "non-unrolled", 1, 128)
}
