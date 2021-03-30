package evmmax_arith

import (
	"testing"
)

// test converting 3 to montgomery and back using MulModMont
func TestMulModMontUnrolled_5limbs(t *testing.T) {
	test_val := make([]uint64, 5, 5)
	test_val[0] = 3

	intermediate_val := make([]uint64, 5, 5)

	mod := MaxModulus(5)
	mont_const := MontConstant_Interleaved(mod)
	r_squared := RSquared(mod)

	_MulModMont_5Limbs(intermediate_val, test_val, r_squared, mod, mont_const)

	if LimbsEq(test_val, intermediate_val) {
		t.Fatalf("conversion to montgomery should modify test val")
	}

	_MulModMont_5Limbs(intermediate_val, intermediate_val, One(5), mod, mont_const)

	if !LimbsEq(test_val, intermediate_val) {
		t.Fatalf("conversion back from montgomery failed")
	}
}

func TestAddMod_5limbs(t *testing.T) {

}

func TestSubMod_5limbs(t *testing.T) {

}

// test converting 3 to montgomery and back using MulModMont
func BenchmarkMulModMontUnrolled_5limbs(b *testing.B) {
	test_val := make([]uint64, 5)
	test_val[0] = 3

	intermediate_val := [5]uint64{0, 0, 0, 0, 0}

	mod := MaxModulus(5)
	mont_const := MontConstant_Interleaved(mod)
	r_squared := RSquared(mod)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_MulModMont_5Limbs(intermediate_val[:], test_val, r_squared[:], mod[:], mont_const)
	}
}
