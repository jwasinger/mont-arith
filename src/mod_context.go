package evmmax_arith

import (
	"math/big"
	"unsafe"
	"errors"
)

type MontArithContext struct {
// TODO make most of these private and the arith operations methods of this struct
    Modulus []uint64
    ModulusNonInterleaved *big.Int // just here for convenience

    MontParamInterleaved uint64
    MontParamNonInterleaved *big.Int

    NumLimbs uint
    addModFunc ModArithFunc
    subModFunc ModArithFunc

    mulModMontWrapperFunc MulModWrapperFunc
    mulModMontFunc MulModMontFunc
}

func MulModMontInterleavedWrapper(out, x, y []byte, m *ModContext) {
	m.mulModMontFunc(out, x, y, m.ModulusInterleaved, m.ModInvInterleaved)
}

func MulModMontNonInterleavedWrapper(out, x, y []byte, m *ModContext) {
	MulModMontNonInterleaved(out, x, y, m.ModulusNonInterleaved, m.MontConstNonInterleaved)
}

func NewModContext() *ModContext {
	result := ModContext {
		nil,
		nil,
		0,
		nil,
		0,
		nil,
		nil,
	}
}

func (m *MontArithContext) AddMod(out, x, y []byte) {
	m.addModFunc(out, x, y, m.Modulus)
}

func (m *MontArithContext) SubMod(out, x, y []byte) {
	m.subModFunc(out, x, y, m.Modulus)
}

func (m *MontArithContext) MulModMont(out, x, y []byte) {
	// pass m explicityl b/c mulModMontWrapperFunc is a struct member
	m.mulModMontWrapperFunc(m, out, x, y)
}

func (m *MontArithContext) ModIsSet() bool {
	return m.NumLimbs != 0
}

func (m *MontArithContext) SetMod(modulus string, base int) (*ModContext, error) {
	mod, success := new(big.Int).SetString(modulus, base)
	if !success {
		return nil, errors.New("SetString failed for modulus/base")
	}

	var x uint
	limbSize := uint(unsafe.Sizeof(x))
	limbCount := uint(len(mod.Bytes())) / limbSize
	_ = x

	if limbCount == 0 {
		return nil, errors.New("can't have 0 as modulus")
	}

	// r val chosen as max representable value for limbCount + 1: 0x1000...000
	rVal := new(big.Int)
	rVal.Lsh(big.NewInt(1), limbCount * limbSize * 8)

	montParamNonInterleaved := new(big.Int)
	montParamNonInterleaved = montParamNonInterleaved.Mul(mod, big.NewInt(-1))
	if montParamNonInterleaved.ModInverse(montParamNonInterleaved, rVal) == nil {
		return nil, errors.New("modinverse failed")
	}

	// mod % (1 << limb_count_bits)  == mod % (1 << limb_count_bytes * 8)
	m.ModulusNonInterleaved = mod
	m.Modulus = IntToLimbs(mod)
	m.MontParamNonInterleaved = rVal
	m.MontParamInterleaved = rVal.Uint64()
	m.NumLimbs = len(m.Modulus)

	addModImpls := AddModImpls()
	subModImpls := SubModImpls()

	m.addModFunc = addModImpls[limbCount - 1]
	m.subModFunc = subModImpls[limbCount - 1]

	// limb count of 10 is the threshold where goff's mulmodmont impl blows up in runtime
	if limbCount < 10 {
		mulModMontImpls := MulModMontImpls()
		m.MulModWrapper = MulModMontInterleavedWrapper
		m.mulModMontFunc = mulModMontImpls[limbCount - 1]
	} else {
		m.MulModWrapper = MulModMontNonInterleavedWrapper
		m.MulModMontFunc = MulModMontNonInterleaved
	}
}
