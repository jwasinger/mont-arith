package mont_arith

// XXX implement the below methods using these types (conversions might make it awkward/slower)
type Word uint
type nat = []Word
const WordSize = 8 // XXX word size in bytes, hardcoded to 64bit limbs

func Eq(n, other nat) bool {
    if len(n) != len(other) {
        panic("unequal lengths")
    }

    for i := 0; i < len(n); i++ {
        if n[i] != other[i] {
            return false
        }
    }
    return true
}
