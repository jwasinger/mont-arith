package arith

import (
	"errors"
	"math/big"
)

func MulModMontNonInterleaved(out_bytes, x_bytes, y_bytes []byte, m *MontArithContext) error {
	// length x == y assumed

	product := new(big.Int)
	x := LEBytesToInt(x_bytes)
	y := LEBytesToInt(y_bytes)

	if x.Cmp(m.ModulusNonInterleaved) > 0 || y.Cmp(m.ModulusNonInterleaved) > 0 {
		return errors.New("x/y >= modulus")
	}

	// m <- ((x*y mod R)N`) mod R
	product.Mul(x, y)
	x.And(product, m.mask)
	x.Mul(x, m.MontParamNonInterleaved)
	x.And(x, m.mask)

	// t <- (T + mN) / R
	x.Mul(x, m.ModulusNonInterleaved)
	x.Add(x, product)
	x.Rsh(x, m.NumLimbs*64)

	copy(out_bytes, LimbsToLEBytes(IntToLimbs(x, m.NumLimbs)))

	return nil
}
