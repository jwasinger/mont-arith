package modext

import (
	"errors"
	"math/bits"
	"unsafe"
)

func AddMod5696(out_bytes, x_bytes, y_bytes []byte, ctx *MontArithContext) error {
	x := (*[89]uint64)(unsafe.Pointer(&x_bytes[0]))[:]
	y := (*[89]uint64)(unsafe.Pointer(&y_bytes[0]))[:]
	z := (*[89]uint64)(unsafe.Pointer(&out_bytes[0]))[:]
	mod := (*[89]uint64)(unsafe.Pointer(&ctx.Modulus[0]))[:]

	var c uint64
	tmp := [89]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if x[88] >= mod[88] || y[88] >= mod[88] {
		panic(errors.New("x/y must be smaller than modulus"))
	}
	z[0], c = bits.Add64(x[0], y[0], 0)
	z[1], c = bits.Add64(x[1], y[1], c)
	z[2], c = bits.Add64(x[2], y[2], c)
	z[3], c = bits.Add64(x[3], y[3], c)
	z[4], c = bits.Add64(x[4], y[4], c)
	z[5], c = bits.Add64(x[5], y[5], c)
	z[6], c = bits.Add64(x[6], y[6], c)
	z[7], c = bits.Add64(x[7], y[7], c)
	z[8], c = bits.Add64(x[8], y[8], c)
	z[9], c = bits.Add64(x[9], y[9], c)
	z[10], c = bits.Add64(x[10], y[10], c)
	z[11], c = bits.Add64(x[11], y[11], c)
	z[12], c = bits.Add64(x[12], y[12], c)
	z[13], c = bits.Add64(x[13], y[13], c)
	z[14], c = bits.Add64(x[14], y[14], c)
	z[15], c = bits.Add64(x[15], y[15], c)
	z[16], c = bits.Add64(x[16], y[16], c)
	z[17], c = bits.Add64(x[17], y[17], c)
	z[18], c = bits.Add64(x[18], y[18], c)
	z[19], c = bits.Add64(x[19], y[19], c)
	z[20], c = bits.Add64(x[20], y[20], c)
	z[21], c = bits.Add64(x[21], y[21], c)
	z[22], c = bits.Add64(x[22], y[22], c)
	z[23], c = bits.Add64(x[23], y[23], c)
	z[24], c = bits.Add64(x[24], y[24], c)
	z[25], c = bits.Add64(x[25], y[25], c)
	z[26], c = bits.Add64(x[26], y[26], c)
	z[27], c = bits.Add64(x[27], y[27], c)
	z[28], c = bits.Add64(x[28], y[28], c)
	z[29], c = bits.Add64(x[29], y[29], c)
	z[30], c = bits.Add64(x[30], y[30], c)
	z[31], c = bits.Add64(x[31], y[31], c)
	z[32], c = bits.Add64(x[32], y[32], c)
	z[33], c = bits.Add64(x[33], y[33], c)
	z[34], c = bits.Add64(x[34], y[34], c)
	z[35], c = bits.Add64(x[35], y[35], c)
	z[36], c = bits.Add64(x[36], y[36], c)
	z[37], c = bits.Add64(x[37], y[37], c)
	z[38], c = bits.Add64(x[38], y[38], c)
	z[39], c = bits.Add64(x[39], y[39], c)
	z[40], c = bits.Add64(x[40], y[40], c)
	z[41], c = bits.Add64(x[41], y[41], c)
	z[42], c = bits.Add64(x[42], y[42], c)
	z[43], c = bits.Add64(x[43], y[43], c)
	z[44], c = bits.Add64(x[44], y[44], c)
	z[45], c = bits.Add64(x[45], y[45], c)
	z[46], c = bits.Add64(x[46], y[46], c)
	z[47], c = bits.Add64(x[47], y[47], c)
	z[48], c = bits.Add64(x[48], y[48], c)
	z[49], c = bits.Add64(x[49], y[49], c)
	z[50], c = bits.Add64(x[50], y[50], c)
	z[51], c = bits.Add64(x[51], y[51], c)
	z[52], c = bits.Add64(x[52], y[52], c)
	z[53], c = bits.Add64(x[53], y[53], c)
	z[54], c = bits.Add64(x[54], y[54], c)
	z[55], c = bits.Add64(x[55], y[55], c)
	z[56], c = bits.Add64(x[56], y[56], c)
	z[57], c = bits.Add64(x[57], y[57], c)
	z[58], c = bits.Add64(x[58], y[58], c)
	z[59], c = bits.Add64(x[59], y[59], c)
	z[60], c = bits.Add64(x[60], y[60], c)
	z[61], c = bits.Add64(x[61], y[61], c)
	z[62], c = bits.Add64(x[62], y[62], c)
	z[63], c = bits.Add64(x[63], y[63], c)
	z[64], c = bits.Add64(x[64], y[64], c)
	z[65], c = bits.Add64(x[65], y[65], c)
	z[66], c = bits.Add64(x[66], y[66], c)
	z[67], c = bits.Add64(x[67], y[67], c)
	z[68], c = bits.Add64(x[68], y[68], c)
	z[69], c = bits.Add64(x[69], y[69], c)
	z[70], c = bits.Add64(x[70], y[70], c)
	z[71], c = bits.Add64(x[71], y[71], c)
	z[72], c = bits.Add64(x[72], y[72], c)
	z[73], c = bits.Add64(x[73], y[73], c)
	z[74], c = bits.Add64(x[74], y[74], c)
	z[75], c = bits.Add64(x[75], y[75], c)
	z[76], c = bits.Add64(x[76], y[76], c)
	z[77], c = bits.Add64(x[77], y[77], c)
	z[78], c = bits.Add64(x[78], y[78], c)
	z[79], c = bits.Add64(x[79], y[79], c)
	z[80], c = bits.Add64(x[80], y[80], c)
	z[81], c = bits.Add64(x[81], y[81], c)
	z[82], c = bits.Add64(x[82], y[82], c)
	z[83], c = bits.Add64(x[83], y[83], c)
	z[84], c = bits.Add64(x[84], y[84], c)
	z[85], c = bits.Add64(x[85], y[85], c)
	z[86], c = bits.Add64(x[86], y[86], c)
	z[87], c = bits.Add64(x[87], y[87], c)
	z[88], c = bits.Add64(x[88], y[88], c)
	tmp[0], c = bits.Sub64(tmp[0], mod[0], 0)
	tmp[1], c = bits.Sub64(tmp[1], mod[1], c)
	tmp[2], c = bits.Sub64(tmp[2], mod[2], c)
	tmp[3], c = bits.Sub64(tmp[3], mod[3], c)
	tmp[4], c = bits.Sub64(tmp[4], mod[4], c)
	tmp[5], c = bits.Sub64(tmp[5], mod[5], c)
	tmp[6], c = bits.Sub64(tmp[6], mod[6], c)
	tmp[7], c = bits.Sub64(tmp[7], mod[7], c)
	tmp[8], c = bits.Sub64(tmp[8], mod[8], c)
	tmp[9], c = bits.Sub64(tmp[9], mod[9], c)
	tmp[10], c = bits.Sub64(tmp[10], mod[10], c)
	tmp[11], c = bits.Sub64(tmp[11], mod[11], c)
	tmp[12], c = bits.Sub64(tmp[12], mod[12], c)
	tmp[13], c = bits.Sub64(tmp[13], mod[13], c)
	tmp[14], c = bits.Sub64(tmp[14], mod[14], c)
	tmp[15], c = bits.Sub64(tmp[15], mod[15], c)
	tmp[16], c = bits.Sub64(tmp[16], mod[16], c)
	tmp[17], c = bits.Sub64(tmp[17], mod[17], c)
	tmp[18], c = bits.Sub64(tmp[18], mod[18], c)
	tmp[19], c = bits.Sub64(tmp[19], mod[19], c)
	tmp[20], c = bits.Sub64(tmp[20], mod[20], c)
	tmp[21], c = bits.Sub64(tmp[21], mod[21], c)
	tmp[22], c = bits.Sub64(tmp[22], mod[22], c)
	tmp[23], c = bits.Sub64(tmp[23], mod[23], c)
	tmp[24], c = bits.Sub64(tmp[24], mod[24], c)
	tmp[25], c = bits.Sub64(tmp[25], mod[25], c)
	tmp[26], c = bits.Sub64(tmp[26], mod[26], c)
	tmp[27], c = bits.Sub64(tmp[27], mod[27], c)
	tmp[28], c = bits.Sub64(tmp[28], mod[28], c)
	tmp[29], c = bits.Sub64(tmp[29], mod[29], c)
	tmp[30], c = bits.Sub64(tmp[30], mod[30], c)
	tmp[31], c = bits.Sub64(tmp[31], mod[31], c)
	tmp[32], c = bits.Sub64(tmp[32], mod[32], c)
	tmp[33], c = bits.Sub64(tmp[33], mod[33], c)
	tmp[34], c = bits.Sub64(tmp[34], mod[34], c)
	tmp[35], c = bits.Sub64(tmp[35], mod[35], c)
	tmp[36], c = bits.Sub64(tmp[36], mod[36], c)
	tmp[37], c = bits.Sub64(tmp[37], mod[37], c)
	tmp[38], c = bits.Sub64(tmp[38], mod[38], c)
	tmp[39], c = bits.Sub64(tmp[39], mod[39], c)
	tmp[40], c = bits.Sub64(tmp[40], mod[40], c)
	tmp[41], c = bits.Sub64(tmp[41], mod[41], c)
	tmp[42], c = bits.Sub64(tmp[42], mod[42], c)
	tmp[43], c = bits.Sub64(tmp[43], mod[43], c)
	tmp[44], c = bits.Sub64(tmp[44], mod[44], c)
	tmp[45], c = bits.Sub64(tmp[45], mod[45], c)
	tmp[46], c = bits.Sub64(tmp[46], mod[46], c)
	tmp[47], c = bits.Sub64(tmp[47], mod[47], c)
	tmp[48], c = bits.Sub64(tmp[48], mod[48], c)
	tmp[49], c = bits.Sub64(tmp[49], mod[49], c)
	tmp[50], c = bits.Sub64(tmp[50], mod[50], c)
	tmp[51], c = bits.Sub64(tmp[51], mod[51], c)
	tmp[52], c = bits.Sub64(tmp[52], mod[52], c)
	tmp[53], c = bits.Sub64(tmp[53], mod[53], c)
	tmp[54], c = bits.Sub64(tmp[54], mod[54], c)
	tmp[55], c = bits.Sub64(tmp[55], mod[55], c)
	tmp[56], c = bits.Sub64(tmp[56], mod[56], c)
	tmp[57], c = bits.Sub64(tmp[57], mod[57], c)
	tmp[58], c = bits.Sub64(tmp[58], mod[58], c)
	tmp[59], c = bits.Sub64(tmp[59], mod[59], c)
	tmp[60], c = bits.Sub64(tmp[60], mod[60], c)
	tmp[61], c = bits.Sub64(tmp[61], mod[61], c)
	tmp[62], c = bits.Sub64(tmp[62], mod[62], c)
	tmp[63], c = bits.Sub64(tmp[63], mod[63], c)
	tmp[64], c = bits.Sub64(tmp[64], mod[64], c)
	tmp[65], c = bits.Sub64(tmp[65], mod[65], c)
	tmp[66], c = bits.Sub64(tmp[66], mod[66], c)
	tmp[67], c = bits.Sub64(tmp[67], mod[67], c)
	tmp[68], c = bits.Sub64(tmp[68], mod[68], c)
	tmp[69], c = bits.Sub64(tmp[69], mod[69], c)
	tmp[70], c = bits.Sub64(tmp[70], mod[70], c)
	tmp[71], c = bits.Sub64(tmp[71], mod[71], c)
	tmp[72], c = bits.Sub64(tmp[72], mod[72], c)
	tmp[73], c = bits.Sub64(tmp[73], mod[73], c)
	tmp[74], c = bits.Sub64(tmp[74], mod[74], c)
	tmp[75], c = bits.Sub64(tmp[75], mod[75], c)
	tmp[76], c = bits.Sub64(tmp[76], mod[76], c)
	tmp[77], c = bits.Sub64(tmp[77], mod[77], c)
	tmp[78], c = bits.Sub64(tmp[78], mod[78], c)
	tmp[79], c = bits.Sub64(tmp[79], mod[79], c)
	tmp[80], c = bits.Sub64(tmp[80], mod[80], c)
	tmp[81], c = bits.Sub64(tmp[81], mod[81], c)
	tmp[82], c = bits.Sub64(tmp[82], mod[82], c)
	tmp[83], c = bits.Sub64(tmp[83], mod[83], c)
	tmp[84], c = bits.Sub64(tmp[84], mod[84], c)
	tmp[85], c = bits.Sub64(tmp[85], mod[85], c)
	tmp[86], c = bits.Sub64(tmp[86], mod[86], c)
	tmp[87], c = bits.Sub64(tmp[87], mod[87], c)
	tmp[88], c = bits.Sub64(tmp[88], mod[88], c)

	if c == 0 {
		copy(z, tmp[:])
	}

	return nil
}
