package main

import "sync"

type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}

	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data {
			println("recv:", x) //接收消息，直到通道被关闭
		}
	}()

	return r
}

func main1() {
	r := newReceiver()
	r.data <- 1
	r.data <- 2

	close(r.data)
	r.Wait()
}
