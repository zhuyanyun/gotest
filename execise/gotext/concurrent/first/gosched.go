package main

import "runtime"

func main88() {
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() { //任务a
		defer close(exit)

		go func() { //  任务b,  放在此处， 是为了确保a 优先执行
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a : ", i)

			if i == 1 { //让出当前线程， 调度执行b
				runtime.Gosched()
			}
		}

	}()
	<-exit
}
