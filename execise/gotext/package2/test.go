package main

//
//import (
//	"fmt"
//	"gotest/execise/gotext/package2/lib"
//	"unsafe"
//
//	//"unsafe"
//)
//
//func main() {
//	d := lib.NewData()
//	d.Y = 200				//直接访问导出字段
//
//	p := (*struct {x int })(unsafe.Pointer(d))	//利用指针转换访问私有字段
//	p.x = 100
//
//	fmt.Printf("%+v\n", *d)
//}
