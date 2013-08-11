package crucible

import (
	"code.google.com/p/gorest"
	"fmt"
	"net/http"
)

func AddOptionsCors(rb *gorest.ResponseBuilder, request *http.Request, sc *ServiceConfigStruct) error {
	origin := request.Header.Get("origin")
	AddAllowOriginsCors(rb, origin, sc.AllowedOrigins)

	method := request.Header.Get("access-control-allow-method")
	AddAllowMethodCors(rb, method, sc.AllowedMethods)
	AddAllowHeadrCors(rb)
	return nil
}

func AddAllowOriginsCors(rb *gorest.ResponseBuilder, origin string, allowedOrigins *map[string]bool) error {
	originAllowed := allowedOrigins[origin]

	if originAllowed {
		rb.AddHeader("Access-Control-Allow-Origin", origin)
	}

	return nil
}

func AddAllowMethodCors(rb *gorest.ResponseBuilder, method string, allowedMethods *map[string]bool) error {
	var allowedMethodsStr string
	for key := range sc.AllowedMethods {
		allowedMethods += fmt.Sprintf("%s,", key)
	}
	allowedMethods = allowedMethods[:len(allowedMethods)-1]

	rb.AddHeader("Access-Control-Allow-Method", allowedMethods)
	return nil
}

func AddAllowHeadrCors(rb *gorest.ResponseBuilder) {
	rb.AddHeader("Access-Control-Allow-Headers", "accept, origin, x-requested-with, content-type")
}
