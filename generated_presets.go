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
		AddModNonUnrolled8192,
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
		SubMod8192,
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

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

		MulModMontNonInterleaved,

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
		AddModNonUnrolled8192,
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
		SubMod8192,
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

		MulModMontNonInterleaved,
	}

	return result
}
