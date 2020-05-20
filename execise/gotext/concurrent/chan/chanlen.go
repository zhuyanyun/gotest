package main

func main55() {
	a, b := make(chan int), make(chan int, 3)

	b <- 1
	b <- 2

	println("a:", len(a), cap(a))
	println("b:", len(b), cap(b))

}
