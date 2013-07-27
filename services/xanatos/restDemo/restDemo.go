package main

import (
    "code.google.com/p/gorest"
    "net/http"
)

func main(){
    gorest.RegisterService(new(OrderService))
    http.Handle("/",gorest.Handle())    
    http.ListenAndServe(":8787",nil)

}

//************************Define Service***************************

type User struct{
    Id int
    Name string
}

type OrderService struct{
    gorest.RestService    `root:"/service/" consumes:"application/json" produces:"application/json"`

    userDetails gorest.EndPoint `method:"GET" path:"/users/{Id:int}" output:"User"`
    helloWorld gorest.EndPoint `method:"GET" path:"/hello/" output:"string"`
    helloName gorest.EndPoint `method:"GET" path:"/hello/{Name:string}" output:"string"`
}

//Handler Methods: Method names must be the same as in config, but exported (starts with uppercase)

func(serv OrderService) HelloWorld() string{
    return "Hello World"
}

func(serv OrderService) HelloName(Name string) string{
    return "Hello " + Name
}

func(serv OrderService) UserDetails(id int) User{
    u := User {id, "Random name"}
    return u
}
