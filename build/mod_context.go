package evmmax_arith

import (
	"errors"
	"math/big"
	"unsafe"
)

type ModContext struct {
	Modulus               []byte
	ModulusNonInterleaved *big.Int // just here for convenience

	MontParamInterleaved    uint64
	MontParamNonInterleaved *big.Int

	NumLimbs uint
}

func NewModContext(modulus string, base int) (*ModContext, error) {
	mod, success := new(big.Int).SetString(modulus, base)
	if !success {
		return nil, errors.New("SetString failed for modulus/base")
	}

	var x uint
	limbSize := uint(unsafe.Sizeof(x))
	limbCount := uint(len(mod.Bytes())) / limbSize
	_ = x

	// r val chosen as max representable value for limbCount + 1: 0x1000...000
	rVal := new(big.Int)
	rVal.Lsh(big.NewInt(1), limbCount*limbSize*8)

	montParamNonInterleaved := new(big.Int)
	montParamNonInterleaved = montParamNonInterleaved.Mul(mod, big.NewInt(-1))
	if montParamNonInterleaved.ModInverse(montParamNonInterleaved, rVal) == nil {
		return nil, errors.New("modinverse failed")
	}

	// mod % (1 << limb_count_bits)  == mod % (1 << limb_count_bytes * 8)
	montParamInterleaved := rVal.Uint64()

	ret := ModContext{
		make([]byte, 32, 32), // TODO ??
		mod,
		montParamInterleaved,
		montParamNonInterleaved,
		limbCount,
	}
	return &ret, nil
}
