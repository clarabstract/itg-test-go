package mortgage

import (
	"math/big"
)

var one = big.NewInt(1)

func bigPow(base *big.Rat, exponent *big.Rat) {
	base.Num().Exp(base.Num(), exponent.Num(), one).Root
	base.Denom().Exp(base.Denom(), exponent.Num(), one)
}
