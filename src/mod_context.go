package modext

import (
	"math/big"
	"errors"
)

type ModArithFunc func (out, x, y, mod []byte) (error)
type MulModMontFunc func (out, x, y, mod []byte, modinv uint64) (error)
type MulModWrapperFunc func(out, x, y []byte, m *MontArithContext)

type MontArithContext struct {
// TODO make most of these private and the arith operations methods of this struct
    Modulus []byte
    ModulusNonInterleaved *big.Int // just here for convenience

    MontParamInterleaved uint64
    MontParamNonInterleaved *big.Int

    NumLimbs uint

    r *big.Int
    rInv *big.Int

    // currently active addmod/submod/mulmodmont for a set NumLimbs
    addModFunc ModArithFunc
    subModFunc ModArithFunc
    mulModMontWrapperFunc MulModWrapperFunc
    mulModMontFunc MulModMontFunc

    // store list of all generated unrolled implementations (TODO use generic implementation at higher bitwidth)
    addmodImpls []ModArithFunc
    submodImpls []ModArithFunc
    mulmodMontImpls []MulModMontFunc
}

func (m *MontArithContext) ToMont(dst, src []uint64) {
	if len(dst) != len(src) || uint(len(dst)) != m.NumLimbs {
		panic("dst and src length must be equal to number of limbs for modulus")
	}

	dst_val := new(big.Int)
	src_val := LimbsToInt(src)
	dst_val.Mul(src_val, m.r)
	dst_val.Mod(dst_val, MAXBytesToInt(m.Modulus))

	copy(dst, IntToLimbs(dst_val, m.NumLimbs))
}

func (m *MontArithContext) ToNorm(dst, src []uint64) {
	if len(dst) != len(src) || uint(len(dst)) != m.NumLimbs {
		panic("dst and src length must be equal to number of limbs for modulus")
	}

	dst_val := new(big.Int)
	src_val := LimbsToInt(src)
	dst_val.Mul(src_val, m.rInv)
	dst_val.Mod(dst_val, MAXBytesToInt(m.Modulus))

	copy(dst, IntToLimbs(dst_val, m.NumLimbs))
}

func MulModMontInterleavedWrapper(out, x, y []byte, m *MontArithContext) {
	m.mulModMontFunc(out, x, y, m.Modulus, m.MontParamInterleaved)
}

func MulModMontNonInterleavedWrapper(out_bytes, x_bytes, y_bytes []byte, m *MontArithContext) {
	// TODO replace with something more performant than big.Int (and with less or no copying)
	// this is just a reference implementation for the higher bitwidths.

	x := MAXBytesToInt(x_bytes)
	y := MAXBytesToInt(y_bytes)
	out := new(big.Int)

	MulModMontNonInterleaved(out, x, y, m.ModulusNonInterleaved, m.MontParamNonInterleaved, m.NumLimbs * 64)

	out_MAX := LimbsToMAXBytes(IntToLimbs(out, m.NumLimbs))

	copy(out_bytes, out_MAX)
}

func NewMontArithContext() *MontArithContext {
	result := MontArithContext {
		nil,
		nil,
		0,
		nil,

		0,
		nil,
		nil,

		nil,
		nil,
		nil,
		nil,

		nil,
		nil,
		nil,
	}

	return &result
}

func (m *MontArithContext) AddMod(out, x, y []byte) {
	m.addModFunc(out, x, y, m.Modulus)
}

func (m *MontArithContext) SubMod(out, x, y []byte) {
	m.subModFunc(out, x, y, m.Modulus)
}

func (m *MontArithContext) MulModMont(out, x, y []byte) {
	// pass m explicityl b/c mulModMontWrapperFunc is a struct member



	m.mulModMontWrapperFunc(out, x, y, m)

}

func (m *MontArithContext) ModIsSet() bool {
	return m.NumLimbs != 0
}

func (m *MontArithContext) SetMod(modulus []uint64) error {
/*
	mod, success := new(big.Int).SetString(modulus, base)
	if !success {
		return nil, errors.New("SetString failed for modulus/base")
	}
*/
	limbCount := uint(len(modulus))
	var limbSize uint = 8

	if limbCount == 0 {
		return errors.New("can't have 0 as modulus")
	}

	mod := LimbsToInt(modulus)


	// r val chosen as max representable value for limbCount + 1: 0x1000...000
	rVal := new(big.Int)
	rVal.Lsh(big.NewInt(1), limbCount * limbSize * 8)

	montParamNonInterleaved := new(big.Int)
	montParamNonInterleaved = montParamNonInterleaved.Mul(mod, big.NewInt(-1))
	montParamNonInterleaved.Mod(montParamNonInterleaved, rVal)

	if montParamNonInterleaved.ModInverse(montParamNonInterleaved, rVal) == nil {
		return errors.New("modinverse failed")
	}

	rVal.Mod(rVal, mod)


	rInv := new(big.Int)
	if rInv.ModInverse(rVal, mod) == nil {
		return errors.New("modinverse to compute rInv failed")
	}

	m.NumLimbs = limbCount
	m.r = rVal
	m.rInv = rInv

	// mod % (1 << limb_count_bits)  == mod % (1 << limb_count_bytes * 8)
	m.ModulusNonInterleaved = mod
	m.Modulus = LimbsToMAXBytes(IntToLimbs(mod, m.NumLimbs))
	m.MontParamNonInterleaved = montParamNonInterleaved
	m.MontParamInterleaved = montParamNonInterleaved.Uint64()

	m.addmodImpls = AddModImpls()
	m.submodImpls = SubModImpls()
	m.mulmodMontImpls = MulModMontImpls()

	// limb count of 10 is the threshold where goff's mulmodmont impl blows up in runtime
	if limbCount < 10 {

		if limbCount < 4 {
			panic("limb counts less than 4 not implemented")
		}

		m.mulModMontWrapperFunc = MulModMontInterleavedWrapper
		m.mulModMontFunc = m.mulmodMontImpls[limbCount - 4]
		m.addModFunc = m.addmodImpls[limbCount - 4]
		m.subModFunc = m.submodImpls[limbCount - 4]
	} else {
		m.mulModMontWrapperFunc = MulModMontNonInterleavedWrapper
		m.mulModMontFunc = nil // non-interleaved mulmodmont called by wrapper has a different function signature

		// TODO generic addmod/submod or more unrolled impls ?
	}

	return nil
}
