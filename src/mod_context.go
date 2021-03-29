package evmmax_arith

import (
	"math/big"
)

type ModContext struct {
    Modulus []byte
    MontParamInterleaved uint64

    MontParamNonInterleaved *big.Int
    ModulusNonInterleaved *big.Int // just here for convenience

    NumLimbs uint
}
