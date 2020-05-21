package main

import (
	"fmt"
	"reflect"
)

type User2 struct {
	Name string
	Age  int
}

func main() {
	u := User2{"张三", 20}
	t := reflect.TypeOf(u)
	fmt.Println(t)

	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)

	v := reflect.ValueOf(u)
	fmt.Println("===", v)

	u1 := v.Interface().(User2)
	fmt.Println("***", u1)

	t1 := v.Type()
	fmt.Println(t1)
}
