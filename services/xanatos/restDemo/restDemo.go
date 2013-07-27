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
    id int
    name string
}

type OrderService struct{
    //Service level config
    gorest.RestService    `root:"/service/" consumes:"application/json" produces:"application/json"`

    //End-Point level configs: Field names must be the same as the corresponding method names,
    // but not-exported (starts with lowercase)

    userDetails gorest.EndPoint `method:"GET" path:"/users/{Id:int}" output:"User"`
    helloWorld gorest.EndPoint `method:"GET" path:"/hello/" output:"string"`
    helloName gorest.EndPoint `method:"GET" path:"/hello/{Name:string}" output:"string"`
    //listItems   gorest.EndPoint `method:"GET" path:"/items/" output:"[]Item"`
    //addItem     gorest.EndPoint `method:"POST" path:"/items/" postdata:"Item"`

    //On a real app for placeOrder below, the POST URL would probably be just /orders/, this is just to
    // demo the ability of mixing post-data parameters with URL mapped parameters.
    //placeOrder  gorest.EndPoint `method:"POST" path:"/orders/new/{UserId:int}/{RequestDiscount:bool}" postdata:"Order"`
    //viewOrder     gorest.EndPoint `method:"GET" path:"/orders/{OrderId:int}" output:"Order"`
    //deleteOrder     gorest.EndPoint `method:"DELETE" path:"/orders/{OrderId:int}"`


}

//Handler Methods: Method names must be the same as in config, but exported (starts with uppercase)

func(serv OrderService) HelloWorld() string{
    return "Hello World"
}

func(serv OrderService) HelloName(Name string) string{
    return "Hello " + Name
}

func(serv OrderService) UserDetails(id int) User{
    return User {id, "Random name"}
}
