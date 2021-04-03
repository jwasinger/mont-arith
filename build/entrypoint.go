package modext

func AddModImpls() []ModArithFunc {
	result := []ModArithFunc{
		AddMod256,
		AddMod320,
		AddMod384,
		AddMod448,
		AddMod512,
		AddMod576,
		AddMod640,
	}

	return result
}

func SubModImpls() []ModArithFunc {
	result := []ModArithFunc{
		SubMod256,
		SubMod320,
		SubMod384,
		SubMod448,
		SubMod512,
		SubMod576,
		SubMod640,
	}

	return result
}

func MulModMontImpls() []MulModMontFunc {
	result := []MulModMontFunc{
		MulModMont256,
		MulModMont320,
		MulModMont384,
		MulModMont448,
		MulModMont512,
		MulModMont576,
		MulModMont640,
	}

	return result
}
