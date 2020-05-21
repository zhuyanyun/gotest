package main

import (
	"fmt"
	"reflect"
)

type X int
type Y int

func main2() {
	var a, b X = 100, 200
	var c Y = 300

	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)

	fmt.Println(ta)
	fmt.Println(ta == tb, ta == tc)
	fmt.Println(ta.Kind() == tc.Kind())

}
