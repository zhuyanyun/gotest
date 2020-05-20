package main

func main3() {
	c := make(chan int, 3) //创建带3个缓冲槽的异步通道

	c <- 1 //缓冲区未满， 不会阻塞
	c <- 2

	println(<-c) //缓冲区尚有数据，不会阻塞
	println(<-c)

}
