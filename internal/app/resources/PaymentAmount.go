package resources

import (
	"fmt"
	"html"
	"net/http"
)

// PaymentAmount is a handler for retrieving the recurring payment amount for the given mortgage.
func PaymentAmount(response http.ResponseWriter, request *http.Request) {
	if assertMethod(response, request, http.MethodGet) {
		fmt.Fprintf(response, "Helzlo, %q", html.EscapeString(request.URL.Path))
	}

}
