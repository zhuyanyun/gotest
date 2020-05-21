package main

//有点意思
func main() {
	done := make(chan struct{})

	data := []chan int{
		make(chan int, 3),
	}

	println("========", len(data))

	go func() {
		defer close(done)

		for i := 0; i < 10; i++ {
			select {
			case data[len(data)-1] <- i:
			default:
				data = append(data, make(chan int, 3))
			}
			println("========", len(data))

		}
	}()

	<-done

	for i := 0; i < len(data); i++ {
		c := data[i]
		close(c)
		for x := range c {
			println(x)
		}
	}

}
