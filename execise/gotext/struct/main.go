package main

import (
	"fmt"
	"unsafe"
)

func main1001() {

	v1 := struct {
		a byte
		b byte
		c int32
	}{}

	v2 := struct {
		a byte
		b byte
	}{}

	v3 := struct {
		a byte
		b []int
		c byte
	}{}

	fmt.Printf("v1:  %d, %d\n", unsafe.Alignof(v1), unsafe.Sizeof(v1))
	fmt.Printf("v2:  %d, %d\n", unsafe.Alignof(v2), unsafe.Sizeof(v2))
	fmt.Printf("v3:  %d, %d\n", unsafe.Alignof(v3), unsafe.Sizeof(v3))
}
