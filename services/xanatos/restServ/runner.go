package main

import (
	"code.google.com/p/gorest"
	_ "flag"
	"fmt"
	"github.com/DewaldV/chain-of-thought/services/xanatos/leads"
	"github.com/DewaldV/crucible"
	"net/http"
)

/*
var (
	validConfigPaths = flag.String("c", "test.conf", "Config path")
)
*/
type Registerable interface {
	Register() bool
}

func main() {
	//flag.Parse()

	/*
		err := crucible.LoadConfig(*validConfigPaths)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	*/
	RegisterServices()

	http.Handle(crucible.Conf.RootContext, crucible.Handle())
	http.ListenAndServe(fmt.Sprintf(":%d", crucible.Conf.HttpPort), nil)
}

func RegisterServices() {
	leads.Register()
	crucible.AddHandler(gorest.Handle())
}
