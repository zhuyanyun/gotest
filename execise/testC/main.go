package main

/*
#include "util.c"
*/
import "C"

import "fmt"

func GoSum(a, b int) int {
	s := C.sum(C.int(a), C.int(b))
	fmt.Println(s)
	return 0
}

func main() {
	GoSum(1, 2)
}
