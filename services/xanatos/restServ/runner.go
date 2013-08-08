package main

import (
	"code.google.com/p/gorest"
	"fmt"
	"github.com/DewaldV/chain-of-thought/services/xanatos/coreServ"
	"github.com/DewaldV/chain-of-thought/services/xanatos/leads"
	"net/http"
)

const validConfigPaths string = "/Users/DewaldV/Dev/test.conf"

type Registerable interface {
	Register() bool
}

func main() {
	leads.Register()
	err := coreServ.LoadConfig(validConfigPaths)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	http.Handle(coreServ.Conf.RootContext, gorest.Handle())
	http.ListenAndServe(string(coreServ.Conf.HttpPort), nil)
}
