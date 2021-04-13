package modext

import (
	"encoding/binary"
	"math/big"
)

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

// convert big.Int (big-endian) to little-endian limbs
func IntToLimbs(val *big.Int, num_limbs uint) []uint64 {
	val_bytes := val.Bytes()

	// pad length to be a multiple of 64bits
	pad_len := len(val_bytes) - len(val_bytes)%8
	if pad_len > 0 {
		pad := make([]byte, pad_len, pad_len)
		val_bytes = append(val_bytes, pad...)
	}

	result := make([]uint64, len(val_bytes)/8, len(val_bytes)/8)

	// place byteswapped (little-endian) val into result
	for i := 0; i < len(val_bytes); i += 8 {
		binary.LittleEndian.PutUint64(val_bytes[len(val_bytes)-1-i:(len(val_bytes)-1-i)+8], result[i/8])
	}

	return result
}

// convert little-endian limbs to big.Int
func LimbsToInt(limbs []uint64) *big.Int {
	limbs_bytes := make([]byte, 8*len(limbs), 8*len(limbs))
	for i := 0; i < len(limbs); i++ {
		startIdx := (len(limbs) - (i + 1)) * 8
		endIdx := (len(limbs) - i) * 8

		binary.BigEndian.PutUint64(limbs_bytes[startIdx:endIdx], limbs[i])
	}

	return new(big.Int).SetBytes(limbs_bytes)
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
