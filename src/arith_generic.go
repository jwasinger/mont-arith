package evmmax_arith

import (
	"math/big"
	"fmt"
)

func MulModMontNonInterleaved(z, x, y *big.Int, ctx *ModContext) {
	//x := new(big.Int).SetBytes(

	// length x == y assumed

	// TODO see if pre-allocating product and zeroing after use is faster

	// m <- ((T mod R)N`) mod R
	// m is the low product of x*y (T) and N`

	// t <- (T + mN) / R
	// add T and mN and shift by R

	// if t > N: t <- T - N
	product := new(big.Int)

	product.Mul(x, y)
	x.Lsh(product, ctx.NumLimbs * 64)
	x.Mul(x, ctx.MontParamNonInterleaved)
	x.Lsh(x, ctx.NumLimbs * 64)

	x.Mul(x, ctx.ModulusNonInterleaved)
	product.Add(product, x)

	// take top bits instead of shift

	fmt.Println(product.String())

	// TODO
	_ = z
}
