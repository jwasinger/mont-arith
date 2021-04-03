package modext

import (
	"encoding/binary"
	"math"
	"math/big"
	"unsafe"
)

func LimbsToInt(limbs []uint64) *big.Int {
	limbs_bytes := make([]byte, 8*len(limbs), 8*len(limbs))
	for i := 0; i < len(limbs); i++ {
		startIdx := (len(limbs) * 8) - ((i + 1) * 8)
		endIdx := (len(limbs) * 8) - (i * 8)

		binary.BigEndian.PutUint64(limbs_bytes[startIdx:endIdx], limbs[i])
	}

	return new(big.Int).SetBytes(limbs_bytes)
}

/*
Methods for converting between Go big-int (64bit little-endian limbs, big-endian limb ordering)
and the bigint representation expected here: 64bit little-endian limbs, little-endian ordered
*/
func MAXBytesToInt(v []byte) *big.Int {
	result := new(big.Int)

	if len(v)%8 != 0 {
		panic("invalid val length for modext bytes")
	}

	val := make([]byte, len(v), len(v))
	copy(val, v)

	// byteswap 8 bytes at a time
	for i := 0; i < len(val)/2; i++ {
		tmp := val[i]
		val[i] = val[len(val)-1-i]
		val[len(val)-1-i] = tmp
	}

	result.SetBytes(val)
	return result
}

func LimbsToMAXBytes(val []uint64) []byte {
	result := make([]byte, len(val)*8)

	for i := 0; i < len(val); i++ {
		for j := 0; j < 8; j++ {
			result[i*8+j] = byte(val[i] >> (j * 8))
		}
	}

	return result
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
		result[num_limbs-(i+1)] = binary.BigEndian.Uint64(val_bytes[i*8 : (i+1)*8])
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

// return a hex-string literal for given limbs
/*
func LimbsToString(limbs []uint64) string {
	result := ""
	for i := 0; i < len(limbs); i++ {
		result += fmt.Sprintf("%06x", limbs[i])
	}
}
*/

func One(limbCount uint) []uint64 {
	one := make([]uint64, limbCount, limbCount)
	one[0] = 1
	return one
}

func RSquared(modulus []uint64) []uint64 {
	mod := LimbsToInt(modulus[:])
	r := new(big.Int)
	r.Exp(big.NewInt(2), big.NewInt(int64(len(modulus))*64), mod)
	r.Mul(r, r)
	r.Mod(r, mod)

	result := IntToLimbs(r, uint(len(modulus)))
	return result
}

// does the Python equivalent of pow(-modulus, -1, 1<<64)
func MontConstant_Interleaved(modulus []uint64) uint64 {
	mod_int := LimbsToInt(modulus)

	// 1<<64
	aux_mod, _ := new(big.Int).SetString("18446744073709551616", 10)
	negative_one, _ := new(big.Int).SetString("-1", 10)

	mod_int.Mul(mod_int, negative_one)
	return mod_int.ModInverse(mod_int, aux_mod).Uint64()
}
