package main

func main2() {
	c := make(chan int, 3)

	c <- 10
	c <- 20

	close(c)

	for i := 0; i < cap(c)+1; i++ {
		x, ok := <-c
		println(i, ":", ok, x)
	}

}
