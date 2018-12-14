package currency

import (
	"math/big"
	"regexp"
	"strconv"
)

const (
	minorCurrencyBase = 100
)

var (
	cleanAmountStringPattern = regexp.MustCompile(`[$.,]`)
)

// Amount is any rational monetary amount without a specified currency.
type Amount struct {
	big.Rat
}

func (amount Amount) String() string {
	return amount.FloatString(2)
}

// Parse a monetary Amount from a string
func Parse(str string) (Amount, error) {
	cleanAmountString := cleanAmountStringPattern.ReplaceAllLiteralString(str, "")
	minorUnitAmount, err := strconv.ParseInt(cleanAmountString, 10, 64)
	if err != nil {
		return FromMinorUnit(0), err
	}
	return FromMinorUnit(minorUnitAmount), nil
}

// FromMinorUnit constructs a new amount from the minor currency unit (e.g. cents, pennies etc).
func FromMinorUnit(count int64) Amount {
	return Amount{Rat: *big.NewRat(count, minorCurrencyBase)}
}
