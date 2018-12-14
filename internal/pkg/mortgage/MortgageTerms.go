package mortgage

import (
	"github.com/rubyruy/itg-test/internal/pkg/currency"
	"time"
)

type MortgageTerms struct {
	askingPrice        currency.Amount
	downPayment        currency.Amount
	paymentSchedule    PaymentType
	amortizationPeriod time.Duration
}
