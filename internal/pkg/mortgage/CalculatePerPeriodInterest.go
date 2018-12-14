package mortgage

import (
	"math/big"
)

var one = big.NewRat(1, 1)

// CalculatePerPeriodInterest returns the per-period interest rate for `paymentsPerYear` periods in a year,
// where the interest compounds `compoundPeriod` times a year.
func CalculatePerPeriodInterest(interest *big.Rat, paymentsPerYear uint32, compoundPeriod uint32) big.Rat {
	ppi := &big.Rat{}
	ppi.Quo(interest, big.NewRat(int64(compoundPeriod), 1))
	ppi.Add(one, ppi)
	bigPow(ppi, big.NewRat(int64(compoundPeriod), int64(paymentsPerYear)))
	ppi.Sub(ppi, one)
	return ppi
}
