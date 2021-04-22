package mont_arith

type ArithPreset struct {
	AddModImpls     []ModArithFunc
	SubModImpls     []ModArithFunc
	MulModMontImpls []ModArithFunc
}

// Preset same as default except it uses blst's go-asm impl of the arithmetic at 384bit widths
func Asm384Preset() *ArithPreset {
	addModImpls := []ModArithFunc{
		AddMod64,
		AddModNonUnrolled128,
		AddModNonUnrolled192,
		AddModNonUnrolled256,
		AddModNonUnrolled320,
		AddMod384_asm,
		AddModNonUnrolled448,
		AddModNonUnrolled512,
		AddModNonUnrolled576,
		AddModNonUnrolled640,
		AddModNonUnrolled704,
		AddModNonUnrolled768,
		AddModNonUnrolled832,
		AddModNonUnrolled896,
		AddModNonUnrolled960,
		AddModNonUnrolled1024,
		AddModNonUnrolled1088,
		AddModNonUnrolled1152,
		AddModNonUnrolled1216,
		AddModNonUnrolled1280,
		AddModNonUnrolled1344,
		AddModNonUnrolled1408,
		AddModNonUnrolled1472,
		AddModNonUnrolled1536,
		AddModNonUnrolled1600,
		AddModNonUnrolled1664,
		AddModNonUnrolled1728,
		AddModNonUnrolled1792,
		AddModNonUnrolled1856,
		AddModNonUnrolled1920,
		AddModNonUnrolled1984,
		AddModNonUnrolled2048,
		AddModNonUnrolled2112,
		AddModNonUnrolled2176,
		AddModNonUnrolled2240,
		AddModNonUnrolled2304,
		AddModNonUnrolled2368,
		AddModNonUnrolled2432,
		AddModNonUnrolled2496,
		AddModNonUnrolled2560,
		AddModNonUnrolled2624,
		AddModNonUnrolled2688,
		AddModNonUnrolled2752,
		AddModNonUnrolled2816,
		AddModNonUnrolled2880,
		AddModNonUnrolled2944,
		AddModNonUnrolled3008,
		AddModNonUnrolled3072,
		AddModNonUnrolled3136,
		AddModNonUnrolled3200,
		AddModNonUnrolled3264,
		AddModNonUnrolled3328,
		AddModNonUnrolled3392,
		AddModNonUnrolled3456,
		AddModNonUnrolled3520,
		AddModNonUnrolled3584,
		AddModNonUnrolled3648,
		AddModNonUnrolled3712,
		AddModNonUnrolled3776,
		AddModNonUnrolled3840,
		AddModNonUnrolled3904,
		AddModNonUnrolled3968,
		AddModNonUnrolled4032,
		AddModNonUnrolled4096,
	}

	subModImpls := []ModArithFunc{
		SubMod64,
		SubMod128,
		SubMod192,
		SubMod256,
		SubMod320,
		SubMod384_asm,
		SubMod448,
		SubMod512,
		SubMod576,
		SubMod640,
		SubMod704,
		SubMod768,
		SubMod832,
		SubMod896,
		SubMod960,
		SubMod1024,
		SubMod1088,
		SubMod1152,
		SubMod1216,
		SubMod1280,
		SubMod1344,
		SubMod1408,
		SubMod1472,
		SubMod1536,
		SubMod1600,
		SubMod1664,
		SubMod1728,
		SubMod1792,
		SubMod1856,
		SubMod1920,
		SubMod1984,
		SubMod2048,
		SubMod2112,
		SubMod2176,
		SubMod2240,
		SubMod2304,
		SubMod2368,
		SubMod2432,
		SubMod2496,
		SubMod2560,
		SubMod2624,
		SubMod2688,
		SubMod2752,
		SubMod2816,
		SubMod2880,
		SubMod2944,
		SubMod3008,
		SubMod3072,
		SubMod3136,
		SubMod3200,
		SubMod3264,
		SubMod3328,
		SubMod3392,
		SubMod3456,
		SubMod3520,
		SubMod3584,
		SubMod3648,
		SubMod3712,
		SubMod3776,
		SubMod3840,
		SubMod3904,
		SubMod3968,
		SubMod4032,
		SubMod4096,
	}
	mulMontImpls := []ModArithFunc{

		MulModMont64,

		MulModMont128,

		MulModMont192,

		MulModMont256,

		MulModMont320,
		MulMont384_asm,

		MulModMont448,

		MulModMont512,

		MulModMont576,

		MulModMont640,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,
	}

	return &ArithPreset{addModImpls, subModImpls, mulMontImpls}
}

