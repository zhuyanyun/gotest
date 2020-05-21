package main

import (
    "fmt"
    "reflect"
)

func main()  {
    var a interface{} = nil
    var b interface{} = (*int)(nil)

    fmt.Println(a == nil)
    fmt.Println(b == nil, reflect.ValueOf(b).IsNil())
}