package main

import "sync"

func main5() {
	var wg sync.WaitGroup
	wg.Add(3)

	a, b := make(chan int), make(chan int)

	go func() { //接收端
		defer wg.Done()

		for {
			select {
			case x, ok := <-a:
				if !ok { //如果通道关闭，则设置为nil,阻塞
					a = nil
					break
				}
				println("a", x)
			case x, ok := <-b:
				if !ok {
					b = nil
					break
				}
				println("b", x)
			}

			if a == nil && b == nil { //全部结束，退出循环
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)

		for i := 0; i < 3; i++ {
			a <- i
		}
	}()

	go func() {
		defer wg.Done()
		defer close(b)

		for i := 0; i < 5; i++ {
			b <- i * 10
		}
	}()

	wg.Wait()
}
