package evmmax_arith

import (
	"encoding/binary"
	"math/big"
	"testing"
)

// **NOTE** naming confusing.  actually the second-largest modulus (largest would have modinv as 1)
func MaxModulus() [4]uint64 {
	mod := [4]uint64{0, 0, 0, 0}

	mod[0] = 0xfffffffffffffffd
	for i := 1; i < 4; i++ {
		mod[i] = 0xffffffffffffffff
	}

	return mod
}

func Eq_4Limbs(x, y [4]uint64) bool {
	if x[0] != y[0] {
		return false
	}
	if x[1] != y[1] {
		return false
	}
	if x[2] != y[2] {
		return false
	}
	if x[3] != y[3] {
		return false
	}

	return true
}

func IntToLimbs_4Limbs(val *big.Int) [4]uint64 {
	val_bytes := make([]byte, 4*8, 4*8)
	result := [4]uint64{0, 0, 0, 0}
	val.FillBytes(val_bytes)

	for i := 0; i < 4; i++ {
		result[4-(i+1)] = binary.BigEndian.Uint64(val_bytes[i*8 : (i+1)*8])
	}

	return result
}

func LimbsToInt_4Limbs(limbs []uint64) big.Int {
	limbs_bytes := make([]byte, 8*len(limbs), 8*len(limbs))
	for i := 0; i < len(limbs); i++ {
		binary.BigEndian.PutUint64(limbs_bytes[i*8:(i+1)*8], limbs[len(limbs)-(i+1)])
	}

	return *(new(big.Int).SetBytes(limbs_bytes))
}

// does the Python equivalent of pow(-modulus, -1, 1<<64)
func MontConstant(modulus [4]uint64) uint64 {
	mod_int := LimbsToInt_4Limbs(modulus[:])

	// 1<<64
	aux_mod, _ := new(big.Int).SetString("18446744073709551616", 10)
	negative_one, _ := new(big.Int).SetString("-1", 10)

	mod_int.Mul(&mod_int, negative_one)
	return mod_int.ModInverse(&mod_int, aux_mod).Uint64()
}

func RSquared(modulus [4]uint64) [4]uint64 {
	mod := LimbsToInt_4Limbs(modulus[:])
	r := new(big.Int)
	r.Exp(big.NewInt(2), big.NewInt(4*64), &mod)
	r.Mul(r, r)
	r.Mod(r, &mod)
	return IntToLimbs_4Limbs(r)
}

func One(size int) *[4]uint64 {
	one := [4]uint64{0, 0, 0, 0}
	one[0] = 1
	return &one
}

// test converting 3 to montgomery and back using MulModMont
func TestMulModMontUnrolled_4limbs(t *testing.T) {
	test_val := [4]uint64{0, 0, 0, 0}
	test_val[0] = 3

	intermediate_val := [4]uint64{0, 0, 0, 0}

	mod := MaxModulus()
	mont_const := MontConstant(mod)
	r_squared := RSquared(mod)

	_MulModMont(&intermediate_val, &test_val, &r_squared, &mod, mont_const)

	if Eq_4Limbs(test_val, intermediate_val) {
		t.Fatalf("conversion to montgomery should modify test val")
	}

	_MulModMont(&intermediate_val, &intermediate_val, One(4), &mod, mont_const)

	if !Eq_4Limbs(test_val, intermediate_val) {
		t.Fatalf("conversion back from montgomery failed")
	}
}

func TestAddMod_4limbs(t *testing.T) {

}

func TestSubMod_4limbs(t *testing.T) {

}
