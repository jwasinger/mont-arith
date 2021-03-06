package mont_arith

import (
	"errors"
	"fmt"
	"math/big"
)

type MontArithContext struct {
	// TODO make most of these private and the arith operations methods of this struct
	Modulus               nat
	ModulusNonInterleaved *big.Int // just here for convenience XXX better naming

	MontParamInterleaved    Word
	MontParamNonInterleaved *big.Int

	NumLimbs uint

	r    *big.Int
	rInv *big.Int

	// mask for mod by R: 0xfff...fff - (1 << NumLimbs * 64) - 1
	mask *big.Int
    montMul mulMontFunc
}

func (m *MontArithContext) RVal() *big.Int {
	return m.r
}

func (m *MontArithContext) RInv() *big.Int {
	return m.rInv
}

func (m *MontArithContext) ToMont(val nat) nat {
	dst_val := new(big.Int)
	src_val := LimbsToInt(val)
	dst_val.Mul(src_val, m.r)
	dst_val.Mod(dst_val, LimbsToInt(m.Modulus))

	//copy(dst, IntToLimbs(dst_val, m.NumLimbs))
    return IntToLimbs(dst_val, m.NumLimbs)
}

func (m *MontArithContext) ToNorm(val nat) nat {
	dst_val := new(big.Int)
	src_val := LimbsToInt(val)
	dst_val.Mul(src_val, m.rInv)
	dst_val.Mod(dst_val, LimbsToInt(m.Modulus))

	return IntToLimbs(dst_val, m.NumLimbs)
}

func NewMontArithContext() *MontArithContext {
	result := MontArithContext{
		nil,
		nil,
		0,
		nil,

		0,
		nil,
		nil,
		nil,

        nil,
	}

	return &result
}

func (m *MontArithContext) MulModMont(out, x, y nat) {
	m.montMul(out, x, y, m.Modulus, m.MontParamInterleaved)
}

func (m *MontArithContext) ModIsSet() bool {
	return m.NumLimbs != 0
}

func (m *MontArithContext) ValueSize() uint {
	return uint(len(m.Modulus))
}

func (m *MontArithContext) SetMod(mod nat) error {
	// XXX proper handling without hardcoding
	if len(mod) == 0 || len(mod) > 12 {
		fmt.Println(len(mod))
		panic("invalid mod length")
	}

	var limbCount uint = uint(len(mod))
	var limbSize uint = 8

	// r val chosen as max representable value for limbCount + 1: 0x1000...000
	rVal := new(big.Int)
	rVal.Lsh(big.NewInt(1), limbCount*limbSize*8)

	rValMask := new(big.Int)
	rValMask.Sub(rVal, big.NewInt(1))

	modInt := LimbsToInt(mod)
	montParamNonInterleaved := new(big.Int)
	montParamNonInterleaved = montParamNonInterleaved.Mul(modInt, big.NewInt(-1))
	montParamNonInterleaved.Mod(montParamNonInterleaved, rVal)

	if montParamNonInterleaved.ModInverse(montParamNonInterleaved, rVal) == nil {
		return errors.New("modinverse failed")
	}

	rInv := new(big.Int)
	if rInv.ModInverse(rVal, modInt) == nil {
		return errors.New("modinverse to compute rInv failed")
	}

	m.NumLimbs = limbCount
	m.r = rVal
	m.rInv = rInv
	m.mask = rValMask

	// mod % (1 << limb_count_bits)  == mod % (1 << limb_count_bytes * 8)
	m.ModulusNonInterleaved = modInt

	m.Modulus = IntToLimbs(modInt, m.NumLimbs)

	m.MontParamNonInterleaved = montParamNonInterleaved
	m.MontParamInterleaved = Word(montParamNonInterleaved.Uint64())
    m.montMul = montgomeryFixedWidth[len(mod) - 1]

	return nil
}
