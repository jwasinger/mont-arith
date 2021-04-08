package modext

func DefaultPreset() ([]ModArithFunc, []ModArithFunc, []ModArithFunc) {
	return MulModMontImpls(), AddModImpls(), SubModImpls()
}

func AddModImpls() []ModArithFunc {
	result := []ModArithFunc{
		AddMod64,
		AddMod128,
		AddMod192,
		AddMod256,
		AddMod320,
		AddMod384,
		AddMod448,
		AddMod512,
		AddMod576,
		AddMod640,
		AddMod704,
		AddMod768,
		AddMod832,
		AddMod896,
		AddMod960,
		AddMod1024,
		AddMod1088,
		AddMod1152,
		AddMod1216,
		AddMod1280,
		AddMod1344,
		AddMod1408,
		AddMod1472,
		AddMod1536,
		AddMod1600,
		AddMod1664,
		AddMod1728,
		AddMod1792,
		AddMod1856,
		AddMod1920,
		AddMod1984,
		AddMod2048,
		AddMod2112,
		AddMod2176,
		AddMod2240,
		AddMod2304,
		AddMod2368,
		AddMod2432,
		AddMod2496,
		AddMod2560,
		AddMod2624,
		AddMod2688,
		AddMod2752,
		AddMod2816,
		AddMod2880,
		AddMod2944,
		AddMod3008,
		AddMod3072,
		AddMod3136,
		AddMod3200,
		AddMod3264,
		AddMod3328,
		AddMod3392,
		AddMod3456,
		AddMod3520,
		AddMod3584,
		AddMod3648,
		AddMod3712,
		AddMod3776,
		AddMod3840,
		AddMod3904,
		AddMod3968,
		AddMod4032,
		AddMod4096,
		AddMod4160,
		AddMod4224,
		AddMod4288,
		AddMod4352,
		AddMod4416,
		AddMod4480,
		AddMod4544,
		AddMod4608,
		AddMod4672,
		AddMod4736,
		AddMod4800,
		AddMod4864,
		AddMod4928,
		AddMod4992,
		AddMod5056,
		AddMod5120,
		AddMod5184,
		AddMod5248,
		AddMod5312,
		AddMod5376,
		AddMod5440,
		AddMod5504,
		AddMod5568,
		AddMod5632,
		AddMod5696,
		AddMod5760,
		AddMod5824,
		AddMod5888,
		AddMod5952,
		AddMod6016,
		AddMod6080,
		AddMod6144,
		AddMod6208,
		AddMod6272,
		AddMod6336,
		AddMod6400,
		AddMod6464,
		AddMod6528,
		AddMod6592,
		AddMod6656,
		AddMod6720,
		AddMod6784,
		AddMod6848,
		AddMod6912,
		AddMod6976,
		AddMod7040,
		AddMod7104,
		AddMod7168,
		AddMod7232,
		AddMod7296,
		AddMod7360,
		AddMod7424,
		AddMod7488,
		AddMod7552,
		AddMod7616,
		AddMod7680,
		AddMod7744,
		AddMod7808,
		AddMod7872,
		AddMod7936,
		AddMod8000,
		AddMod8064,
		AddMod8128,
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
