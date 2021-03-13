package arith256

// /!\ WARNING /!\
// this code has not been audited and is provided as-is. In particular,
// there is no security guarantees such as constant time implementation
// or side-channel attack resistance
// /!\ WARNING /!\

import (
	"math/big"
	"math/bits"
	"unsafe"
	"errors"
)

type Element{{limb_count}} [{{limb_count}}]uint64
var ZeroElement{{limb_count}} = make(Element{{limb_count}}, {{limb_count}}, {{limb_count}})

func (z *Element{{limb_count}}) MulModMont(x, y *Element{{limb_count}}, modCtx *ModCtx) (error) {

	//var t [4]uint64
    var t Element{{limb_count}}
	var c [{{limb_count}}]uint64
	var sub_val *Element = modCtx.Modulus

	if x[{{limb_count - 1}}] > mod[{{limb_count - 1}}] || y[{{limb_count - 1}}] > mod[{{limb_count - 1}}] {
		return errors.New("x/y must be smaller than modulus")
	}

	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, y[0])
		m := c[0] * modinv
		c[2] = madd0(m, mod[0], c[0])
		c[1], c[0] = madd1(v, y[1], c[1])
		c[2], t[0] = madd2(m, mod[1], c[2], c[0])
		c[1], c[0] = madd1(v, y[2], c[1])
		c[2], t[1] = madd2(m, mod[2], c[2], c[0])
		c[1], c[0] = madd1(v, y[3], c[1])
		t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * modinv
		c[2] = madd0(m, mod[0], c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, mod[1], c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, mod[2], c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * modinv
		c[2] = madd0(m, mod[0], c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, mod[1], c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, mod[2], c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, mod[3], c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * modinv
		c[2] = madd0(m, mod[0], c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], z[0] = madd2(m, mod[1], c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], z[1] = madd2(m, mod[2], c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		z[3], z[2] = madd3(m, mod[3], c[0], c[2], c[1])
	}

	_, c[1] = bits.Sub64(z[0], mod[0], 0)
	_, c[1] = bits.Sub64(z[1], mod[1], c[1])
	_, c[1] = bits.Sub64(z[2], mod[2], c[1])
	_, c[1] = bits.Sub64(z[3], mod[3], c[1])

	if c[1] != 0 { // unnecessary sub
		sub_val = &ZeroElement
	}

	z[0], c[0] = bits.Sub64(z[0], sub_val[0], 0)
	z[1], c[0] = bits.Sub64(z[1], sub_val[1], c[0])
	z[2], c[0] = bits.Sub64(z[2], sub_val[2], c[0])
	z[3], c[0] = bits.Sub64(z[3], sub_val[3], c[0])

	return nil
}

/*
	Modular Addition
*/
func (out *Element) AddMod(x, y, mod *Element) (error) {
	var c uint64
	var tmp Element

	if x[3] > mod[3] || y[3] > mod[3] {
		return errors.New("x/y must be smaller than modulus")
	}

	tmp[0], c = bits.Add64(x[0], y[0], 0)
	tmp[1], c = bits.Add64(x[1], y[1], c)
	tmp[2], c = bits.Add64(x[2], y[2], c)
	tmp[3], c = bits.Add64(x[3], y[3], c)

	out[0], c = bits.Sub64(tmp[0], mod[0], 0)
	out[1], c = bits.Sub64(tmp[1], mod[1], c)
	out[2], c = bits.Sub64(tmp[2], mod[2], c)
	out[3], c = bits.Sub64(tmp[3], mod[3], c)

	if c != 0 { // unnecessary sub
		*out = tmp
	}

	return nil
}

/*
	Modular Subtraction
*/
func (out *Element) SubMod(x, y, mod *Element) (error) {
	var c, c1 uint64
	var tmp Element

	if x[3] > mod[3] || y[3] > mod[3] {
		return errors.New("x/y must be smaller than modulus")
	}

	tmp[0], c1 = bits.Sub64(x[0], y[0], 0)
	tmp[1], c1 = bits.Sub64(x[1], y[1], c1)
	tmp[2], c1 = bits.Sub64(x[2], y[2], c1)
	tmp[3], c1 = bits.Sub64(x[3], y[3], c1)

	out[0], c = bits.Add64(tmp[0], mod[0], 0)
	out[1], c = bits.Add64(tmp[1], mod[1], c)
	out[2], c = bits.Add64(tmp[2], mod[2], c)
	out[3], c = bits.Add64(tmp[3], mod[3], c)

	if c1 == 0 { // unnecessary add
		*out = tmp
	}

	return nil
}