func DefaultPreset() *ArithPreset {
	return &ArithPreset{AddModNonUnrolledImpls(), SubModNonUnrolledImpls(), MulModMontImpls()}
}

func AddModNonUnrolledImpls() []ModArithFunc {
	result := []ModArithFunc{
		AddMod64,
		AddModNonUnrolled128,
		AddModNonUnrolled192,
		AddModNonUnrolled256,
		AddModNonUnrolled320,
		AddModNonUnrolled384,
		AddModNonUnrolled448,
		AddModNonUnrolled512,
		AddModNonUnrolled576,
		AddModNonUnrolled640,
		AddModNonUnrolled704,
		AddModNonUnrolled768,
		AddModNonUnrolled832,
		AddModNonUnrolled896,
		AddModNonUnrolled960,
		AddModNonUnrolled1024,
		AddModNonUnrolled1088,
		AddModNonUnrolled1152,
		AddModNonUnrolled1216,
		AddModNonUnrolled1280,
		AddModNonUnrolled1344,
		AddModNonUnrolled1408,
		AddModNonUnrolled1472,
		AddModNonUnrolled1536,
		AddModNonUnrolled1600,
		AddModNonUnrolled1664,
		AddModNonUnrolled1728,
		AddModNonUnrolled1792,
		AddModNonUnrolled1856,
		AddModNonUnrolled1920,
		AddModNonUnrolled1984,
		AddModNonUnrolled2048,
		AddModNonUnrolled2112,
		AddModNonUnrolled2176,
		AddModNonUnrolled2240,
		AddModNonUnrolled2304,
		AddModNonUnrolled2368,
		AddModNonUnrolled2432,
		AddModNonUnrolled2496,
		AddModNonUnrolled2560,
		AddModNonUnrolled2624,
		AddModNonUnrolled2688,
		AddModNonUnrolled2752,
		AddModNonUnrolled2816,
		AddModNonUnrolled2880,
		AddModNonUnrolled2944,
		AddModNonUnrolled3008,
		AddModNonUnrolled3072,
		AddModNonUnrolled3136,
		AddModNonUnrolled3200,
		AddModNonUnrolled3264,
		AddModNonUnrolled3328,
		AddModNonUnrolled3392,
		AddModNonUnrolled3456,
		AddModNonUnrolled3520,
		AddModNonUnrolled3584,
		AddModNonUnrolled3648,
		AddModNonUnrolled3712,
		AddModNonUnrolled3776,
		AddModNonUnrolled3840,
		AddModNonUnrolled3904,
		AddModNonUnrolled3968,
		AddModNonUnrolled4032,
		AddModNonUnrolled4096,
	}

	return result
}

func SubModNonUnrolledImpls() []ModArithFunc {
	result := []ModArithFunc{
		SubMod64,
		SubMod128,
		SubMod192,
		SubMod256,
		SubMod320,
		SubMod384,
		SubMod448,
		SubMod512,
		SubMod576,
		SubMod640,
		SubMod704,
		SubMod768,
		SubMod832,
		SubMod896,
		SubMod960,
		SubMod1024,
		SubMod1088,
		SubMod1152,
		SubMod1216,
		SubMod1280,
		SubMod1344,
		SubMod1408,
		SubMod1472,
		SubMod1536,
		SubMod1600,
		SubMod1664,
		SubMod1728,
		SubMod1792,
		SubMod1856,
		SubMod1920,
		SubMod1984,
		SubMod2048,
		SubMod2112,
		SubMod2176,
		SubMod2240,
		SubMod2304,
		SubMod2368,
		SubMod2432,
		SubMod2496,
		SubMod2560,
		SubMod2624,
		SubMod2688,
		SubMod2752,
		SubMod2816,
		SubMod2880,
		SubMod2944,
		SubMod3008,
		SubMod3072,
		SubMod3136,
		SubMod3200,
		SubMod3264,
		SubMod3328,
		SubMod3392,
		SubMod3456,
		SubMod3520,
		SubMod3584,
		SubMod3648,
		SubMod3712,
		SubMod3776,
		SubMod3840,
		SubMod3904,
		SubMod3968,
		SubMod4032,
		SubMod4096,
	}

	return result
}

func MulModMontImpls() []ModArithFunc {
	result := []ModArithFunc{

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

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,
	}

	return result
}
