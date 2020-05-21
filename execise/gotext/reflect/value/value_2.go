package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name string
	code int
}

func main2() {
	p := new(User)
	v := reflect.ValueOf(p).Elem()

	name := v.FieldByName("Name")
	code := v.FieldByName("code")

	fmt.Printf("name: canaddr = %v, canset = %v\n", name.CanAddr(), name.CanSet())
	fmt.Printf("code: canaddr = %v, canset = %v\n", code.CanAddr(), code.CanSet())

	if name.CanSet() {
		name.SetString("Tom")
	}

	if code.CanAddr() {
		*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100 //Pointer 返回字段所保存的地址，UnsafeAddr返回该字段自身的地址
	}

	fmt.Printf("%+v\n", *p)
}