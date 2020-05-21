package main

import (
    "reflect"
    "fmt"
)

func main3()  {
    type user struct {
	Name string
	Age int
    }

    u := user{
	"q.yuan",
	60,
    }

    v := reflect.ValueOf(&u)

    if !v.CanInterface(){
	println("CanInterface: fail.")
	return
    }

    p, ok := v.Interface().(*user)
    if !ok {
	println("Interface:fale.")
	return
    }

    p.Age++
    fmt.Printf("%+v\n", u)


}


