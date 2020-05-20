package main

import "fmt"

type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("p: %p, %v\n", n, *n)
}

func main() {
	var a N = 5

	a.value()
	a.pointer()

	fmt.Printf("a: %p, %v\n", &a, a)
}
