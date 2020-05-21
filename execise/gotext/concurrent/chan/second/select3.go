package main

import "sync"

//0 为什么有两个？？？
func main6() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)

	go func() { //发送端
		defer wg.Done()

		for {
			var v int
			var ok bool

			select { //随机选择case
			case v, ok = <-c:
				println("a1:", v)
			case v, ok = <-c:
				println("a2:", v)
			}

			if !ok {
				return
			}
		}
	}()

	go func() { //发送端
		defer wg.Done()
		defer close(c)

		for i := 0; i < 10; i++ {
			select { //随机选择case
			case c <- i:
			case c <- i * 10:
			}
		}
	}()

	wg.Wait()
}
