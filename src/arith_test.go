package modext

import (
	"testing"
	"math/big"
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

	out_bytes := make([]byte, montCtx.NumLimbs * 8)

	x_bytes := LimbsToMAXBytes(IntToLimbs(x, limbCount))
	y_bytes := LimbsToMAXBytes(IntToLimbs(y, limbCount))

	montCtx.MulModMont(out_bytes, x_bytes, y_bytes)

	result := MAXBytesToInt(out_bytes)

	if result.Cmp(expected) != 0 {
		t.Fatal()
	}
}

func TestMulModMont(t *testing.T) {
	testMulModMont(t, 4)
}
