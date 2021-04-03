package modext

/*
import (
	"math/big"
	"testing"
)
*/

/*
func BenchmarkMulModMontNonInterleaved(b *testing.B) {
	modCtx, err := NewModContext("1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab", 16)
	_ = modCtx
	_ = err

	// 5 limb - val = (1<<512) - 1
	x, _ := new(big.Int).SetString("13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084095", 10)
	y, _ := new(big.Int).SetString("13407807929942597099574024998205846127479365820592393377723561443721764030073546976801874298166903427690031858186486050853753882811946569946433649006084095", 10)
	z := new(big.Int)

	z.Mul(x, y)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		z := new(big.Int)
		z.SetBytes(x.Bytes())

		z = new(big.Int)
		z.SetBytes(x.Bytes())

		z = new(big.Int)
		z.SetBytes(x.Bytes())

		z.Mul(x, y)
		z.Mul(x, y)
		z.Mul(x, y)
		z.Mul(x, y)
	}
}
*/

/*
func rSquared(limbCount uint) Element {

}

func testcase(limbBits uint) bool {
	mod := big.NewInt(1).Lsh(limbBits * 8).Sub(big.NewInt(3)).String()
	modCtx := NewModContext(mod)
	limbCount := uint(math.Ceil(float(limbBits) / 8.0))
	x_limbs := make([]uint64, limbCount, limbCount)
	one_limbs := make([]uint64, limbCount, limbCount)

	x_limbs[0] = 3
	one[0] = 1

	return true
}

func TestMulModMontNonInterleaved(t *testing.T) {
	for i := uint(1); i < 256; i++ {
		testcase(i * 64)
	}
}
*/
