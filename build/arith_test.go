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

	x_bytes := LimbsToMAXBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToMAXBytes(IntToLimbs(y, limbCount))

	montCtx.MulModMont(out_bytes, x_bytes, y_bytes)

	result := MAXBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {
		t.Fatal()
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

	// convert x/y to montgomery

	expected := new(big.Int)
	expected.Add(x, y)
	expected.Mod(expected, LimbsToInt(mod))

	out_bytes := make([]byte, montCtx.NumLimbs*8)

	x_bytes := LimbsToMAXBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToMAXBytes(IntToLimbs(y, limbCount))

	montCtx.AddMod(out_bytes, x_bytes, y_bytes)

	result := MAXBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {
		t.Fatal()
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

	x_bytes := LimbsToMAXBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToMAXBytes(IntToLimbs(y, limbCount))

	montCtx.SubMod(out_bytes, x_bytes, y_bytes)

	result := MAXBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {
		t.Fatal()
	}

}

func TestMulModMont(t *testing.T) {
	for i := 4; i < 10; i++ {
		// TODO
		// test x/y >= modulus
		testMulModMont(t, uint(i))
	}
}

func TestAddMod(t *testing.T) {
	for i := 4; i < 10; i++ {
		// TODO:
		// test overflow
		// test value in middle of the range
		// test x/y = 0
		// test no-clobber x/y
		testAddMod(t, uint(i))
	}
}

func TestSubMod(t *testing.T) {
	for i := 4; i < 10; i++ {
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

	x_bytes := LimbsToMAXBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToMAXBytes(IntToLimbs(y, limbCount))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		montCtx.MulModMont(out_bytes, x_bytes, y_bytes)
	}
}

func BenchmarkMulModMont(b *testing.B) {
	for i := 4; i < 10; i++ {
		// TODO
		// test x/y >= modulus
		b.Run(fmt.Sprintf("num_limbs=%d", i), func(b *testing.B) {
			benchmarkMulModMont(b, uint(i))
		})
	}
}
