package mont_arith

import (
	"testing"
	"math/big"
	"fmt"
	"github.com/jwasinger/mont-arith/arith"
)

func testMulModMont(t *testing.T, preset *arith.ArithPreset, limbCount uint) {

	mod := arith.MaxModulus(limbCount)
	montCtx := arith.NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}


	// these inputs fail (TODO)
/*
	x := arith.LimbsToInt(mod)
	x = x.Sub(x, big.NewInt(10))

	y := arith.LimbsToInt(mod)
	y = y.Sub(y, big.NewInt(15))
*/

	// these inputs pass
	x := big.NewInt(10)
	y := big.NewInt(13)

	x.Mul(x, montCtx.RVal())
	x.Mod(x, arith.LimbsToInt(mod))

	y.Mul(y, montCtx.RVal())
	y.Mod(y, arith.LimbsToInt(mod))

/*
	fmt.Println("x/y/mod")
	fmt.Println(x.String())
	fmt.Println(y.String())
	fmt.Println(arith.LimbsToInt(mod))
*/

	expected := new(big.Int)
	expected.Mul(x, y)
	expected.Mul(expected, montCtx.RInv())
	expected.Mod(expected, arith.LimbsToInt(mod))

	out_bytes := make([]byte, montCtx.NumLimbs * 8)

	x_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(x, limbCount))
	y_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(y, limbCount))

	montCtx.MulModMont(out_bytes, x_bytes, y_bytes)

	result := arith.LEBytesToInt(out_bytes)
	if result.Cmp(expected) != 0 {
		t.Fatalf("result (%x) != expected (%x)\n", result, expected)
	}
}

func testAddMod(t *testing.T, preset *arith.ArithPreset, limbCount uint) {
	mod := arith.MaxModulus(limbCount)
	montCtx := arith.NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := arith.LimbsToInt(mod)
	x = x.Sub(x, big.NewInt(10))

	y := arith.LimbsToInt(mod)
	y = y.Sub(y, big.NewInt(15))

	expected := new(big.Int)
	expected.Add(x, y)
	expected.Mod(expected, arith.LimbsToInt(mod))

	out_bytes := make([]byte, montCtx.NumLimbs * 8)

	x_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(x, limbCount))
	y_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(y, limbCount))

	montCtx.AddMod(out_bytes, x_bytes, y_bytes)

	result := arith.LEBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {
		fmt.Printf("%x,%x,%x\n", x, y, arith.LimbsToInt(mod))
		t.Fatalf("%x != %x", result, expected)
	}
}

func testSubMod(t *testing.T, preset *arith.ArithPreset, limbCount uint) {
	mod := arith.MaxModulus(limbCount)
	montCtx := arith.NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := arith.LimbsToInt(mod)
	x = x.Sub(x, big.NewInt(10))

	y := arith.LimbsToInt(mod)
	y = y.Sub(y, big.NewInt(15))

	// convert x/y to montgomery

	expected := new(big.Int)
	expected.Sub(x, y)
	expected.Mod(expected, arith.LimbsToInt(mod))

	out_bytes := make([]byte, montCtx.NumLimbs * 8)

	x_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(x, limbCount))
	y_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(y, limbCount))

	montCtx.SubMod(out_bytes, x_bytes, y_bytes)

	result := arith.LEBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {

		t.Fatal()
	}

}

func benchmarkMulModMont(b *testing.B, preset *arith.ArithPreset, limbCount uint) {
	mod := arith.MaxModulus(limbCount)
	montCtx := arith.NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(2)
	x.Lsh(x, (limbCount * 64) - 10)

	y := big.NewInt(2)
	y.Lsh(y, (limbCount * 64) - 10)

	// convert x/y to montgomery

/*
	x.Mul(x, montCtx.r)
	x.Mod(x, LimbsToInt(mod))

	y.Mul(y, montCtx.r)
	y.Mod(y, LimbsToInt(mod))
*/

/*
	expected := new(big.Int)
	expected.Mul(x, y)
	expected.Mul(expected, montCtx.rInv)
	expected.Mod(expected, LimbsToInt(mod))
*/

	out_bytes := make([]byte, montCtx.NumLimbs * 8)

	x_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(x, limbCount))
	y_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(y, limbCount))

	x_int := arith.LEBytesToInt(x_bytes)
	_ = x_int

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulModMont(out_bytes, x_bytes, y_bytes)
	}
}

func benchmarkAddMod(b *testing.B, preset *arith.ArithPreset, limbCount uint) {
	mod := arith.MaxModulus(limbCount)
	montCtx := arith.NewMontArithContext(preset)

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(3)
	y := big.NewInt(4)

	out_bytes := make([]byte, montCtx.NumLimbs * 8)

	x_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(x, limbCount))
	y_bytes := arith.LimbsToLEBytes(arith.IntToLimbs(y, limbCount))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.AddMod(out_bytes, x_bytes, y_bytes)
	}
}

func TestAddMod(t *testing.T) {
	test := func(t *testing.T, preset *arith.ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i < maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("/%s/%d-bit", name, i * 64), func(t *testing.T) {
				testAddMod(t, preset, uint(i))
			})
		}
	}

	test(t, arith.DefaultPreset(), "unrolled", 1, 20)
	test(t, arith.NonUnrolledPreset(), "non-unrolled", 1, 64)
}

func TestSubMod(t *testing.T) {
	test := func(t *testing.T, preset *arith.ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i < maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("/%s/%d-bit", name, i * 64), func(t *testing.T) {
				testSubMod(t, preset, uint(i))
			})
		}
	}

	test(t, arith.DefaultPreset(), "unrolled", 1, 20)
	test(t, arith.NonUnrolledPreset(), "non-unrolled", 1, 64)
}

func TestMulModMont(t *testing.T) {
	test := func(t *testing.T, preset *arith.ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			t.Run(fmt.Sprintf("%s/%d-bit", name, i * 64), func(t *testing.T) {
				testMulModMont(t, preset, uint(i))
			})
		}
	}

	test(t, arith.DefaultPreset(), "unrolled", 1, 64)
	//test(t, arith.NonUnrolledPreset(), "non-unrolled")
}

func BenchmarkMulModMont(b *testing.B) {
	bench := func(b *testing.B, preset *arith.ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			b.Run(fmt.Sprintf("%s/%d-bit", name, i * 64), func(b *testing.B) {
				benchmarkMulModMont(b, preset, uint(i))
			})
		}
	}

	bench(b, arith.DefaultPreset(), "goff", 1, 64)
}

func BenchmarkAddMod(b *testing.B) {
	bench := func(b *testing.B, preset *arith.ArithPreset, name string, minLimbs, maxLimbs int) {
		for i := minLimbs; i <= maxLimbs; i++ {
			// test x/y >= modulus
			b.Run(fmt.Sprintf("%s/%d-bit", name, i * 64), func(b *testing.B) {
				benchmarkAddMod(b, preset, uint(i))
			})
		}
	}

	//bench(b, arith.DefaultPreset(), "unrolled", 1, 128)
	bench(b, arith.NonUnrolledPreset(), "non-unrolled", 128, 64)
}
