package modext

import (
	"fmt"
	"math/big"
	"testing"
)

func testMulModMont(t *testing.T, limbCount uint) {

	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext()

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(3)
	y := big.NewInt(4)

	// convert x/y to montgomery

	x.Mul(x, montCtx.r)
	x.Mod(x, LimbsToInt(mod))

	y.Mul(y, montCtx.r)
	y.Mod(y, LimbsToInt(mod))

	expected := new(big.Int)
	expected.Mul(x, y)
	expected.Mul(expected, montCtx.rInv)
	expected.Mod(expected, LimbsToInt(mod))

	out_bytes := make([]byte, montCtx.NumLimbs*8)

	fmt.Println("testMulModMont")
	fmt.Println(montCtx.Modulus)
	fmt.Println(IntToLimbs(x, limbCount))
	x_bytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToLEBytes(IntToLimbs(y, limbCount))
	fmt.Println(x_bytes)

	montCtx.MulModMont(out_bytes, x_bytes, y_bytes)

	result := LEBytesToInt(out_bytes)
	if result.Cmp(expected) != 0 {
		fmt.Println(result.String())
		fmt.Println(expected.String())
		t.Fail()
	}
}

func testAddMod(t *testing.T, limbCount uint) {
	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext()

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(3)
	y := big.NewInt(4)

	expected := new(big.Int)
	expected.Add(x, y)
	expected.Mod(expected, LimbsToInt(mod))

	out_bytes := make([]byte, montCtx.NumLimbs*8)

	x_bytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

	montCtx.AddMod(out_bytes, x_bytes, y_bytes)

	result := LEBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {
		t.Fail()
	}
}

func testSubMod(t *testing.T, limbCount uint) {
	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext()

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(3)
	y := big.NewInt(4)

	// convert x/y to montgomery

	expected := new(big.Int)
	expected.Sub(x, y)
	expected.Mod(expected, LimbsToInt(mod))

	out_bytes := make([]byte, montCtx.NumLimbs*8)

	x_bytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

	montCtx.SubMod(out_bytes, x_bytes, y_bytes)

	result := LEBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {

		t.Fatal()
	}

}

func TestMulModMont(t *testing.T) {
	for i := 1; i < 2; i++ {
		testMulModMont(t, uint(i))
	}
}

func TestAddMod(t *testing.T) {
	for i := 1; i < 128; i++ {
		// TODO:
		// test overflow
		// test value in middle of the range
		// test x/y = 0
		// test no-clobber x/y
		testAddMod(t, uint(i))
	}
}

func TestSubMod(t *testing.T) {
	for i := 1; i < 128; i++ {
		// TODO:
		// test underflow
		// test value in middle of the range
		// test x/y = 0
		// test no-clobber x/y
		testSubMod(t, uint(i))
	}
}

func benchmarkMulModMont(b *testing.B, limbCount uint) {

	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext()

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(2)
	x.Lsh(x, (limbCount*64)-10)

	y := big.NewInt(2)
	y.Lsh(y, (limbCount*64)-10)

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

	out_bytes := make([]byte, montCtx.NumLimbs*8)

	x_bytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

	x_int := LEBytesToInt(x_bytes)
	_ = x_int

	b.ResetTimer()

	//fmt.Println()

	for i := 0; i < b.N; i++ {
		montCtx.MulModMont(out_bytes, x_bytes, y_bytes)
	}
}

func benchmarkAddMod(b *testing.B, limbCount uint) {
	mod := MaxModulus(limbCount)
	montCtx := NewMontArithContext()

	err := montCtx.SetMod(mod)
	if err != nil {
		panic("error")
	}

	x := big.NewInt(3)
	y := big.NewInt(4)

	out_bytes := make([]byte, montCtx.NumLimbs*8)

	x_bytes := LimbsToLEBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToLEBytes(IntToLimbs(y, limbCount))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.AddMod(out_bytes, x_bytes, y_bytes)
	}
}

/*
func BenchmarkAddMod(b *testing.B) {
	for i := 1; i < 128; i++ {
		// TODO
		// test x/y >= modulus
		b.Run(fmt.Sprintf("num_limbs=%d", i), func(b *testing.B) {
			benchmarkAddMod(b, uint(i))
		})
	}

}
*/

func BenchmarkMulModMont(b *testing.B) {
	for i := 100; i < 128; i++ {
		// TODO
		// test x/y >= modulus
		b.Run(fmt.Sprintf("num_limbs=%d", i), func(b *testing.B) {
			benchmarkMulModMont(b, uint(i))
		})
	}
}
