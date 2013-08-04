package main

import (
    "net/http"
	"code.google.com/p/gorest"
	"github.com/DewaldV/chain-of-thought/services/xanatos/endDemo"
)

func main(){
	endDemo.Register()
    http.Handle("/",gorest.Handle())    
    http.ListenAndServe(":8787",nil)
}