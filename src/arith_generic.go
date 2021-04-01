package modext

import (
	"math/big"
	"fmt"
)

func MulModMontNonInterleaved(z, x, y, mod, montParam *big.Int, limbCount uint) {
	//x := new(big.Int).SetBytes(

	// length x == y assumed

	// TODO see if pre-allocating product and zeroing after use is faster

	// m <- ((T mod R)N`) mod R
	// m is the low product of x*y (T) and N`

	// t <- (T + mN) / R
	// add T and mN and shift by R

	// if t > N: t <- T - N
	product := new(big.Int)
	var limbBits uint = 64

	product.Mul(x, y)
	x.Lsh(product, limbCount * limbBits)
	x.Mul(x, montParam)
	x.Lsh(x, limbCount * limbBits)

	x.Mul(x, mod)
	product.Add(product, x)

	// take top bits instead of shift

	fmt.Println(product.String())

	// TODO
	_ = z
}
