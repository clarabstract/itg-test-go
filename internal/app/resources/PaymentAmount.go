package resources

import (
	"fmt"
	"github.com/rubyruy/itg-test/internal/pkg/currency"
	"github.com/rubyruy/itg-test/internal/pkg/mortgage"
	"net/http"
)

// PaymentAmount is a handler for retrieving the recurring payment amount for the given mortgage.
func PaymentAmount(response http.ResponseWriter, request *http.Request) {
	if assertMethod(response, request, http.MethodGet) {
		params := request.URL.Query()
		principal, err := currency.Parse(params.Get("askingPrice"))
		if err != nil {
			fmt.Fprintf(response, "Invalid asking price %v (reason: %v)", params.Get("askingPrice"), err)
		}
		paymentYears, err := currency.Parse(params.Get("paymentSchedule"))
		if err != nil {
			fmt.Fprintf(response, "Invalid payment schedule %v (reason: %v)", params.Get("paymentSchedule"), err)
		}

		var yearlyInterest float32 = 0.05
		// yearDuration := time.Hour * 24 * 365
		periods := 12 * uint32(paymentYears)
		response.Header().Add("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(response, "<tt>mortgage.CalculatePayments(%v, %v, %v)</tt>", currency.UAmount(principal), yearlyInterest/float32(periods), periods)
		payment := mortgage.CalculatePayments(currency.UAmount(principal), yearlyInterest/12, periods)
		fmt.Fprintf(response, "You are paying <tt>%v</tt> over <tt>%v</tt> periods.", payment, periods)
	}

}
