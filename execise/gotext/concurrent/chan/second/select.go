package main

import "sync"

func main4() {
	var wg sync.WaitGroup
	wg.Add(2)

	a, b := make(chan int), make(chan int)

	go func() { //接收端
		defer wg.Done()

		for {
			var (
				name string
				x    int
				ok   bool
			)

			select { //随机选择可用 channel 接收数据
			case x, ok = <-a:
				name = "a"
			case x, ok = <-b:
				name = "b"
			}

			if !ok { //如果任一通道关闭,则终止接收
				return
			}

			println(name, x) //输出接收的数据信息
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i := 0; i < 10; i++ {
			select { //随机选择发送channel
			case a <- i:
			case b <- i * 10:
			}

		}
	}()

	wg.Wait()
}
