package main

import (
	"code.google.com/p/gorest"
	"fmt"
	_ "github.com/DewaldV/chain-of-thought/services/xanatos/leads"
	"github.com/DewaldV/crucible"
	"net/http"
)

type Registerable interface {
	Register() bool
}

func main() {
	RegisterServices()

	http.Handle(crucible.Conf.RootContext, crucible.Handle())
	http.ListenAndServe(fmt.Sprintf(":%d", crucible.Conf.HttpPort), nil)
}

func RegisterServices() {
	crucible.AddFilter(crucible.CorsFilter())
	//crucible.AddFilter(crucible.CachingFilter())
	crucible.AddHandler(gorest.Handle())
}
