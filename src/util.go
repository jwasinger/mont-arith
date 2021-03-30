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



func IntToLimbs(val *big.Int) []uint64 {
	// TODO move to global
	var foo uint
	LimbSize := uint(unsafe.Sizeof(foo))
	_ = foo

	num_limbs := uint(math.Ceil(float64(len(val.Bytes())) / float64(unsafe.Sizeof(foo))))
	num_bytes := num_limbs * LimbSize

	val_bytes := make([]byte, num_bytes, num_bytes)
	result := make([]uint64, num_bytes, num_bytes)
	val.FillBytes(val_bytes)

	for i := uint(0); i < num_limbs; i++ {
		result[num_limbs - (i + 1)] = binary.BigEndian.Uint64(val_bytes[i * 8: (i + 1) * 8])
	}

	return result
}
