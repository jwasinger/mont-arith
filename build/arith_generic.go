package modext

import (
	"math/big"
)

func MulModMontNonInterleaved(z, x, y, mod, montParam *big.Int, limbBits uint) {
	//x := new(big.Int).SetBytes(

	// length x == y assumed

	// TODO see if pre-allocating product space and zeroing after use is faster for higher widths

	// m <- ((T mod R)N`) mod R
	// m is the low product of x*y (T) and N`

	// t <- (T + mN) / R
	// add T and mN and shift by R

	// if t > N: t <- T - N
	product := new(big.Int)

	product.Mul(x, y)
	x.Lsh(product, limbBits)
	x.Mul(x, montParam)
	x.Lsh(x, limbBits)

	x.Mul(x, mod)
	product.Add(product, x)

	// take top bits instead of shift

	// TODO
	_ = z
}
