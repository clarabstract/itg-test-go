package app

import (
	"github.com/rubyruy/itg-test/internal/app/resources"
	"log"
	"net/http"
)

// ItgTestServer runs the ITGlue Test app server.
func ItgTestServer(port string) {
	http.HandleFunc("/payment-amount", resources.PaymentAmount)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
