package mortgage

import (
	"github.com/rubyruy/itg-test/internal/pkg/currency"
	"math/big"
)

var one = big.NewRat(1, 1)

// CalculatePeriodicPayment returns the per-period payment for a given principal, subject to given interestRate over
// given number of periods.
func CalculatePeriodicPayment(principal currency.Amount, interestRate *big.Rat, periods uint32) currency.Amount {
	bigPeriods := big.NewRat(int64(periods), 1)
	// "big" operations all send their results to their receiever, in-place, so this works best if we accumulate the
	// calculation rather than a single expression
	numerator := &big.Rat{}
	numerator.Add(one, interestRate)
	bigPow(numerator, bigPeriods)
	numerator.Mul(interestRate, numerator)

	denominator := &big.Rat{}
	denominator.Add(one, interestRate)
	bigPow(denominator, bigPeriods)
	denominator.Sub(denominator, one)

	payment := &big.Rat{}
	payment.Quo(numerator, denominator)
	payment.Mul(&principal.Rat, payment)

	return currency.Amount{Rat: *payment}
}
