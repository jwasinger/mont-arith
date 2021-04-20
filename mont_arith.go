package mont_arith

import (
	"github.com/jwasinger/mont-arith/arith"
)

type MontArithContext = arith.MontArithContext

func NewMontArithContext() *MontArithContext {
	return arith.NewMontArithContext(arith.DefaultPreset())
}
