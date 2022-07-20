package mont_arith

import (
	"errors"
	"fmt"
	"math/big"
)

type ModArithFunc func(out, x, y []uint, m *MontArithContext) error

type MontArithContext struct {
	// TODO make most of these private and the arith operations methods of this struct
	Modulus               []uint
	ModulusNonInterleaved *big.Int // just here for convenience XXX better naming

	MontParamInterleaved    uint
	MontParamNonInterleaved *big.Int

	NumLimbs uint

	r    *big.Int
	rInv *big.Int

	// mask for mod by R: 0xfff...fff - (1 << NumLimbs * 64) - 1
	mask *big.Int

	// currently active addmod/submod/mulmodmont for a set NumLimbs
	addModFunc     ModArithFunc
	subModFunc     ModArithFunc
	mulModMontFunc ModArithFunc

	arithImpl ArithPreset
}

func (m *MontArithContext) RVal() *big.Int {
	return m.r
}

func (m *MontArithContext) RInv() *big.Int {
	return m.rInv
}

func (m *MontArithContext) ToMont(dst, src []uint) {
	if len(dst) != len(src) || uint(len(dst)) != m.NumLimbs {
		panic("dst and src length must be equal to number of limbs for modulus")
	}

	dst_val := new(big.Int)
	src_val := LimbsToInt(src)
	dst_val.Mul(src_val, m.r)
	dst_val.Mod(dst_val, LimbsToInt(m.Modulus))

	copy(dst, IntToLimbs(dst_val, m.NumLimbs))
}

func (m *MontArithContext) ToNorm(dst, src []uint) {
	if len(dst) != len(src) || uint(len(dst)) != m.NumLimbs {
		panic("dst and src length must be equal to number of limbs for modulus")
	}

	dst_val := new(big.Int)
	src_val := LimbsToInt(src)
	dst_val.Mul(src_val, m.rInv)
	dst_val.Mod(dst_val, LimbsToInt(m.Modulus))

	copy(dst, IntToLimbs(dst_val, m.NumLimbs))
}

func NewMontArithContext(preset *ArithPreset) *MontArithContext {
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
		nil,
		nil,

		*preset,
	}

	return &result
}

func (m *MontArithContext) AddMod(out, x, y []uint) {
	// pass m explicitly b/c mulModMontWrapperFunc is a struct member
	m.addModFunc(out, x, y, m)
}

func (m *MontArithContext) SubMod(out, x, y []uint) {
	m.subModFunc(out, x, y, m)
}

func (m *MontArithContext) MulModMont(out, x, y []uint) {
	m.mulModMontFunc(out, x, y, m)
}

func (m *MontArithContext) ModIsSet() bool {
	return m.NumLimbs != 0
}

func (m *MontArithContext) ValueSize() uint {
	return uint(len(m.Modulus))
}

func (m *MontArithContext) SetMod(mod []uint) error {
	// XXX proper handling
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
	m.MontParamInterleaved = uint(montParamNonInterleaved.Uint64())

	m.mulModMontFunc = m.arithImpl.MulModMontImpls[limbCount-1]
	m.addModFunc = nil // m.arithImpl.AddModImpls[limbCount-1]
	m.subModFunc = nil // m.arithImpl.SubModImpls[limbCount-1]

	return nil
}
