package main

// First-response-wins strategy
func request() int {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}

	return <-ch
}
