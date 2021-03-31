package evmmax_arith

import (
	"testing"
)

// test converting 3 to montgomery and back using MulModMont
func TestMulModMontUnrolled_6limbs(t *testing.T) {
	test_val := make([]uint64, 6, 6)
	test_val[0] = 3

	intermediate_val := make([]uint64, 6, 6)

	mod := MaxModulus(6)
	mont_const := MontConstant_Interleaved(mod)
	r_squared := RSquared(mod)

	_MulModMont_6Limbs(intermediate_val, test_val, r_squared, mod, mont_const)

	if LimbsEq(test_val, intermediate_val) {
		t.Fatalf("conversion to montgomery should modify test val")
	}

	_MulModMont_6Limbs(intermediate_val, intermediate_val, One(6), mod, mont_const)

	if !LimbsEq(test_val, intermediate_val) {
		t.Fatalf("conversion back from montgomery failed")
	}
}

func TestAddMod_6limbs(t *testing.T) {

}

func TestSubMod_6limbs(t *testing.T) {

}

// test converting 3 to montgomery and back using MulModMont
func BenchmarkMulModMontUnrolled_6limbs(b *testing.B) {
	test_val := make([]uint64, 6)
	test_val[0] = 3

	intermediate_val := [6]uint64{0, 0, 0, 0, 0, 0}

	mod := MaxModulus(6)
	mont_const := MontConstant_Interleaved(mod)
	r_squared := RSquared(mod)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_MulModMont_6Limbs(intermediate_val[:], test_val, r_squared[:], mod[:], mont_const)
	}
}