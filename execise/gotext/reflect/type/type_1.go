package main

import (
	"fmt"
	"reflect"
)

type X2 int

func main1() {
	var a X2 = 100
	t := reflect.TypeOf(a)

	fmt.Println(t.Name(), t.Kind()) //1,真实类型    2，底层类型

}
