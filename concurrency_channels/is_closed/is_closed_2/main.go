package main

import "fmt"

func IsClosed(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(IsClosed(ch))
	close(ch)
	fmt.Println(IsClosed(ch))
}
