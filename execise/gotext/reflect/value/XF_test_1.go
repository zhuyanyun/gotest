package main

import (
	"fmt"
	"reflect"
    "time"
)

type User2 struct {
	Name string
	Age  int
}

func (u User2) Print(prfix string){
    fmt.Printf("%s ;  Name is %s Age is %d", prfix, u.Name, u.Age)
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

    time.Sleep(time.Second *3)
    println("================","动态调用方法")


    mPrint := v.MethodByName("Print")
    args := []reflect.Value{reflect.ValueOf("前缀")}
    fmt.Println(mPrint.Call(args))
}
