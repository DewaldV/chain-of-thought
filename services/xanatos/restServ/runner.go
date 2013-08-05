package main

import (
    "net/http"
	"code.google.com/p/gorest"
	"github.com/DewaldV/chain-of-thought/services/xanatos/leads"
)

func main(){
	leads.Register()
    http.Handle("/",gorest.Handle())    
    http.ListenAndServe(":8787",nil)
}