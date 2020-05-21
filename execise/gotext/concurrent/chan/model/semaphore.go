package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main3() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup

	sem := make(chan struct{}, 2) //最多允许2个并发同时执行

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        //acquire:获取信号
			defer func() { <-sem }() //release:释放信号

			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}

	wg.Wait()
}
