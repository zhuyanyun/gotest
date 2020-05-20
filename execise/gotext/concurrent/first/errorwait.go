package main

import "sync"

func main66() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Add(1) //来不及设置

		//defer wg.Done()
		println("hi!")
	}()

	wg.Wait()
	println("exit.")

}
