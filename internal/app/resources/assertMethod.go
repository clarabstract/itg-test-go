package resources

import (
	"fmt"
	"net/http"
)

func assertMethod(response http.ResponseWriter, request *http.Request, method string) bool {
	if request.Method == method {
		return true
	}
	response.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(response, "<h1>Method %s is not allowed</h1>", method)
	return false

}
