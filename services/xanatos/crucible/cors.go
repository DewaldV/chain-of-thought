package crucible

import (
	"code.google.com/p/gorest"
	"fmt"
	"net/http"
)

func AddOptionsCors(rb *gorest.ResponseBuilder, request *http.Request, sc *ServiceConfigStruct) error {
	origin := request.Header.Get("origin")
	originAllowed := sc.AllowedOrigins[origin]

	if !originAllowed {
		origin = ""
	}

	var allowedMethods string
	for key := range sc.AllowedMethods {
		allowedMethods += fmt.Sprintf("%s,", key)
	}
	allowedMethods = allowedMethods[:len(allowedMethods)-1]

	rb.AddHeader("Access-Control-Allow-Origin", origin)
	rb.AddHeader("Access-Control-Allow-Method", allowedMethods)
	rb.AddHeader("Access-Control-Allow-Headers", "accept, origin, x-requested-with, content-type")
	return nil
}

func AddAllowOriginsCors(rb *gorest.ResponseBuilder, request *http.Request, sc *ServiceConfigStruct) error {
	origin := request.Header.Get("origin")
	originAllowed := sc.AllowedOrigins[origin]

	if !originAllowed {
		origin = ""
	}

	rb.AddHeader("Access-Control-Allow-Origin", origin)
	return nil
}
