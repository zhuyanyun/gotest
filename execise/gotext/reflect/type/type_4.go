package main

import (
	"fmt"
	"reflect"
)

func main4() {
	x := 100
	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)

	fmt.Println(tx, tp, tx == tp)
	fmt.Println(tx.Kind(), tp.Kind())
	fmt.Println(tx == tp.Elem()) //Elem返回指针， 数组，切片，字典（值） 或通道的基类型

	fmt.Println(reflect.TypeOf(map[string]int{}).Elem())
	fmt.Println(reflect.TypeOf([]int32{}).Elem())
}
