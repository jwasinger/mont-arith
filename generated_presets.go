package mont_arith






type ArithPreset struct {
	AddModImpls []ModArithFunc
	SubModImpls []ModArithFunc
	MulModMontImpls []ModArithFunc
}

func MulModMontImpls() []ModArithFunc {
	result := []ModArithFunc {
        MulModMont64,
            MulModMont128,
            MulModMont192,
            MulModMont256,
            MulModMont320,
            MulModMont384,
            MulModMont448,
            MulModMont512,
            MulModMont576,
            MulModMont640,
            MulModMont704,
            MulModMont768,
	}

	return result
}

func DefaultPreset() *ArithPreset {
	return &ArithPreset{nil, nil, MulModMontImpls()}
}
