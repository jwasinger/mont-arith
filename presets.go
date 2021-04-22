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

func UnrolledPreset() *ArithPreset {
	return &ArithPreset{AddModUnrolledImpls(), SubModUnrolledImpls(), MulModMontImpls()}
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

func AddModUnrolledImpls() []ModArithFunc {
	result := []ModArithFunc{
		AddMod64,
		AddModUnrolled128,
		AddModUnrolled192,
		AddModUnrolled256,
		AddModUnrolled320,
		AddModUnrolled384,
		AddModUnrolled448,
		AddModUnrolled512,
		AddModUnrolled576,
		AddModUnrolled640,
		AddModUnrolled704,
		AddModUnrolled768,
		AddModUnrolled832,
		AddModUnrolled896,
		AddModUnrolled960,
		AddModUnrolled1024,
		AddModUnrolled1088,
		AddModUnrolled1152,
		AddModUnrolled1216,
		AddModUnrolled1280,
		AddModUnrolled1344,
		AddModUnrolled1408,
		AddModUnrolled1472,
		AddModUnrolled1536,
		AddModUnrolled1600,
		AddModUnrolled1664,
		AddModUnrolled1728,
		AddModUnrolled1792,
		AddModUnrolled1856,
		AddModUnrolled1920,
		AddModUnrolled1984,
		AddModUnrolled2048,
		AddModUnrolled2112,
		AddModUnrolled2176,
		AddModUnrolled2240,
		AddModUnrolled2304,
		AddModUnrolled2368,
		AddModUnrolled2432,
		AddModUnrolled2496,
		AddModUnrolled2560,
		AddModUnrolled2624,
		AddModUnrolled2688,
		AddModUnrolled2752,
		AddModUnrolled2816,
		AddModUnrolled2880,
		AddModUnrolled2944,
		AddModUnrolled3008,
		AddModUnrolled3072,
		AddModUnrolled3136,
		AddModUnrolled3200,
		AddModUnrolled3264,
		AddModUnrolled3328,
		AddModUnrolled3392,
		AddModUnrolled3456,
		AddModUnrolled3520,
		AddModUnrolled3584,
		AddModUnrolled3648,
		AddModUnrolled3712,
		AddModUnrolled3776,
		AddModUnrolled3840,
		AddModUnrolled3904,
		AddModUnrolled3968,
		AddModUnrolled4032,
		AddModUnrolled4096,
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

func SubModUnrolledImpls() []ModArithFunc {
	result := []ModArithFunc{
		SubMod64,
		SubModUnrolled128,
		SubModUnrolled192,
		SubModUnrolled256,
		SubModUnrolled320,
		SubModUnrolled384,
		SubModUnrolled448,
		SubModUnrolled512,
		SubModUnrolled576,
		SubModUnrolled640,
		SubModUnrolled704,
		SubModUnrolled768,
		SubModUnrolled832,
		SubModUnrolled896,
		SubModUnrolled960,
		SubModUnrolled1024,
		SubModUnrolled1088,
		SubModUnrolled1152,
		SubModUnrolled1216,
		SubModUnrolled1280,
		SubModUnrolled1344,
		SubModUnrolled1408,
		SubModUnrolled1472,
		SubModUnrolled1536,
		SubModUnrolled1600,
		SubModUnrolled1664,
		SubModUnrolled1728,
		SubModUnrolled1792,
		SubModUnrolled1856,
		SubModUnrolled1920,
		SubModUnrolled1984,
		SubModUnrolled2048,
		SubModUnrolled2112,
		SubModUnrolled2176,
		SubModUnrolled2240,
		SubModUnrolled2304,
		SubModUnrolled2368,
		SubModUnrolled2432,
		SubModUnrolled2496,
		SubModUnrolled2560,
		SubModUnrolled2624,
		SubModUnrolled2688,
		SubModUnrolled2752,
		SubModUnrolled2816,
		SubModUnrolled2880,
		SubModUnrolled2944,
		SubModUnrolled3008,
		SubModUnrolled3072,
		SubModUnrolled3136,
		SubModUnrolled3200,
		SubModUnrolled3264,
		SubModUnrolled3328,
		SubModUnrolled3392,
		SubModUnrolled3456,
		SubModUnrolled3520,
		SubModUnrolled3584,
		SubModUnrolled3648,
		SubModUnrolled3712,
		SubModUnrolled3776,
		SubModUnrolled3840,
		SubModUnrolled3904,
		SubModUnrolled3968,
		SubModUnrolled4032,
		SubModUnrolled4096,
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
