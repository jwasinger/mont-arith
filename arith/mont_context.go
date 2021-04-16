package arith

import (
	"errors"
	"math/big"
)

type ModArithFunc func(out, x, y []byte, m *MontArithContext) error

type MontArithContext struct {
	// TODO make most of these private and the arith operations methods of this struct
	Modulus               []byte
	ModulusNonInterleaved *big.Int // just here for convenience

	MontParamInterleaved    uint64
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

	// store list of all generated unrolled implementations (TODO use generic implementation at higher bitwidth)
	addmodImpls     []ModArithFunc
	submodImpls     []ModArithFunc
	mulmodMontImpls []ModArithFunc
}

func (m *MontArithContext) RVal() *big.Int {
	return m.r
}

func (m *MontArithContext) RInv() *big.Int {
	return m.rInv
}

func (m *MontArithContext) ToMont(dst, src []uint64) {
	if len(dst) != len(src) || uint(len(dst)) != m.NumLimbs {
		panic("dst and src length must be equal to number of limbs for modulus")
	}

	dst_val := new(big.Int)
	src_val := LimbsToInt(src)
	dst_val.Mul(src_val, m.r)
	dst_val.Mod(dst_val, LEBytesToInt(m.Modulus))

	copy(dst, IntToLimbs(dst_val, m.NumLimbs))
}

func (m *MontArithContext) ToNorm(dst, src []uint64) {
	if len(dst) != len(src) || uint(len(dst)) != m.NumLimbs {
		panic("dst and src length must be equal to number of limbs for modulus")
	}

	dst_val := new(big.Int)
	src_val := LimbsToInt(src)
	dst_val.Mul(src_val, m.rInv)
	dst_val.Mod(dst_val, LEBytesToInt(m.Modulus))

	copy(dst, IntToLimbs(dst_val, m.NumLimbs))
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
		nil,
		nil,

		nil,
		nil,
		nil,
	}

	return &result
}

func (m *MontArithContext) AddMod(out, x, y []byte) {
	m.addModFunc(out, x, y, m)
}

func (m *MontArithContext) SubMod(out, x, y []byte) {
	m.subModFunc(out, x, y, m)
}

func (m *MontArithContext) MulModMont(out, x, y []byte) {
	// pass m explicitly b/c mulModMontWrapperFunc is a struct member
	m.mulModMontFunc(out, x, y, m)
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
	/*
		fmt.Println(mod.String())
		fmt.Println(IntToLimbs(mod, m.NumLimbs))
		fmt.Println(LimbsToLEBytes(IntToLimbs(mod, m.NumLimbs)))
	*/

	// r val chosen as max representable value for limbCount + 1: 0x1000...000
	rVal := new(big.Int)
	rVal.Lsh(big.NewInt(1), limbCount*limbSize*8)

	rValMask := new(big.Int)
	rValMask.Sub(rVal, big.NewInt(1))

	montParamNonInterleaved := new(big.Int)
	montParamNonInterleaved = montParamNonInterleaved.Mul(mod, big.NewInt(-1))
	montParamNonInterleaved.Mod(montParamNonInterleaved, rVal)

	if montParamNonInterleaved.ModInverse(montParamNonInterleaved, rVal) == nil {
		return errors.New("modinverse failed")
	}

	//rVal.Mod(rVal, mod)

	rInv := new(big.Int)
	if rInv.ModInverse(rVal, mod) == nil {
		return errors.New("modinverse to compute rInv failed")
	}

	m.NumLimbs = limbCount
	m.r = rVal
	m.rInv = rInv
	m.mask = rValMask

	// mod % (1 << limb_count_bits)  == mod % (1 << limb_count_bytes * 8)
	m.ModulusNonInterleaved = mod

	m.Modulus = LimbsToLEBytes(IntToLimbs(mod, m.NumLimbs))

	m.MontParamNonInterleaved = montParamNonInterleaved
	m.MontParamInterleaved = montParamNonInterleaved.Uint64()

	m.mulmodMontImpls, m.addmodImpls, m.submodImpls = NonUnrolledPreset()

	m.mulModMontFunc = m.mulmodMontImpls[limbCount-1]
	m.addModFunc = m.addmodImpls[limbCount-1]
	m.subModFunc = m.submodImpls[limbCount-1]

	return nil
}
