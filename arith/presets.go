package arith

func NonUnrolledPreset() ([]ModArithFunc, []ModArithFunc, []ModArithFunc) {
	return MulModMontImpls(), AddModNonUnrolledImpls(), SubModImpls()
}

func DefaultPreset() ([]ModArithFunc, []ModArithFunc, []ModArithFunc) {
	return MulModMontImpls(), AddModUnrolledImpls(), SubModImpls()
}

func AddModUnrolledImpls() []ModArithFunc {
	result := []ModArithFunc{
		AddModNonUnrolled64,
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
		AddModNonUnrolled4160,
		AddModNonUnrolled4224,
		AddModNonUnrolled4288,
		AddModNonUnrolled4352,
		AddModNonUnrolled4416,
		AddModNonUnrolled4480,
		AddModNonUnrolled4544,
		AddModNonUnrolled4608,
		AddModNonUnrolled4672,
		AddModNonUnrolled4736,
		AddModNonUnrolled4800,
		AddModNonUnrolled4864,
		AddModNonUnrolled4928,
		AddModNonUnrolled4992,
		AddModNonUnrolled5056,
		AddModNonUnrolled5120,
		AddModNonUnrolled5184,
		AddModNonUnrolled5248,
		AddModNonUnrolled5312,
		AddModNonUnrolled5376,
		AddModNonUnrolled5440,
		AddModNonUnrolled5504,
		AddModNonUnrolled5568,
		AddModNonUnrolled5632,
		AddModNonUnrolled5696,
		AddModNonUnrolled5760,
		AddModNonUnrolled5824,
		AddModNonUnrolled5888,
		AddModNonUnrolled5952,
		AddModNonUnrolled6016,
		AddModNonUnrolled6080,
		AddModNonUnrolled6144,
		AddModNonUnrolled6208,
		AddModNonUnrolled6272,
		AddModNonUnrolled6336,
		AddModNonUnrolled6400,
		AddModNonUnrolled6464,
		AddModNonUnrolled6528,
		AddModNonUnrolled6592,
		AddModNonUnrolled6656,
		AddModNonUnrolled6720,
		AddModNonUnrolled6784,
		AddModNonUnrolled6848,
		AddModNonUnrolled6912,
		AddModNonUnrolled6976,
		AddModNonUnrolled7040,
		AddModNonUnrolled7104,
		AddModNonUnrolled7168,
		AddModNonUnrolled7232,
		AddModNonUnrolled7296,
		AddModNonUnrolled7360,
		AddModNonUnrolled7424,
		AddModNonUnrolled7488,
		AddModNonUnrolled7552,
		AddModNonUnrolled7616,
		AddModNonUnrolled7680,
		AddModNonUnrolled7744,
		AddModNonUnrolled7808,
		AddModNonUnrolled7872,
		AddModNonUnrolled7936,
		AddModNonUnrolled8000,
		AddModNonUnrolled8064,
		AddModNonUnrolled8128,
	}

	return result
}

func AddModUnrolledImpls() []ModArithFunc {
	result := []ModArithFunc{
		AddModUnrolled64,
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
		AddModUnrolled4160,
		AddModUnrolled4224,
		AddModUnrolled4288,
		AddModUnrolled4352,
		AddModUnrolled4416,
		AddModUnrolled4480,
		AddModUnrolled4544,
		AddModUnrolled4608,
		AddModUnrolled4672,
		AddModUnrolled4736,
		AddModUnrolled4800,
		AddModUnrolled4864,
		AddModUnrolled4928,
		AddModUnrolled4992,
		AddModUnrolled5056,
		AddModUnrolled5120,
		AddModUnrolled5184,
		AddModUnrolled5248,
		AddModUnrolled5312,
		AddModUnrolled5376,
		AddModUnrolled5440,
		AddModUnrolled5504,
		AddModUnrolled5568,
		AddModUnrolled5632,
		AddModUnrolled5696,
		AddModUnrolled5760,
		AddModUnrolled5824,
		AddModUnrolled5888,
		AddModUnrolled5952,
		AddModUnrolled6016,
		AddModUnrolled6080,
		AddModUnrolled6144,
		AddModUnrolled6208,
		AddModUnrolled6272,
		AddModUnrolled6336,
		AddModUnrolled6400,
		AddModUnrolled6464,
		AddModUnrolled6528,
		AddModUnrolled6592,
		AddModUnrolled6656,
		AddModUnrolled6720,
		AddModUnrolled6784,
		AddModUnrolled6848,
		AddModUnrolled6912,
		AddModUnrolled6976,
		AddModUnrolled7040,
		AddModUnrolled7104,
		AddModUnrolled7168,
		AddModUnrolled7232,
		AddModUnrolled7296,
		AddModUnrolled7360,
		AddModUnrolled7424,
		AddModUnrolled7488,
		AddModUnrolled7552,
		AddModUnrolled7616,
		AddModUnrolled7680,
		AddModUnrolled7744,
		AddModUnrolled7808,
		AddModUnrolled7872,
		AddModUnrolled7936,
		AddModUnrolled8000,
		AddModUnrolled8064,
		AddModUnrolled8128,
	}

	return result
}

func SubModImpls() []ModArithFunc {
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
		SubMod4160,
		SubMod4224,
		SubMod4288,
		SubMod4352,
		SubMod4416,
		SubMod4480,
		SubMod4544,
		SubMod4608,
		SubMod4672,
		SubMod4736,
		SubMod4800,
		SubMod4864,
		SubMod4928,
		SubMod4992,
		SubMod5056,
		SubMod5120,
		SubMod5184,
		SubMod5248,
		SubMod5312,
		SubMod5376,
		SubMod5440,
		SubMod5504,
		SubMod5568,
		SubMod5632,
		SubMod5696,
		SubMod5760,
		SubMod5824,
		SubMod5888,
		SubMod5952,
		SubMod6016,
		SubMod6080,
		SubMod6144,
		SubMod6208,
		SubMod6272,
		SubMod6336,
		SubMod6400,
		SubMod6464,
		SubMod6528,
		SubMod6592,
		SubMod6656,
		SubMod6720,
		SubMod6784,
		SubMod6848,
		SubMod6912,
		SubMod6976,
		SubMod7040,
		SubMod7104,
		SubMod7168,
		SubMod7232,
		SubMod7296,
		SubMod7360,
		SubMod7424,
		SubMod7488,
		SubMod7552,
		SubMod7616,
		SubMod7680,
		SubMod7744,
		SubMod7808,
		SubMod7872,
		SubMod7936,
		SubMod8000,
		SubMod8064,
		SubMod8128,
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
