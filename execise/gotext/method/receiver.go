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

func main1002() {
	var a N = 5
	p := &a

	a.value()
	a.pointer()

	p.value()
	p.pointer()

	p.pointer()
	p.value()
	p.value()
	p.value()

	fmt.Printf("a: %p, %v\n", &a, a)
}
