package evmmax_arith

import (
	"unsafe"
	"math/big"
	"math"
	"encoding/binary"
)

func LimbsToInt(limbs []uint64) big.Int {
	limbs_bytes := make([]byte, 8 * len(limbs), 8 * len(limbs))
	for i := 0; i < len(limbs); i++ {
		binary.BigEndian.PutUint64(limbs_bytes[i * 8:(i + 1) * 8], limbs[len(limbs) - (i + 1)])
	}

	return *(new(big.Int).SetBytes(limbs_bytes))
}

func IntToLEBytes_Uint64Limbs(val *big.Int) []byte {

}

func IntToLimbs(val *big.Int, num_limbs uint) []uint64 {
	// TODO move to global
	var foo uint
	LimbSize := uint(unsafe.Sizeof(foo))
	_ = foo

	val_limbs := uint(math.Ceil(float64(len(val.Bytes())) / float64(unsafe.Sizeof(foo))))
	if val_limbs > num_limbs {
		panic("val cannot fit inside specified limb count")
	}

	num_bytes := num_limbs * LimbSize

	val_bytes := make([]byte, num_bytes, num_bytes)
	result := make([]uint64, num_limbs, num_limbs)
	val.FillBytes(val_bytes)

	for i := uint(0); i < num_limbs; i++ {
		result[num_limbs - (i + 1)] = binary.BigEndian.Uint64(val_bytes[i * 8: (i + 1) * 8])
	}

	return result
}

// **NOTE** naming confusing.  actually the second-largest modulus (largest would have modinv as 1)
func MaxModulus(limbCount uint) []uint64 {
	mod := make([]uint64, limbCount, limbCount)

	mod[0] = 0xfffffffffffffffd
	for i := uint(1); i < limbCount; i++ {
		mod[i] = 0xffffffffffffffff
	}

	return mod
}

func LimbsEq(x, y []uint64) bool {
	if len(x) != len(y) {
		panic("unequally-sized elements")
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func One(limbCount uint) []uint64 {
	one := make([]uint64, limbCount, limbCount)
	one[0] = 1
	return one
}

func RSquared(modulus []uint64) []uint64 {
	mod := LimbsToInt(modulus[:])
	r := new(big.Int)
	r.Exp(big.NewInt(2), big.NewInt(int64(len(modulus)) * 64), &mod)
	r.Mul(r, r)
	r.Mod(r, &mod)

	result := IntToLimbs(r, uint(len(modulus)))
	return result
}

// does the Python equivalent of pow(-modulus, -1, 1<<64)
func MontConstant_Interleaved(modulus []uint64) uint64 {
	mod_int := LimbsToInt(modulus)

	// 1<<64
	aux_mod, _ := new(big.Int).SetString("18446744073709551616", 10)
	negative_one, _ := new(big.Int).SetString("-1", 10)

	mod_int.Mul(&mod_int, negative_one)
	return mod_int.ModInverse(&mod_int, aux_mod).Uint64()
}
