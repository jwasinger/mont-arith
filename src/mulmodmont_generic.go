package modext

import (
	"math/big"
	"errors"
)

func MulModMontNonInterleaved(out_bytes, x_bytes, y_bytes []byte, m *MontArithContext) (error) {
	// length x == y assumed

	product := new(big.Int)
	x := MAXBytesToInt(x_bytes)
	y := MAXBytesToInt(y_bytes)

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
	x.Rsh(x, m.NumLimbs * 64)

	copy(out_bytes, LimbsToMAXBytes(IntToLimbs(x, m.NumLimbs)))

	return nil
}
