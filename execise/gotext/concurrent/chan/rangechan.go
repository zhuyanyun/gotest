package main

func main7() {
	done := make(chan struct{})
	c := make(chan int)

	go func() {
		defer close(done)

		for x := range c {
			println(x)
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	<-done
}
